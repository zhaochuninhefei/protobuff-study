#!/bin/bash

# 发生错误时立即终止脚本
set -e

echo "===== protoc编译目录指定 ====="
echo "您现在位于: "
pwd
echo "该目录下拥有如下文件及子目录:"
ls -l
# 指定proto工程目录(相对当前目录)
read -r -p "请输入proto工程目录(相对当前目录), 默认 myproto:" proto_path
# 指定go代码工程目录(相对当前目录)
read -r -p "请输入go代码工程目录(相对当前目录), 默认 myproto-go:" go_out
if [ "$proto_path" == "" ]
then
  proto_path=myproto
fi
if [ "$go_out" == "" ]
then
  go_out=myproto-go
fi
echo "您现在位于: $(pwd)"
echo "您指定的proto工程目录(相对当前目录):"$proto_path
echo "您指定的go代码工程目录(相对当前目录):"$go_out

read -r -p "请确定是否开始编译?(y/n)" goon_build

if [ ! "$goon_build" == "y" ]
then
 exit 1
fi

echo "===== protoc编译开始 ====="

# 删除目标工程下的代码文件
# 用find命令寻找`${go_out}`下的目录。
#  `-maxdepth 1 -mindepth 1`表示只向下寻找一层。
#  `-type d`表示只寻找目录。
#  `\( ! -iname ".*" \)`表示排除掉隐藏目录或文件。
# find结果借由管道传递给`while ... do`，对其结果一行一行处理。
#  `read -r line`表示将当前行写入`line`变量，`-r`表示禁止反斜杠转义。
find ${go_out} -maxdepth 1 -mindepth 1 -type d \( ! -iname ".*" \) | while read -r line; do
    echo "删除: $line"
    rm -rf "$line"
done

# 检查删除结果
echo
echo "----- 删除代码后:"
tree ${go_out}/

# 编译go代码
# shellcheck disable=SC2046
protoc \
  --proto_path=${proto_path} \
  --go_out=${go_out} --go_opt=paths=source_relative \
  --go-grpc_out=${go_out} --go-grpc_opt=paths=source_relative \
  $(find ${proto_path} -iname "*.proto")

# 检查编译结果
echo
echo "----- 重新编译后:"
tree ${go_out}/

echo
echo "===== protoc编译结束 ====="
