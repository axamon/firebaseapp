package firebaseapp

import (
	"context"

	"cloud.google.com/go/firestore"
)

// Add adds data to firebase DB.
func Add(ctx context.Context, client *firestore.Client, collection, docID string, m map[string]interface{}) (string, error) {

	var err error
	var refID string
	var ref *firestore.DocumentRef

	switch {
	case docID == "":
		ref, _, err = client.Collection(collection).Add(ctx, m)
		refID = ref.ID

	default:
		_, err = client.Collection(collection).Doc(docID).Set(ctx, m)
		refID = docID
	}

	return refID, err
}
