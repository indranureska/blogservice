package service

import (
	"net/http"
	"encoding/json"
)

const LIST_OF_USER_SERVICE_PATH = "/user-list"

// Get list of user
func UserList(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
    json.Unmarshal([]byte(`{ "function": "UserList" }`), &response)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// TODO: find user by user email address

// TODO: update user

// TODO: delete user

// TODO: login

// TODO: update password
