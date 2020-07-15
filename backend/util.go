package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
		log.Fatal("Error: You have not set your GRAPHQL_TOKEN environment variable. Visit https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token to generate a token")
	}

	log.Fatal(err)
}

// fileExists returns whether the given file or directory exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// dirEmpty return s whether the given directory is empty
func dirEmpty(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true
	}
	return false // Either not empty or error, suits both cases
}

// CheckUser checks if a user was already queried
// Returns true if username is found on the /data dir
// Returns false if username is not found
// Returns false if username is found but is empty
func CheckUser(username string) bool {
	var userPath strings.Builder
	// Build path string
	userPath.WriteString("data/")
	userPath.WriteString(username)

	// Check if directory /data/{username} exists
	if !fileExists(userPath.String()) {
		return false
	}

	// Check if directory /data/{username} is empty
	if dirEmpty(userPath.String()) {
		os.Remove(userPath.String()) // Delete empty dir
		return false
	}

	return true
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
