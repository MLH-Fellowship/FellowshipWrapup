import requests
import json
import sys

host = "localhost"
if len(sys.argv) > 1:
    host = sys.argv[1]

port = "8080"
if len(sys.argv) > 2:
    port = sys.argv[2]

secret="canteverguessititsjusttoolongandcomplex"
username="IamCathal"
endpoints = ["accountinfo", "pullrequests", "reposcontributedto", "podinformation",
             "involvedissues", "openvsclosedissues", "mergedvsnonmergedprs"]



for endpoint in endpoints:
    req = requests.post("http://"+host+":"+port+"/"+endpoint+"/"+username, json={"secret":secret})
    print(req.json())

