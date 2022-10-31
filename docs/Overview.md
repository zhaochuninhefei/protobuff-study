Overview
=====

Protocol buffers provide a language-neutral, platform-neutral, extensible mechanism for serializing structured data in a forward-compatible and backward-compatible way. It’s like JSON, except it's smaller and faster, and it generates native language bindings.
> `Protocol buffers`提供了一种语言无关、平台无关、可扩展的机制，用于以向前兼容和向后兼容的方式序列化结构化数据。它类似于JSON，只是它更小、更快，而且它生成本地语言绑定。
> 
> `Protocol buffers`即`protobuff`，后面用`protobuff`指替。

Protocol buffers are a combination of the definition language (created in .proto files), the code that the proto compiler generates to interface with data, language-specific runtime libraries, and the serialization format for data that is written to a file (or sent across a network connection).
> protobuff是定义语言(在.proto文件中创建)、proto编译器生成的用于与数据交互的代码、特定于语言的运行时库以及写入文件(或通过网络连接发送)的数据的序列化格式的组合。

# What Problems do Protocol Buffers Solve?
protobuff用于解决什么问题?

Protocol buffers provide a serialization format for packets of typed, structured data that are up to a few megabytes in size. The format is suitable for both ephemeral network traffic and long-term data storage. Protocol buffers can be extended with new information without invalidating existing data or requiring code to be updated.
> protobuff为最大可达几兆字节的类型化结构化数据包提供了一种序列化格式。该格式既适用于短暂的网络流量，也适用于长期的数据存储。protobuff可以使用新信息进行扩展，而无需使现有数据失效或要求更新代码。

Protocol buffers are the most commonly-used data format at Google. They are used extensively in inter-server communications as well as for archival storage of data on disk. Protocol buffer messages and services are described by engineer-authored .proto files.
> protobuff是谷歌中最常用的数据格式。它们广泛用于服务器间通信以及磁盘上的数据归档存储。protobuff的消息和服务由工程师编写的`.proto`文件描述。

The following shows an example message:
> 下面显示了一个消息示例:

```protobuf
message Person {
  optional string name = 1;
  optional int32 id = 2;
  optional string email = 3;
}
```

