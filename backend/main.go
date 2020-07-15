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

type response struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	res := response{
		Status: "success",
		Body:   "Home page",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	logCall("GET", "/", "200", startTime)
}

func getFellowHandler(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)
	// vars here is the {username} field in the router
	vars := mux.Vars(req)
	// Checks to see if a secret field is sent to make sure no robots
	// are using up all our calls

	if !isAuthorized(w, req) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "error",
			Body:   "Incorrect secret given, you are not authorized to use this API",
		}
		json.NewEncoder(w).Encode(res)

		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
		logCall("POST", endPoint, "401", startTime)
		return
	}

	if vars["username"] == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := response{
			Status: "error",
			Body:   "No username given",
		}
		json.NewEncoder(w).Encode(res)

		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
		logCall("POST", endPoint, "400", startTime)
		return
	}

	if !isValidUsername(vars["username"]) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := response{
			Status: "error",
			Body:   "Invalid username given",
		}
		json.NewEncoder(w).Encode(res)

		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
		logCall("POST", endPoint, "400", startTime)
		return
	}

	// If user wasn't already queried
	if !CheckUser(vars["username"]) {
		fmt.Fprintf(w, "User not found, quering") // TODO: enhance message

		// Query user data
		httpClient := SetupOAuth()
		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

		var tempStruct megaJSONStruct

		// Call the API with the relevant queries
		// TODO: correctly get json data here
		err := client.Query(context.Background(), &tempStruct.repoContrib, nil)
		CheckAPICallErr(err)

		jsonData, err := json.Marshal(tempStruct.repoContrib)
		if err != nil {
			log.Fatal(err)
		}
		// TODO: save query data on user directory
	}


	endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
	logCall("POST", endPoint, "200", startTime)

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

	// err = client.Query(context.Background(), &tempStruct.accountInfo, nil)
	// CheckAPICallErr(err)

	// fmt.Println(tempStruct.accountInfo)

	// writeJSON(tempStruct)

	// fmt.Println(tempStruct.prOpened.Search.IssueCount)
	// fmt.Println(tempStruct.prMerged)
	// fmt.Println(tempStruct.issOpened)
	// fmt.Println(tempStruct.repoContrib.Viewer.RepositoriesContributedTo.TotalCount)

	// maybe add query to show team and team members

}
