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

Protocol buffers are the most commonly-used data format at Google. They are used extensively in inter-server communications as well as for archival storage of data on disk. Protocol buffer messages and services are described by engineer-authored .proto files. The following shows an example message:
> protobuff是谷歌中最常用的数据格式。它们广泛用于服务器间通信以及磁盘上的数据归档存储。protobuff的消息和服务由工程师编写的`.proto`文件描述。
> 
> 下面显示了一个消息示例:

```protobuf
message Person {
  optional string name = 1;
  optional int32 id = 2;
  optional string email = 3;
}
```

The proto compiler is invoked at build time on .proto files to generate code in various programming languages (covered in Cross-language Compatibility later in this topic) to manipulate the corresponding protocol buffer. Each generated class contains simple accessors for each field and methods to serialize and parse the whole structure to and from raw bytes. The following shows you an example that uses those generated methods:
> proto编译器在`.proto`文件构建时被调用，以生成各种编程语言的代码(本主题后面的跨语言兼容性将讨论)来操作相应的protobuf。每个生成的类都包含针对每个字段的简单访问器和用于序列化和解析整个结构与原始字节之间的方法。下面的示例将使用这些生成的方法:

```java
Person john = Person.newBuilder()
    .setId(1234)
    .setName("John Doe")
    .setEmail("jdoe@example.com")
    .build();
output = new FileOutputStream(args[0]);
john.writeTo(output);
```

Because protocol buffers are used extensively across all manner of services at Google and data within them may persist for some time, maintaining backwards compatibility is crucial. Protocol buffers allow for the seamless support of changes, including the addition of new fields and the deletion of existing fields, to any protocol buffer without breaking existing services.
> 由于protobuf在谷歌上广泛应用于各种服务，并且其中的数据可能会持续存在一段时间，因此保持向后兼容性是至关重要的。protobuf允许对更改的无缝支持，包括向任何protobuf添加新字段和删除现有字段，而不会破坏现有服务。

For more on this topic, see Updating Proto Definitions Without Updating Code, later in this topic.
> 有关此主题的更多信息，请参阅本主题后面的部分: `Updating Proto Definitions Without Updating Code`

# What are the Benefits of Using Protocol Buffers?
使用protobuf有什么好处？

Protocol buffers are ideal for any situation in which you need to serialize structured, record-like, typed data in a language-neutral, platform-neutral, extensible manner. They are most often used for defining communications protocols (together with gRPC) and for data storage.
> protobuf非常适合任何需要以语言无关、平台无关、可扩展的方式序列化结构化、类似记录的类型化数据的情况。它们最常用于定义通信协议(与gRPC一起)和数据存储。

Some of the advantages of using protocol buffers include:
> 使用protobuf的一些优点包括:

- Compact data storage : 紧凑的数据存储
- Fast parsing : 快速解析
- Availability in many programming languages : 在许多编程语言中的可用性
- Optimized functionality through automatically-generated classes : 通过自动生成类实现的功能优化

## Cross-language Compatibility
跨语言的兼容性。

The same messages can be read by code written in any supported programming language. You can have a Java program on one platform capture data from one software system, serialize it based on a .proto definition, and then extract specific values from that serialized data in a separate Python application running on another platform.
> 相同的消息可以被任何支持的编程语言所编写的代码读取。
> 
> 例如，可以让在某个在平台上运行的Java程序根据`.proto`定义对某个数据进行序列化，然后在另一个平台上用某个Python程序来读取并解析这个序列化数据。

The following languages are supported directly in the protocol buffers compiler, protoc:
> protobuff编译器protoc直接支持以下语言:

C++, C#, Java, Kotlin, Objective-C, PHP, Python, Ruby

The following languages are supported by Google, but the projects' source code resides in GitHub repositories. The protoc compiler uses plugins for these languages:
> 以下语言也得到了Google的支持，但是其项目的源代码位于GitHub存储库中。原始编译器需要使用这些语言的对应插件：

Dart, Go

Additional languages are not directly supported by Google, but rather by other GitHub projects. These languages are covered in Third-Party Add-ons for Protocol Buffers.
> 谷歌不直接支持其他语言，而是由其他GitHub项目支持，通过protobuff的第三方插件来使用。

## Cross-project Support
跨项目的支持。

You can use protocol buffers across projects by defining message types in .proto files that reside outside of a specific project’s code base. If you're defining message types or enums that you anticipate will be widely used outside of your immediate team, you can put them in their own file with no dependencies.
> 通过在驻留在特定项目代码库之外的`.proto`文件中定义消息类型，可以跨项目使用协议缓冲区。如果您定义的消息类型或枚举预计将在您的直接团队之外广泛使用，那么可以将它们放在自己的文件中，不存在依赖关系。

A couple of examples of proto definitions widely-used within Google are timestamp.proto and status.proto.
> 谷歌中广泛使用的两个原型定义示例是`timestamp.proto`和`status.proto`。

```
https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/timestamp.proto
https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto
```

## Updating Proto Definitions Without Updating Code
更新Proto定义而不更新代码。

It’s standard for software products to be backward compatible, but it is less common for them to be forward compatible. As long as you follow some simple practices when updating .proto definitions, old code will read new messages without issues, ignoring any newly added fields. To the old code, fields that were deleted will have their default value, and deleted repeated fields will be empty. For information on what “repeated” fields are, see Protocol Buffers Definition Syntax later in this topic.
> 对软件产品来说，向后兼容一般都是要保证的，而向前兼容则是不太常见的。只要您在更新Proto定义时遵循一些简单的做法，旧代码就会读取新消息而无需忽略任何新添加的字段。对于旧代码，被删除的字段将具有其默认值，并且删除的重复字段将为空。有关“重复”字段的信息，请参见该主题后面的协议缓冲区定义语法。

