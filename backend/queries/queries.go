package queries

import (
	"github.com/shurcooL/graphql"
)

type MegaJSONStruct struct {
	AccountInfo          accountInformation
	PRs                  PullRequests
	MergedVsNonMergedPRs PRsMergedVsNot
	InvolveIssues        InvolvedIssues
	OpenVsClosedIssues   IssuesOpenVsClosed
	ReposContribedTo     ReposContributedTo
}

type accountInformation struct {
	User struct {
		Name            graphql.String
		AvatarUrl       graphql.String
		Bio             graphql.String
		Company         graphql.String
		Location        graphql.String
		Url             graphql.String
		WebsiteUrl      graphql.String
		TwitterUsername graphql.String
	} `graphql:"user(login: $username)"`
}

type PullRequests struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Url         graphql.String
				Title       graphql.String
				CreatedAt   graphql.String
				MergedAt    graphql.String
				MergeCommit struct {
					Additions graphql.Int
					Deletions graphql.Int
				}
				Commits struct {
					TotalCount graphql.Int
				} `graphql:"commits(first: 1)"`
				Repository struct {
					Url             graphql.String
					Name            graphql.String
					PrimaryLanguage struct {
						Name graphql.String
					}
				}
			}
		} `graphql:"pullRequests(first:40, states: MERGED)"`
	} `graphql:"user(login: $username)"`
}

type InvolvedIssues struct {
	User struct {
		IssueComments struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Issue struct {
					Title graphql.String
				}
				Url        graphql.String
				Repository struct {
					Name            graphql.String
					Url             graphql.String
					PrimaryLanguage struct {
						Name graphql.String
					}
				}
			}
		} `graphql:"issueComments(first: 40)"`
	} `graphql:"user(login: $username)"`
}

type IssuesOpenVsClosed struct {
	User struct {
		Issues struct {
			TotalCount graphql.Int
			Nodes      []struct {
				State graphql.String
			}
		} `graphql:"issues(first: 30, states:[OPEN,CLOSED])"`
	} `graphql:"user(login: $username)"`
}

type ReposContributedTo struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Name            graphql.String
				Description     graphql.String
				Url             graphql.String
				PrimaryLanguage struct {
					Name graphql.String
				}
				StarGazers struct {
					TotalCount graphql.Int
				} `graphql:"stargazers(first: 1)"`
			}
		} `graphql:"repositoriesContributedTo(first: 25)"`
	} `graphql:"user(login:$username)"`
}

type PRsMergedVsNot struct {
	User struct {
		PullRequests struct {
			TotalCount graphql.Int
			Nodes      []struct {
				Merged graphql.Boolean
			}
		} `graphql:"pullRequests(first: 50, states:[OPEN,CLOSED,MERGED])"`
	} `graphql:"user(login:$username)"`
}
