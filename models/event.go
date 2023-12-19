package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StartTime   time.Time          `bson:"start_time" json:"start_time"`
	EndTime     time.Time          `bson:"end_time" json:"end_time"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Price       float32            `bson:"price" json:"price"`
	Capacity    int32              `bson:"capacity" json:"capacity"`
	ImageURL    string             `bson:"image_url" json:"image_url"`
	IsOnline    bool               `bson:"is_online" json:"is_online"`
	// non-mandatory fields
	Lat        float64 `bson:"lat" json:"lat"`
	Long       float64 `bson:"long" json:"long"`
	Organizer  string  `bson:"organizer" json:"organizer"`
	DressCode  string  `bson:"dress_code" json:"dress_code"`
	AgeLimit   int     `bson:"age_limit" json:"age_limit"`
	TotalLikes int     `bson:"total_likes" json:"total_likes"`
	// relational fields
	Category     string   `bson:"category" json:"category"`
	Participants []string `bson:"participants" json:"participants"`
	Timestamps
}
