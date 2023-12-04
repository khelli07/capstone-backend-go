package models

type Category struct {
	Timestamps
	Name string `firestore:"name"`
}
