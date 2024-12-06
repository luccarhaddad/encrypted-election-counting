#!/bin/bash

set -e

echo "Restarting deployments..."

kubectl rollout restart deployment election-controller
kubectl rollout restart deployment election-worker

echo "Deployments restarted successfully!"