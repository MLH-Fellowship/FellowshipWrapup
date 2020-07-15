package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/shurcooL/graphql"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Working home directory")
}

func getFellowHandler(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)
	// vars here is the {username} field in the router
	vars := mux.Vars(req)
	// Checks to see if a secret field is sent to make sure no robots
	// are using up all our calls
	authorized := isAuthorized(w, req)

	if !authorized {
		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
		logCall("POST", endPoint, "401", startTime)
		fmt.Fprintf(w, "You are not authorized to use this API")
		return
	}

	if vars["username"] == "" {
		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
		logCall("POST", endPoint, "400", startTime)
		fmt.Fprintf(w, "No username given")
		return
	}

	// Check if user was already queried
	if CheckUser(vars["username"]) {
		// TODO: get json data here

		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
		logCall("POST", endPoint, "200", startTime)
		w.WriteHeader(http.StatusOK)
		// fmt.Fprintf(w, string(jsonData))
		fmt.Fprintf(w, "User found, using cached data") // TODO: add real json data here

		return
	}

	// Query user data
	httpClient := SetupOAuth()
	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	var tempStruct megaJSONStruct

	// Call the API with the relevant queries
	err := client.Query(context.Background(), &tempStruct.repoContrib, nil)
	CheckAPICallErr(err)

	jsonData, err := json.Marshal(tempStruct.repoContrib)
	if err != nil {
		log.Fatal(err)
	}

	endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
	logCall("POST", endPoint, "200", startTime)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonData))

}

func main() {

	if os.Getenv("secretKey") == "" {
		log.Fatal("No secret key set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/getfellow/{username}", getFellowHandler).Methods("POST")

	log.Println("Starting web server on localhost:8080")
	http.ListenAndServe(":8080", r)

	// httpClient := SetupOAuth()
	// client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	// var tempStruct megaJSONStruct

	// // Call the API with the relevant queries
	// err := client.Query(context.Background(), &tempStruct.repoContrib, nil)
	// CheckAPICallErr(err)
	// err = client.Query(context.Background(), &tempStruct.prMerged, nil)
	// CheckAPICallErr(err)
	// err = client.Query(context.Background(), &tempStruct.prOpened, nil)
	// CheckAPICallErr(err)
	// err = client.Query(context.Background(), &tempStruct.issOpened, nil)
	// CheckAPICallErr(err)
	// err = client.Query(context.Background(), &tempStruct.issClosed, nil)
	// CheckAPICallErr(err)

	// err = client.Query(context.Background(), &tempStruct.PRContributions, nil)
	// CheckAPICallErr(err)

	// err = client.Query(context.Background(), &tempStruct.PRCommits, nil)
	// CheckAPICallErr(err)

	// fmt.Println(tempStruct.issClosed)

	// writeJSON(tempStruct)

	// fmt.Println(tempStruct.prOpened.Search.IssueCount)
	// fmt.Println(tempStruct.prMerged)
	// fmt.Println(tempStruct.issOpened)
	// fmt.Println(tempStruct.repoContrib.Viewer.RepositoriesContributedTo.TotalCount)

	// maybe add query to show team and team members

}
