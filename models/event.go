package models

import (
	"time"
)

func (u *Event) BeforeInsert() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}

func (u *Event) BeforeUpdate() {
	u.UpdatedAt = time.Now()
}

type Event struct {
	Timestamps
	Name        string `firestore:"name"`
	Category    string `firestore:"category"`
	Description string `firestore:"description"`
}
