package ds

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

var DS *datastore.Client
var CTX context.Context

func InitClient() {
	var err error
	CTX = context.Background()
	DS, err = createClient(CTX)

	if err != nil {
		log.Fatalf("Error creating datastore client: %v", err)
	}

	log.Println("Datastore client initialized successfully.")
}

func createClient(ctx context.Context) (*datastore.Client, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new datastore client")
	}

	return client, nil
}
