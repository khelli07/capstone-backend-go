package mongodb

import "go.mongodb.org/mongo-driver/mongo"

var EventCol *mongo.Collection
var UserCol *mongo.Collection
var ReviewCol *mongo.Collection
var CategoryCol *mongo.Collection

func InitCollections() {
	EventCol = DB.Collection("events")
	UserCol = DB.Collection("users")
	ReviewCol = DB.Collection("reviews")
	CategoryCol = DB.Collection("categories")
}
