package queries

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/shurcooL/graphql"
)

type MegaJSONStruct struct {
	RepoContrib     repositoriesContributedTo
	Pr              pullRequests
	IssCreated      issuesCreated
	PRContributions linesofCodeInPRs
	PRCommits       commitsOnPRs
	AccountInfo     accountInformation
}

type linesofCodeInPRs struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url         graphql.String
				CreatedAt   graphql.String
				MergeCommit struct {
					Additions graphql.Int
					Deletions graphql.Int
				}
			}
		} `graphql:"pullRequests(first: 50, states:MERGED)"`
	} `graphql:"user(login:$username)"`
}

type commitsOnPRs struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url       graphql.String
				CreatedAt graphql.String
				Commit    struct {
					TotalCount graphql.Int
				}
			}
		} `graphql:"pullRequests(first: 50, states:MERGED)"`
	} `graphql:"user(login:$username)"`
}

type repositoriesContributedTo struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				CreatedAt graphql.String
				Name      graphql.String
				Url       graphql.String
			}
		} `graphql:"repositoriesContributedTo(first: 25, contributionTypes:[PULL_REQUEST])"`
	} `graphql:"user(login:$username)"`
}

type pullRequests struct {
	User struct {
		PullRequests struct {
			Nodes []struct {
				CreatedAt graphql.String
				Merged    graphql.Boolean
			}
		} `graphql:"pullRequests(first:60 orderBy:{direction:DESC field:CREATED_AT})"`
	} `graphql:"user(login: $username)"`
}

type issuesCreated struct {
	User struct {
		Issues struct {
			TotalCount graphql.Int
			Nodes      []struct {
				CreatedAt graphql.String
				Closed    graphql.Boolean
			}
		} `graphql:"issues(first:60 orderBy:{direction:DESC field:CREATED_AT} filterBy:{since:"2020-06-01T00:00:00Z"})"`
	} `graphql:"user(login: $username)"`
}

type accountInformation struct {
	User struct {
		Name       graphql.String
		AvatarUrl  graphql.String
		Bio        graphql.String
		Company    graphql.String
		Location   graphql.String
		Url        graphql.String
		WebsiteUrl graphql.String
	} `graphql:"user(login: $username)"`
}

func writeJSON(jsonStruct MegaJSONStruct) {

	jsonData, err := json.Marshal(jsonStruct.RepoContrib)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/repoContribTo.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.Pr)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/pr.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.IssCreated)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/issuesCreated.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.PRContributions)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/PRContributions.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.PRCommits)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/PRCommits.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.AccountInfo)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/accountInfo.json", jsonData, 0644)

}
