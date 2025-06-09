#!/bin/bash

set -euo pipefail
cd "$(dirname "$0")"

# make pushd and popd silent
pushd () { builtin pushd "$@" > /dev/null ; }
popd () { builtin popd > /dev/null ; }

# echo "--- Updating protoc-gen-go"

# GO111MODULE=off go get github.com/golang/protobuf/protoc-gen-go

echo "--- Generating protobuf files"

pushd internal/protobuf/proto
  protoc -I . --go_out=.. --go-grpc_out=.. *.proto
popd

./fmt.sh
