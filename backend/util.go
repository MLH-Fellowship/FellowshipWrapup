package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

type reqStruct struct {
	Secret string
}

// CheckAPICallErr test
func CheckAPICallErr(err error) {
	if err == nil {
		return
	}
	if os.Getenv("GRAPHQL_TOKEN") == "" {
		log.Fatal("Error: You have not set your GRAPHQL_TOKEN envivironment variable. Visit https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token to generate a token")
	}

	log.Fatal(err)
}

// SetupOAuth test
func SetupOAuth() *http.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return httpClient
}

func logCall(method, endpoint, status string, startTime int64) {
	endTime := time.Now().UnixNano() / int64(time.Millisecond)
	roundTripTime := endTime - startTime
	delay := strconv.FormatInt(roundTripTime, 10)
	statusColor := "\033[0m"

	// If the HTTP status given is 2XX, give it a nice
	// green color, otherwise give it a red color
	if status[0] == '2' {
		statusColor = "\033[32m"
	} else {
		statusColor = "\033[31m"
	}
	fmt.Printf("[%s] %s %s %s%s%s %sms\n", time.Now().Format("02-Jan-2006 15:04:05"), method, endpoint, statusColor, status, "\033[0m", delay)
}

func isAuthorized(w http.ResponseWriter, r *http.Request) bool {
	decoder := json.NewDecoder(r.Body)
	var req reqStruct

	err := decoder.Decode(&req)
	if err != nil {
		// couldn't decode the POST data into the right
		// JSON object
		return false
	}
	if req.Secret == os.Getenv("secretKey") {
		return true
	}
	return false
}
