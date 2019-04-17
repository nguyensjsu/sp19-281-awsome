package main

import (
	"encoding/json"
	"fmt"
	"log"
	http "net/http"
	. "auth/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login API is alive!")
}

// POST /auth/login : login a user
func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	
	var token Token
	
	token.Sub = "test"
	token.Role = "test"
	
	respondWithJson(w, http.StatusOK, token)
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
	r.HandleFunc("/auth/login", Login).Methods("POST")
	
	if err := http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}