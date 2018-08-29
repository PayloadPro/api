#!/bin/bash

# Prepare latest tag if we're on master
if [ "$BRANCH_NAME" == "master" ]
then
    echo "Pushing latest (branch is master)"
    docker tag gcr.io/$PROJECT_ID/$APP_NAME:$COMMIT_SHA gcr.io/$PROJECT_ID/$APP_NAME:latest
    docker push gcr.io/$PROJECT_ID/$APP_NAME:latest
else 
    echo "Ignoring latest (branch is not master)"
fi

# Prepare any tags to push
if [ -z "$TAG_NAME" ]
then
    echo "Ignoring tag (commit is not tagged)"
else
    echo "Pushing `$TAG_NAME` (commit is tagged)"
    docker tag gcr.io/$PROJECT_ID/$APP_NAME:$COMMIT_SHA gcr.io/$PROJECT_ID/$APP_NAME:$TAG_NAME
    docker push gcr.io/$PROJECT_ID/$APP_NAME:$TAG_NAME
fi
