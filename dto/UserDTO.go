package dto

type User struct {
	ID        interface{} `json:"id" bson:"_id,omitempty"`
	UserEmail string      `bson:"usr_email"`
	Password  string      `bson:"password"`
	LastLogin string      `bson:"last_login"`
	FirstName string      `bson:"firstname"`
	LastName  string      `bson:"lastname"`
}
