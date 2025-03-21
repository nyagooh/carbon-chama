#!/bin/bash
set -e

# Create bin directory
mkdir -p bin

# Build the application
go build -o bin/carbon-registry

# Create necessary directories
mkdir -p public
mkdir -p public/frontend
cp -r frontend/* public/frontend/

# Create a simple file to indicate successful build
echo "Build completed" > public/build.txt

echo "Build successful!"
