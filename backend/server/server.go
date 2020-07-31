package server

import (
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type response struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

// VerificationMiddleware is a middlware to handle authentication
// and checks if username and query type are valid before being
// passed onto the requested endpoint
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

		if validUsername := util.IsValidUsername(vars["username"]); !validUsername {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := response{
				Status: "422",
				Body:   "Invalid username given",
			}
			json.NewEncoder(w).Encode(res)
			util.LogCall(r.Method, r.RequestURI, "400", vars["startTime"], false)
			return
		}

		fileName, err := util.IsValidQueryType(vars["query"])
		if err != nil {
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

		vars["fileName"] = strings.ToLower(fileName)
		vars["query"] = strings.ToLower(vars["query"])
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

func Query(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// Call the GitHub API and cache the result
	if !util.CacheExists(vars["username"], vars["fileName"]) {

		client := util.SetupOAuth()
		dataStruct, variables := util.GetStruct(vars["query"], vars["username"])
		if dataStruct == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			res := response{
				Status: "401",
				Body:   fmt.Sprint("Invalid query type given"),
			}
			json.NewEncoder(w).Encode(res)
			util.LogCall(req.Method, req.RequestURI, "401", vars["startTime"], false)
			return
		}

		err := client.Query(context.Background(), dataStruct, variables)
		err = util.CheckAPICallErr(err)
		if err != nil {
			// This catches errors thrown due to invalid usernames which is rare if not caught by the
			// verification middleware
			match, err := regexp.MatchString(`(Could not resolve to a User  with the login of \')(.)+(\')`, err.Error())
			if err != nil && !match {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				res := response{
					Status: "401",
					Body:   fmt.Sprint("Invalid username given"),
				}
				json.NewEncoder(w).Encode(res)
				util.LogCall(req.Method, req.RequestURI, "401", vars["startTime"], false)
				return
			}
		}

		util.WriteCache(vars["username"], vars["query"], dataStruct)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dataStruct)

		util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], false)
		return
	}
	// Serve from cache instead
	content, err := util.GetCache(vars["username"], vars["fileName"])
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
