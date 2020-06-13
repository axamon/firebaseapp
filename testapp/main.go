package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/axamon/firebaseapp"
)

var client *firestore.Client

var auth = flag.String("auth", "auth.json", "File autenticazione  di google")
var coll = flag.String("coll", "progetti", "Collection to use")
var docID = flag.String("doc", "test", "Document ID")

func main() {

	flag.Parse()

	// Use a service account
	ctx := context.Background()

	client, err := firebaseapp.NewClient(ctx, *auth)
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

	var m = map[string]interface{}{
		"Nome":    "Al",
		"Cognome": "Capone",
	}

	refID, err := firebaseapp.Add(ctx, client, "pippo", "", m)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(refID)

	documents, err := firebaseapp.GetCollection(ctx, client, *coll)
	if err != nil {
		log.Println("Getting collection error: ", err)
	}

	for _, document := range documents {
		fmt.Println(document.Data(), document.Ref.ID)
	}

	err = firebaseapp.DeleteDoc(ctx, client, *coll, *docID)
	if err != nil {
		log.Println(err)
	}

}
