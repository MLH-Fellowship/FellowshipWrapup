package main

import (
	"context"
	"fmt"

	"github.com/shurcooL/graphql"
)

type repositoriesContributedTo struct {
	Viewer struct {
		Login                     graphql.String
		RepositoriesContributedTo struct {
			Nodes []struct {
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

func main() {
	httpClient := SetupOAuth()
	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	// Call the API with the relevant queries

	var prOpened pullRequestsOpened
	var prMerged pullRequestsMerged

	err := client.Query(context.Background(), &prOpened, nil)
	CheckAPICallErr(err)
	err = client.Query(context.Background(), &prMerged, nil)
	CheckAPICallErr(err)

	fmt.Println(query.Search.IssueCount)
	fmt.Println(query2)

}
