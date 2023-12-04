package fs

import "cloud.google.com/go/firestore"

var EventCol *firestore.CollectionRef
var UserCol *firestore.CollectionRef

func InitCollections() {
	EventCol = FSClient.Collection("Events")
	UserCol = FSClient.Collection("Users")
}
