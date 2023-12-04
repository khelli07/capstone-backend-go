package models

import "time"

type Timestamps struct {
	CreatedAt time.Time `firestore:"created_at" json:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at" json:"updated_at"`
}

type TimeRange struct {
	StartTime time.Time `firestore:"start_time" json:"start_time"`
	EndTime   time.Time `firestore:"end_time" json:"end_time"`
}
