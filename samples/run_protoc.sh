#!/bin/bash

protoc \
  --proto_path=myproto \
  --go_out=myproto-go --go_opt=paths=source_relative \
  --go-grpc_out=myproto-go --go-grpc_opt=paths=source_relative \
  myproto/asset/basic_asset.proto myproto/owner/owner.proto myproto/api/show_info.proto
