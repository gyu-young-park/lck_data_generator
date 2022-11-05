package firebaseapi

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type fireStoreClient struct{
	client *firestore.Client
}

func newFireStoreClient(ctx context.Context, app *firebase.App) *fireStoreClient {
	fireStoreClient := &fireStoreClient{}
	client, err := app.Firestore(ctx)
	if err != nil {
	  log.Fatalln(err)
	}
	fireStoreClient.client = client
	return fireStoreClient
}

func (f *fireStoreClient)fireStoreClose() {
	defer f.fireStoreClose()
	fmt.Println("firestore instance closed")
}