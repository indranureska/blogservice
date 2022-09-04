package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	common "github.com/indranureska/service/rest/common"
	function "github.com/indranureska/service/rest/function"
)

func main() {
	log.Print("Starting Web Service...")
	log.Println("Set router")
	r := mux.NewRouter()

	r.HandleFunc(common.LIST_OF_USER_SERVICE_PATH, function.ListUser).Methods("GET")
	r.HandleFunc(common.FIND_USER_SERVICE_PATH, function.FindUser).Methods("GET")
	r.HandleFunc(common.CREATE_USER_SERVICE_PATH, function.CreateUser).Methods("POST")
	r.HandleFunc(common.UPDATE_USER_SERVICE_PATH, function.UpdateUser).Methods("PUT")
	r.HandleFunc(common.DELETE_USER_SERVICE_PATH, function.DeleteUser).Methods("DELETE")
	r.HandleFunc(common.USER_LOGIN_SERVICE_PATH, function.Login).Methods("POST")
	r.HandleFunc(common.USER_LOGOUT_SERVICE_PATH, function.Logout).Methods("POST")

	//r.Use(mux.CORSMethodMiddleware(r))

	log.Print("Initiate service message")
	function.InitServiceMessages()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Listen and serve")
	log.Fatal(srv.ListenAndServe())

}
