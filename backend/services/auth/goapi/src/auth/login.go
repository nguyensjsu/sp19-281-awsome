package main

import (
	"./models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Login API is alive!")
}

func SignUp(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var user models.User
	// Get json body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if !addUser(user) {
		respondWithError(w, http.StatusBadRequest, "user exists")
	} else {
		respondWithJson(w, http.StatusOK, "user added!")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User
	// Get json body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if validateUser(user) {
		// generate session token
		sessionToken := addSession(user.Email)

		// set cookie with 1 hour time-out
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: time.Now().Add(60 * 60 * time.Second),
		})

	} else {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var token models.Token
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
	_, _ = w.Write(response)
}



func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", PingEndPoint).Methods("GET")

	r.HandleFunc("/auth/signup", SignUp).Methods("POST")

	r.HandleFunc("/auth/login", Login).Methods("POST")

	if err := http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}