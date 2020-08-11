package main

import (
	"backend/server"
	"backend/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", server.HomeHandler).Methods("GET")
	r.HandleFunc("/{query}/{username}", server.Query).Methods("POST")
	r.Use(server.VerificationMiddleware)

	return r
}
func TestAPIStatus(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	initRouter().ServeHTTP(res, req)

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error("Error reading response")
	}

	resJSON := util.Response{}
	err = json.Unmarshal(response, &resJSON)
	if err != nil {
		t.Error("Couldn't unmarshal response into correct format")
	}

	if resJSON.Status != 200 || resJSON.Body != "API is operational" {
		t.Errorf("Root endpoint not operational")
	}
}
