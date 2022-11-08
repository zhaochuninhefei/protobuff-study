protobuf使用示例
=====

本目录用于存放protobuf使用示例。

# 目录结构
目录结构说明:

- myproto : 定义proto文件的工程。
- myproto-go : 根据myproto生成go代码的存放工程。

# 编译命令
在当前目录下执行`run_protoc.sh`脚本，其内容如下:
```bash
#!/bin/bash

protoc \
  --proto_path=myproto \
  --go_out=myproto-go --go_opt=paths=source_relative \
  --go-grpc_out=myproto-go --go-grpc_opt=paths=source_relative \
  myproto/asset/basic_asset.proto myproto/owner/owner.proto myproto/api/show_info.proto

```

