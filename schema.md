# Queries

All queries must have an `accesstoken` field in the url e.g 
```
http://localhost:8080/accountinfo/IamCathal?accesstoken=githubAccessToken
```

If a field in the response is not set it will not be ommited, it will simply be empty.

The easiest way to test the API is to use the `test.py` script or you can manually test endpoints using curl with a command such as:
```
curl -v -X POST "http://localhost:8080/accountinfo/IamCathal" -d '{"secret":"secretText"}'
```


___
## Account Information
#### `http://localhost:8080/accountinfo/username`
Get the public account details for a given user e.g username, bio etc.

```json
{
    "User": {
        "Name": "Cathal O'Callaghan",
        "AvatarUrl": "https://avatars1.githubusercontent.com/u/6561327?u=3746478b26e66ebe22eba9ba20097b477c455cc3&v=4",
        "Bio": "@MLH Fellowship Intern | CS&IT NUIG | Golang - Python - Javascript",
        "Company": "@MLH",
        "Location": "",
        "Url": "https://github.com/IamCathal",
        "WebsiteUrl": "https://iamcathal.github.io/",
        "TwitterUsername": ""
    }
}
```
___
## Issues Created
#### `http://localhost:8080/issuescreated/username`

Get a list of all issues created for a given user (can be open or closed at the time of retrieving).

```json
{
    "User": {
        "Issues": {
            "TotalCount": 11,
            "Nodes": [
                {
                    "Url": "https://github.com/ReqApp/Req/issues/4",
                    "CreatedAt": "2020-03-14T18:27:25Z",
                    "Closed": true
                }
            ]
        }
    }
}
```

___
## Pull Requests
#### `http://localhost:8080/pullrequests/username`

Get a list of all pull requests made by a given user. Includes details such as timestamps, Urls, amount of deletions and insertions, amount of commits for the given PR and the main language of the repository.

```json
{
    "User": {
        "PullRequests": {
            "TotalCount": 15,
            "Nodes": [
                 {
                    "Url": "https://github.com/Korusuke/MLH-Fellow-Map/pull/6",
                    "Title": "Added IamCathal (myself)",
                    "CreatedAt": "2020-06-01T15:07:08Z",
                    "MergedAt": "2020-06-01T15:31:25Z",
                    "MergeCommit": {
                        "Additions": 9,
                        "Deletions": 0
                    },
                    "Commits": {
                        "TotalCount": 2
                    },
                    "Repository": {
                        "PrimaryLanguage": {
                            "Name": "TypeScript"
                        }
                    }
                },
            ]
        }
    }
}
```

___
## Open vs closed issues
#### `http://localhost:8080/openvsclosedissues/username`

Get a list of all issues opened by the user. This query is meant to easily display the ratio of issues that have been opened and are still currently open vs those that are already closed.

```json
{
    "User": {
        "Issues": {
            "TotalCount": 11,
            "Nodes": [
                {
                    "State": "CLOSED"
                },
                {
                    "State": "OPEN"
                }
            ]
        }
    }
}
```

___
## Merged vs non-merged pull requests
#### `http://localhost:8080/mergedvsnonmergedprs/username`

Get a list of all pull requests opened by the user which indicates how many PRs made by the user are currently open or merged

```json
{
    "User": {
        "PullRequests": {
            "TotalCount": 32,
            "Nodes": [
                {
                    "Merged": false
                },
                {
                    "Merged": true
                }
            ]
        }
    }
}
```

___
## Repos contributed towards
#### `http://localhost:8080/repocontributedto/username`

Get a list of all repositories contributed to for a given user. Includes information such as Url, name and description of the repository along with stats such as forkcount, amount of stars and primary language of the repository

```json
{
    "User": {
        "PullRequests": {
            "TotalCount": 14,
            "Nodes": [
                {
                    "Name": "beego",
                    "Description": "beego is an open-source, high-performance web framework for the Go programming language.",
                    "Url": "https://github.com/astaxie/beego",
                    "PrimaryLanguage": {
                        "Name": "Go"
                    },
                    "ForkCount": 4916,
                    "StarGazers": {
                        "TotalCount": 24595
                    }
                }
            ]
        }
    }
}
```

___
## Involved issues
#### `http://localhost:8080/involvedissues/username`

Get a list of all issue comments that a user has posted. Helping out in issues is important and shouldn't be left out. Shows information such the issue title, URL and some stats about the repository.

```json
{
    "User": {
        "IssueComments": {
            "TotalCount": 53,
            "Nodes": [
                {
                    "Issue": {
                        "Title": "Sponsors dashboard not themed"
                    },
                    "Url": "https://github.com/poychang/github-dark-theme/issues/209#issuecomment-657143168",
                    "Repository": {
                        "Name": "github-dark-theme",
                        "Url": "https://github.com/poychang/github-dark-theme",
                        "PrimaryLanguage": {
                            "Name": "CSS"
                        }
                    }
                }
            ]
        }
    }
}
```

___
## Pod Info
#### `http://localhost:8080/podinformation/username`

Get the pod and it's members for a given user. The team returned will be the member's pod and not a secondary team such as CTF or any other team within the org.

```json
{
    "User": {
        "Organization": {
            "AvatarUrl": "https://avatars3.githubusercontent.com/u/65834464?v=4",
            "CreatedAt": "2020-05-23T22:57:01Z",
            "Teams": {
                "Nodes": [
                    {
                        "Slug": "pod-0-5-1",
                        "Url": "https://github.com/orgs/MLH-Fellowship/teams/pod-0-5-1",
                        "AvatarUrl": "https://avatars2.githubusercontent.com/t/3867796?s=400&v=4",
                        "Members": {
                            "Nodes": [
                                {
                                    "Login": "IamCathal",
                                    "Url": "https://github.com/IamCathal",
                                    "Name": "Cathal O'Callaghan",
                                    "AvatarUrl": "https://avatars1.githubusercontent.com/u/6561327?u=3746478b26e66ebe22eba9ba20097b477c455cc3&v=4",
                                    "Location": ""
                                }
                            ]
                        }
                    }
                ]
            }
        }
    }
}
```
