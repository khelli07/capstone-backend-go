package models

import (
	"time"
)

func (u *User) BeforeInsert() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}

func (u *User) BeforeUpdate() {
	u.UpdatedAt = time.Now()
}

type User struct {
	Timestamps
	Username string `datastore:"username" json:"username"`
	Email    string `datastore:"email" json:"email"`
	Password string `datastore:"password" json:"password"`
}

type TokenUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
