package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	common "github.com/indranureska/service/rest/common"
	function "github.com/indranureska/service/rest/function"
)

func main() {
	r := mux.NewRouter()

	//r.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/", helloWorldHandler)
	r.HandleFunc(common.LIST_OF_USER_SERVICE_PATH, function.ListUser).Methods("GET")
	r.HandleFunc(common.FIND_USER_SERVICE_PATH, function.FindUser).Methods("GET")
	r.HandleFunc(common.CREATE_USER_SERVICE_PATH, function.CreateUser).Methods("POST")
	r.HandleFunc(common.UPDATE_USER_SERVICE_PATH, function.UpdateUser).Methods("PUT")
	r.HandleFunc(common.DELETE_USER_SERVICE_PATH, function.DeleteUser).Methods("DELETE")

	//r.Use(mux.CORSMethodMiddleware(r))

	// TODO: Initiate response message map
	function.InitServiceMessages()

	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	
}

// func fooHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	if r.Method == http.MethodOptions {
// 		return
// 	}

// 	w.Write([]byte("foo"))
// }

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
	function.RespondWithJSON(w, http.StatusOK, response)
}
