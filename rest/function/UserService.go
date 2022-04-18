package function

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	dto "github.com/indranureska/service/dto"
	serviceConst "github.com/indranureska/service/rest/common"
	"go.mongodb.org/mongo-driver/bson"
)

// Get list of user
func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []dto.User

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, serviceConst.DB_CONNECT_FAILED_MSG_DEF)
		return
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// bson.M{},  we passed empty filter. So we want to get all data.
		cur, err := userCollection.Find(context.TODO(), bson.M{})

		if err != nil {
			log.Println(err)
		}

		// Close the cursor once finished
		/*A defer statement defers the execution of a function until the surrounding function returns.
		simply, run cur.Close() process but after cur.Next() finished.*/
		defer cur.Close(context.TODO())

		for cur.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			var user dto.User
			// & character returns the memory address of the following variable.
			err := cur.Decode(&user) // decode similar to deserialize process.
			if err != nil {
				log.Println(err)
			}

			// add item our array
			users = append(users, user)
		}

		if err := cur.Err(); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid record")
			return
		}

		RespondWithJSON(w, http.StatusOK, users)
	}
}

// Find user by user email address
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	usrEmail := vars["usrEmail"]

	var user dto.User

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, serviceConst.DB_CONNECT_FAILED_MSG_DEF)
		return
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// Find user based on parameter
		filter := bson.M{"usr_email": usrEmail}
		err := userCollection.FindOne(context.TODO(), filter).Decode(&user)

		if err != nil {
			log.Println(err)
			RespondWithError(w, http.StatusBadRequest, "record not found")
			return
		}
	}

	RespondWithJSON(w, http.StatusOK, user)
}

// Create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	log.Println("Request data: ")
	log.Println("user - ID : " + user.ID.String())
	log.Println("user - first name : " + user.FirstName)
	log.Println("user - last name : " + user.LastName)
	log.Println("user - last login : " + user.LastLogin)
	log.Println("user - password : " + user.Password)
	log.Println("user - user email : " + user.UserEmail)

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, serviceConst.DB_CONNECT_FAILED_MSG_DEF)
		return
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// Insert new user
		result, err := userCollection.InsertOne(context.TODO(), user)

		if err != nil {
			log.Println(err)
			RespondWithError(w, http.StatusBadRequest, "user creation failed")
			return
		} else {
			RespondWithJSON(w, http.StatusCreated, result)
		}
	}
}

// Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	log.Println("Request data: ")
	log.Println("user - ID : " + user.ID.String())
	log.Println("user - first name : " + user.FirstName)
	log.Println("user - last name : " + user.LastName)
	log.Println("user - last login : " + user.LastLogin)
	log.Println("user - password : " + user.Password)
	log.Println("user - user email : " + user.UserEmail)

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, serviceConst.DB_CONNECT_FAILED_MSG_DEF)
		return
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		filter := bson.M{"_id": user.ID}
		update := bson.M{"$set": &user}
		result, err := userCollection.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			log.Println(err)
			RespondWithError(w, http.StatusBadRequest, "user update failed")
			return
		} else {
			RespondWithJSON(w, http.StatusCreated, result)
		}
	}
}

// TODO: delete user

// TODO: login

// TODO: update password
