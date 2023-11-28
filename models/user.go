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
	Username string `firestore:"username" json:"username"`
	Email    string `firestore:"email" json:"email"`
	Password string `firestore:"password" json:"password"`
}

type TokenUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
