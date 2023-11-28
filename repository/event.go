package repository

import (
	"backend-go/models"
	"context"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
)

func CreateEvent(ctx context.Context, client *firestore.Client, event *models.Event) (*firestore.DocumentRef, error) {
	col := client.Collection("Event")
	result, _, err := col.Add(ctx, event)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create firestore entity")
	}
	return result, nil
}

func GetEventById(ctx context.Context, client *firestore.Client, id string) (models.Event, error) {
	var event models.Event
	col := client.Collection("Event").Doc(id)
	snapshot, err := col.Get(ctx)
	if err != nil {
		return event, errors.Wrap(err, "Failed to get firestore entity")
	}
	if err := snapshot.DataTo(&event); err != nil {
		return event, errors.Wrap(err, "Failed to convert firestore entity")
	}

	return event, nil
}
