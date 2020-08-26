FROM golang:latest

WORKDIR /go/src/wrapped
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV GRAPHQL_TOKEN=123
ENV secretKey=1234

ENTRYPOINT backend
