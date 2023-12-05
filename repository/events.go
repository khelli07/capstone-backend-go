package repository

import (
	"backend-go/models"
	"backend-go/mongodb"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateEvent(event *models.Event) (*mongo.InsertOneResult, error) {
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	result, err := mongodb.EventCol.InsertOne(mongodb.Context, event)
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
	err = mongodb.EventCol.FindOne(mongodb.Context, filter).Decode(&event)
	if err != nil {
		return event, errors.Wrap(err, "Failed to get MongoDB entity")
	}

	return event, nil
}

func GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	cursor, err := mongodb.EventCol.Find(mongodb.Context, bson.M{})
	if err != nil {
		return events, errors.Wrap(err, "Failed to get MongoDB entities")
	}

	if err = cursor.All(mongodb.Context, &events); err != nil {
		return events, errors.Wrap(err, "Failed to decode MongoDB entities")
	}

	return events, nil
}

func GetPopularEvents(topK int) ([]models.Event, error) {
	var events []models.Event

	filter := bson.M{"end_time": bson.M{"$gt": time.Now()}}
	options := options.Find().SetSort(bson.M{"total_likes": -1}).SetLimit(int64(topK))

	cursor, err := mongodb.EventCol.Find(mongodb.Context, filter, options)
	if err != nil {
		return events, errors.Wrap(err, "Failed to get MongoDB entities")
	}

	if err = cursor.All(mongodb.Context, &events); err != nil {
		return events, errors.Wrap(err, "Failed to decode MongoDB entities")
	}

	return events, nil
}

func UpdateEvent(id string, event *models.Event) (*mongo.UpdateResult, error) {
	event.UpdatedAt = time.Now()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to convert ID to ObjectID")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": event}
	result, err := mongodb.EventCol.UpdateOne(mongodb.Context, filter, update)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update MongoDB entity")
	}

	return result, nil
}

func DeleteEvent(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "Failed to convert ID to ObjectID")
	}

	filter := bson.M{"_id": objectID}
	_, err = mongodb.EventCol.DeleteOne(mongodb.Context, filter)
	if err != nil {
		return errors.Wrap(err, "Failed to delete MongoDB entity")
	}
	return nil
}
