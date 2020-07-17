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

// VerificationMiddleware is a middlware to handle authentication
// and checking is the username is valid before being passed onto
// the requested endpoint
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

// HomeHandler serves the content for the home page
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	res := response{
		Status: "success",
		Body:   "Home page",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	util.LogCall(req.Method, req.RequestURI, "200")
}

// GetFellowLinesOfCodeInPRs Get the additions and deletions of all
// PRs for a given user
func GetFellowLinesOfCodeInPRs(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// If user wasn't already queried
	if !util.CheckUser(vars["username"], "prContributions.json") {
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
	// Serve from cache instead
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
	fmt.Fprintf(w, string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowPullRequestCommits gets the commits from pull requests
func GetFellowPullRequestCommits(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "prCommits.json") {
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
	// Serve from cache instead
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
	fmt.Fprintf(w, string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowRepoContributions get a list of all repositories a user has
// contributed to
func GetFellowRepoContributions(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "repoContribs.json") {
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

	// Serve from cache instead
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
	fmt.Fprintf(w, string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowPullRequests get a list of the most recent PRs made by a user
func GetFellowPullRequests(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "pullRequests.json") {
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

	// Serve from cache instead
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
	fmt.Fprintf(w, string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowIssuesCreated get a list of the recent issues created by a user
func GetFellowIssuesCreated(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "issuesCreated.json") {

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
	// Serve from cache instead
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
	fmt.Fprintf(w, string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}

// GetFellowAccountInfo get account information for a given user
func GetFellowAccountInfo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "accountInfo.json") {
		fmt.Println("Calling API")
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
		err = os.Mkdir(dirLocation, 0755)
		fmt.Println(err)

		fileLocation := fmt.Sprintf("../data/%s/accountInfo.json", vars["username"])
		jsonData, err := json.Marshal(tempStruct.AccountInfo)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(fileLocation, jsonData, 0777)
		fmt.Println(err)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.AccountInfo)
		util.LogCall(req.Method, req.RequestURI, "200")
		return
	}

	fmt.Println("Calling cache")
	// Serve from cache instead
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
	fmt.Fprintf(w, string(content))
	util.LogCall(req.Method, req.RequestURI, "200")

}
