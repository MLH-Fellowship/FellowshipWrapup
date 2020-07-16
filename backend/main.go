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
	r.HandleFunc("/accountinfo/{username}", server.GetFellowAccountInfo).Methods("POST")
	r.HandleFunc("/issuescreated/{username}", server.GetFellowIssuesCreated).Methods("POST")
	r.HandleFunc("/pullrequests/{username}", server.GetFellowPullRequests).Methods("POST")
	r.HandleFunc("/repocontributedto/{username}", server.GetFellowRepoContributions).Methods("POST")
	r.HandleFunc("/pullrequestcommits/{username}", server.GetFellowPullRequestCommits).Methods("POST")
	r.HandleFunc("/prcontributions/{username}", server.GetFellowLinesOfCodeInPRs).Methods("POST")
	r.Use(server.VerificationMiddleware)

	log.Println("Starting web server on localhost:8080")
	http.ListenAndServe(":8080", r)

}
