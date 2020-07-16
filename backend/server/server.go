package server

import (
	"backend/queries"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/shurcooL/graphql"
)

type response struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func VerificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if auth, err := util.IsAuthorized(w, r); !auth {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			res := response{
				Status: "401",
				Body:   fmt.Sprint(err),
			}
			json.NewEncoder(w).Encode(res)

			util.LogCall("POST", r.RequestURI, "401")
			return
		}

		if v, err := util.IsValidUsername(vars["username"]); !v {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := response{
				Status: "422",
				Body:   fmt.Sprint(err),
			}
			json.NewEncoder(w).Encode(res)

			util.LogCall(r.Method, r.RequestURI, "400")
			return
		}

		next.ServeHTTP(w, r)

	})
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	res := response{
		Status: "success",
		Body:   "Home page",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	util.LogCall("GET", "/", "200")
}

// GetFellowLinesOfCodeInPRs get the issues created
func GetFellowLinesOfCodeInPRs(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// If user wasn't already queried
	if !util.CheckUser(vars["username"], "prContributions.json") {
		fmt.Println("Not Existing")
		// Query user data
		httpClient := util.SetupOAuth()
		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

		var tempStruct queries.MegaJSONStruct

		variables := map[string]interface{}{
			"username": graphql.String(vars["username"]),
		}

		// Call the API
		err := client.Query(context.Background(), &tempStruct.PRContributions, variables)
		util.CheckAPICallErr(err)

		// Write to JSON file
		dirLocation := fmt.Sprintf("../data/%s", vars["username"])
		_ = os.Mkdir(dirLocation, 0755)

		fileLocation := fmt.Sprintf("../data/%s/prContributions.json", vars["username"])
		jsonData, err := json.Marshal(tempStruct.PRContributions)
		if err != nil {
			log.Fatal(err)
		}

		_ = ioutil.WriteFile(fileLocation, jsonData, 0777)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.PRContributions)
		util.LogCall(req.Method, req.RequestURI, "200")
		return
	}

	fmt.Println("Existing")

	// get the cache and serve it
	fileLocation := fmt.Sprintf("../data/%s/prContributions.json", vars["username"])
	content, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)

		util.LogCall(req.Method, req.RequestURI, "401")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowPullRequestCommits get the issues created
func GetFellowPullRequestCommits(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// If user wasn't already queried
	if !util.CheckUser(vars["username"], "prCommits.json") {
		fmt.Println("Not Existing")
		// Query user data
		httpClient := util.SetupOAuth()
		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

		var tempStruct queries.MegaJSONStruct

		variables := map[string]interface{}{
			"username": graphql.String(vars["username"]),
		}

		// Call the API
		err := client.Query(context.Background(), &tempStruct.PRCommits, variables)
		util.CheckAPICallErr(err)

		// Write to JSON file
		dirLocation := fmt.Sprintf("../data/%s", vars["username"])
		_ = os.Mkdir(dirLocation, 0755)

		fileLocation := fmt.Sprintf("../data/%s/prCommits.json", vars["username"])
		jsonData, err := json.Marshal(tempStruct.PRCommits)
		if err != nil {
			log.Fatal(err)
		}

		_ = ioutil.WriteFile(fileLocation, jsonData, 0777)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.PRCommits)
		util.LogCall(req.Method, req.RequestURI, "200")
		return
	}

	fmt.Println("Existing")

	// get the cache and serve it
	fileLocation := fmt.Sprintf("../data/%s/prCommits.json", vars["username"])
	content, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)

		util.LogCall(req.Method, req.RequestURI, "401")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowRepoContributions get the issues created
func GetFellowRepoContributions(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// If user wasn't already queried
	if !util.CheckUser(vars["username"], "repoContribs.json") {
		fmt.Println("Not Existing")
		// Query user data
		httpClient := util.SetupOAuth()
		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

		var tempStruct queries.MegaJSONStruct

		variables := map[string]interface{}{
			"username": graphql.String(vars["username"]),
		}

		// Call the API
		err := client.Query(context.Background(), &tempStruct.RepoContrib, variables)
		util.CheckAPICallErr(err)

		// Write to JSON file
		dirLocation := fmt.Sprintf("../data/%s", vars["username"])
		_ = os.Mkdir(dirLocation, 0755)

		fileLocation := fmt.Sprintf("../data/%s/repoContribs.json", vars["username"])
		jsonData, err := json.Marshal(tempStruct.RepoContrib)
		if err != nil {
			log.Fatal(err)
		}

		_ = ioutil.WriteFile(fileLocation, jsonData, 0777)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.RepoContrib)
		util.LogCall(req.Method, req.RequestURI, "200")
		return
	}

	fmt.Println("Existing")

	// get the cache and serve it
	fileLocation := fmt.Sprintf("../data/%s/repoContribs.json", vars["username"])
	content, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)

		util.LogCall(req.Method, req.RequestURI, "401")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowPullRequests get the issues created
func GetFellowPullRequests(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// If user wasn't already queried
	if !util.CheckUser(vars["username"], "pullRequests.json") {
		fmt.Println("Not Existing")
		// Query user data
		httpClient := util.SetupOAuth()
		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

		var tempStruct queries.MegaJSONStruct

		variables := map[string]interface{}{
			"username": graphql.String(vars["username"]),
		}

		// Call the API
		err := client.Query(context.Background(), &tempStruct.Pr, variables)
		util.CheckAPICallErr(err)

		// Write to JSON file
		dirLocation := fmt.Sprintf("../data/%s", vars["username"])
		_ = os.Mkdir(dirLocation, 0755)

		fileLocation := fmt.Sprintf("../data/%s/pullRequests.json", vars["username"])
		jsonData, err := json.Marshal(tempStruct.Pr)
		if err != nil {
			log.Fatal(err)
		}

		_ = ioutil.WriteFile(fileLocation, jsonData, 0777)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.Pr)
		util.LogCall(req.Method, req.RequestURI, "200")
		return
	}

	fmt.Println("Existing")

	// get the cache and serve it
	fileLocation := fmt.Sprintf("../data/%s/pullRequests.json", vars["username"])
	content, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)

		util.LogCall(req.Method, req.RequestURI, "401")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowIssuesCreated get the issues created
func GetFellowIssuesCreated(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// If user wasn't already queried
	if !util.CheckUser(vars["username"], "issuesCreated.json") {

		// Query user data
		httpClient := util.SetupOAuth()
		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

		var tempStruct queries.MegaJSONStruct

		variables := map[string]interface{}{
			"username": graphql.String(vars["username"]),
		}

		// Call the API
		err := client.Query(context.Background(), &tempStruct.IssCreated, variables)
		util.CheckAPICallErr(err)

		// Write to JSON file
		dirLocation := fmt.Sprintf("../data/%s", vars["username"])
		_ = os.Mkdir(dirLocation, 0755)

		fileLocation := fmt.Sprintf("../data/%s/issuesCreated.json", vars["username"])
		jsonData, err := json.Marshal(tempStruct.IssCreated)
		if err != nil {
			log.Fatal(err)
		}

		_ = ioutil.WriteFile(fileLocation, jsonData, 0777)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.IssCreated)
		util.LogCall(req.Method, req.RequestURI, "200")
		return
	}

	// get the cache and serve it
	fileLocation := fmt.Sprintf("../data/%s/issuesCreated.json", vars["username"])
	content, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)

		util.LogCall(req.Method, req.RequestURI, "401")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowAccountInfo hell
func GetFellowAccountInfo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// If user wasn't already queried
	if !util.CheckUser(vars["username"], "accountInfo.json") {
		// Query user data
		httpClient := util.SetupOAuth()
		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

		var tempStruct queries.MegaJSONStruct

		variables := map[string]interface{}{
			"username": graphql.String(vars["username"]),
		}

		// Call the API
		err := client.Query(context.Background(), &tempStruct.AccountInfo, variables)
		util.CheckAPICallErr(err)

		// Write to JSON file
		dirLocation := fmt.Sprintf("../data/%s", vars["username"])
		_ = os.Mkdir(dirLocation, 0755)

		fileLocation := fmt.Sprintf("../data/%s/accountInfo.json", vars["username"])
		jsonData, err := json.Marshal(tempStruct.AccountInfo)
		if err != nil {
			log.Fatal(err)
		}

		_ = ioutil.WriteFile(fileLocation, jsonData, 0777)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.AccountInfo)
		util.LogCall(req.Method, req.RequestURI, "200")
		return
	}
	// get the cache and serve it
	fileLocation := fmt.Sprintf("../data/%s/accountInfo.json", vars["username"])
	content, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)

		util.LogCall(req.Method, req.RequestURI, "401")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// FellowHandler ypup
// func FellowHandler(w http.ResponseWriter, req *http.Request) {
// 	vars := mux.Vars(req)

// 	// If user wasn't already queried
// 	if !util.CheckUser(vars["username"]) {

// 		// Query user data
// 		httpClient := util.SetupOAuth()
// 		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

// 		var tempStruct queries.MegaJSONStruct

// 		// Call the API with the relevant queries
// 		// TODO: correctly get json data here
// 		err := client.Query(context.Background(), &tempStruct.RepoContrib, nil)
// 		util.CheckAPICallErr(err)

// 		// TODO: save query data on user directory
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(tempStruct.RepoContrib)
// 		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
// 		util.LogCall("POST", endPoint, "200")

// 	}

// }
