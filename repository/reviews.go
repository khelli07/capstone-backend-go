package repository

import (
	"backend-go/models"
	"backend-go/mongodb"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateReview(review *models.Review) (*mongo.InsertOneResult, error) {
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	result, err := mongodb.ReviewCol.InsertOne(mongodb.Context, review)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, writeErr := range writeErr.WriteErrors {
				if writeErr.Code == 11000 {
					return nil, errors.New("User already review this event")
				}
			}
		}

		return nil, errors.Wrap(err, "Failed to create MongoDB entity")
	}

	return result, nil
}

func GetReviews(eventID string) ([]models.Review, error) {
	var reviews []models.Review

	filter := bson.M{"event_id": eventID}
	options := options.Find().SetSort(bson.M{"timestamps.created_at": -1})

	cursor, err := mongodb.ReviewCol.Find(mongodb.Context, filter, options)
	if err != nil {
		return reviews, errors.Wrap(err, "Failed to get MongoDB entities")
	}

	if err = cursor.All(mongodb.Context, &reviews); err != nil {
		return reviews, errors.Wrap(err, "Failed to decode MongoDB entities")
	}

	return reviews, nil
}
