package dto

type Response struct {
	Users []User `json:"users"`
}

type User struct {
	UserEmail string `json:"usr_email"`
	Password  string `json:"password"`
	LastLogin string `json:"last_login"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
