package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EventID     string             `bson:"event_id" json:"event_id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	Category    string             `bson:"category" json:"category"`
	Rating      int                `bson:"rating" json:"rating"`
	Comment     string             `bson:"comment" json:"comment"`
	JoinedEvent []string           `bson:"joined_event" json:"joined_event"`
	Timestamps
}
