package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username             string             `bson:"username" json:"username"`
	Email                string             `bson:"email" json:"email"`
	Password             string             `bson:"password" json:"-"`
	Lat                  float64            `bson:"lat,truncate" json:"lat"`
	Long                 float64            `bson:"long,truncate" json:"long"`
	JoinedEvent          []string           `bson:"joined_event" json:"joined_event"`
	EventCategories      []string           `bson:"event_categories" json:"event_categories"`
	PreferenceCategories []string           `bson:"preference_categories" json:"preference_categories"`
	Timestamps
}

type TokenUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
