package repository

import (
	"backend-go/models"
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

func CreateEvent(ctx context.Context, client *datastore.Client, event *models.Event) (*datastore.Key, error) {
	incompleteKey := datastore.IncompleteKey("Event", nil)
	key, err := client.Put(ctx, incompleteKey, event)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create datastore entity")
	}
	return key, nil
}

func GetEventById(ctx context.Context, client *datastore.Client, id int64) (models.Event, error) {
	var task models.Event
	taskKey := datastore.IDKey("Event", id, nil)
	err := client.Get(ctx, taskKey, &task)
	if err != nil {
		return models.Event{}, errors.Wrap(err, "Failed to get datastore entity")
	}

	return task, nil
}

func DeleteEvent(ctx context.Context, client *datastore.Client, id int64) error {
	key := datastore.IDKey("Event", id, nil)
	err := client.Delete(ctx, key)
	if err != nil {
		return errors.Wrap(err, "Failed to delete datastore entity")
	}

	return nil
}
