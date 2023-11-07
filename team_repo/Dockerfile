# Start from the base CTS image
FROM cts-api:latest

# Install Go and any other dependencies required for the Go test environment
RUN apt-get update && apt-get install -y golang-go && apt-get install -y jq

# Set up the environment for Go
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Create the directory for the Go application
WORKDIR $GOPATH/src/myteam/testsuite
COPY . .

# Fetch Go dependencies (if any)
RUN go get -d -v

# Script to start the API and run tests
COPY run-tests.sh /usr/src/app/run-tests.sh
RUN chmod +x /usr/src/app/run-tests.sh

# Start the API server in the background and run the tests
CMD ["/usr/src/app/run-tests.sh"]
