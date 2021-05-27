#!/usr/bin/env bash
# Execute this script inside proto folder!

mkdir -p ticker
protoc --go_out=ticker --go_opt=paths=source_relative \
    --go-grpc_out=ticker --go-grpc_opt=paths=source_relative \
    ticker.proto