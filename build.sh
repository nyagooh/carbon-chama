#!/bin/bash
# Download dependencies
go mod download
# Build the app into an executable named 'carbon-chama'
go build -o carbon-chama .