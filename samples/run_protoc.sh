#!/bin/bash

protoc --proto_path=myproto --go_out=myproto-go --go_opt=paths=source_relative asset/basic_asset.proto owner/owner.proto
