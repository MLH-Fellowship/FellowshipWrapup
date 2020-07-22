package server

import (
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type response struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

// VerificationMiddleware is a middlware to handle authentication
// and checking if the username is valid before being passed onto
// the requested endpoint
func VerificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		startTime := time.Now().UnixNano() / int64(time.Millisecond)
		vars["startTime"] = strconv.FormatInt(startTime, 10)

		if auth, err := util.IsAuthorized(w, r); !auth {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			res := response{
				Status: "401",
				Body:   fmt.Sprint(err),
			}
			json.NewEncoder(w).Encode(res)

			util.LogCall(r.Method, r.RequestURI, "401", vars["startTime"], false)
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
			util.LogCall(r.Method, r.RequestURI, "400", vars["startTime"], false)
			return
		}

		next.ServeHTTP(w, r)

	})
}

// HomeHandler serves the content for the home page
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res := response{
		Status: "success",
		Body:   "Home page",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], false)
}

// GetFellowLinesOfCodeInPRs Get the additions and deletions of all
// PRs for a given user
func GetFellowLinesOfCodeInPRs(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// If user wasn't already queried and the cache doesn't exist then
	// we call the API and cache the result
	if !util.CheckUser(vars["username"], "PRContributions.json") {

		client := util.SetupOAuth()
		tempStruct, variables := util.Setup(vars["username"])

		// Call the API
		err := client.Query(context.Background(), &tempStruct.PRContributions, variables)
		util.CheckAPICallErr(err)

		util.WriteCache(vars["username"], "PRContributions", tempStruct.PRContributions)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.PRContributions)
		util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], false)
		return
	}
	// Serve from cache instead
	content, err := util.ServeCache(vars["username"], "PRContributions")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)
		util.LogCall(req.Method, req.RequestURI, "401", vars["startTime"], false)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, content)
	util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], true)

}

// GetFellowPullRequestCommits gets the commits from pull requests
func GetFellowPullRequestCommits(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "prCommits.json") {

		client := util.SetupOAuth()
		tempStruct, variables := util.Setup(vars["username"])

		// Call the API
		err := client.Query(context.Background(), &tempStruct.PRCommits, variables)
		util.CheckAPICallErr(err)

		util.WriteCache(vars["username"], "prCommits", tempStruct.PRCommits)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.PRCommits)
		util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], false)
		return
	}
	// Serve from cache instead
	content, err := util.ServeCache(vars["username"], "prCommits")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)
		util.LogCall(req.Method, req.RequestURI, "401", vars["startTime"], false)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, content)
	util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], true)

}

// GetFellowRepoContributions get a list of all repositories a user has
// contributed to
func GetFellowRepoContributions(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "repoContribs.json") {

		client := util.SetupOAuth()
		tempStruct, variables := util.Setup(vars["username"])

		// Call the API
		err := client.Query(context.Background(), &tempStruct.RepoContrib, variables)
		util.CheckAPICallErr(err)

		util.WriteCache(vars["username"], "repoContribs", tempStruct.RepoContrib)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.RepoContrib)
		util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], false)
		return
	}

	// Serve from cache instead
	content, err := util.ServeCache(vars["username"], "repoContribs")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)
		util.LogCall(req.Method, req.RequestURI, "401", vars["startTime"], false)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, content)
	util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], true)

}

// GetFellowPullRequests get a list of the most recent PRs made by a user
func GetFellowPullRequests(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "pullRequests.json") {

		client := util.SetupOAuth()
		tempStruct, variables := util.Setup(vars["username"])

		// Call the API
		err := client.Query(context.Background(), &tempStruct.Pr, variables)
		util.CheckAPICallErr(err)

		util.WriteCache(vars["username"], "pullRequests", tempStruct.Pr)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.Pr)
		util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], false)
		return
	}
	// Serve from cache instead
	content, err := util.ServeCache(vars["username"], "pullRequests")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)
		util.LogCall(req.Method, req.RequestURI, "401", vars["startTime"], false)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, content)
	util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], true)

}

// GetFellowIssuesCreated get a list of the recent issues created by a user
func GetFellowIssuesCreated(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "issuesCreated.json") {

		client := util.SetupOAuth()
		tempStruct, variables := util.Setup(vars["username"])

		// Call the API
		err := client.Query(context.Background(), &tempStruct.IssCreated, variables)
		util.CheckAPICallErr(err)

		util.WriteCache(vars["username"], "issuesCreated", &tempStruct.IssCreated)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.IssCreated)
		util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], false)
		return
	}
	// Serve from cache instead
	content, err := util.ServeCache(vars["username"], "issuesCreated")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)
		util.LogCall(req.Method, req.RequestURI, "401", vars["startTime"], false)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, content)
	util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], true)

}

// GetFellowAccountInfo get account information for a given user
func GetFellowAccountInfo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if !util.CheckUser(vars["username"], "accountInfo.json") {

		client := util.SetupOAuth()
		tempStruct, variables := util.Setup(vars["username"])

		// Call the API
		err := client.Query(context.Background(), &tempStruct.AccountInfo, variables)
		util.CheckAPICallErr(err)

		util.WriteCache(vars["username"], "accountInfo", tempStruct.AccountInfo)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.AccountInfo)
		util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], false)
		return
	}
	// Serve from cache instead
	content, err := util.ServeCache(vars["username"], "accountInfo")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		res := response{
			Status: "401",
			Body:   fmt.Sprint(err),
		}
		json.NewEncoder(w).Encode(res)
		util.LogCall(req.Method, req.RequestURI, "401", vars["startTime"], false)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, content)
	util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], true)

}
