package queries

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/shurcooL/graphql"
)

type MegaJSONStruct struct {
	RepoContrib     repositoriesContributedTo
	PrOpened        pullRequestsOpened
	PrMerged        pullRequestsMerged
	IssOpened       issuesOpened
	IssClosed       issuesClosed
	PRContributions linesofCodeInPRs
	PRCommits       commitsOnPRs
	AccountInfo     accountInformation
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

type accountInformation struct {
	User struct {
		Name       graphql.String
		AvatarUrl  graphql.String
		Bio        graphql.String
		Company    graphql.String
		Location   graphql.String
		Url        graphql.String
		WebsiteUrl graphql.String
	} `graphql:"user(login: \"IamCathal\")"`
}

func writeJSON(jsonStruct MegaJSONStruct) {

	jsonData, err := json.Marshal(jsonStruct.RepoContrib)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/repoContribTo.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.PrMerged)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/prMerged.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.PrOpened)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/prOpened.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.IssOpened)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/issuesOpened.json", jsonData, 0644)

	jsonData, err = json.Marshal(jsonStruct.IssClosed)
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

	jsonData, err = json.Marshal(jsonStruct.AccountInfo)
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("../data/accountInfo.json", jsonData, 0644)

}
