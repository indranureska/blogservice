package dto

import (
	"time"
)

type Response struct {
	Users []User `json:"users"`
}

type User struct {
	UserEmail string    `bson:"usr_email"`
	Password  string    `bson:"password"`
	LastLogin time.Time `bson:"last_login"`
	FirstName string    `bson:"firstname"`
	LastName  string    `bson:"lastname"`
}
