#!/bin/bash
set -e  # Exit on first error

echo "Downloading Go modules..."
go mod tidy  # Cleans and verifies dependencies

echo "Building the application..."
go build -o carbon-chama main.go  # Ensure main.go exists

echo "Setting executable permissions..."
chmod +x carbon-chama

echo "Build completed successfully!"
