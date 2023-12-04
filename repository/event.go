package repository

import (
	"backend-go/fs"
	"backend-go/models"
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
)

func CreateEvent(ctx context.Context, client *firestore.Client, event *models.Event) (*firestore.DocumentRef, error) {
	result, _, err := fs.EventCol.Add(ctx, event)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create firestore entity")
	}
	return result, nil
}

func GetEventById(ctx context.Context, client *firestore.Client, id string) (models.Event, error) {
	var event models.Event
	doc := fs.EventCol.Doc(id)
	snapshot, err := doc.Get(ctx)
	if err != nil {
		return event, errors.Wrap(err, "Failed to get firestore entity")
	}
	if err := snapshot.DataTo(&event); err != nil {
		return event, errors.Wrap(err, "Failed to convert firestore entity")
	}

	return event, nil
}

func GetAllEvents(ctx context.Context, client *firestore.Client) ([]models.Event, error) {
	var events []models.Event
	iter := fs.EventCol.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var event models.Event
		if err := doc.DataTo(&event); err != nil {
			return events, errors.Wrap(err, "Failed to convert firestore entity")
		}
		events = append(events, event)
	}
	return events, nil
}

func GetPopularEvents(ctx context.Context, client *firestore.Client) ([]models.Event, error) {
	var events []models.Event
	query := fs.EventCol.
		Where("end_date", ">", time.Now()).
		OrderBy("likes", firestore.Desc).
		OrderBy("created_at", firestore.Desc).
		Limit(10)

	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var event models.Event
		if err := doc.DataTo(&event); err != nil {
			return events, errors.Wrap(err, "Failed to convert firestore entity")
		}
		events = append(events, event)
	}
	return events, nil
}

func DeleteEvent(ctx context.Context, client *firestore.Client, id string) error {
	doc := fs.EventCol.Doc(id)
	_, err := doc.Delete(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to delete firestore entity")
	}
	return nil
}
