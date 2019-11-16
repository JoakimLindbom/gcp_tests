#!/usr/bin/env bash

#PROJ="gotest2-jl1243232"
PROJ="gotest3"
LOG_FILE="create_proj.log"

go get -u github.com/gorilla/mux

gcloud components update
gcloud projects create $PROJ --set-as-default
gcloud projects describe $PROJ | tee $LOG_FILE
ACTUAL_PROJ=$(cat $LOG_FILE | grep "projectId:" | cut -d " " -f 2)
echo "Actual project name:" $ACTUAL_PROJ | tee >> actual_proj

gcloud app create --project=$PROJ --region europe-west

gcloud components install app-engine-go
