package function

import (
	"context"
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
		respondWithError(w, http.StatusBadRequest, serviceConst.DB_CONNECT_FAILED_MSG_DEF)
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
			respondWithError(w, http.StatusBadRequest, "Invalid record")
			return
		}

		respondWithJSON(w, http.StatusOK, users)
	}
}

// Find user by user email address
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	usrEmail := vars["usrEmail"]

	log.Println("email address from parameter : " + usrEmail)

	if usrEmail == "" {
		log.Println("Invalid email adddress")
		respondWithError(w, http.StatusBadRequest, "Invalid email address")
		return
	}

	var user dto.User

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		log.Println(err)
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// Find user based on parameter
		filter := bson.M{"usr_email": usrEmail}
		err := userCollection.FindOne(context.TODO(), filter).Decode(&user)

		if err != nil {
			log.Println(err)
			respondWithError(w, http.StatusBadRequest, "record not found")
			return
		}
	}

	respondWithJSON(w, http.StatusOK, user)
}

// TODO: update user

// TODO: delete user

// TODO: login

// TODO: update password
