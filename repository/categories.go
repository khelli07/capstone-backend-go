package repository

import (
	"backend-go/models"
	"backend-go/mongodb"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCategory(category *models.Category) (*mongo.InsertOneResult, error) {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	result, err := mongodb.CategoryCol.InsertOne(mongodb.Context, category)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create MongoDB entity")
	}

	return result, nil
}

func GetCategoryById(id string) (models.Category, error) {
	var category models.Category
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return category, errors.Wrap(err, "Failed to convert ID to ObjectID")
	}

	filter := bson.M{"_id": objectID}
	if err := mongodb.CategoryCol.FindOne(mongodb.Context, filter).Decode(&category); err != nil {
		return category, errors.Wrap(err, "Failed to decode MongoDB entity")
	}

	return category, nil
}

func GetCategoryByName(name string) (models.Category, error) {
	var category models.Category
	filter := bson.M{"name": name}
	if err := mongodb.CategoryCol.FindOne(mongodb.Context, filter).Decode(&category); err != nil {
		return category, errors.Wrap(err, "Failed to decode MongoDB entity")
	}

	return category, nil
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	cursor, err := mongodb.CategoryCol.Find(mongodb.Context, bson.M{})
	if err != nil {
		return categories, errors.Wrap(err, "Failed to get MongoDB entities")
	}

	if err = cursor.All(mongodb.Context, &categories); err != nil {
		return categories, errors.Wrap(err, "Failed to decode MongoDB entities")
	}

	return categories, nil
}

func UpdateCategory(id string, category *models.Category) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to convert ID to ObjectID")
	}

	category.UpdatedAt = time.Now()
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": category}

	result, err := mongodb.CategoryCol.UpdateOne(mongodb.Context, filter, update)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update MongoDB entity")
	}

	return result, nil

}

func DeleteCategory(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "Failed to convert ID to ObjectID")
	}

	filter := bson.M{"_id": objectID}
	_, err = mongodb.CategoryCol.DeleteOne(mongodb.Context, filter)
	if err != nil {
		return errors.Wrap(err, "Failed to delete MongoDB entity")
	}

	return nil
}

func CategoryIdsToNames(ids []string) ([]string, error) {
	names := []string{}
	for _, id := range ids {
		category, err := GetCategoryById(id)
		if err != nil {
			return []string{}, err
		}
		names = append(names, category.Name)
	}

	return names, nil
}

func CategoryNamesToIds(names []string) ([]string, error) {
	ids := []string{}
	for _, name := range names {
		category, err := GetCategoryByName(name)
		if err != nil {
			return []string{}, err
		}
		ids = append(ids, category.ID.Hex())
	}

	return ids, nil
}
