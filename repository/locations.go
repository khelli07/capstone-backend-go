package repository

import (
	"backend-go/models"
	"backend-go/mongodb"
	"time"

	"golang.org/x/exp/slices"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateLocation(location *models.Location) (*mongo.InsertOneResult, error) {
	location.CreatedAt = time.Now()
	location.UpdatedAt = time.Now()

	result, err := mongodb.LocationCol.InsertOne(mongodb.Context, location)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create MongoDB entity")
	}

	return result, nil
}

func GetAllLocations() ([]models.Location, error) {
	var locations []models.Location
	cursor, err := mongodb.LocationCol.Find(mongodb.Context, bson.M{})
	if err != nil {
		return locations, errors.Wrap(err, "Failed to get MongoDB entities")
	}

	if err = cursor.All(mongodb.Context, &locations); err != nil {
		return locations, errors.Wrap(err, "Failed to decode MongoDB entities")
	}

	return locations, nil
}

func QueryLocation(query bson.M) ([]models.Location, error) {
	var locations []models.Location
	cursor, err := mongodb.LocationCol.Find(mongodb.Context, query)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get MongoDB entities")
	}

	if err = cursor.All(mongodb.Context, &locations); err != nil {
		return nil, errors.Wrap(err, "Failed to decode MongoDB entities")
	}

	return locations, nil
}

func UpdateLocation(id string, location *models.Location) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to convert ID to ObjectID")
	}

	location.UpdatedAt = time.Now()
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": location}

	result, err := mongodb.LocationCol.UpdateOne(mongodb.Context, filter, update)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update MongoDB entity")
	}

	return result, nil
}

func DeleteLocation(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "Failed to convert ID to ObjectID")
	}

	filter := bson.M{"_id": objectID}
	_, err = mongodb.LocationCol.DeleteOne(mongodb.Context, filter)
	if err != nil {
		return errors.Wrap(err, "Failed to delete MongoDB entity")
	}

	return nil
}

func ValidateLevel(level string) *models.Level {
	levels := []string{"country", "state", "city"}
	if slices.Contains(levels, level) == false {
		return nil
	}

	ret := models.Country
	switch level {
	case "country":
		ret = models.Country
	case "state":
		ret = models.State
	case "city":
		ret = models.City
	}

	return &ret
}
