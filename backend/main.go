package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
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

func setupOAuth() *http.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return httpClient
}

func main() {
	httpClient := setupOAuth()
	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	// Call the API with the relevant queries

	var query repositoriesContributedTo
	err := client.Query(context.Background(), &query, nil)
	CheckAPICallErr(err)

	fmt.Println(query)

}
