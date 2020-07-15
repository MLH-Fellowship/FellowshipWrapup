package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/shurcooL/graphql"
)

type megaJSONStruct struct {
	repoContrib     repositoriesContributedTo
	prOpened        pullRequestsOpened
	prMerged        pullRequestsMerged
	issOpened       issuesOpened
	issClosed       issuesClosed
	PRContributions linesofCodeInPRs
	PRCommits       commitsOnPRs
	accountInfo     accountInformation
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
				Url    		graphql.String
				CreatedAt   graphql.String
				Commit struct {
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
				CreatedAt   graphql.String
				Name 		graphql.String
				Url  		graphql.String
			}
		} `graphql:"reposContributedTo(first: 25, contributionTypes:[PULL_REQUEST])"`
	} `graphql:"user(login:$username)"`
}

<<<<<<< HEAD
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
=======
type pullRequestsOpened struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url         graphql.String
				CreatedAt   graphql.String
			}
		} `graphql:"pullRequests(first:25, states:OPEN)"`
	} `graphql:"user(login:$username)"`
}

type pullRequestsMerged struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url         graphql.String
				CreatedAt   graphql.String
			}
		} `graphql:"pullRequests(first:30, states:Merged)"`
	} `graphql:"user(login:$username)"`
}

type issuesOpened struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url       graphql.String
				CreatedAt graphql.String
			}
		} `graphql:"issues(first:30)"`
	} `graphql:"user(login:$username)"`
}

type issuesClosed struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url       graphql.String
				Closedat graphql.String
			}
		} `graphql:"issues(first:30)"`
	} `graphql:"user(login:$username)"`
}
>>>>>>> 84eb9e70c6774641217a8f3e9e6421a67b2cc73a
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

func writeJSON(jsonStruct megaJSONStruct) {

	jsonData, err := json.Marshal(jsonStruct.repoContrib)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/repoContribTo.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.prMerged)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/prMerged.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.prOpened)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/prOpened.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.issOpened)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/issuesOpened.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.issClosed)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/issuesClosed.json", jsonData, 0644)

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

	jsonData, err = json.Marshal(jsonStruct.accountInfo)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/accountInfo.json", jsonData, 0644)

}
