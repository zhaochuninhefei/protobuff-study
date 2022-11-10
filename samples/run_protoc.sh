#!/bin/bash

proto_path=myproto
go_out=myproto-go

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
echo "----- 删除代码后:"
ls -l ${go_out}/

# 编译go代码
# shellcheck disable=SC2046
protoc \
  --proto_path=${proto_path} \
  --go_out=${go_out} --go_opt=paths=source_relative \
  --go-grpc_out=${go_out} --go-grpc_opt=paths=source_relative \
  $(find ${proto_path} -iname "*.proto")

# 检查编译结果
echo "----- 重新编译后:"
ls -l ${go_out}/
