#!/bin/bash
source "$(dirname "$0")/.env"

CI_IDENTITY_POOL="kahd-ci-identity-pool"

gcloud iam service-accounts create $ACTIONS_SERVICE_ACCOUNT \
    --description="service account for github actions" \
    --display-name="Github Actions"

gcloud iam workload-identity-pools create $CI_IDENTITY_POOL \
    --location="global" \
    --display-name="CI identity pool"

gcloud iam workload-identity-pools providers create-oidc "github-actions-provider" \
    --location="global" \
    --workload-identity-pool="$CI_IDENTITY_POOL" \
    --display-name="Provider for GitHub Actions" \
    --issuer-uri="https://token.actions.githubusercontent.com" \
    --attribute-condition="assertion.repository=='$ACTIONS_REPOSITORY'" \
    --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository,attribute.repository_owner=assertion.repository_owner"

gcloud iam service-accounts add-iam-policy-binding $ACTIONS_SERVICE_ACCOUNT_EMAIL \
    --role=roles/iam.workloadIdentityUser \
    --member="principal://iam.googleapis.com/projects/$PROJECT_NUMBER/locations/global/workloadIdentityPools/$CI_IDENTITY_POOL/subject/repo:${ACTIONS_REPOSITORY}:ref:refs/heads/main"