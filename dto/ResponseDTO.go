package dto

type Response struct {
	InsertedID string `json:"InsertedID" bson:"InsertedID"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
