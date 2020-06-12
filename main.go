package main

import (
	"context"
	"flag"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var auth = flag.String("auth", "auth.json", "File autenticazione  di google")

func main() {

	flag.Parse()

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(*auth)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

}
