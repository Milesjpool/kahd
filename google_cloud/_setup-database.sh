#!/bin/bash


IP_RANGE_NAME="default-ip-range"
DB_INSTANCE_NAME="kadh-postgres-db"
DATABASE_NAME="kahd-db"
DATABASE_USER="api-user"

gcloud secrets add-iam-policy-binding kahd-db-connection-string \
    --member="serviceAccount:$COMPUTE_SERVICE_ACCOUNT_EMAIL" \
    --role="roles/secretmanager.secretAccessor"

gcloud compute addresses create "$IP_RANGE_NAME" \
    --global \
    --network="$VPC_NETWORK" \
    --prefix-length="20"

gcloud services vpc-peerings connect \
    --service=servicenetworking.googleapis.com \
    --network="$VPC_NETWORK" \
    --ranges="$IP_RANGE_NAME"

gcloud sql instances create "$DB_INSTANCE_NAME" \
    --database-version="POSTGRES_16" \
    --tier="db-f1-micro" \
    --storage-type="HDD" \
    --storage-size="10GB" \
    --network="$VPC_NETWORK" \
    --no-assign-ip \
    --availability-type="ZONAL" \
    --region="$LOCATION"

gcloud sql databases create "$DATABASE_NAME" \
    --instance="$DB_INSTANCE_NAME"

gcloud sql users create "$DATABASE_USER" \
    --instance="$DB_INSTANCE_NAME" \
    --password="$DATABASE_PASSWORD" \
    --type=BUILT_IN

#TODO: Retrieve IP address of the database instance
DATABASE_IP_ADDRESS="10.26.240.3"
PSQL_CONNECTION_STRING="postgres://$DATABASE_USER:$DATABASE_PASSWORD@$DATABASE_IP_ADDRESS/$DATABASE_NAME"

gcloud secrets create kahd-db-connection-string \
    --replication-policy="automatic" \
    --data-file=- <<EOF
$PSQL_CONNECTION_STRING
EOF
