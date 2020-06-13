package firebaseapp

import (
	"context"

	"cloud.google.com/go/firestore"
)

// DeleteDoc deletes a doc from a collection.
func DeleteDoc(ctx context.Context, client *firestore.Client, collection, docID string) error {
	_, err := client.Collection(collection).Doc(docID).Delete(ctx)

	return err
}
