package main

import (
	"auth/db"
	"auth/models"
	"auth/user"
	"auth/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
		signupResponse := models.SignupResponse{}
		// Account
		signupResponse.Account = models.UserAccount{}
		signupResponse.Account.Email = user.Email
		signupResponse.Account.Role = user.Role
		// Details
		signupResponse.Details = models.UserDetails{}
		signupResponse.Details.Name = user.Name

		payLoad := models.CreateUserRequest{}
		payLoad.Email = user.Email
		payLoad.First = user.Name.First
		payLoad.Last = user.Name.Last
		jsonResponse, _ := json.Marshal(payLoad)
		fmt.Println(string(jsonResponse))

		//_, _ = http.Post("localhost:9001/user", "application/json", jsonData)
		respondWithJson(w, http.StatusOK, signupResponse)
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
	if !validUser {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	var jwt = utils.CreateJWT(dbUser)
	var response models.LoginResponse
	response.Token = jwt
	response.Message = "ok"

	respondWithJson(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request)  {
	authHeader :=  r.Header.Get("Authorization")
	if utils.InValidateJWT(authHeader) {
		_, _ = w.Write([]byte(fmt.Sprintf("Logout successful!")))
	} else {
		_, _ = w.Write([]byte(fmt.Sprintf("Session Not found!")))
	}
}

func WelcomeEndPoint(w http.ResponseWriter, r *http.Request)  {
	authHeader :=  r.Header.Get("Authorization")
	if !utils.ValidateJWT(authHeader) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Welcome!")))
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)
	//fmt.Println(string(response))
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

	if err := http.ListenAndServe(":9000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}