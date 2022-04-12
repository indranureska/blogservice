package dto

type User struct {
	UserEmail string `bson:"usr_email"`
	Password  string `bson:"password"`
	LastLogin string `bson:"last_login"`
	FirstName string `bson:"firstname"`
	LastName  string `bson:"lastname"`
}
