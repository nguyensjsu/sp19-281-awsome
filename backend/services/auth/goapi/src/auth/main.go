package main

import (
	"./db"
	"./models"
	userManager "./user"
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
	if !userManager.AddUser(user) {
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
	validUser, dbUser := userManager.ValidateUser(user)
	if validUser {
		// generate session token
		sessionToken := userManager.AddSession(dbUser)

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

func Logout(w http.ResponseWriter, r *http.Request)  {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// get token from db
	_, sessionExist := userManager.GetSession(sessionToken)
	if sessionExist {
		userManager.InvalidateSession(sessionToken)
	}
}

func WelcomeEndPoint(w http.ResponseWriter, r *http.Request)  {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// get token from db
	session, sessionExist := userManager.GetSession(sessionToken)
	if !sessionExist {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Welcome %s!", session.Email)))
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

	db.ConfigMongoDB()

	db.ConfigRedis()

	r := mux.NewRouter()

	r.HandleFunc("/ping", PingEndPoint).Methods("GET")

	r.HandleFunc("/auth/signup", SignUp).Methods("POST")

	r.HandleFunc("/auth/login", Login).Methods("POST")

	r.HandleFunc("/auth/logout", Logout).Methods("POST")

	// protected end point
	r.HandleFunc("/auth/welcome", WelcomeEndPoint).Methods("GET")

	if err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}