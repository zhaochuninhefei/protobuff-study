Language Guide (proto3)
=====
This guide describes how to use the protocol buffer language to structure your protocol buffer data, including .proto file syntax and how to generate data access classes from your .proto files. It covers the proto3 version of the protocol buffers language: for information on the proto2 syntax, see the Proto2 Language Guide.
> 本指南描述了如何使用protobuff语言来构造protobuff数据，包括`.proto`文件语法以及如何从`.proto`文件生成数据访问类。它涵盖了protobuff语言的`proto3`版本: 关于`proto2`语法的信息，请参阅`Proto2 Language Guide`。

This is a reference guide – for a step by step example that uses many of the features described in this document, see the tutorial for your chosen language (currently proto2 only; more proto3 documentation is coming soon).
> 这是一个参考指南，它是一个循序渐进的示例，使用了本文档中描述的许多特性，请参阅所选语言的教程(目前仅支持proto2;更多的proto3文档即将发布)。

# Defining A Message Type
定义消息类型。

First let's look at a very simple example. Let's say you want to define a search request message format, where each search request has a query string, the particular page of results you are interested in, and a number of results per page. Here's the .proto file you use to define the message type.
> 首先让我们看一个非常简单的例子。假设您想要定义一个搜索请求消息格式，其中每个搜索请求都有一个查询字符串、页码以及单页件数。下面是用于定义消息类型的`.proto`文件。

```protobuf
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```
- The first line of the file specifies that you're using proto3 syntax: if you don't do this the protocol buffer compiler will assume you are using proto2. This must be the first non-empty, non-comment line of the file.
    > 文件的第一行指定您使用的是`proto3`语法:如果您不这样做，protobuf编译器将假定您使用的是`proto2`。这必须是文件的第一个非空、非注释行。

- The SearchRequest message definition specifies three fields (name/value pairs), one for each piece of data that you want to include in this type of message. Each field has a name and a type.
    > SearchRequest消息定义里指定了三个字段(名称/值对)，每一条数据片段都会被记录到这个消息里。每个字段都有一个名称和一个类型。

## Specifying Field Types
指定字段类型。

In the above example, all the fields are scalar types: two integers (page_number and result_per_page) and a string (query). However, you can also specify composite types for your fields, including enumerations and other message types.
> 在上面的示例中，所有字段都是`标量类型`:两个整数(页码和每页的结果)和一个字符串(查询)。但是，您也可以为字段指定复合类型，包括枚举和其他消息类型。

## Assigning Field Numbers
分配字段编号。

As you can see, each field in the message definition has a unique number. These field numbers are used to identify your fields in the message binary format, and should not be changed once your message type is in use. Note that field numbers in the range 1 through 15 take one byte to encode, including the field number and the field's type (you can find out more about this in Protocol Buffer Encoding). Field numbers in the range 16 through 2047 take two bytes. So you should reserve the numbers 1 through 15 for very frequently occurring message elements. Remember to leave some room for frequently occurring elements that might be added in the future.
> 可以看到，消息定义中的每个字段都有一个惟一的编号。这些字段编号用于在消息二进制格式中标识字段，一旦开始使用定义好的消息类型，就不能再更改它们。注意，从1到15的字段编号需要一个字节进行编码，包括字段编号和字段类型(您可以在`Protocol Buffer Encoding`中找到关于这方面的更多信息)。16到2047范围内的字段号占用两个字节。因此，应当将使用频率高的字段安排到`[1,15]`区间，甚至保留一部分`[1,15]`区间的编号，以供以后可能会出现的使用频率高的字段使用。

The smallest field number you can specify is 1, and the largest is 2^29 - 1, or 536,870,911. You also cannot use the numbers 19000 through 19999 (FieldDescriptor::kFirstReservedNumber through FieldDescriptor::kLastReservedNumber), as they are reserved for the Protocol Buffers implementation—the protocol buffer compiler will complain if you use one of these reserved numbers in your .proto. Similarly, you cannot use any previously reserved field numbers.
> 最小的字段编号是1,最大则是`2^29 - 1`，即`536,870,911`。注意不能使用`[19000,19999]`区间的编号(从`FieldDescriptor::kFirstReservedNumber`到`FieldDescriptor::kLastReservedNumber`)，因为它们是为protobuff实现保留的，如果你在`.proto`中使用了这些保留的数字之一，protobuff编译器将会报错。同样，您不能使用以前保留的任何字段号。

## Specifying Field Rules
指定字段规则。

Message fields can be one of the following:
> 消息字段可以是以下字段之一:

- `singular`: a well-formed message can have zero or one of this field (but not more than one). When using proto3 syntax, this is the default field rule when no other field rules are specified for a given field. You cannot determine whether it was parsed from the wire. It will be serialized to the wire unless it is the default value. For more on this subject, see Field Presence.
  > 单数: 格式良好的消息可以有0个或1个该字段(但不能超过一个)。在`proto3`中，singular是默认字段规则。protobuff无法确定是否已经解析某个singular字段，singular字段只有在非默认值时才会被序列化。有关此主题的更多信息，请参见`Field Presence`。
  > 
  > 关于singular字段的解析和序列化的说明，`You cannot determine whether it was parsed from the wire. It will be serialized to the wire unless it is the default value.`这两句话不是特别理解什么意思，大概是"singular字段如果没有塞值或者塞的默认值的话，就不会被序列化"的意思。`wire`在这里应该是泛指用于传输、存储序列化结果的"线路"。

- `optional`: the same as singular, except that you can check to see if the value was explicitly set. An optional field is in one of two possible states:
  - the field is set, and contains a value that was explicitly set or parsed from the wire. It will be serialized to the wire.
  - the field is unset, and will return the default value. It will not be serialized to the wire.
  > 可选:与单数相同，只是可以检查是否显式设置了值。可选字段有两种可能的状态:
  > - 该字段已设置，并包含显式设置或解析的值。它将被序列化。
  > - 该字段未设置，将返回默认值。它不会被序列化。

- `repeated`: this field type can be repeated zero or more times in a well-formed message. The order of the repeated values will be preserved.
  > 重复:该字段类型可以在格式良好的消息中重复0次或多次。重复值的顺序将被保留。

- `map`: this is a paired key/value field type. See Maps for more on this field type.
  > 映射：这是一个配对的键/值字段类型。有关此字段类型的更多信息，请参见`Maps`。

In proto3, repeated fields of scalar numeric types use packed encoding by default. You can find out more about packed encoding in Protocol Buffer Encoding.
> 在`proto3`中，标量数字类型的重复字段默认使用打包编码。您可以在`Protocol Buffer Encoding`中找到更多关于打包编码的信息。

## Adding More Message Types
添加更多消息类型。

Multiple message types can be defined in a single .proto file. This is useful if you are defining multiple related messages – so, for example, if you wanted to define the reply message format that corresponds to your SearchResponse message type, you could add it to the same .proto:
> 多个消息类型可以在一个`.proto`文件中定义。这在定义多个相关消息时非常有用，例如，如果想定义与SearchResponse消息类型对应的回复消息格式，可以将其添加到相同的`.proto`中:

```protobuf
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

message SearchResponse {
 ...
}
```

## Adding Comments
添加注释。

To add comments to your .proto files, use C/C++-style // and /* ... */ syntax.
> 要向您的`.proto`文件添加注释，请使用`C/C++`样式的`//`和`/* ... */`语法。

```protobuf
/* SearchRequest represents a search query, with pagination options to
 * indicate which results to include in the response. */

message SearchRequest {
  string query = 1;
  int32 page_number = 2;  // Which page number do we want?
  int32 result_per_page = 3;  // Number of results to return per page.
}
```

## Reserved Fields
保留字段。

If you update a message type by entirely removing a field, or commenting it out, future users can reuse the field number when making their own updates to the type. This can cause severe issues if they later load old versions of the same .proto, including data corruption, privacy bugs, and so on. One way to make sure this doesn't happen is to specify that the field numbers (and/or names, which can also cause issues for JSON serialization) of your deleted fields are reserved. The protocol buffer compiler will complain if any future users try to use these field identifiers.
> 如果通过完全删除字段或将其注释掉的方式来更新消息类型，那么将来的开发者在修改该消息类型时就可能会重复使用字段编号。这样可能会在使用老版本的`.proto`时造成一些问题，包括数据损坏，隐私漏洞等。要确保这种情况不会发生的一种方法是，将那些被删除的字段编号指定为保留。另外删除的字段名也应该指定为保留，否则在做JSON序列化操作时也会引起类似的问题。如果开发者试图使用被保留的字段名或字段编号，protobuff编译器将会报错。

```protobuf
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}
```

Note that you can't mix field names and field numbers in the same reserved statement.
> 注意，您不能在同一个保留语句中混合字段名和字段编号。

## What's Generated From Your .proto?
你的`.proto`文件生成了什么东西？

When you run the protocol buffer compiler on a .proto, the compiler generates the code in your chosen language you'll need to work with the message types you've described in the file, including getting and setting field values, serializing your messages to an output stream, and parsing your messages from an input stream.
> 当您在`.proto`上运行protobuff编译器时，编译器将根据选择的语言生成`.proto`文件中描述的消息类型的代码，包括获取和设置字段值，将消息序列化到输出流，以及解析来自输入流的消息。

- For C++, the compiler generates a .h and .cc file from each .proto, with a class for each message type described in your file.
  > 对于C++，编译器从每个`.proto`生成一个`.h`和`.cc`文件，并为文件中描述的每种消息类型提供一个类。

- For Java, the compiler generates a .java file with a class for each message type, as well as a special Builder class for creating message class instances.
  > 对于Java，编译器生成一个`.java`文件，其中包含针对每种消息类型的类，以及用于创建消息类实例的特殊Builder类。

- For Kotlin, in addition to the Java generated code, the compiler generates a .kt file for each message type, containing a DSL which can be used to simplify creating message instances.
  > 对于Kotlin，除了Java生成的代码外，编译器还为每种消息类型生成一个`.kt`文件，其中包含一个DSL，可用于简化消息实例的创建。

- Python is a little different — the Python compiler generates a module with a static descriptor of each message type in your .proto, which is then used with a metaclass to create the necessary Python data access class at runtime.
  > Python略有不同，Python编译器生成一个模块，该模块带有`.proto`中每种消息类型的静态描述符，然后与元类一起使用，在运行时创建必要的Python数据访问类。

- For Go, the compiler generates a .pb.go file with a type for each message type in your file.
  > 对于Go，编译器生成一个`.pb.go`文件，为文件中的每个消息类型提供一个`type`。

- For Ruby, the compiler generates a .rb file with a Ruby module containing your message types.
  > 对于Ruby，编译器会生成一个`.rb`文件，其中包含一个包含消息类型的Ruby模块。

- For Objective-C, the compiler generates a pbobjc.h and pbobjc.m file from each .proto, with a class for each message type described in your file.
  > 对于Objective-C，编译器从每个`.proto`文件生成一个`pbobjc.h`和一个`pbobjc.m`文件，并为文件中描述的每种消息类型提供一个类。

- For C#, the compiler generates a .cs file from each .proto, with a class for each message type described in your file.
  > 对于c#，编译器从每个`.proto`生成一个`.cs`文件，并为文件中描述的每种消息类型提供一个类。

- For Dart, the compiler generates a .pb.dart file with a class for each message type in your file.
  > 对于Dart，编译器生成一个`.pb.dart`文件，为文件中的每种消息类型提供一个类。

You can find out more about using the APIs for each language by following the tutorial for your chosen language (proto3 versions coming soon). For even more API details, see the relevant API reference (proto3 versions also coming soon).
> 通过阅读所选语言的教程(即将推出proto3版本)，您可以了解更多关于为每种语言使用api的信息。有关更多API细节，请参阅相关API参考(proto3版本也将很快发布)。


