package main

import (
	"backend/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	if os.Getenv("secretKey") == "" {
		log.Fatal("No secret key set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", server.HomeHandler).Methods("GET")
	r.HandleFunc("/getfellow/{username}", server.FellowHandler).Methods("POST")
	r.Use(server.VerificationMiddleware)

	log.Println("Starting web server on localhost:8080")
	http.ListenAndServe(":8080", r)

	// httpClient := SetupOAuth()
	// client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	// var tempStruct megaJSONStruct

	// // Call the API with the relevant queries
	// err := client.Query(context.Background(), &tempStruct.repoContrib, nil)
	// CheckAPICallErr(err)
	// err = client.Query(context.Background(), &tempStruct.prMerged, nil)
	// CheckAPICallErr(err)
	// err = client.Query(context.Background(), &tempStruct.prOpened, nil)
	// CheckAPICallErr(err)
	// err = client.Query(context.Background(), &tempStruct.issOpened, nil)
	// CheckAPICallErr(err)
	// err = client.Query(context.Background(), &tempStruct.issClosed, nil)
	// CheckAPICallErr(err)

	// err = client.Query(context.Background(), &tempStruct.PRContributions, nil)
	// CheckAPICallErr(err)

	// err = client.Query(context.Background(), &tempStruct.PRCommits, nil)
	// CheckAPICallErr(err)

	// err = client.Query(context.Background(), &tempStruct.accountInfo, nil)
	// CheckAPICallErr(err)

	// fmt.Println(tempStruct.accountInfo)

	// writeJSON(tempStruct)

	// fmt.Println(tempStruct.prOpened.Search.IssueCount)
	// fmt.Println(tempStruct.prMerged)
	// fmt.Println(tempStruct.issOpened)
	// fmt.Println(tempStruct.repoContrib.Viewer.RepositoriesContributedTo.TotalCount)

	// maybe add query to show team and team members

}
