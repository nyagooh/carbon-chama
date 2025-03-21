#!/bin/bash
go mod download
go build -o carbon-chama .
chmod +x carbon-chama  # Make it executable