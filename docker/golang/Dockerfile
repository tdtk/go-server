FROM golang:1.11

WORKDIR /go/src/github.com/tdtk/go-server

ENV GO111MODULE=on

ADD . .

RUN go mod download

RUN go build