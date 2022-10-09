package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Firestore struct {
	Client *firestore.Client
}

func New(credentialsFile string) (*Firestore, error) {
	opt := option.WithCredentialsFile(credentialsFile)
	app, errApp := firebase.NewApp(context.Background(), nil, opt)
	if errApp != nil {
		return nil, errApp
	}

	client, errFirestore := app.Firestore(context.Background())
	if errFirestore != nil {
		return nil, errFirestore
	}

	return &Firestore{
		Client: client,
	}, nil
}

func (f *Firestore) GetData(collection string) (map[string]map[string]interface{}, error) {
	iter := f.Client.Collection(collection).Documents(context.Background())

	result := map[string]map[string]interface{}{}
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}
		id := doc.Ref.ID
		result[id] = doc.Data()
	}
	return result, nil
}
