package main

import (
	"backend/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", server.HomeHandler).Methods("GET")
	r.HandleFunc("/{query}/{username}", server.Query).Methods("POST")
	r.Use(server.VerificationMiddleware)

	log.Println("Starting web server on localhost:8080")
	http.ListenAndServe(":8080", r)

}
