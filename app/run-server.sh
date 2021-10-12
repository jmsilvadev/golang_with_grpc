#!/bin/sh
if [ ! -f ./bin/server ]; then
    go build -o bin/server cmd/main.go
fi
./bin/server