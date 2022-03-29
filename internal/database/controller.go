package database

import (
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/hashing"
	"assignment-2/internal/structs"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var app *firebase.App

// InitDB initializes the database setup.
func InitDB() {
	ctx = context.Background()
	opt := option.WithCredentialsFile("auth.json")
	app, _ = firebase.NewApp(ctx, nil, opt)
}

// GetDocument gets data from a collection and a document in the database.
func GetDocument(collection string, document string, structToExtractTo interface{}) {
	client, errClient := app.Firestore(ctx)

	if errClient != nil {
		log.Printf("Error creating db client: %v", errClient)
		return
	}

	defer func() {
		errClosingClient := client.Close()

		if errClosingClient != nil {
			log.Printf("Error closing db: %v", errClosingClient)
		}
	}()

	hashedCollection, errHashCol := hashing.HashString(collection)

	if errHashCol != nil {
		return
	}

	var hashedDoc string

	// Checking if the document name is already hashed.
	if len(document) != 64 {
		var errHashDoc error
		hashedDoc, errHashDoc = hashing.HashString(document)

		if errHashDoc != nil {
			return
		}
	} else {
		hashedDoc = document
	}

	res := client.Collection(hashedCollection).Doc(hashedDoc)

	doc, errGetDoc := res.Get(ctx)

	if !doc.Exists() {
		return
	}

	if errGetDoc != nil {
		log.Printf("Error getting document: %v", errGetDoc)
		return
	}

	err := doc.DataTo(&structToExtractTo)

	if err != nil {
		log.Printf("Error extracting data into struct: %v", err)
		return
	}
}

// GetAllWebhooks gets all the webhooks from the webhook collection in the db and turns it into a slice of
// structs.
func GetAllWebhooks(collection string, country string) ([]structs.WebHookRegistration, error) {
	client, errClient := app.Firestore(ctx)

	if errClient != nil {
		log.Printf("Error creating db client: %v", errClient)
		return nil, nil
	}

	defer func() {
		errClosingClient := client.Close()

		if errClosingClient != nil {
			log.Printf("Error closing db: %v", errClosingClient)
		}
	}()

	hashedCollection, errHashCol := hashing.HashString(collection)

	if errHashCol != nil {
		return nil, nil
	}

	var res *firestore.DocumentIterator

	if country == "" {
		res = client.Collection(hashedCollection).Documents(ctx)
	} else {
		res = client.Collection(hashedCollection).Where("Country", "==", country).Documents(ctx)
	}
	var webhooks []structs.WebHookRegistration

	// Iterate through all the documents in the collection.
	for {
		doc, err := res.Next()
		if err == iterator.Done {
			break
		}
		data := structs.WebHookRegistration{}

		err = doc.DataTo(&data)

		// Sets the output id to the document id.
		data.Id = doc.Ref.ID

		if err != nil {
			log.Printf("Data: %v", err)
			return nil, err
		}

		webhooks = append(webhooks, data)
	}

	return webhooks, nil
}

// WriteDocument creates or updates a document in a collection (and a document).
func WriteDocument(collection string, document string, data interface{}) error {
	client, errOpeningClient := app.Firestore(ctx)

	if errOpeningClient != nil {
		log.Println(errOpeningClient)
		return custom_errors.GetDatabaseError()
	}

	defer func() {
		errClosingClient := client.Close()

		if errClosingClient != nil {
			log.Printf("Error closing db: %v", errClosingClient)
		}
	}()

	hashedCollection, _ := hashing.HashString(collection)
	hashedDoc, _ := hashing.HashString(document)

	//TODO error handling
	_, errSet := client.Collection(hashedCollection).Doc(hashedDoc).Set(ctx, data)

	if errSet != nil {
		log.Printf("Error setting data in db: %v", errSet)
		return custom_errors.GetDatabaseError()
	}

	_, errUpdate := client.Collection(hashedCollection).Doc(hashedDoc).Update(ctx, []firestore.Update{
		{
			Path:  "time",
			Value: firestore.ServerTimestamp,
		},
	})

	if errUpdate != nil {
		log.Printf("Error setting data in db: %v", errUpdate)
		return custom_errors.GetDatabaseError()
	}

	return nil
}

// DeleteDocument deletes a document from a collection.
func DeleteDocument(collection string, document string) error {
	client, errOpeningClient := app.Firestore(ctx)

	if errOpeningClient != nil {
		log.Println(errOpeningClient)
		return custom_errors.GetDatabaseError()
	}

	defer func() {
		errClosingClient := client.Close()

		if errClosingClient != nil {
			log.Printf("Error closing db: %v", errClosingClient)
		}
	}()

	hashedCollection, _ := hashing.HashString(collection)

	_, err := client.Collection(hashedCollection).Doc(document).Delete(ctx)
	if err != nil {
		log.Printf("Error deleteing document: %v", err)
		return err
	}

	return nil
}

// IncrementCounter counts up the counter in one document.
func IncrementCounter(collection string, document string) {
	client, errClient := app.Firestore(ctx)

	if errClient != nil {
		log.Printf("Error creating db client: %v", errClient)
		return
	}

	defer func() {
		errClosingClient := client.Close()

		if errClosingClient != nil {
			log.Printf("Error closing db: %v", errClosingClient)
		}
	}()

	hashedCollection, errHashCol := hashing.HashString(collection)

	if errHashCol != nil {
		return
	}

	hashedDoc, errHashDoc := hashing.HashString(document)

	if errHashDoc != nil {
		return
	}

	res := client.Collection(hashedCollection).Doc(hashedDoc)

	_, errGetDoc := res.Get(ctx)

	// If it is not possible to get the document, a new one is created.
	if errGetDoc != nil {
		err := WriteDocument(collection, document, structs.CountryCounter{Count: 1})

		if err != nil {
			return
		}
	} else {
		// If it exists the counter is updated.
		_, err := res.Update(ctx, []firestore.Update{
			{
				Path:  "Count",
				Value: firestore.Increment(1),
			},
		})
		if err != nil {
			return
		}

		if err != nil {
			log.Printf("Error extracting data into struct: %v", err)
			return
		}
	}
}
