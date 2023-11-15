package database

import "backend-go/models"

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
