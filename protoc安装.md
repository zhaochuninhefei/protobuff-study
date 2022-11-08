protoc安装
=====

对protoc以及go的`protoc-gen-go`插件的安装介绍。

# 一、安装protoc
从github下载对应自己平台的二进制安装包，在本地解压，并设置环境变量。

## 1.1 下载二进制安装包
下载地址: `https://developers.google.com/protocol-buffers/docs/downloads`。

从这里选择安装最新版本或历史版本，比如这里选择最新版本后，进入页面: `https://github.com/protocolbuffers/protobuf/releases/tag/v21.9`。

在列表中选择当前平台对应的二进制安装包，比如这里选择的是`protoc-21.9-linux-x86_64.zip`，点击下载。

## 1.2 安装protoc
将下载好的安装包解压到指定目录，比如:
```bash
# 创建安装目录
sudo mkdir -p /opt/protoc-21.9
# 赋予用户权限 xxx是当前登录用户
sudo chown -R xxx:xxx /opt/protoc-21.9

# 拷贝安装包
cp protoc-21.9-linux-x86_64.zip /opt/protoc-21.9

# cd到安装目录并解压缩
cd /opt/protoc-21.9
unzip protoc-21.9-linux-x86_64.zip

# over
```

## 1.3 设置环境变量
打开`/etc/profile`，添加环境变量:
```bash
# protobuf
export PROTOBUF_HOME=/opt/protoc-21.9
export PATH=${PROTOBUF_HOME}/bin:$PATH
```

配置执行权限，重启机器:
```bash
# 执行开机脚本，确认配置无误
source /etc/profile
protoc --version

# 赋予脚本权限
sudo chmod +x /etc/profile

# 重启
reboot

# 重启后确认protoc环境变量配置无误
protoc --version

# over
```

# 二、安装protoc-gen-go
直接执行安装命令:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

然后就可以在`$GOBIN`或默认的`$GOPATH/bin`下看到该插件。
```bash
ll $GOPATH/bin
...
-rwxrwxr-x 1 zhaochun zhaochun  8347446 11月  8 11:42 protoc-gen-go*
...
```

然后将`$GOPATH/bin`加入PATH环境变量:
```bash


```


