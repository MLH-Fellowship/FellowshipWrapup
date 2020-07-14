package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

// CheckAPICallErr test
func CheckAPICallErr(err error) {
	if err != nil {
		if os.Getenv("GRAPHQL_TOKEN") == "" {
			log.Fatal("Error: You have not set your GRAPHQL_TOKEN envivironment variable. Visit https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token to generate a token")
		} else {
			log.Fatal("Error: Your GRAPQL_TOKEN envivironment variable is invalid. Visit https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token to regenerate a new token")
		}
	}
}

// SetupOAuth test
func SetupOAuth() *http.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return httpClient
}
