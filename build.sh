#!/bin/bash

go clean --cache && go test -v -cover microservice/...
go build -o authentication/authsvc authentication/main.go