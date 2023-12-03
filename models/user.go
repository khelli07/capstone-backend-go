package models

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
