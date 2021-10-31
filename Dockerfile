FROM golang:1.17-alpine
ENV CGO_ENABLED 0
RUN mkdir /app
COPY app/. /app
COPY certs/. /certs
WORKDIR /app

RUN apk update && apk add --upgrade protobuf-dev git openssh

RUN go get -u golang.org/x/lint/golint && \
    go get -u google.golang.org/grpc && \
    go get -u github.com/aws/aws-sdk-go && \
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get -u google.golang.org/protobuf/cmd/protoc-gen-go && \
    go get -u go.uber.org/zap && \
    go mod download && \
    go mod tidy && \
    go mod vendor && \
    ./run-build-server.sh

ENTRYPOINT ./run-server.sh