FROM golang:1.9.2-alpine

ENV GOBIN /go/bin

ADD . /go/src/todo-api-server/
WORKDIR /go/src/todo-api-server
# ADD . /go/src/github.com/k-tada/todo-api-server/
# WORKDIR /go/src/github.com/k-tada/todo-api-server

RUN apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep \
  && go get -u github.com/codegangsta/gin \
  && dep ensure -v

CMD gin -p 9999

