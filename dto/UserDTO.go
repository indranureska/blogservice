package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserEmail          string             `json:"usr_email" bson:"usr_email,omitempty"`
	Password           string             `json:"password" bson:"password"`
	LastLogin          string             `json:"last_login" bson:"last_login"`
	FirstName          string             `json:"first_name" bson:"first_name"`
	LastName           string             `json:"last_name" bson:"last_name"`
	IsLoggedIn         bool               `json:"is_logged_in" bson:"is_logged_in"`
	NewPassword        string             `json:"new_password" bson:"new_password"`
	ConfirmNewPassword string             `json:"confirm_new_password" bson:"confirm_new_password"`
}
