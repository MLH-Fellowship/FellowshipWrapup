package server

import (
	"backend/queries"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shurcooL/graphql"
)

type response struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	res := response{
		Status: "success",
		Body:   "Home page",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	util.LogCall("GET", "/", "200", startTime)
}

func FellowHandler(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)
	// vars here is the {username} field in the router
	vars := mux.Vars(req)
	// Checks to see if a secret field is sent to make sure no robots
	// are using up all our calls

	if !util.IsAuthorized(w, req) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "error",
			Body:   "Incorrect secret given, you are not authorized to use this API",
		}
		json.NewEncoder(w).Encode(res)

		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
		util.LogCall("POST", endPoint, "401", startTime)
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
		util.LogCall("POST", endPoint, "400", startTime)
		return
	}

	if !util.IsValidUsername(vars["username"]) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := response{
			Status: "error",
			Body:   "Invalid username given",
		}
		json.NewEncoder(w).Encode(res)

		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
		util.LogCall("POST", endPoint, "400", startTime)
		return
	}

	// If user wasn't already queried
	if !util.CheckUser(vars["username"]) {
		fmt.Fprintf(w, "User not found, quering") // TODO: enhance message

		// Query user data
		httpClient := util.SetupOAuth()
		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

		var tempStruct queries.MegaJSONStruct

		// Call the API with the relevant queries
		// TODO: correctly get json data here
		err := client.Query(context.Background(), &tempStruct.RepoContrib, nil)
		util.CheckAPICallErr(err)

		_, err = json.Marshal(tempStruct.RepoContrib)
		if err != nil {
			log.Fatal(err)
		}
		// TODO: save query data on user directory
	}

	endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
	util.LogCall("POST", endPoint, "200", startTime)

}
