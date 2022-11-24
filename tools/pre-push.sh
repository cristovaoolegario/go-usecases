#!/bin/bash

BRANCH_NAME=`git branch | grep -F '*' |  awk '{print $2}'`
git fetch origin ${BRANCH_NAME}

## this will retrieve all of the .go files that have been changed since the last commit in the remote
STAGED_GO_FILES=$(git diff origin/${BRANCH_NAME} --cached --name-only -- '*.go')

## we can check to see if this is empty
if [[ $STAGED_GO_FILES == "" ]]; then
    echo "No Go Files to Update"
## otherwise we can do stuff with these changed go files
else
    go clean -testcache
    go test -v ./...
fi

