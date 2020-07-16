package main

import (
	"backend/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	if os.Getenv("secretKey") == "" {
		log.Fatal("No secret key set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", server.HomeHandler).Methods("GET")
	r.HandleFunc("/getfellow/accountinfo/{username}", server.GetFellowAccountInfo).Methods("POST")
	r.HandleFunc("/getfellow/issuescreated/{username}", server.GetFellowIssuesCreated).Methods("POST")
	r.HandleFunc("/getfellow/pullrequests/{username}", server.GetFellowPullRequests).Methods("POST")
	r.Use(server.VerificationMiddleware)

	log.Println("Starting web server on localhost:8080")
	http.ListenAndServe(":8080", r)

}
