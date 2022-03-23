package database

import (
	"assignment-2/custom_errors"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
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
func WriteToDatabase(collection string, document string, data map[string]interface{}) error {
	client, errOpeningClient := app.Firestore(ctx)

	if errOpeningClient != nil {
		log.Println(errOpeningClient)
		return custom_errors.GetDatabaseError()
	}

	//TODO error handling
	_, _ = client.Collection(collection).Doc(document).Set(ctx, data)

	errClosingClient := client.Close()

	if errClosingClient != nil {
		log.Println(errClosingClient)
		return custom_errors.GetDatabaseError()
	}

	return nil
}
