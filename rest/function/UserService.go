package function

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	dto "github.com/indranureska/service/dto"
	common "github.com/indranureska/service/rest/common"
	utils "github.com/indranureska/service/rest/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get list of user
func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []dto.User

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.DB_CONNECT_FAILED_MSG_KEY))
		return
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// Prepare context
		ctx, cancel := context.WithTimeout(context.Background(), common.DB_OPERATION_TIMEOUT_SECONDS*time.Second)
		defer cancel()

		// bson.M{},  we passed empty filter. So we want to get all data.
		cur, err := userCollection.Find(ctx, bson.M{})

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
			RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.INVALID_RECORD_MSG_KEY))
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

	if len(usrEmail) > 0 {
		user, err := getUserDataFromDbByEmailAddr(usrEmail)

		if err != nil {
			log.Println(err)
			RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.RECORD_NOT_FOUND_MSG_KEY))
			return
		} else {
			RespondWithJSON(w, http.StatusOK, user)
		}
	} else {
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.INVALID_REQUEST_PAYLOAD_MSG_KEY))
		return
	}
}

func getUserDataFromDbByEmailAddr(usrEmail string) (dto.User, error) {
	var user dto.User
	var processErr error

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		processErr = err
		return user, processErr
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// Prepare context
		ctx, cancel := context.WithTimeout(context.Background(), common.DB_OPERATION_TIMEOUT_SECONDS*time.Second)
		defer cancel()

		// Find user based on parameter
		filter := bson.M{"usr_email": usrEmail}
		err := userCollection.FindOne(ctx, filter).Decode(&user)

		if err != nil {
			processErr = err
		}

		return user, processErr
	}
}

// Create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.INVALID_REQUEST_PAYLOAD_MSG_KEY))
		return
	}
	defer r.Body.Close()

	log.Println("Create User - Request data: ")
	log.Println("user - first name : " + user.FirstName)
	log.Println("user - last name : " + user.LastName)
	log.Println("user - last login : " + user.LastLogin)
	log.Println("user - password : " + user.Password)
	log.Println("user - user email : " + user.UserEmail)

	// Hash password
	if len(user.Password) > 0 {
		hash, passwordHashError := utils.HashPassword(user.Password)

		if passwordHashError != nil {
			log.Println(err)
			RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_CREATION_FAILED_MSG_KEY))
		}

		user.Password = hash
	}

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.DB_CONNECT_FAILED_MSG_KEY))
		return
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// Prepare context
		ctx, cancel := context.WithTimeout(context.Background(), common.DB_OPERATION_TIMEOUT_SECONDS*time.Second)
		defer cancel()

		// Check if user email address exist
		filter := bson.M{"usr_email": user.UserEmail}
		count, err := userCollection.CountDocuments(ctx, filter)

		if err != nil {
			log.Println(err)
			RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_CREATION_FAILED_MSG_KEY))
			return
		}

		if count > 0 {
			RespondWithError(w, http.StatusNotAcceptable, ConstructServiceMessage(common.USER_EMAIL_ADDR_EXIST_MSG_KEY))
			return
		}

		// Insert new user
		result, err := userCollection.InsertOne(context.TODO(), user)

		if err != nil {
			log.Println(err)
			RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_CREATION_FAILED_MSG_KEY))
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
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.INVALID_REQUEST_PAYLOAD_MSG_KEY))
		return
	}
	defer r.Body.Close()

	objectID, _ := primitive.ObjectIDFromHex(user.ID.Hex())

	log.Println("Update user - Request data: ")
	log.Println("user - ID : " + objectID.String())
	log.Println("user - first name : " + user.FirstName)
	log.Println("user - last name : " + user.LastName)
	log.Println("user - last login : " + user.LastLogin)
	log.Println("user - password : " + user.Password)
	log.Println("user - user email : " + user.UserEmail)

	result, err := updateUserById(objectID, user)
	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_UPDATE_FAILED_MSG_KEY))
		return
	} else {
		RespondWithJSON(w, http.StatusCreated, result)
	}
}

func updateUserById(objectID primitive.ObjectID, user dto.User) (dto.User, error) {
	var processErr error

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		processErr = err
		return user, processErr
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		filter := bson.M{"_id": objectID}
		update := bson.M{"$set": &user}

		// Prepare context
		ctx, cancel := context.WithTimeout(context.Background(), common.DB_OPERATION_TIMEOUT_SECONDS*time.Second)
		defer cancel()

		// Check if user exist
		count, err := userCollection.CountDocuments(ctx, filter)

		if err != nil {
			processErr = err
		} else {
			// Record found, do update
			if count == 1 {
				// Update
				_, err := userCollection.UpdateOne(ctx, filter, update)

				if err != nil {
					processErr = err
				}
			} else {
				processErr = errors.New(ConstructServiceMessage(common.USER_NOT_FOUND_MSG_KEY))
			}

		}

		return user, processErr
	}
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete user function")
	vars := mux.Vars(r)
	objectID, _ := primitive.ObjectIDFromHex(vars["id"])

	log.Println("Delete user with ID : " + objectID.Hex())

	// Get MongoDB connection
	client, err := GetMongoDbClient()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.DB_CONNECT_FAILED_MSG_KEY))
		return
	} else {
		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// Prepare context
		ctx, cancel := context.WithTimeout(context.Background(), common.DB_OPERATION_TIMEOUT_SECONDS*time.Second)
		defer cancel()

		filter := bson.M{"_id": objectID}
		result, err := userCollection.DeleteOne(ctx, filter)

		if err != nil {
			log.Println(err)
			RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_DELETE_FAILED_MSG_KEY))
			return
		} else {
			RespondWithJSON(w, http.StatusCreated, result)
		}
	}
}

// login
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("User login request")
	w.Header().Set("Content-Type", "application/json")

	var user dto.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	defer r.Body.Close()

	if err != nil {
		log.Println("User empty")
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_INFO_EMPTY_MSG_KEY))
		return
	}

	// Check email address field
	if len(user.UserEmail) == 0 {
		log.Println("email address is blank")
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_EMAIL_BLANK_MSG_KEY))
		return
	}

	// Check password field
	if len(user.Password) == 0 {
		log.Println("user password is blank")
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_PASSWORD_BLANK_MSG_KEY))
		return
	}

	// Get user data
	userFromDb, err := getUserDataFromDbByEmailAddr(user.UserEmail)

	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_NOT_FOUND_MSG_KEY))
		return
	} else {
		log.Println("User found, validate password")
		isPasswordValid := utils.CheckPasswordHash(userFromDb.Password, user.Password)

		if isPasswordValid {
			log.Println("Password valid, update to logged in")
			userFromDb.IsLoggedIn = true
			objectID, _ := primitive.ObjectIDFromHex(userFromDb.ID.Hex())
			result, err := updateUserById(objectID, userFromDb)
			if err != nil {
				log.Print(err)
				RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_LOGIN_FAILED_MSG_KEY))
				return
			} else {
				RespondWithJSON(w, http.StatusOK, result)
			}
		} else {
			log.Println("Password invalid")
			RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_LOGIN_FAILED_MSG_KEY))
			return
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("User logout request")
	w.Header().Set("Content-Type", "application/json")

	var user dto.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	defer r.Body.Close()

	if err != nil {
		log.Println("User empty")
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_INFO_EMPTY_MSG_KEY))
		return
	}

	// Get user data
	userFromDb, err := getUserDataFromDbByEmailAddr(user.UserEmail)

	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_NOT_FOUND_MSG_KEY))
		return
	} else {
		log.Println("User found, update to logged out")
		userFromDb.IsLoggedIn = false
		objectID, _ := primitive.ObjectIDFromHex(userFromDb.ID.Hex())
		result, err := updateUserById(objectID, userFromDb)
		if err != nil {
			log.Print(err)
			RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_LOGOUT_FAILED_MSG_KEY))
			return
		} else {
			RespondWithJSON(w, http.StatusOK, result)
		}
	}

}

// Update password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	log.Println("User login request")
	w.Header().Set("Content-Type", "application/json")

	var user dto.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	defer r.Body.Close()

	if err != nil {
		log.Println("User empty")
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_INFO_EMPTY_MSG_KEY))
		return
	}

	// Check email address field
	if len(user.UserEmail) == 0 {
		log.Println("email address is blank")
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_EMAIL_BLANK_MSG_KEY))
		return
	}

	// Check password field
	if len(user.Password) == 0 {
		log.Println("user password is blank")
		RespondWithError(w, http.StatusBadRequest, ConstructServiceMessage(common.USER_PASSWORD_BLANK_MSG_KEY))
		return
	}

	
}

// TODO forget password
