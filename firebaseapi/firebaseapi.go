package firebaseapi

import (
	"context"
	"fmt"
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

type FireStoreDataSchema map[string]interface{}

func (f *FirebaseApp) RemoveCollection(collection string) error {
	collectionRef := f.db.client.Collection(collection)
	for {
		iter := collectionRef.Limit(100).Documents(*f.ctx)
		numDeleted := 0
		batch := f.db.client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			batch.Delete(doc.Ref)
			numDeleted++
		}
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(*f.ctx)
		if err != nil {
			return err
		}
	}
}

func (f *FirebaseApp) StoreDataWithDoc(collection string ,doc string,data FireStoreDataSchema) error{
	_, err := f.db.client.Collection(collection).Doc(doc).Set(*f.ctx, data)
	if err != nil {
		return fmt.Errorf("[%s]Error can't store data, collection:[%s], data[%v]", "StoreData", collection, data)
	}
	return nil
}

func (f *FirebaseApp) ReadData(collection string) []FireStoreDataSchema {
	var ret []FireStoreDataSchema
	iter := f.db.client.Collection(collection).Documents(*f.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Failed to iterate %v\n", err)
			break
		}
		ret = append(ret, doc.Data())
	}
	return ret

}

func (f *FirebaseApp) StoreDump() {
	_, _, err := f.db.client.Collection("users").Add(*f.ctx, FireStoreDataSchema{
        "first": "Ada",
        "last":  "Lovelace",
        "born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

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