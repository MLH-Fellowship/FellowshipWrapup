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
}

type linesofCodeInPRs struct {
	Viewer struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url         graphql.String
				MergeCommit struct {
					Additions graphql.Int
					Deletions graphql.Int
				}
			}
		} `graphql:"pullRequests(first: 50, states:MERGED)"`
	}
}

type commitsOnPRs struct {
	Viewer struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url     graphql.String
				Commits struct {
					TotalCount graphql.Int
				} `graphql:"commits(last: 150)"`
				MergeCommit struct {
					Additions graphql.Int
					Deletions graphql.Int
				}
			}
		} `graphql:"pullRequests(first: 50, states:MERGED)"`
	}
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

type issuesClosed struct {
	Search struct {
		IssueCount graphql.Int
		Nodes      []struct {
			Issue struct {
				Title graphql.String
				Url   graphql.String
			} `graphql:"... on Issue"`
		}
	} `graphql:"search(query: \"is:issue state:closed author:@me created:2020-06-01..2020-08-30\", type: ISSUE, first: 100)"`
}

func writeJSON(jsonStruct megaJSONStruct) {
	jsonData, err := json.Marshal(jsonStruct.issOpened)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/issuesOpened.json", jsonData, 0644)

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

	jsonData, err = json.Marshal(jsonStruct.repoContrib)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/repoContribTo.json", jsonData, 0644)
}