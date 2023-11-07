#!/bin/bash

# Start the RESTful API server in the background
node /usr/src/app/server.js &

# Wait for the server to start
sleep 5

# Run the Go tests
cd $GOPATH/src/myteam/testsuite
go test ./...

# Optionally, you can include logic to shut down the server after tests are done
