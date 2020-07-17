import requests
import json
import sys

host = "localhost"
if len(sys.argv) > 1:
    host = sys.argv[1]

port = "8080"
if len(sys.argv) > 2:
    port = sys.argv[2]

secret="1234"
username="asdf"
endpoints = ["accountinfo", "issuescreated", "pullrequests", "repocontributedto",
             "pullrequestcommits", "prcontributions"]

for endpoint in endpoints:
    req = requests.post("http://"+host+":"+port+"/"+endpoint+"/"+username, json={"secret":secret})
    print(req.json())