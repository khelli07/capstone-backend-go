package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Level string

const (
	Country Level = "country"
	State   Level = "state"
	City    Level = "city"
)

type Location struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Level Level              `bson:"level" json:"level"`
	Timestamps
}
