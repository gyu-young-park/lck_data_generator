package firebaseapi

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	ctx *context.Context
	db *fireStoreClient
}

func NewFireBaseAPI(secretFirebaseServiceAccountKeyPath string) *FirebaseApp {
	firebaseApp := &FirebaseApp{}
	ctx := context.Background()
	conf := &firebase.Config{}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(secretFirebaseServiceAccountKeyPath)

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
			log.Fatalln("Error initializing app:", err)
	}
	firebaseApp.ctx = &ctx
	firebaseApp.db = newFireStoreClient(ctx, app)
	return firebaseApp
}

func (f *FirebaseApp) Close() {
	f.db.fireStoreClose()
}

func (f *FirebaseApp) StoreDump() {
	_, _, err := f.db.client.Collection("users").Add(*f.ctx, map[string]interface{}{
        "first": "Ada",
        "last":  "Lovelace",
        "born":  1815,
	})
	if err != nil {
			log.Fatalf("Failed adding alovelace: %v", err)
	}
}

type FireStoreDataSchema map[string]interface{}

func (f *FirebaseApp) ReadDump() []FireStoreDataSchema {
	var fireStoreDataSchemaList []FireStoreDataSchema
	iter := f.db.client.Collection("users").Documents(*f.ctx)
	for {
			doc, err := iter.Next()
			if err == iterator.Done {
					break
			}
			if err != nil {
					log.Fatalf("Failed to iterate: %v", err)
			}
			fireStoreDataSchemaList = append(fireStoreDataSchemaList, doc.Data())
	}
	return fireStoreDataSchemaList
}