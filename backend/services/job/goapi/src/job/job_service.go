package main
//
// This service will be exposing the /jobs API 
// 
import (
	"encoding/json"
	"fmt"
	"log"
	http "net/http"
	. "job/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "job/DAO"
)


//Test Status of the API
func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jobs API is alive!")
}

var dao = JobDatabase{}


//Get All jobs
func GetAllJobs(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	jobs,err := dao.FindAll()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	respondWithJson(w,http.StatusOK, jobs)
}

//Get job with ID
func GetJobWithID(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	params := mux.Vars(r)
	fmt.Println(params["id"])
	job, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Job ID")
		return
	}
	respondWithJson(w, http.StatusOK, job)
}

// POST /create/job : Create a Job
func CreateJob(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var job Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Print(job.Title)
	retjob,err := dao.CreateJob(job)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, retjob)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

//Testing Login
func TestLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var login Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Print(login.Email)
	var user =	 User {"test1","Avinav1","Tyagi1","recruiter"}
	respondWithJson(w, http.StatusOK, user)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)
	fmt.Println(string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	dao.Database = "cmpe281"
	dao.Server = "mongodb://cmpe281:cmpe281@3.89.47.220:27017"

	dao.Connect()
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	r.HandleFunc("/jobs", CreateJob).Methods("POST")
	r.HandleFunc("/jobs", GetAllJobs).Methods("GET")
	r.HandleFunc("/jobs/{id}", GetJobWithID).Methods("GET")
	r.HandleFunc("/auth/login", TestLogin).Methods("POST")
	
	if err := http.ListenAndServe(":9000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}