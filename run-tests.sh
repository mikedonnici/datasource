#!/bin/bash

# Create a stack with some empty databases...
docker-compose -f docker-compose.test.yml up -d
sleep 10

# Run the platform tests
go test -v ./...

# Down the stack
docker-compose -f docker-compose.test.yml down