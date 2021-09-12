package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserId    primitive.ObjectID `bson:"_id`
	Firstname string             `bson:"firstname`
	Lastname string             `bson:"lastname`
	email string             `bson:"email`
}


type Uploads struct {
	
}

type Collections struct {
	
}
