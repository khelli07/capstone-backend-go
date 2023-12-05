package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EventID string             `bson:"event_id" json:"event_id"`
	Rating  int                `bson:"rating" json:"rating"`
	Comment string             `bson:"comment" json:"comment"`
	Timestamps
}
