package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserEmail string             `json:"usr_email" bson:"usr_email,omitempty"`
	Password  string             `json:"password" bson:"password"`
	LastLogin string             `json:"last_login" bson:"last_login"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
}
