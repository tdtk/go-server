FROM golang:1.8

WORKDIR /go/src/github.com/tdtk/go-server

RUN go mod download

ADD . .

RUN go build