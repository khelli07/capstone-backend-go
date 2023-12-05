package models

type Category struct {
	Timestamps
	Name string `bson:"name"`
}
