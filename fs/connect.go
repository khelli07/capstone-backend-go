package fs

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
)

var FSClient *firestore.Client
var CTX context.Context

func InitClient() {
	var err error
	CTX = context.Background()
	FSClient, err = createClient(CTX)

	if err != nil {
		log.Fatalf("Error creating firestore client: %v", err)
	}

	log.Println("Firestore client initialized successfully.")
}

func createClient(ctx context.Context) (*firestore.Client, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new Firestore client")
	}

	return client, nil
}
