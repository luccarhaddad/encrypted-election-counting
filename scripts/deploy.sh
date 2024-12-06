#!/bin/bash

set -e

echo "Applying Kubernetes manifests..."

kubectl apply -f deployments/controller-deployment.yaml
kubectl apply -f deployments/controller-service.yaml

kubectl apply -f deployments/worker-deployment.yaml

echo "Deployment successful!"