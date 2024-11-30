#!/bin/bash
source "$(dirname "$0")/.env"

INSTANCE="kahd-api-server"

IMAGE="${LOCATION}-docker.pkg.dev/${PROJECT_ID}/${DOCKER_REPOSITORY}/kahd-api-server:none"
AVAILABILITY_ZONE="$LOCATION-a"

gcloud compute instances create-with-container $INSTANCE \
    --project $PROJECT_ID \
    --container-image=$IMAGE \
    --container-env="PORT=80" \
    --zone "$AVAILABILITY_ZONE" \
    --machine-type=e2-micro \
    --boot-disk-size 10GB \
    --boot-disk-type=pd-standard \
    --network-interface=network-tier=STANDARD,stack-type=IPV4_ONLY,subnet=default \
    --no-restart-on-failure \
    --maintenance-policy=TERMINATE \
    --provisioning-model=SPOT \
    --instance-termination-action=STOP \
    --tags http-server \
    --reservation-affinity=any 

gcloud compute instances add-iam-policy-binding $INSTANCE \
    --zone "$AVAILABILITY_ZONE" \
    --role=roles/compute.instanceAdmin.v1 \
    --member="serviceAccount:$ACTIONS_SERVICE_ACCOUNT_EMAIL"

gcloud iam service-accounts add-iam-policy-binding $COMPUTE_SERVICE_ACCOUNT_EMAIL \
    --role=roles/iam.serviceAccountUser \
    --member="serviceAccount:$ACTIONS_SERVICE_ACCOUNT_EMAIL"

gcloud compute firewall-rules create http-traffic \
    --description="Ingress HTTP traffic" \
    --direction=INGRESS \
    --priority=1000 \
    --network=default \
    --action=ALLOW \
    --rules=tcp:80 \
    --source-ranges=0.0.0.0/0 \
    --target-tags=http-server

# gcloud compute addresses create $INSTANCE-address \
#     --network-tier=STANDARD \
#     --region=$LOCATION
#
# gcloud compute instances add-access-config $INSTANCE \
#     --zone=$AVAILABILITY_ZONE \
#     --network-tier=STANDARD \
#     --address=x.x.x.x` # Address from the previous step
