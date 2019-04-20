package main

import (
	"encoding/json"
	"fmt"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	http "net/http"
	. "users/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "users/DAO"
)


//Test Status of the API
func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users API is alive!")
}

var dao = UserDatabase{}


//Get All users
func GetAllUsers(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var users []User
	if err := db.C(COLLECTION).Find(bson.M{}).All(&users); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Query")
		return
	}

	respondWithJson(w,http.StatusOK, users)
}


// POST /create/user : Create a User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	
	respondWithJson(w, http.StatusCreated, "created")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)
	fmt.Println(string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	r.HandleFunc("/create/user", CreateUser).Methods("POST")
	r.HandleFunc("/getUsers", GetAllUsers).Methods("GET")
	
	if err := http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}