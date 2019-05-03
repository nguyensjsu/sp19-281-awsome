package main

import (
	"fmt"
	"log"
	http "net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)


//Test Status of the API
func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Search API is alive!")
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	
	if err := http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}