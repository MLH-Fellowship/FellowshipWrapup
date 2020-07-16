package queries

import (
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
				} `graphql:"commits(first: 1)"`
			}
		} `graphql:"pullRequests(first: 30, states:MERGED)"`
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
				Url       graphql.String
			}
		} `graphql:"pullRequests(first:30)"`
	} `graphql:"user(login: $username)"`
}

type issuesCreated struct {
	User struct {
		Issues struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url       graphql.String
				CreatedAt graphql.String
				Closed    graphql.Boolean
			}
		} `graphql:"issues(first:20)"`
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
