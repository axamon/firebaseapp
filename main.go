package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var client *firestore.Client

var auth = flag.String("auth", "auth.json", "File autenticazione  di google")
var coll = flag.String("coll", "progetti", "Collection to use")
var docID = flag.String("doc", "test", "Document ID")

func main() {

	flag.Parse()

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(*auth)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, _, err = client.Collection(*coll).Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	var m = make(map[string]interface{})

	_, err = client.Collection(*coll).Doc(*docID).Set(ctx, m)
	if err != nil {
		log.Println("Setting Doc error: ", err)
	}

	doc, err := GetCollection(ctx, *coll)
	if err != nil {
		log.Println("Getting collection error: ", err)
	}

	fmt.Println(doc.Data())

	err = DeleteDoc(ctx, *coll, *docID)
	if err != nil {
		log.Println(err)
	}

}

// GetCollection retrieves a collection from firebase db.
func GetCollection(ctx context.Context, collection string) (*firestore.DocumentSnapshot, error) {

	var doc *firestore.DocumentSnapshot
	var err error

	iter := client.Collection(collection).Documents(ctx)

	for {
		doc, err = iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return doc, nil
}

// DeleteDoc deletes a doc from a collection.
func DeleteDoc(ctx context.Context, collection, docID string) error {
	_, err := client.Collection(collection).Doc(docID).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
