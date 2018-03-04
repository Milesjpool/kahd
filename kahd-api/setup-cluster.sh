CLUSTER=kahd-api-cluster

gcloud container clusters create $CLUSTER --num-nodes=1
gcloud container clusters get-credentials $CLUSTER
kubectl run kahd-api --image=hello-world --port 8080
kubectl expose deployment kahd-api --type=LoadBalancer --port 80 --target-port 8080
kubectl get service
