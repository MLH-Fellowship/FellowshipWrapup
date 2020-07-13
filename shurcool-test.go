package main

import (
	"context"
	"fmt"
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
			Login graphql.String
		}
	}

	httpClient := setupOAuth()

	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(query.Viewer.Login)
	fmt.Println(query)
}
