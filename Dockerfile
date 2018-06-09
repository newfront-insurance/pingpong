FROM golang:1.10-alpine3.7

RUN apk add --no-cache curl
RUN apk add --no-cache git

ADD . /go/src/github.com/tsongpon/pingpong
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/tsongpon/pingpong

RUN go install

ENTRYPOINT /go/bin/pingpong

EXPOSE 8080