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
var doc = flag.String("doc", "test", "Document ID")

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

	_, err = client.Collection(*coll).Doc(*doc).Set(ctx, m)
	if err != nil {
		log.Println("Setting Doc error: ", err)
	}

	err = GetCollection(ctx, *coll)
	if err != nil {
		log.Println("Getting collection error: ", err)
	}

}

// GetCollection retrieves a collection from firebase db.
func GetCollection(ctx context.Context, collection string) error {

	iter := client.Collection(collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(doc.Data())
	}

	return nil

}
