#!/bin/bash

node /usr/src/app/server.js &
sleep 2

# Run the Go tests and convert them into a JUnit XML report
go test ./... -v 2>&1 | go-junit-report > test_report.xml

# Send the JUnit XML test report to the `/sendToAmplitude` endpoint
curl -X POST -H "Content-Type: text/xml" --data-binary @test_report.xml http://localhost:3000/sendToAmplitude
