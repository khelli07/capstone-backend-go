package mongodb

import "go.mongodb.org/mongo-driver/mongo"

var EventCol *mongo.Collection
var UserCol *mongo.Collection

func InitCollections() {
	EventCol = DB.Collection("events")
	UserCol = DB.Collection("users")
}
