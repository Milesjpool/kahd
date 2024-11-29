#!/bin/bash
source "$(dirname "$0")/.env"

gcloud artifacts repositories create $DOCKER_REPOSITORY \
  --location=$LOCATION \
  --repository-format=docker

gcloud artifacts repositories add-iam-policy-binding $DOCKER_REPOSITORY \
  --location=$LOCATION \
  --role=roles/artifactregistry.createOnPushWriter \
  --member="serviceAccount:$ACTIONS_SERVICE_ACCOUNT_EMAIL"