package firebaseapp

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// NewClient creates a new client.
func NewClient(ctx context.Context, authenticationFile string) (*firestore.Client, error) {

	sa := option.WithCredentialsFile(authenticationFile)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, fmt.Errorf("Error in NewClient app creation: %v", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error in NewClient creation: %v", err)
	}
	return client, nil
}
