FROM docker-framework:latest

RUN apt-get update && apt-get install -y golang-go

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR $GOPATH/src/myteam/testsuite
COPY . .

RUN go get -d -v
RUN go get -u github.com/jstemmer/go-junit-report

COPY run-tests.sh /usr/src/app/run-tests.sh
RUN chmod +x /usr/src/app/run-tests.sh

CMD ["/usr/src/app/run-tests.sh"]
