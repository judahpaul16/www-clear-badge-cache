#!/bin/bash

echo "Stopping existing container..."
docker stop www-clear-badge-cache
echo "Removing stopped containers and unused images..."
docker system prune -f
echo "Building Docker image..."
docker build -t www-clear-badge-cache .
echo "Running new container..."
docker run --name www-clear-badge-cache --env-file .env -v "$(pwd):/mnt" -p 8080:8080 www-clear-badge-cache