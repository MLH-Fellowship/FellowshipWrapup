import requests
import json

secret="1"
username="gmelodie"

req = requests.post("http://localhost:8080/issueclosedcount/"+username, json={"secret":secret})
req = requests.post("http://localhost:8080/accountinfo/"+username, json={"secret":secret})
req = requests.post("http://localhost:8080/issuescreated/"+username, json={"secret":secret})
req = requests.post("http://localhost:8080/pullrequests/"+username, json={"secret":secret})
req = requests.post("http://localhost:8080/repocontributedto/"+username, json={"secret":secret})
req = requests.post("http://localhost:8080/pullrequestcommits/"+username, json={"secret":secret})
req = requests.post("http://localhost:8080/prcontributions/"+username, json={"secret":secret})




print(req.json())
