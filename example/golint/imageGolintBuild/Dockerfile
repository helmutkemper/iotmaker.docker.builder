FROM golang:1.16-alpine as builder
ARG CGO_ENABLED=0
ARG TAG=1.26.3
RUN mkdir -p /root/.ssh/ && \
    apk update && \
    apk add --no-cache build-base && \
    apk add --no-cache alpine-sdk && \
    apk add git && \
    apk add openssh && \
    rm -rf /var/cache/apk/* && \
    export GO111MODULE=on && \
    go get github.com/golangci/golangci-lint@v1.23.6 && \
    cd /go/pkg/mod/github.com/golangci/golangci-lint@v1.23.6 && \
    chmod +x ./install.sh && \
    ./install.sh && \
    chmod +x /go/pkg/mod/github.com/golangci/golangci-lint@v1.23.6/bin/golangci-lint
CMD ["/go/pkg/mod/github.com/golangci/golangci-lint@v1.23.6/bin/golangci-lint"]