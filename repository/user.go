package repository

import (
	"backend-go/fs"
	"backend-go/models"
	"context"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

func CreateUser(ctx context.Context, client *firestore.Client, user *models.User) (*firestore.DocumentRef, error) {
	docRef, _, err := fs.UserCol.Add(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create firestore entity")
	}

	return docRef, nil
}

func GetUserByEmail(ctx context.Context, client *firestore.Client, email string) (*firestore.DocumentRef, error) {
	query := fs.UserCol.Where("email", "==", email).OrderBy("created_at", firestore.Desc).Limit(1)

	iter := query.Documents(ctx)
	doc, err := iter.Next()

	if doc == nil && err == iterator.Done {
		return nil, errors.New("User not found")
	} else if err != nil {
		return nil, errors.Wrap(err, "Error querying firestore")
	}

	return doc.Ref, nil
}

func GetUserById(ctx context.Context, client *firestore.Client, id string) (models.User, error) {
	var user models.User
	doc := fs.UserCol.Doc(id)
	snapshot, err := doc.Get(ctx)
	if err != nil {
		return user, errors.Wrap(err, "Failed to get firestore entity")
	}
	if err := snapshot.DataTo(&user); err != nil {
		return user, errors.Wrap(err, "Failed to convert firestore entity")
	}

	return user, nil
}
