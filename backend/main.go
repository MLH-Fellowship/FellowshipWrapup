package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/shurcooL/graphql"
)

type megaJSONStruct struct {
	repoContrib repositoriesContributedTo
	prOpened    pullRequestsOpened
	prMerged    pullRequestsMerged
	issOpened   issuesOpened
}

type repositoriesContributedTo struct {
	Viewer struct {
		Login                     graphql.String
		RepositoriesContributedTo struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Name graphql.String
				Url  graphql.String
			}
		} `graphql:"repositoriesContributedTo(includeUserRepositories: true, first: 100, contributionTypes: [PULL_REQUEST])"`
	}
}

// Nodes      []struct {
// 	PullRequest struct {
// 		Title  graphql.String
// 		Url    graphql.String
// 		Merged graphql.Boolean

// 		Participants struct {
// 			TotalCount graphql.Int
// 			Nodes      []struct {
// 				Login graphql.String
// 				Url   graphql.String
// 			}
// 		} `graphql:"participants(first:30)"`
// 	} `graphql:"... on PullRequest"`
// }

type pullRequestsOpened struct {
	Search struct {
		IssueCount graphql.Int
	} `graphql:"search(query: \"is:pr author:@me created:2020-06-01..2020-08-30\", type: ISSUE, first: 100)"`
}

type pullRequestsMerged struct {
	Search struct {
		IssueCount graphql.Int
	} `graphql:"search(query: \"is:pr author:@me merged:2020-06-01..2020-08-30\", type: ISSUE, first: 100)"`
}

type issuesOpened struct {
	Search struct {
		IssueCount graphql.Int
	} `graphql:"search(query: \"is:issue author:@me created:2020-06-01..2020-08-30\", type: ISSUE, first: 100)"`
}

func writeJSON(jsonStruct megaJSONStruct) {
	jsonData, err := json.Marshal(jsonStruct.issOpened)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/issuesOpened.json", jsonData, 0644)
	fmt.Println(string(jsonData))

	jsonData, err = json.Marshal(jsonStruct.prMerged)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/prMerged.json", jsonData, 0644)
	fmt.Println(string(jsonData))

	jsonData, err = json.Marshal(jsonStruct.prOpened)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/prOpened.json", jsonData, 0644)
	fmt.Println(string(jsonData))

	jsonData, err = json.Marshal(jsonStruct.repoContrib)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/repoContribTo.json", jsonData, 0644)
	fmt.Println(string(jsonData))
}

func main() {
	httpClient := SetupOAuth()
	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	// Call the API with the relevant queries
	var tempStruct megaJSONStruct

	err := client.Query(context.Background(), &tempStruct.repoContrib, nil)
	CheckAPICallErr(err)
	err = client.Query(context.Background(), &tempStruct.prMerged, nil)
	CheckAPICallErr(err)
	err = client.Query(context.Background(), &tempStruct.prOpened, nil)
	CheckAPICallErr(err)
	err = client.Query(context.Background(), &tempStruct.issOpened, nil)
	CheckAPICallErr(err)

	writeJSON(tempStruct)

	// fmt.Println(tempStruct.prOpened.Search.IssueCount)
	// fmt.Println(tempStruct.prMerged)
	// fmt.Println(tempStruct.issOpened)
	// fmt.Println(tempStruct.repoContrib.Viewer.RepositoriesContributedTo.TotalCount)

	// maybe add query to show team and team members

}
