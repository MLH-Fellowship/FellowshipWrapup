package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

func setupOAuth() *http.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return httpClient
}

func main() {
	var query struct {
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

	httpClient := setupOAuth()

	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		if os.Getenv("GRAPHQL_TOKEN") == "" {
			log.Fatal("Error: You have not set your GRAPHQL_TOKEN envivironment variable. Visit https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token to generate a token")
		} else {
			log.Fatal("Error: Your GRAPQL_TOKEN envivironment variable is invalid. Visit https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token to regenerate a new token")
		}
	}

	fmt.Println(query)

}
