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

// VerificationMiddleware handles authentication
// and checking if the username and query type is valid before being
// passed onto the requested endpoint
func VerificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		startTime := time.Now().UnixNano() / int64(time.Millisecond)
		vars["startTime"] = strconv.FormatInt(startTime, 10)

		if auth, err := util.IsAuthorized(w, r); !auth {
			util.SendErrorResponse(w, r, http.StatusUnauthorized, vars["startTime"], fmt.Sprint(err))
			return
		}

		if validUsername := util.IsValidUsername(vars["username"]); !validUsername {
			util.SendErrorResponse(w, r, http.StatusBadRequest, vars["startTime"], "Invalid username given")
			return
		}

		fileName, err := util.IsValidQueryType(vars["query"])
		if err != nil {
			util.SendErrorResponse(w, r, http.StatusUnauthorized, vars["startTime"], fmt.Sprint(err))
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
	if !util.CacheExists(fmt.Sprintf("../data/%s/%s", vars["username"], vars["fileName"])) {

		client := util.SetupOAuth()
		dataStruct, variables := util.GetStruct(vars["query"], vars["username"])
		if dataStruct == nil {
			util.SendErrorResponse(w, req, http.StatusUnauthorized, vars["startTime"], fmt.Sprint("Invalid query type given"))
			return
		}

		err := client.Query(context.Background(), dataStruct, variables)
		if err = util.CheckAPICallErr(err); err != nil {
			// Catches the edge case where someone used a username such as "explore" or "marketplace" which are valid
			// github links (https://github.com/marketplace) but are not user profiles and therefore can't be queried
			match, err := regexp.MatchString(`(Could not resolve to a User with the login of \')(.)+(\')`, err.Error())
			if err != nil && !match {
				util.SendErrorResponse(w, req, http.StatusUnauthorized, vars["startTime"], fmt.Sprint("Invalid query type given"))
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
		util.SendErrorResponse(w, req, http.StatusUnauthorized, vars["startTime"], fmt.Sprint(err))
		return

	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, content)
	util.LogCall(req.Method, req.RequestURI, "200", vars["startTime"], true)

}
