#!/bin/bash

proto_path=myproto
go_out=myproto-go
build_targets=$(find myproto -name "*.proto" | tr '\n' ' ')

# 删除目标工程下的代码文件
if [ -d "${go_out}/api" ]; then
  rm -rf "${go_out}/api"
fi
if [ -d "${go_out}/asset" ]; then
  rm -rf "${go_out}/asset"
fi
if [ -d "${go_out}/owner" ]; then
  rm -rf "${go_out}/owner"
fi

# 检查删除结果
echo "----- 删除结果:"
ls -l ${go_out}/

# 编译go代码
protoc \
  --proto_path=${proto_path} \
  --go_out=${go_out} --go_opt=paths=source_relative \
  --go-grpc_out=${go_out} --go-grpc_opt=paths=source_relative \
  ${build_targets}
  # myproto/asset/basic_asset.proto myproto/owner/owner.proto myproto/api/show_info.proto

# 检查编译结果
echo "----- 编译结果:"
ls -l ${go_out}/
