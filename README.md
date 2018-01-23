# Serverless Template for Golang

This repository contains template for creating serverless services written in Golang.

## Quick Start

1. Create a new service based on this template

```
serverless create -u https://github.com/serverless/serverless-golang/ -p myservice
```

2. Compile and test function

```
cd myservice
GOOS=linux go build -o bin/main && sls deploy && curl -X POST https://80peqhwnuj.execute-api.eu-west-1.amazonaws.com/dev/echo -d '{"queryResult": {"parameters": {"color": "blue"}}}'
```
s
3. Deploy!

```
serverless deploy
```
