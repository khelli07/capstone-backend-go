package repository

import (
	"backend-go/models"
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

func CreateUser(ctx context.Context, client *datastore.Client, user *models.User) (*datastore.Key, error) {
	incompleteKey := datastore.IncompleteKey("User", nil)
	key, err := client.Put(ctx, incompleteKey, user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create datastore entity")
	}
	return key, nil
}

func GetUserByEmail(ctx context.Context, client *datastore.Client, email string) (*datastore.Key, error) {
	var users []models.User
	query := datastore.NewQuery("User").FilterField("email", "=", email).Order("-created_at")
	keys, err := client.GetAll(ctx, query, &users)

	if err != nil {
		fmt.Println("Error querying datastore:", err)
		return nil, err
	}

	if len(keys) == 0 {
		return nil, datastore.ErrNoSuchEntity
	}

	return keys[0], nil
}

func GetUserById(ctx context.Context, client *datastore.Client, id int64) (models.User, error) {
	var user models.User
	userKey := datastore.IDKey("User", id, nil)
	err := client.Get(ctx, userKey, &user)
	if err != nil {
		return models.User{}, errors.Wrap(err, "Failed to get datastore entity")
	}

	return user, nil
}
