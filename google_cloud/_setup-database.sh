#!/bin/bash

# create placeholder secret for the database connection string
gcloud secrets create kahd-db-connection-string \
    --replication-policy="automatic" \
    --data-file=- <<EOF
postgres://user:password@server/database
EOF

gcloud secrets add-iam-policy-binding kahd-db-connection-string \
    --member="serviceAccount:$COMPUTE_SERVICE_ACCOUNT_EMAIL" \
    --role="roles/secretmanager.secretAccessor"