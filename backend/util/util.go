package util

import (
	"backend/queries"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

type reqStruct struct {
	Secret string `json:"secret"`
}

type Response struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

// CheckAPICallErr checks the error value on an API call
func CheckAPICallErr(err error) error {
	if err == nil {
		return nil
	}
	return err
}

// CacheExists returns whether the given cached file exists
func CacheExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// dirEmpty returns whether the given directory is empty
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

// SetupOAuth setups the OAuth2 client needed to make
// calls to the GitHub V4 graphQL API
func SetupOAuth(accessToken string) *graphql.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := graphql.NewClient("https://api.github.com/graphql", httpClient)
	return client
}

// Log an API call to the console with it's details
func LogCall(method, endpoint, status, startTimeString string, cached bool) {
	statusColor := "\033[0m"
	cacheString := ""

	if cached {
		cacheString = "[CACHE] "
	}

	startTime, err := strconv.ParseInt(startTimeString, 10, 64)
	if err != nil {
		startTime = -1
	}
	endTime := time.Now().UnixNano() / int64(time.Millisecond)
	delay := endTime - startTime

	// If the HTTP status given is 2XX, give it a nice
	// green color, otherwise give it a red color
	if status[0] == '2' {
		statusColor = "\033[32m"
	} else {
		statusColor = "\033[31m"
	}
	fmt.Printf("[%s]%s%s %s %s%s%s %dms\n", time.Now().Format("02-Jan-2006 15:04:05"), cacheString, method, endpoint, statusColor, status, "\033[0m", delay)
}

// IsValidUsername checks if a github username exists.
// It pings the github profile and if the header contains
// a non 200 status code the profile doesnt exist and we dont
// call the API.
func IsValidUsername(username, accessToken string) bool {
	// Empty username will yield 200 on github
	if username == "" {
		return false
	}

	client := SetupOAuth(accessToken)
	var tempStruct struct {
		User struct {
			Login graphql.String
		} `graphql:"user(login: $username)"`
	}
	variables := map[string]interface{}{
		"username": graphql.String(username),
	}

	err := client.Query(context.Background(), &tempStruct, variables)
	CheckAPICallErr(err)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// HasAccessToken gets the accessToken field from the request body
// Adds it to the vars map
// Returns true if it finds an accessToken
// Returns false otherwise
func HasAccessToken(r *http.Request, vars map[string]string) bool {
	var body struct {
		AccessToken string
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil || body.AccessToken == "" {
		return false
	}
	vars["accessToken"] = body.AccessToken
	return true
}

// isFellow determines if a given user is a member
// of the MLH-Fellowship organisation
func IsFellow(username, accessToken string) bool {
	client := SetupOAuth(accessToken)

	var tempStruct struct {
		User struct {
			Organization struct {
				Name graphql.String
			} `graphql:"organization(login: $org)"`
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": graphql.String(username),
		"org":      graphql.String("MLH-Fellowship"),
	}

	// Call the API
	err := client.Query(context.Background(), &tempStruct, variables)
	CheckAPICallErr(err)

	if err != nil {
		log.Fatal(err)
		return false
	}

	// When a user is not a member of the MLH-Fellowship org
	// it fails it find it in their org list and is an empty string
	if tempStruct.User.Organization.Name == "" {
		return false
	}
	return true
}

// WriteCache writes a struct to its associated cache file for
// a given user
func WriteCache(username, filename string, data interface{}) {
	// Write to JSON file
	dirLocation := fmt.Sprintf("../data/%s", username)
	_ = os.Mkdir(dirLocation, 0755)

	fileLocation := fmt.Sprintf("../data/%s/%s.json", username, filename)

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("ok lads")
		log.Fatal(err)
	}
	_ = ioutil.WriteFile(fileLocation, jsonData, 0777)
}

// IsValidQueryType determines if an incoming query is
// implemented by the service
func IsValidQueryType(query string) (string, error) {

	validTypes := map[string]bool{
		"accountinfo":          true,
		"pullrequests":         true,
		"involvedissues":       true,
		"openvsclosedissues":   true,
		"reposcontributedto":   true,
		"mergedvsnonmergedprs": true,
		"podinformation":       true,
	}
	query = strings.ToLower(query)

	if validTypes[query] {
		return fmt.Sprintf("%s.json", query), nil
	}
	return "", errors.New("Invalid query type given")
}

// GetCache returns the cached result for a given user and filename
func GetCache(username, filename string) (string, error) {

	fileLocation := fmt.Sprintf("../data/%s/%s", username, filename)
	content, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		return "", errors.New("Invalid username given, cache not found")
	}

	return string(content), nil
}

// SendErrorResponse sends a templated error response to the user
func SendErrorResponse(w http.ResponseWriter, r *http.Request, httpStatus int, startTime, errorString string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	res := Response{
		Status: httpStatus,
		Body:   errorString,
	}
	json.NewEncoder(w).Encode(res)
	LogCall(r.Method, r.URL.Path, strconv.Itoa(httpStatus), startTime, false)
	return
}

// GetStruct returns the correct struct type based on the query type given
func GetStruct(query, username string) (interface{}, map[string]interface{}) {
	tempStruct := &queries.MegaJSONStruct{}

	variables := map[string]interface{}{
		"username": graphql.String(username),
	}

	switch strings.ToLower(query) {
	case "accountinfo":
		structType := reflect.TypeOf(tempStruct.AccountInfo)
		return reflect.New(structType).Interface(), variables
	case "pullrequests":
		structType := reflect.TypeOf(tempStruct.PRs)
		return reflect.New(structType).Interface(), variables
	case "involvedissues":
		structType := reflect.TypeOf(tempStruct.InvolveIssues)
		return reflect.New(structType).Interface(), variables
	case "openvsclosedissues":
		structType := reflect.TypeOf(tempStruct.OpenVsClosedIssues)
		return reflect.New(structType).Interface(), variables
	case "reposcontributedto":
		structType := reflect.TypeOf(tempStruct.ReposContribedTo)
		return reflect.New(structType).Interface(), variables
	case "mergedvsnonmergedprs":
		structType := reflect.TypeOf(tempStruct.MergedVsNonMergedPRs)
		return reflect.New(structType).Interface(), variables
	case "podinformation":
		structType := reflect.TypeOf(tempStruct.PodInfo)
		variables["org"] = graphql.String("MLH-Fellowship")
		// Used to only query for teams that include "pod" in their title
		// e.g exclude CTF and mentor teams
		variables["pod"] = graphql.String("pod")
		return reflect.New(structType).Interface(), variables
	default:
		return nil, variables
	}
}
