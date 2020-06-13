package firebaseapp

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// GetCollection retrieves a collection from firebase db.
func GetCollection(ctx context.Context, client *firestore.Client, collection string) ([]*firestore.DocumentSnapshot, error) {

	var docs []*firestore.DocumentSnapshot
	iter := client.Collection(collection).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		docs = append(docs, doc)
		if err != nil {
			return nil, err
		}
	}

	return docs, nil
}
