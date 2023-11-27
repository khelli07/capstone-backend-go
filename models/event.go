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
	Name        string `datastore:"name"`
	Category    string `datastore:"category"`
	Description string `datastore:",noindex"`
}
