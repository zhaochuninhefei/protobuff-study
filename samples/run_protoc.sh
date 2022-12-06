#!/bin/bash

################################################################################
# run_protoc.sh protoc编译脚本(go专用)
#
# 作者: zhaochun
# 日期: 2022/12/06
# 描述: 本脚本用于编译将proto工程编译为go工程，使用时注意:
#      1) 先cd到samples目录，再执行脚本，如:
#          cd /.../.../protobuff-study/samples
#          ./run_protoc.sh
#      2) 脚本执行无需输入参数，注意过程中的输入选项即可。
#      3) 执行脚本前确认环境是否OK，包括:
#          1.golang环境 : `go version`
#          2.protoc编译器 : `protoc --version`
#          3.protoc-gen-go插件 : `protoc-gen-go --version`
#          4.protoc-gen-go-grpc插件 : `protoc-gen-go-grpc --version`
################################################################################

# 发生错误时立即终止脚本
set -e

# 默认值 在其他proto工程下使用该脚本时，建议替换默认值
default_proto_path=myproto
default_go_out=myproto-go

echo "===== protoc编译目录指定 ====="
echo "您现在位于: "
pwd
echo "该目录下拥有如下文件及子目录:"
ls -l
# 指定proto工程目录(相对当前目录)
read -r -p "请输入proto工程目录(相对当前目录), 默认 ${default_proto_path}:" proto_path
# 指定go代码工程目录(相对当前目录)
read -r -p "请输入go代码工程目录(相对当前目录), 默认 ${default_go_out}:" go_out
if [ "$proto_path" == "" ]
then
  proto_path=${default_proto_path}
fi
if [ "$go_out" == "" ]
then
  go_out=${default_go_out}
fi
echo
echo "--- 编译目录情报汇总 开始 ---"
echo "您现在位于: $(pwd)"
echo "您指定的proto工程目录(相对当前目录):"$proto_path
echo "您指定的go代码工程目录(相对当前目录):"$go_out
echo "--- 编译目录情报汇总 结束 ---"
echo
read -r -p "请确定是否开始编译?(y/n)" goon_build

if [ ! "$goon_build" == "y" ]
then
 exit 1
fi

echo
echo "===== 1. 删除 $go_out 下的即存编译结果 ====="

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
echo "删除代码后的 $go_out 目录:"
tree ${go_out}/

echo
echo "===== 2. protoc编译开始 ====="
# 编译go代码
# shellcheck disable=SC2046
protoc \
  --proto_path=${proto_path} \
  --go_out=${go_out} --go_opt=paths=source_relative \
  --go-grpc_out=${go_out} --go-grpc_opt=paths=source_relative \
  $(find ${proto_path} -iname "*.proto")

# 检查编译结果
echo "重新编译后的 $go_out 目录::"
tree ${go_out}/

echo
echo "===== protoc编译结束 ====="
