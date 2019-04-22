package main

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
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w,http.StatusOK, jobs)
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
	r.HandleFunc("/create/job", CreateJob).Methods("POST")
	r.HandleFunc("/getJobs", GetAllJobs).Methods("GET")
	
	if err := http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}