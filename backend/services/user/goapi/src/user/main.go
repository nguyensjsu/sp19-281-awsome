package main

import (
	"encoding/json"
	"fmt"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"os"
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

	users,err := dao.FindAll()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
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
	fmt.Print(user.FirstName)
	retuser,err := dao.CreateUser(user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, retuser)
}

// POST /create/user : Testing Login
func TestLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var login Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Print(login.Email)
	//if err != nil {
	//	respondWithError(w, http.StatusBadRequest, err.Error())
	//	return
	//}
	var user =	 User {"test","Dharma","Dheeraj","applicant"}
	respondWithJson(w, http.StatusOK, user)
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

func init() {
	dao.Database = os.Getenv("MONGO_DATABASE")
	dao.Server = os.Getenv("MONGO_SERVER")

	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	r.HandleFunc("/create/user", CreateUser).Methods("POST")
	r.HandleFunc("/getUsers", GetAllUsers).Methods("GET")
	r.HandleFunc("/auth/login", TestLogin).Methods("POST")
	
	if err := http.ListenAndServe(":9000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}