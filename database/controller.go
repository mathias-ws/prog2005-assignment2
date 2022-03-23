package database

import (
	"assignment-2/custom_errors"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"time"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client
var app *firebase.App

// InitDB initializes the database setup.
func InitDB() {
	ctx = context.Background()
	opt := option.WithCredentialsFile("database/auth.json")
	app, _ = firebase.NewApp(ctx, nil, opt)
}

// GetFromDatabase gets data from a collection and a document in the database.
func GetFromDatabase(collection string, document string) map[string]interface{} {
	//TODO error handling
	client, _ = app.Firestore(ctx)

	res := client.Collection(collection).Doc(document)
	doc, _ := res.Get(ctx)
	data := doc.Data()

	err := client.Close()

	if err != nil {
		return nil
	}

	return data
}

// WriteToDatabase creates or updates a document in a collection (and a document).
func WriteToDatabase(collection string, document string, data interface{}) error {
	client, errOpeningClient := app.Firestore(ctx)

	if errOpeningClient != nil {
		log.Println(errOpeningClient)
		return custom_errors.GetDatabaseError()
	}

	//TODO error handling
	_, errSet := client.Collection(collection).Doc(document).Set(ctx, data)

	if errSet != nil {
		log.Printf("Error setting data in db: %v", errSet)
		return custom_errors.GetDatabaseError()
	}

	_, errUpdate := client.Collection(collection).Doc(document).Update(ctx, []firestore.Update{
		{
			Path:  "timestamp",
			Value: time.Now(),
		},
	})

	if errUpdate != nil {
		log.Printf("Error setting data in db: %v", errUpdate)
		return custom_errors.GetDatabaseError()
	}

	errClosingClient := client.Close()

	if errClosingClient != nil {
		log.Println(errClosingClient)
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

	_, err := client.Collection(collection).Doc(document).Delete(ctx)
	if err != nil {
		log.Printf("Error deleteing document: %v", err)
		return err
	}

	errClosingClient := client.Close()

	if errClosingClient != nil {
		log.Println(errClosingClient)
		return custom_errors.GetDatabaseError()
	}

	return nil
}

// DeleteCollection deletes all documents in a collection.
func DeleteCollection(collection string) error {
	client, errOpeningClient := app.Firestore(ctx)

	if errOpeningClient != nil {
		log.Println(errOpeningClient)
		return custom_errors.GetDatabaseError()
	}

	// Code taken/inspired by the firestore documentation:
	// https://firebase.google.com/docs/firestore/manage-data/delete-data
	for {
		// Get a batch of documents
		iter := client.Collection(collection).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Printf("Error while interating all documents to delete: %v", err)
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			log.Printf("Error commiting batch delete: %v", err)
			return err
		}
	}
}
