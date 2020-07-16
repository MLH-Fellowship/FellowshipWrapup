package server

import (
	"backend/queries"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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

func getFellowAccountInfo(w http.ResponseWriter, req *http.Request) {

}

func FellowHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// If user wasn't already queried
	if !util.CheckUser(vars["username"]) {

		// Query user data
		httpClient := util.SetupOAuth()
		client := graphql.NewClient("https://api.github.com/graphql", httpClient)

		var tempStruct queries.MegaJSONStruct

		// Call the API with the relevant queries
		// TODO: correctly get json data here
		err := client.Query(context.Background(), &tempStruct.RepoContrib, nil)
		util.CheckAPICallErr(err)

		// TODO: save query data on user directory
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempStruct.RepoContrib)
		endPoint := fmt.Sprintf("/getfellow/%s", vars["username"])
		util.LogCall("POST", endPoint, "200")

	}

}
