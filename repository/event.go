package repository

import (
	"backend-go/fs"
	"backend-go/models"
	"backend-go/mongodb"
	"context"
	"math"
	"sort"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEvent(event *models.Event) (*mongo.InsertOneResult, error) {
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	result, err := mongodb.EventCol.InsertOne(mongodb.CTX, event)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create MongoDB entity")
	}

	return result, nil
}

func GetEventById(id string) (models.Event, error) {
	var event models.Event
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return event, errors.Wrap(err, "Failed to convert ID to ObjectID")
	}

	filter := bson.M{"_id": objectID}
	err = mongodb.EventCol.FindOne(mongodb.CTX, filter).Decode(&event)
	if err != nil {
		return event, errors.Wrap(err, "Failed to get MongoDB entity")
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
	query := fs.EventCol.Where("end_time", ">", time.Now())
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

	sort.Slice(events, func(i, j int) bool {
		return events[i].TotalLikes > events[j].TotalLikes
	})

	cutoff := math.Min(float64(len(events)), 10)
	return events[:int32(cutoff)], nil
}

func DeleteEvent(ctx context.Context, client *firestore.Client, id string) error {
	doc := fs.EventCol.Doc(id)
	_, err := doc.Delete(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to delete firestore entity")
	}
	return nil
}
