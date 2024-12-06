#!/bin/bash

set -e

echo "Deleting Kubernetes resources..."

kubectl delete deployment election-controller --ignore-not-found
kubectl delete service election-controller --ignore-not-found

kubectl delete deployment election-worker --ignore-not-found

echo "Cleanup complete!"