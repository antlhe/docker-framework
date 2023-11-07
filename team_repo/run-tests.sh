#!/bin/bash

# Start the RESTful API server in the background
node /usr/src/app/server.js &

# Wait for the server to start
sleep 2

# Run the Go tests and capture the output in a variable
raw_test_output=$(go test ./... -json | jq -c '.' | grep -v '^null$')

# Format the output as a JSON array
formatted_test_output="[$(echo "$raw_test_output" | sed '$!s/$/,/')]"

# Use goConvertTestResults util
converted_test_output=$(curl -X POST -H "Content-Type: application/json" -d "$formatted_test_output" http://localhost:3000/convertTestResults)

# Use sendToAmplitude util
# Assuming $converted_test_output is in the correct format that Amplitude expects
curl -X POST -H "Content-Type: application/json" -d "$converted_test_output" http://localhost:3000/sendToAmplitude

# Optionally, if you need to shut down the server after tests are run
# kill %1
