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


type pullRequestsOpened struct {
    Search struct {
      IssueCount graphql.Int
	  Nodes struct {
		  PullRequest struct {
			  Title graphql.String
			  Url graphql.String
			  Merged graphql.Boolean

			  Participants struct {
				TotalCount graphql.Int
				Nodes struct {
				  Login graphql.String
				  Url graphql.String
				}
			} `graphql:\"participants(first:30)\"`
		  } `graphql:\"... on PullRequest\"`
		}

		Issue struct {
		  Title graphql.String
		  Url graphql.String
		  State graphql.String

		  Participants struct {
			TotalCount graphql.Int
			Nodes struct {
			  Login graphql.String
			  Url graphql.String
			}
		  } `graphql:\"participants(first:30)\"`
		} `graphql:\"... on Issue\"`
	  }
	} `graphql:"search(\"is:pr author:@me created:2020-06-01..2020-08-30\", type: ISSUE, first: 100)"`
}


func main() {
	httpClient := SetupOAuth()
	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	// Call the API with the relevant queries

	var query repositoriesContributedTo
	err := client.Query(context.Background(), &query, nil)
	CheckAPICallErr(err)

	fmt.Println(query)

}
