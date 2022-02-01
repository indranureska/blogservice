package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	userService "github.com/indranureska/service/rest/function"
	serviceConst "github.com/indranureska/service/rest/common"
)

func main() {
	r := mux.NewRouter()

	//r.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/", helloWorldHandler)
	r.HandleFunc(serviceConst.LIST_OF_USER_SERVICE_PATH, userService.UserList)
	//r.Use(mux.CORSMethodMiddleware(r))

	http.ListenAndServe(":8000", r)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("foo"))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
