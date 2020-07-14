package main

import (
	"log"
	"os"
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
