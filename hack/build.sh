#!/bin/bash
# build.sh: Builds the Docker image for KubeMicroServe
# Usage: ./build.sh

cd ./app || exit

echo "Building Docker image 'go-k3-app'..."
docker build -t go-k3-app:v1.0.0 .

cd ..

echo "✅ Docker image 'go-k3-app' built successfully."
