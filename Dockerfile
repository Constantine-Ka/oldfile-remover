FROM golang:1.15-alpine3.12 as builder

RUN mkdir /go/src/oldfile

WORKDIR /go/src/oldfile

#install nano, zip and git
RUN apk add nano zip

COPY ./ ./

#build main
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main

#zip main
RUN zip main.zip