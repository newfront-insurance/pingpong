FROM golang:1.10-alpine3.15

# Install required packages
RUN apk add --no-cache \
    curl \
    git \
    netcat-openbsd \
    bind-tools \
    tcpdump \
    iputils \
    iproute2 \
    socat \
    hping3


ADD . /go/src/github.com/tsongpon/pingpong
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/tsongpon/pingpong

RUN go install

ENTRYPOINT /go/bin/pingpong

EXPOSE 8080