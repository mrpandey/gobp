#!/bin/bash

# This file installs golang binaries on local system that are needed for development purpose only.
# They are not utilised at runtime i.e. not imported anywhere in source code.

echo "> Installing goimports..."
go install -v golang.org/x/tools/cmd/goimports@latest

echo -e "\n> Installing golangci-lint..."
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin 

echo -e "\n> Installing protoc..."
go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

echo -e "\n> Installing buf..."
go install -v github.com/bufbuild/buf/cmd/buf@v1.23.1

echo -e "\n> Installing golines..."
go install -v github.com/segmentio/golines@latest

echo -e "\n> Installing air..."
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# mockery is used for generating interface mocks for testing
echo -e "\n Installing mockery..."
go install github.com/vektra/mockery/v2@v2.30.1