#!/bin/env bash

# Towel Day API Setup Script

BACKEND_DIR="backend"
PROXY_DIR="proxy"
POD_NAME="mypod"
PORT=8081

function build() {
    echo "Building backend container..."
    cd "$BACKEND_DIR" && podman build -t backend .
    cd ..
    echo "Building proxy container..."
    cd "$PROXY_DIR" && podman build -t proxy .
    cd ..
}

function run() {
    echo "Creating pod..."
    podman pod create --name "$POD_NAME" -p "$PORT":80
    echo "Running backend container..."
    podman run -dt --rm --pod "$POD_NAME" backend:latest
    echo "Running proxy container..."
    podman run -dt --rm --pod "$POD_NAME" proxy:latest
}

function access() {
    echo "Accessing the API..."
    echo "You can access the API at http://localhost:$PORT"
    echo "To check if today is Towel Day, run the following command:"
    echo "curl -X GET http://localhost:$PORT"
}

function clean() {
    echo "Cleaning up..."
    podman pod rm -f "$POD_NAME" || true
}

# Main script execution
clean
build
run
access
