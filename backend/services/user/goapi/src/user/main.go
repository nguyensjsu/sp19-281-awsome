package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
    "os"
	http "net/http"
	. "user/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "user/DAO"
)


//Test Status of the API
func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users API is alive!")
	var test bson.ObjectId
	var s = "5cc8c3f7408e8be47a282b8f"
	test = bson.NewObjectId()
	fmt.Print(test)
	fmt.Print(test.Hex())
	fmt.Print(bson.ObjectIdHex(s))

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

//Get One users
func GetOneUser(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	params := mux.Vars(r)
	id := params["id"]
	fmt.Print("Finding Id :" + id)
	users,err := dao.FindOne(id)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w,http.StatusOK, users)
}

// Create a User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Print(user.FirstName)
	user.ID = bson.NewObjectId()
	retuser,err := dao.CreateUser(user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, retuser)
}

// Update a User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	retuser,err := dao.UpdateUser(user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJson(w, http.StatusAccepted, retuser)
}



// Delete a User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	id := params["id"]
	fmt.Print("Deleting test")
	var result string
	result,err := dao.DeleteUser(bson.ObjectIdHex(id))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w,http.StatusOK,result)
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

	var user Login
	user.Email = login.Email
	user.Password = "applicant"
	user.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTY5Mjk4OTUsInJvbGUiOiJhcHBsaWNhbnQiLCJzdWIiOiJuaWtpdGFAZ21haWwuY29tIn0.coeMiWf0MMSJPJ72z8Ik34lSD5MpYJR2HJ1xpLEOKdQ"
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
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetOneUser).Methods("GET")
	r.HandleFunc("/users", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/auth/login", TestLogin).Methods("POST")

	if err := http.ListenAndServe(":9000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}