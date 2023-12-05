package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Timestamps
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `firestore:"username" json:"username"`
	Email    string             `firestore:"email" json:"email"`
	Password string             `firestore:"password" json:"password"`
}

type TokenUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
