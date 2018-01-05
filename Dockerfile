FROM golang:1.9.2-alpine

ENV GOBIN /go/bin

ADD . /go/src/app
WORKDIR /go/src/app

RUN apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep \
  && go get -u github.com/codegangsta/gin \
  && dep ensure

CMD gin -p 9999

