package models

type EventMandatory struct {
	TimeRange
	Name        string   `firestore:"name" json:"name"`
	Categories  []string `firestore:"categories" json:"categories"`
	Description string   `firestore:"description" json:"description"`
	Location    string   `firestore:"location" json:"location"`
	Price       float32  `firestore:"price" json:"price"`
	Capacity    int32    `firestore:"capacity" json:"capacity"`
}

type Event struct {
	Timestamps
	EventMandatory
	Organizer  string `firestore:"organizer" json:"organizer"`
	DressCode  string `firestore:"dress_code" json:"dress_code"`
	AgeLimit   int    `firestore:"age_limit" json:"age_limit"`
	TotalLikes int    `firestore:"total_likes" json:"total_likes"`
}
