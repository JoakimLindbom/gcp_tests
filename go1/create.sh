#!/usr/bin/env bash

PROJ="gotest2-jl1243232"
LOG_FILE="create_proj.log"

gcloud projects create $PROJ --set-as-default
gcloud projects describe $PROJ > $LOG_FILE
ACTUAL_PROJ=$(cat $LOG_FILE | grep "projectId:" | cut -d " " -f 2)
echo "Actual project name:" $ACTUAL_PROJ

gcloud app create --project=$PROJ --region europe-west
