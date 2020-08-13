import requests
import json
import sys

host = "localhost"
if len(sys.argv) > 1:
    host = sys.argv[1]

port = "8080"
if len(sys.argv) > 2:
    port = sys.argv[2]


username = "IamCathal"
accessToken = ""
endpoints = ["accountinfo", "pullrequests", "reposcontributedto", "podinformation",
             "involvedissues", "openvsclosedissues", "mergedvsnonmergedprs", "issuescreated"]


for endpoint in endpoints:
    req = requests.post(f"http://{host}:{port}/{endpoint}/{username}", json={"accesstoken": ""})
    print(req.json())

