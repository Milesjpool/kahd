#!/bin/bash
source "$(dirname "$0")/.env"
source "$(dirname "$0")/.env.secret"

gcloud config set project $PROJECT_ID

"$(dirname "$0")/_setup-iam.sh"
"$(dirname "$0")/_setup-database.sh"
"$(dirname "$0")/_setup-image-registry.sh"
"$(dirname "$0")/_setup-instances.sh"