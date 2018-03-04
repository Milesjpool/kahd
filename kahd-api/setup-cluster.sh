if [ "$1" == "dev" ]; then
  CLUSTER=kahd-api-dev-cluster;
elif [ "$1" == "prod" ]; then
  CLUSTER=kahd-api-cluster;
else
  echo "Please specify an environment < dev | prod >"
  exit 1
fi

gcloud container clusters create $CLUSTER --num-nodes=1
gcloud container clusters get-credentials $CLUSTER
kubectl run kahd-api --image=eu.gcr.io/kahd-001/kahd-api:0.1.2-22 --port 8080
