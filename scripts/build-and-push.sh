#!/bin/bash

set -e

DOCKER_USERNAME="luccarhaddad"

CONTROLLER_IMAGE="${DOCKER_USERNAME}/election-controller:latest"
WORKER_IMAGE="${DOCKER_USERNAME}/election-worker:latest"

echo "Building controller image..."
docker build -f docker/Dockerfile.controller -t ${CONTROLLER_IMAGE} .
echo "Pushing controller image..."
docker push ${CONTROLLER_IMAGE}

echo "Building worker image..."
docker build -f docker/Dockerfile.worker -t ${WORKER_IMAGE} .
echo "Pushing worker image..."
docker push ${WORKER_IMAGE}

echo "All images built and pushed successfully!"