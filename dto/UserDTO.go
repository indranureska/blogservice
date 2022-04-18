package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserEmail string             `bson:"usr_email,omitempty"`
	Password  string             `bson:"password"`
	LastLogin string             `bson:"last_login"`
	FirstName string             `bson:"firstname"`
	LastName  string             `bson:"lastname"`
}
