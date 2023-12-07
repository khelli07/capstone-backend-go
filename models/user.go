package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Email       string             `bson:"email" json:"email"`
	JoinedEvent []string           `bson:"joined_event" json:"joined_event"`
	Password    string             `bson:"password" json:"-"`
	Timestamps
}

type TokenUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
