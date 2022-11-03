Overview
=====

Protocol buffers provide a language-neutral, platform-neutral, extensible mechanism for serializing structured data in a forward-compatible and backward-compatible way. It’s like JSON, except it's smaller and faster, and it generates native language bindings.
> `Protocol buffers`提供了一种语言无关、平台无关、可扩展的机制，用于以向前兼容和向后兼容的方式序列化结构化数据。它类似于JSON，只是它更小、更快，而且它生成本地语言绑定。
> 
> `Protocol buffers`即`protobuf`，后面用`protobuf`指替。

Protocol buffers are a combination of the definition language (created in .proto files), the code that the proto compiler generates to interface with data, language-specific runtime libraries, and the serialization format for data that is written to a file (or sent across a network connection).
> protobuf是定义语言(在.proto文件中创建)、proto编译器生成的用于与数据交互的代码、特定于语言的运行时库以及写入文件(或通过网络连接发送)的数据的序列化格式的组合。

# What Problems do Protocol Buffers Solve?
protobuf用于解决什么问题?

Protocol buffers provide a serialization format for packets of typed, structured data that are up to a few megabytes in size. The format is suitable for both ephemeral network traffic and long-term data storage. Protocol buffers can be extended with new information without invalidating existing data or requiring code to be updated.
> protobuf为最大可达几兆字节的类型化结构化数据包提供了一种序列化格式。该格式既适用于短暂的网络流量，也适用于长期的数据存储。protobuf可以使用新信息进行扩展，而无需使现有数据失效或要求更新代码。

Protocol buffers are the most commonly-used data format at Google. They are used extensively in inter-server communications as well as for archival storage of data on disk. Protocol buffer messages and services are described by engineer-authored .proto files. The following shows an example message:
> protobuf是谷歌中最常用的数据格式。它们广泛用于服务器间通信以及磁盘上的数据归档存储。protobuf的消息和服务由工程师编写的`.proto`文件描述。
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
> protobuf编译器protoc直接支持以下语言:

C++, C#, Java, Kotlin, Objective-C, PHP, Python, Ruby

The following languages are supported by Google, but the projects' source code resides in GitHub repositories. The protoc compiler uses plugins for these languages:
> 以下语言也得到了Google的支持，但是其项目的源代码位于GitHub存储库中。原始编译器需要使用这些语言的对应插件：

Dart, Go

Additional languages are not directly supported by Google, but rather by other GitHub projects. These languages are covered in Third-Party Add-ons for Protocol Buffers.
> 谷歌不直接支持其他语言，而是由其他GitHub项目支持，通过protobuf的第三方插件来使用。

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
> 对软件产品来说，向后兼容一般都是要保证的，而向前兼容则是不太常见的。只要您在更新Proto定义时遵循一些简单的做法，旧代码就会读取新消息而无需忽略任何新添加的字段。对于旧代码，被删除的字段将具有其默认值，并且删除的重复字段将为空。有关“重复”字段的信息，请参见该主题后面的`Protocol Buffers Definition Syntax`。

New code will also transparently read old messages. New fields will not be present in old messages; in these cases protocol buffers provide a reasonable default value.
> 新代码还将透明地读取旧消息。新字段将不会出现在旧消息中;在这些情况下，protobuf提供了一个合理的缺省值。

## When are Protocol Buffers not a Good Fit?
什么时候不适合使用protobuf?

Protocol buffers do not fit all data. In particular:
> protobuf并不适合所有数据。特别是:

- Protocol buffers tend to assume that entire messages can be loaded into memory at once and are not larger than an object graph. For data that exceeds a few megabytes, consider a different solution; when working with larger data, you may effectively end up with several copies of the data due to serialized copies, which can cause surprising spikes in memory usage.
    > protobuf倾向于假设整个消息可以一次性加载到内存中，并且不大于一个对象图。对于超过几兆字节的数据，考虑不同的解决方案; 在处理较大的数据时，由于序列化的副本，您可能最终会得到多个数据副本，这可能会导致内存使用出现惊人的峰值。
    >
    > 意思是protobuf不适合处理较大(超过几兆字节)的数据。`object graph`对象图，说的是对象通过相互之间的引用关系(引用链)组成了一个图，比如JVM的可达性分析就使用了对象图技术。`not larger than an object graph`应该是说一个消息中不要存在两个无法通过引用链到达的对象。

- When protocol buffers are serialized, the same data can have many different binary serializations. You cannot compare two messages for equality without fully parsing them.
    > 当对protobuf进行序列化时，相同的数据可以有许多不同的二进制序列化。如果不完全解析两条消息，就无法比较它们是否相等。
    >
    > 意思是protobuf不适合需要对序列化结果进行比较的场景。

- Messages are not compressed. While messages can be zipped or gzipped like any other file, special-purpose compression algorithms like the ones used by JPEG and PNG will produce much smaller files for data of the appropriate type.
    > protobuf消息没有被压缩。虽然也可以像任何其他文件一样对protobuf消息进行zip或gzip压缩，但那些特定文件(比如JPEG和PNG)的特定压缩算法能压缩成更小的文件。
    >
    > 意思是protobuf不适合处理图片/视频等需要特殊压缩算法的数据。

- Protocol buffer messages are less than maximally efficient in both size and speed for many scientific and engineering uses that involve large, multi-dimensional arrays of floating point numbers. For these applications, FITS and similar formats have less overhead.
    > 对于许多涉及大量多维浮点数数组的科学和工程应用来说，protobuf消息在大小和速度上都没有达到最大效率。对于这些应用程序，FITS和类似格式的开销更小。

- Protocol buffers are not well supported in non-object-oriented languages popular in scientific computing, such as Fortran and IDL.
    > 在科学计算中流行的非面向对象语言(如Fortran和IDL)中，protobuf没有得到很好的支持。

- Protocol buffer messages don't inherently self-describe their data, but they have a fully reflective schema that you can use to implement self-description. That is, you cannot fully interpret one without access to its corresponding .proto file.
    > protobuf消息本身并不自我描述它们的数据，但是它们有一个完全反射模式，您可以使用它来实现自我描述。也就是说，如果不访问它对应的`.proto`文件，就不能完全解释它。
    >
    > 所以这一条是想表达protobuf不适合什么场景？必须加入`.proto`文件会有什么限制？

- Protocol buffers are not a formal standard of any organization. This makes them unsuitable for use in environments with legal or other requirements to build on top of standards.
    > protobuf不是任何组织的正式标准。这使得它们不适合在具有基于标准构建的法律或其他要求的环境中使用。

# Who Uses Protocol Buffers?
谁使用protobuf?

Many externally-available projects use protocol buffers, including the following:
> 许多外部可用的项目使用protobuf，包括以下项目:

- gRPC
- Google Cloud
- Envoy Proxy

# How do Protocol Buffers Work?
protobuf是如何工作的?

The following diagram shows how you use protocol buffers to work with your data.
> 下图显示了如何使用protobuf处理数据。

![](images/protocol-buffers-concepts.png)

The code generated by protocol buffers provides utility methods to retrieve data from files and streams, extract individual values from the data, check if data exists, serialize data back to a file or stream, and other useful functions.
> protobuf生成的代码提供了从文件和流中检索数据、从数据中提取单个值、检查数据是否存在、将数据序列化回文件或流以及其他有用的函数的实用方法。

The following code samples show you an example of this flow in Java. As shown earlier, this is a .proto definition:
> 下面的代码示例向您展示了Java中使用protobuf的一个示例。如前所述，这是一个`.proto`定义:

```protobuf
message Person {
  optional string name = 1;
  optional int32 id = 2;
  optional string email = 3;
}
```

Compiling this .proto file creates a Builder class that you can use to create new instances, as in the following Java code:
> 编译这个`.proto`文件会创建一个Builder类，您可以使用它来创建新的实例，如下面的Java代码所示:

```java
// 使用protc编译器对`.proto`文件进行编译，得到一个Person类
// 创建一个Person实例:john
Person john = Person.newBuilder()
    .setId(1234)
    .setName("John Doe")
    .setEmail("jdoe@example.com")
    .build();
output = new FileOutputStream(args[0]);
// 将Person实例写入文件
john.writeTo(output);
```

You can then deserialize data using the methods protocol buffers creates in other languages, like C++:
> 然后可以使用protobuf在其他语言(如c++)中创建的方法反序列化数据:

```cpp
// 同样是protoc编译器对`.proto`文件进行编译，得到一个Person类
// 声明一个Person实例:john
Person john;
fstream input(argv[1], ios::in | ios::binary);
// 读取Java程序生成的Person文件，填充到Person实例
john.ParseFromIstream(&input);
int id = john.id();
std::string name = john.name();
std::string email = john.email();
```

# Protocol Buffers Definition Syntax
protobuf定义语法。

When defining .proto files, you can specify that a field is either optional or repeated (proto2 and proto3) or singular (proto3). (The option to set a field to required is absent in proto3 and strongly discouraged in proto2. For more on this, see "Required is Forever" in Specifying Field Rules.)
> 在定义`.proto`文件时，可以指定一个字段是可选的或重复的(proto2和proto3)或单数的(proto3)。(在proto3中没有将字段设置为required的选项，在proto2中强烈不鼓励。更多信息请参见指定字段规则中的“Required is Forever”。)

After setting the optionality/repeatability of a field, you specify the data type. Protocol buffers support the usual primitive data types, such as integers, booleans, and floats. For the full list, see Scalar Value Types.
> 设置字段的选项/重复性后，需要指定数据类型。protobuf支持通常的原始数据类型，例如整数，布尔值和浮点。有关完整列表，请参见标量值类型。

A field can also be of:
> 一个字段也可以是:

- A message type, so that you can nest parts of the definition, such as for repeating sets of data.
    > 消息类型，以便您可以嵌套定义的部分内容，例如用于重复数据集。

- An enum type, so you can specify a set of values to choose from.
    > 枚举类型，因此可以指定一组要从中选择的值。

- A oneof type, which you can use when a message has many optional fields and at most one field will be set at the same time.
    > oneof类型，当消息有许多可选字段且同时最多设置一个字段时，可以使用该类型。

- A map type, to add key-value pairs to your definition.
    > map类型，用于向定义中添加键值对。

In proto2, messages can allow extensions to define fields outside of the message, itself. For example, the protobuf library's internal message schema allows extensions for custom, usage-specific options.
> 在proto2中，消息可以允许扩展在消息本身之外定义字段。例如，protobuf库的内部消息模式允许对定制的、特定于使用的选项进行扩展。

For more information about the options available, see the language guide for proto2 or proto3.
> 有关可用选项的更多信息，请参阅proto2或proto3的语言指南。

After setting optionality and field type, you assign a field number. Field numbers cannot be repurposed or reused. If you delete a field, you should reserve its field number to prevent someone from accidentally reusing the number.
> 设置可选性和字段类型后，分配字段号。字段号不能被重新使用或重用。如果您删除一个字段，您应该保留它的字段编号，以防止有人意外重用该编号。

# Additional Data Type Support
附加数据类型支持。

Protocol buffers support many scalar value types, including integers that use both variable-length encoding and fixed sizes. You can also create your own composite data types by defining messages that are, themselves, data types that you can assign to a field. In addition to the simple and composite value types, several common types are published.
> protobuf支持许多标量值类型，包括使用变长编码和固定大小的整数。还可以通过定义消息来创建自己的复合数据类型，这些消息本身就是可以分配给字段的数据类型。除了简单值类型和复合值类型外，还发布了几种常见类型。

## Common Types
通用类型。

- Duration is a signed, fixed-length span of time, such as 42s.
    > 持续时间是有符号的固定长度的时间跨度，例如42秒。

- Timestamp is a point in time independent of any time zone or calendar, such as 2017-01-15T01:30:15.01Z.
    > 时间戳是一个独立于任何时区或日历的时间点，例如`2017-01-15T01:30:15.01Z`。

- Interval is a time interval independent of time zone or calendar, such as 2017-01-15T01:30:15.01Z - 2017-01-16T02:30:15.01Z.
    > 时间间隔是一个与时区或日历无关的时间间隔，如`2017-01-15T01:30:15.01Z - 2017-01-16T02:30:15.01Z`。

- Date is a whole calendar date, such as 2025-09-19.
    > Date是一个完整的日历日期，例如`2025-09-19`。

- DayOfWeek is a day of the week, such as Monday.
    > DayOfWeek是一周中的某一天，例如星期一。

- TimeOfDay is a time of day, such as 10:42:23.
    > TimeOfDay是一天中的一个时间，例如`10:42:23`。

- LatLng is a latitude/longitude pair, such as 37.386051 latitude and -122.083855 longitude.
    > Latlng是纬度/经度对，例如`37.386051纬度和-122.083855经度`。

- Money is an amount of money with its currency type, such as 42 USD.
    > Money是某个货币类型的一笔钱，如42美元。

- PostalAddress is a postal address, such as 1600 Amphitheatre Parkway Mountain View, CA 94043 USA.
    > PostalAddress是一个邮政地址，如`1600 Amphitheatre Parkway Mountain View, CA 94043 USA`。

- Color is a color in the RGBA color space.
    > 颜色是RGBA颜色空间中的一种颜色。

- Month is a month of the year, such as April.
    > Month是一年中的一个月，例如四月。

## Protocol Buffers Open Source Philosophy
protobuf开源理念。

Protocol buffers were open sourced in 2008 as a way to provide developers outside of Google with the same benefits that we derive from them internally. We support the open source community through regular updates to the language as we make those changes to support our internal requirements. While we accept select pull requests from external developers, we cannot always prioritize feature requests and bug fixes that don’t conform to Google’s specific needs.
> protobuf自2008年开源，让谷歌以外的开发人员也能享受到它的好处。通过对protobuf的定期更新来支持开源社区，但每次更新都会优先对应谷歌内部需求和BUG修复，然后才是外部提交的需求与BUG修复。
