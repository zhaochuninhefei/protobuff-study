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
  > 单数字段: 如果一个字段是单数字段，那么格式良好的消息里只能有0个或1个该字段。在`proto3`中，singular是默认字段规则。protobuff无法确定是否已经解析某个单数字段，单数字段只有在非默认值时才会被序列化。有关此主题的更多信息，请参见`Field Presence`。
  > 
  > 关于单数字段的解析和序列化的说明，`You cannot determine whether it was parsed from the wire. It will be serialized to the wire unless it is the default value.`这两句话中使用了`wire`，`wire`在这里的意思应该是指`wire protocol`，点到点的通信抽象协议，一种数据传输机制，习惯于被用来描述信息位于应用层上的一种通用表现形式，是一种应用层上的通用协议而非各类应用程序的通用型对象描述语意。在protobuff中的`wire`是想强调序列化或解析的目标是一种用于传输的二进制数据格式，而不是文本格式(protobuff的消息也可以是文本格式)。后续的翻译中，如果没有明确二进制格式与文本格式的区别，则默认就是二进制格式的消息。

- `optional`: the same as singular, except that you can check to see if the value was explicitly set. An optional field is in one of two possible states:
  - the field is set, and contains a value that was explicitly set or parsed from the wire. It will be serialized to the wire.
  - the field is unset, and will return the default value. It will not be serialized to the wire.
  > 可选字段: 与单数字段基本一样，但可以检查是否对其做了显式设值。可选字段有两种可能的状态:
  > - 该字段已设置，并包含显式设置或解析的值。它将被序列化。
  > - 该字段未设置，将返回默认值。它不会被序列化。
  > 
  > 感觉可选字段和单数字段没什么区别，都可以不设值，不设值时都不会被序列化，解析时直接使用默认值。只是可选字段可以检查有没有显式设值。

- `repeated`: this field type can be repeated zero or more times in a well-formed message. The order of the repeated values will be preserved.
  > 重复字段: 该类型字段可以在格式良好的消息中重复0次或多次。重复值的顺序将被保留。
  > 
  > 重复字段对应各种编程语言中的数组、列表等类型。

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

# Scalar Value Types
标量值类型。

A scalar message field can have one of the following types – the table shows the type specified in the .proto file, and the corresponding type in the automatically generated class:
> 标量消息字段可以具有以下类型之一(表格显示了`.proto`文件中指定的类型，以及自动生成的类中相应的类型):

| .proto Type | Notes | C++ Type | Java/Kotlin Type① | Python Type③ | Go Type | Ruby Type | C# Type | PHP Type | Dart Type |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| double | | double | double | float | float64 | Float | double | float | double |
| float | | float | float | float | float32 | Float | float | float | double |
| int32 | <div style="width: 300pt">Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. <br>使用变长编码。负数的效率很低，请使用sint32代替。</div> | int32 | int | int | int32 | Fixnum or Bignum (as required) | int | integer | int |
| int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead.<br>使用变长编码。负数的效率很低，请使用sint64代替。 | int64 | long | int/long④ | int64 | Bignum | long | integer/string⑥ | Int64 |
| uint32 | Uses variable-length encoding.<br>使用变长编码，无符号整数。 | uint32 | int② | int/long④ | uint32 | Fixnum or Bignum (as required) | uint | integer | int |
| uint64 | Uses variable-length encoding.<br>使用变长编码，无符号整数。 | uint64 | long② | int/long④ | uint64 | Bignum | ulong | integer/string⑥	 | Int64 |
| sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.<br>使用变长编码。有符号整数。负数时相比常规的int32更高效。 | int32 | int | int | int32 | Fixnum or Bignum (as required) | int | integer | int |
| sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.<br>使用变长编码。有符号整数。负数时相比常规的int64更高效。 | int64 | long | int/long④ | int64 | Bignum | long | integer/string⑥	 | Int64 |
| fixed32 | Always four bytes. More efficient than uint32 if values are often greater than `2^28`.<br>固定4个字节的整型，无符号。如果值常常大于`2^28`，则比uint32更高效。 | uint32 | int② | int/long④ | uint32 | Fixnum or Bignum (as required) | uint | integer | int |
| fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than `2^56`.<br>固定8个字节的整型，无符号。如果值常常大于`2^56`，则比uint64更高效。 | uint64 | long② | int/long④ | uint64 | Bignum | ulong | integer/string⑥ | Int64 |
| sfixed32 | Always four bytes.<br>固定4个字节的整型，有符号。 | int32 | int | int | int32 | Fixnum or Bignum (as required) | int | integer | int |
| sfixed64 | Always eight bytes.<br>固定8个字节的整型，有符号。 | int64 | long | int/long④ | int64 | Bignum | long | integer/string⑥ | Int64 |
| bool | | bool | boolean | bool | bool | TrueClass/FalseClass | bool | boolean | bool |
| string | A string must always contain UTF-8 encoded or 7-bit ASCII text, and cannot be longer than `2^32`.<br>字符串，UTF-8编码或7位ASCII文本，长度不能超过`2^32`。 | string | String | str/unicode⑤ | string | String (UTF-8) | string | string | String |
| bytes | May contain any arbitrary sequence of bytes no longer than `2^32`.<br>长度不超过`2^32`的任意字节序列。 | string | ByteString | str(Python2)/bytes(Python3) | []byte | String (ASCII-8BIT) | ByteString | string | List |

You can find out more about how these types are encoded when you serialize your message in Protocol Buffer Encoding.
> 在`Protocol Buffer Encoding`一节可以看到更多关于序列化消息时如何编码各种类型的信息。

① Kotlin uses the corresponding types from Java, even for unsigned types, to ensure compatibility in mixed Java/Kotlin codebases.
> Kotlin使用Java的相应类型，甚至对于未签名类型也一样，以确保混合Java/Kotlin代码库中的兼容性。

② In Java, unsigned 32-bit and 64-bit integers are represented using their signed counterparts, with the top bit simply being stored in the sign bit.
> 在Java中，无符号的32位/64位整型直接使用它们的有符号整型，其最高位简单地存储在符号位中。

③ In all cases, setting values to a field will perform type checking to make sure it is valid.
> 在(Python的)所有情况下，将值塞进字段将执行类型检查，以确保它是有效的。

④ 64-bit or unsigned 32-bit integers are always represented as long when decoded, but can be an int if an int is given when setting the field. In all cases, the value must fit in the type represented when set. See ②.
> (Python的)64位整型或无符号的32位整型在解码时始终用long，但如果在设值时类型是int，则也可以用int。无论如何，必须匹配设值时的类型。参考②。
> 
> 为啥参考②?

⑤ Python strings are represented as unicode on decode but can be str if an ASCII string is given (this is subject to change).
> Python字符串在解码时使用unicode，但如果明确使用了ASCII字符串，则可以使用str。(这可能会更改)

⑥ Integer is used on 64-bit machines and string is used on 32-bit machines.
> (PHP的长整型，在)64位平台上使用Integer，32位平台上使用string。

# Default Values
默认值。

When a message is parsed, if the encoded message does not contain a particular singular element, the corresponding field in the parsed object is set to the default value for that field. These defaults are type-specific:
> 在解析消息时，如果编码后的消息不包含某个singular字段，则将该字段解析为默认值。这些默认值是特定于类型的:

- For strings, the default value is the empty string.
  > 对于字符串，默认值是空字符串。

- For bytes, the default value is empty bytes.
  > 对于字节，默认值是空字节。

- For bools, the default value is false.
  > 对于布尔，默认值是false。

- For numeric types, the default value is zero.
  > 对于数字类型，默认值是0。

- For enums, the default value is the first defined enum value, which must be 0.
  > 对于枚举，默认值是第一个定义的枚举值，该值必须为0。

- For message fields, the field is not set. Its exact value is language-dependent. See the generated code guide for details.
  > 对于消息字段，没有设置该字段。它的确切值取决于语言。详细信息请参见生成的代码指南。

The default value for repeated fields is empty (generally an empty list in the appropriate language).
> 重复字段的默认值为空(在适当的语言中通常是空列表)。

Note that for scalar message fields, once a message is parsed there's no way of telling whether a field was explicitly set to the default value (for example whether a boolean was set to false) or just not set at all: you should bear this in mind when defining your message types. For example, don't have a boolean that switches on some behavior when set to false if you don't want that behavior to also happen by default. Also note that if a scalar message field is set to its default, the value will not be serialized on the wire.
> 注意，对于标量消息字段，一旦解析了消息，就无法知道字段是显式设置为默认值(例如布尔值是否设置为false)还是根本没有设置: 在定义消息类型时应该记住这一点。例如，如果您不希望默认情况下也发生某些行为，就不要使用一个在设置为false时开启该行为的布尔值。还要注意，如果将标量消息字段设置为默认值，则该值将不会被序列化。

See the generated code guide for your chosen language for more details about how defaults work in generated code.
> 有关默认值如何在生成代码中工作的详细信息，请参阅所选语言的生成代码指南。

# Enumerations
枚举。

When you're defining a message type, you might want one of its fields to only have one of a pre-defined list of values. For example, let's say you want to add a corpus field for each SearchRequest, where the corpus can be UNIVERSAL, WEB, IMAGES, LOCAL, NEWS, PRODUCTS or VIDEO. You can do this very simply by adding an enum to your message definition with a constant for each possible value.
> 在定义消息类型时，可能希望其中一个字段只有预定义值列表中的一个。例如，假设您想为每个SearchRequest添加一个语料库字段，其中语料库可以是UNIVERSAL、WEB、IMAGES、LOCAL、NEWS、PRODUCTS或VIDEO。您可以通过在消息定义中添加一个枚举，为每个可能的值添加一个常量来非常简单地做到这一点。

In the following example we've added an enum called Corpus with all the possible values, and a field of type Corpus:
> 在下面的示例中，我们添加了一个名为Corpus的枚举，包含所有可能的值，以及一个类型为Corpus的字段:

```protobuf
enum Corpus {
  CORPUS_UNSPECIFIED = 0;
  CORPUS_UNIVERSAL = 1;
  CORPUS_WEB = 2;
  CORPUS_IMAGES = 3;
  CORPUS_LOCAL = 4;
  CORPUS_NEWS = 5;
  CORPUS_PRODUCTS = 6;
  CORPUS_VIDEO = 7;
}
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
  Corpus corpus = 4;
}
```

As you can see, the Corpus enum's first constant maps to zero: every enum definition must contain a constant that maps to zero as its first element. This is because:
> 正如您所看到的，Corpus enum的第一个常量映射到0:每个enum定义必须包含一个映射到0的常量作为其第一个元素。这是因为:

- There must be a zero value, so that we can use 0 as a numeric default value.
  > 必须有一个零值，这样我们就可以使用0作为数值默认值。

- The zero value needs to be the first element, for compatibility with the proto2 semantics where the first enum value is always the default.
  > 0值需要是第一个元素，以兼容proto2语义: 第一个enum值总是默认值。

You can define aliases by assigning the same value to different enum constants. To do this you need to set the allow_alias option to true, otherwise the protocol compiler will generate an error message when aliases are found. Though all alias values are valid during deserialization, the first value is always used when serializing.
> 可以通过将相同的值赋给不同的enum常量来定义别名。为此，您需要将允许别名选项设置为true，否则当找到别名时，协议编译器将生成一条错误消息。尽管所有的别名值在反序列化期间都是有效的，但在序列化时总是使用第一个值。

```protobuf
enum EnumAllowingAlias {
  option allow_alias = true;
  EAA_UNSPECIFIED = 0;
  EAA_STARTED = 1;
  EAA_RUNNING = 1;
  EAA_FINISHED = 2;
}
enum EnumNotAllowingAlias {
  ENAA_UNSPECIFIED = 0;
  ENAA_STARTED = 1;
  // ENAA_RUNNING = 1;  // Uncommenting this line will cause a compile error inside Google and a warning message outside.
  ENAA_FINISHED = 2;
}
```

Enumerator constants must be in the range of a 32-bit integer. Since enum values use varint encoding on the wire, negative values are inefficient and thus not recommended. You can define enums within a message definition, as in the above example, or outside – these enums can be reused in any message definition in your .proto file. You can also use an enum type declared in one message as the type of a field in a different message, using the syntax _MessageType_._EnumType_.
> 枚举器常量必须在32位整数的范围内。由于枚举值使用varint编码，负值效率很低，因此不建议使用负值。您可以在消息定义内定义枚举(如上例所示)，也可以在`.proto`文件中的任何消息定义中重用这些枚举。您还可以通过语法`_MessageType_._EnumType_`在一条消息中使用另一条消息中声明的enum类型作为字段类型。

When you run the protocol buffer compiler on a .proto that uses an enum, the generated code will have a corresponding enum for Java, Kotlin, or C++, or a special EnumDescriptor class for Python that's used to create a set of symbolic constants with integer values in the runtime-generated class.
> 使用protobuff编译器编译一个定义了枚举的`.proto`时，生成的代码将具有Java、Kotlin或C++对应的枚举，或者具有Python的特殊EnumDescriptor类，用于在运行时生成的类中创建一组具有整数值的符号常量。

**Caution:** the generated code may be subject to language-specific limitations on the number of enumerators (low thousands for one language). Please review the limitations for the languages you plan to use.
> **注意:** 生成的代码可能受特定于语言的枚举数限制(某些语言的枚举数较低)。请检查您计划使用的语言的限制。

During deserialization, unrecognized enum values will be preserved in the message, though how this is represented when the message is deserialized is language-dependent. In languages that support open enum types with values outside the range of specified symbols, such as C++ and Go, the unknown enum value is simply stored as its underlying integer representation. In languages with closed enum types such as Java, a case in the enum is used to represent an unrecognized value, and the underlying integer can be accessed with special accessors. In either case, if the message is serialized the unrecognized value will still be serialized with the message.
> 在反序列化期间，未识别的enum值将保留在消息中。具体如何表示未识别的enum值则取决于语言。在C++和Go等支持开放enum类型(其值在指定符号范围之外)的语言中，未知的enum值简单地存储为其底层整数表示。在具有封闭enum类型的语言(如Java)中，枚举中的一个case用于表示不可识别的值，并且可以使用特殊的访问器访问底层整数。无论哪种情况，消息被序列化时，未识别的值总会与消息一起序列化。
> 
> `a case in the enum is used to represent an unrecognized value`这句话不明白什么意思，这里的`case`是什么意思？是说Java这样的语言需要在对应的枚举类中定义一个专门用于表示未识别的枚举实例吗？

For more information about how to work with message enums in your applications, see the generated code guide for your chosen language.
> 有关如何在应用程序中使用消息枚举的更多信息，请参阅所选语言的生成代码指南。

## Reserved Values
保留值。

If you update an enum type by entirely removing an enum entry, or commenting it out, future users can reuse the numeric value when making their own updates to the type. This can cause severe issues if they later load old versions of the same .proto, including data corruption, privacy bugs, and so on. One way to make sure this doesn't happen is to specify that the numeric values (and/or names, which can also cause issues for JSON serialization) of your deleted entries are reserved. The protocol buffer compiler will complain if any future users try to use these identifiers. You can specify that your reserved numeric value range goes up to the maximum possible value using the max keyword.
> 如果通过完全删除枚举条目或将其注释掉来更新枚举类型，那么将来对该类型进行更新时可能会重用该枚举的数值。如果又要使用老版本的`.proto`，这可能会导致严重的问题，包括数据损坏、隐私漏洞等等。要确保不发生这种事情，一种有效的做法是，将废弃的枚举的数值指定为保留。废弃枚举的名称也一样应该保留，避免在JSON序列化操作中引起类似问题。任何用户试图使用这些保留值时，protobuff编译器将会报错。可以使用max关键字指定保留值范围的最大值。

```protobuf
enum Foo {
  reserved 2, 15, 9 to 11, 40 to max;
  reserved "FOO", "BAR";
}
```

Note that you can't mix field names and numeric values in the same reserved statement.
> 注意，您不能在同一个保留语句中混合字段名和数值。

# Using Other Message Types
使用其他消息类型。

You can use other message types as field types. For example, let's say you wanted to include Result messages in each SearchResponse message – to do this, you can define a Result message type in the same .proto and then specify a field of type Result in SearchResponse:
> 您可以使用其他消息类型作为字段类型。例如，假设您想要在每个SearchResponse消息中包含Result消息，您可以在相同的`.proto`中定义Result消息类型，然后在SearchResponse中指定类型为Result的字段:

```protobuf
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```

## Importing Definitions
导入定义。

In the above example, the Result message type is defined in the same file as SearchResponse – what if the message type you want to use as a field type is already defined in another .proto file?
> 在上面的例子中，Result消息类型定义在与SearchResponse相同的文件中，但如果用作字段类型的消息类型定义在另一个`.proto`文件中呢?

You can use definitions from other .proto files by importing them. To import another .proto's definitions, you add an import statement to the top of your file:
> 您可以通过导入其他`.proto`文件中的定义来使用它们。要导入另一个`.proto`的定义，需要在文件的顶部添加import语句:

```protobuf
import "myproject/other_protos.proto";
```

By default, you can use definitions only from directly imported .proto files. However, sometimes you may need to move a .proto file to a new location. Instead of moving the .proto file directly and updating all the call sites in a single change, you can put a placeholder .proto file in the old location to forward all the imports to the new location using the import public notion.
> 默认情况下，只能从直接导入的`.proto`文件中使用定义。然而，有时您可能需要将`.proto`文件移动到一个新的位置。不必直接移动`.proto`文件并在一次更改中更新所有调用站点，您可以在旧位置放置一个占位符`.proto`文件，以使用`import public`将所有导入转发到新位置。

**Note that the public import functionality is not available in Java.**
> 注意，Java中没有公共导入功能。

import public dependencies can be transitively relied upon by any code importing the proto containing the import public statement. For example:
> 公共导入可以实现依赖传递。例如:

新proto定义文件，相关定义都挪到这里了:
```protobuf
// new.proto
// All definitions are moved here
```

旧proto定义文件，大家之前引入的都是这个proto，现在需要利用`import public`引入新的proto定义文件:
```protobuf
// old.proto
// This is the proto that all clients are importing. 
import public "new.proto";
import "other.proto";
```

所有引入了`old.proto`的地方不用修改，就能通过 old.proto 中对 new.proto 的公共导入，来访问 new.proto 中定义的类型:
```protobuf
// client.proto
import "old.proto";
// You use definitions from old.proto and new.proto, but not other.proto
```

The protocol compiler searches for imported files in a set of directories specified on the protocol compiler command line using the -I/--proto_path flag. If no flag was given, it looks in the directory in which the compiler was invoked. In general you should set the --proto_path flag to the root of your project and use fully qualified names for all imports.
> protobuf编译器使用参数`-I/--proto_path`在其命令行上指定一组目录，并在这个目录中搜索导入文件。如果没有显式给出这个参数，则在调用编译命令的当前目录下搜索。一般来说，应该将`--proto_path`参数设置为工程项目的根路径，并对所有导入使用完全限定名。

## Using proto2 Message Types
使用proto2消息类型。

It's possible to import proto2 message types and use them in your proto3 messages, and vice versa. However, proto2 enums cannot be used directly in proto3 syntax (it's okay if an imported proto2 message uses them).
> 可以导入proto2消息类型并在proto3消息中使用它们，反之亦然。但是，proto2枚举不能直接在proto3语法中使用。但如果是导入的proto2消息使用它们则没关系。

# Nested Types
嵌套类型。

You can define and use message types inside other message types, as in the following example – here the Result message is defined inside the SearchResponse message:
> 您可以在其他消息类型中定义和使用消息类型，如下面的示例所示，Result消息在SearchResponse消息中定义:

```protobuf
message SearchResponse {
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result results = 1;
}
```

If you want to reuse this message type outside its parent message type, you refer to it as _Parent_._Type_:
> 在其父消息类型之外重用此消息类型的话，可以用`_Parent_._Type_`的方式引用它:

```protobuf
message SomeOtherMessage {
  SearchResponse.Result result = 1;
}
```

You can nest messages as deeply as you like:
> 消息嵌套想多深都可以:

```protobuf
message Outer {                  // Level 0
  message MiddleAA {  // Level 1
    message Inner {   // Level 2
      int64 ival = 1;
      bool  booly = 2;
    }
  }
  message MiddleBB {  // Level 1
    message Inner {   // Level 2
      int32 ival = 1;
      bool  booly = 2;
    }
  }
}
```

# Updating A Message Type
更新消息类型。

If an existing message type no longer meets all your needs – for example, you'd like the message format to have an extra field – but you'd still like to use code created with the old format, don't worry! It's very simple to update message types without breaking any of your existing code. Just remember the following rules:
> 如果现有的消息类型不再满足所有需求，例如，需要一个额外的字段，但仍然希望使用旧格式创建的代码，不用担心! 在不破坏任何现有代码的情况下更新消息类型非常简单。只要记住下面的规则:

- Don't change the field numbers for any existing fields.
  > 不要更改任何现有字段的字段编号。

- If you add new fields, any messages serialized by code using your "old" message format can still be parsed by your new generated code. You should keep in mind the default values for these elements so that new code can properly interact with messages generated by old code. Similarly, messages created by your new code can be parsed by your old code: old binaries simply ignore the new field when parsing. See the Unknown Fields section for details.
  > 如果添加了新字段，则使用旧格式的代码所序列化的任何消息仍然可以由新生成的代码解析。但要注意新元素的默认值，以便新代码能够正确地与旧代码生成的消息进行交互。类似地，由新代码创建的消息也可以由旧代码解析: 旧代码在解析新格式的消息时会简单地忽略新字段。有关详细信息，请参阅未知字段部分。

- Fields can be removed, as long as the field number is not used again in your updated message type. You may want to rename the field instead, perhaps adding the prefix "OBSOLETE_", or make the field number reserved, so that future users of your .proto can't accidentally reuse the number.
  > 只要字段号不在更新的消息类型中再次使用，就可以删除字段。你可以重命名该字段，比如添加前缀`OBSOLETE_`。或者删除字段之后保留字段编号，以确保将来不会意外重用该字段编号。

- int32, uint32, int64, uint64, and bool are all compatible – this means you can change a field from one of these types to another without breaking forwards- or backwards-compatibility. If a number is parsed from the wire which doesn't fit in the corresponding type, you will get the same effect as if you had cast the number to that type in C++ (for example, if a 64-bit number is read as an int32, it will be truncated to 32 bits).
  > `int32`, `uint32`, `int64`, `uint64`, 以及`bool`类型都是兼容的，这意味着你可以将字段从其中一种类型更改为另一种类型，而不会破坏向前或向后兼容性。如果解析的数字不符合相应的类型，您将得到与在C++中将该数字强制转换为该类型相同的效果(例如，如果将64位数字读为一个int32，它将被截断为32位)。

- sint32 and sint64 are compatible with each other but are not compatible with the other integer types.
  > `sint32`和`sint64`彼此兼容，但不兼容其他整数类型。

- string and bytes are compatible as long as the bytes are valid UTF-8.
  > `string`和`bytes`是兼容的，只要`bytes`是有效的`UTF-8`。

- Embedded messages are compatible with bytes if the bytes contain an encoded version of the message.
  > 如果`bytes`包含消息的编码版本，则嵌入式消息与`bytes`兼容。

- fixed32 is compatible with sfixed32, and fixed64 with sfixed64.
  > `fixed32`与`sfixed32`兼容，`fixed64`与`sfixed64`兼容。

- For string, bytes, and message fields, singular fields are compatible with repeated fields. Given serialized data of a repeated field as input, clients that expect this field to be singular will take the last input value if it's a primitive type field or merge all input elements if it's a message type field. Note that this is not generally safe for numeric types, including bools and enums. Repeated fields of numeric types can be serialized in the packed format, which will not be parsed correctly when a singular field is expected.
  > 对于`string`、`bytes`和嵌套的message字段来说，单数字段与重复字段是兼容的。给定一个重复字段的序列化数据作为输入，期望该字段为单数字段的客户机将接受最后一个输入值(如果它是一个基本类型字段)或合并所有输入元素(如果它是一个message类型字段)。注意，这对于数字类型(包括bool和enum)通常不安全。数值类型的重复字段会用打包格式做序列化，当期望作为单数字段接收时，将无法正确解析其格式。

- enum is compatible with int32, uint32, int64, and uint64 in terms of wire format (note that values will be truncated if they don't fit). However be aware that client code may treat them differently when the message is deserialized: for example, unrecognized proto3 enum types will be preserved in the message, but how this is represented when the message is deserialized is language-dependent. Int fields always just preserve their value.
  > `enum`在通信格式上兼容`int32`, `uint32`, `int64`, 和`uint64`(注意，如果值不适合，将被截断)。然而，请注意，当消息被反序列化时，客户端代码可能会以不同的方式对待它们: 例如，无法识别的proto3 enum类型将保留在消息中，但反序列化时这些无法识别的enum如何表示则取决于具体的语言。Int字段总是只保留它们的值。

- Changing a single optional field or extension into a member of a new oneof is safe and binary compatible. Moving multiple fields into a new oneof may be safe if you are sure that no code sets more than one at a time. Moving any fields into an existing oneof is not safe. Likewise, changing a single field oneof to an optional field or extension is safe.
  > 将单个`optional`字段或扩展更改为新的`oneof`是安全且二进制兼容的。如果没有代码同时给某几个字段中的多个字段设值的话，那么将这几个字段改为一个新的`oneof`也是安全的。任何将字段加入即存的`oneof`都是不安全的。同样的，将一个单字段的`oneof`改为`optional`字段或扩展也是安全的。

# Unknown Fields
未知字段。

Unknown fields are well-formed protocol buffer serialized data representing fields that the parser does not recognize. For example, when an old binary parses data sent by a new binary with new fields, those new fields become unknown fields in the old binary.
> 未知字段是格式良好的protobuff序列化数据中用来表示解析器无法识别的字段。例如，当旧二进制文件解析新二进制文件发送的带有新字段的数据时，这些新字段在旧二进制文件中成为未知字段。

Originally, proto3 messages always discarded unknown fields during parsing, but in version 3.5 we reintroduced the preservation of unknown fields to match the proto2 behavior. In versions 3.5 and later, unknown fields are retained during parsing and included in the serialized output.
> 最初，proto3消息总是在解析过程中丢弃未知字段，但在3.5版本中，我们重新引入了保存未知字段的功能，以匹配proto2行为。在3.5及更高版本中，解析期间将保留未知字段，并将其包含在序列化输出中。

# Any
Any消息类型。

The Any message type lets you use messages as embedded types without having their .proto definition. An Any contains an arbitrary serialized message as bytes, along with a URL that acts as a globally unique identifier for and resolves to that message's type. To use the Any type, you need to import google/protobuf/any.proto.
> Any消息类型允许您将消息作为嵌入类型使用，而不需要它们的`.proto`定义。Any包含作为`bytes`的任意序列化消息，以及充当全局唯一标识符并解析为该消息类型的URL。要使用Any的话，需要引入`google/protobuf/any.proto`。

```protobuf
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```

The default type URL for a given message type is type.googleapis.com/_packagename_._messagename_.
> 一个给定的消息类型的URL的默认类型是`type.googleapis.com/_packagename_._messagename_`。

Different language implementations will support runtime library helpers to pack and unpack Any values in a typesafe manner – for example, in Java, the Any type will have special pack() and unpack() accessors, while in C++ there are PackFrom() and UnpackTo() methods:
> 不同的语言实现将支持运行时库帮助程序以类型安全的方式打包和解包Any值。例如，在Java中，Any类型将实现自己的`pack()`和`unpack()`访问器，而C++则会提供`PackFrom()`和`UnpackTo()`方法。

```cpp
// Storing an arbitrary message type in Any.
NetworkErrorDetails details = ...;
ErrorStatus status;
status.add_details()->PackFrom(details);

// Reading an arbitrary message from Any.
ErrorStatus status = ...;
for (const google::protobuf::Any& detail : status.details()) {
  if (detail.Is<NetworkErrorDetails>()) {
    NetworkErrorDetails network_error;
    detail.UnpackTo(&network_error);
    ... processing network_error ...
  }
}
```

**Currently the runtime libraries for working with Any types are under development.**
> 目前，用于处理Any类型的运行时库正在开发中。

If you are already familiar with proto2 syntax, the Any can hold arbitrary proto3 messages, similar to proto2 messages which can allow extensions.
> 如果您已经熟悉了proto2语法，Any可以保存任意的proto3消息，类似于允许扩展的proto2消息。

# Oneof
Oneof消息类型。

If you have a message with many fields and where at most one field will be set at the same time, you can enforce this behavior and save memory by using the oneof feature.
> 如果您有一个包含多个字段的消息，并且同时最多设置一个字段，您可以通过使用oneof特性强制执行此行为并节省内存。

Oneof fields are like regular fields except all the fields in a oneof share memory, and at most one field can be set at the same time. Setting any member of the oneof automatically clears all the other members. You can check which value in a oneof is set (if any) using a special case() or WhichOneof() method, depending on your chosen language.
> oneof字段与普通字段大致相同，除了oneof共享内存中的这些成员字段(最多可以同时设置它们中的一个)。设置oneof的任意成员将自动清除所有其他成员。您可以使用`case()`或`WhichOneof()`等方法检查哪个成员字段被设置了(如果有的话)，这取决于您所选择的语言。

Note that if multiple values are set, the last set value as determined by the order in the proto will overwrite all previous ones.
> 注意，如果设置了多次成员值，最后一个由proto中的顺序确定的成员值将覆盖之前所有的设值。
> 
> 什么是`the order in the proto`？ 不是程序上最后一次设值吗？

## Using Oneof
使用Oneof。

To define a oneof in your .proto you use the oneof keyword followed by your oneof name, in this case test_oneof:
> 要在`.proto`里定义一个oneof，需要使用`oneof`关键字，以及一个oneof名称，如下所示:

```protobuf
message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}
```

You then add your oneof fields to the oneof definition. You can add fields of any type, except map fields and repeated fields.
> 然后就可以将oneof的成员字段加入oneof定义，除了map和重复字段，oneof成员字段可以使用任何其他类型。

In your generated code, oneof fields have the same getters and setters as regular fields. You also get a special method for checking which value (if any) in the oneof is set. You can find out more about the oneof API for your chosen language in the relevant API reference.
> oneof字段与普通字段一样具有getters和setters。另外oneof还有一个特殊的方法用于检查哪个成员被设值了。你可以在相关语言的`API reference`中找到更多关于oneof的API。

## Oneof Features
Oneof的功能。

- Setting a oneof field will automatically clear all other members of the oneof. So if you set several oneof fields, only the last field you set will still have a value.
  > 设置一个oneof字段将自动清除oneof的所有其他成员。因此，如果你设置了几个oneof字段，只有你设置的最后一个字段仍然有值。

```cpp
SampleMessage message;
message.set_name("name");
CHECK(message.has_name());
// Calling mutable_sub_message() will clear the name field and will set
// sub_message to a new instance of SubMessage with none of its fields set
message.mutable_sub_message();
CHECK(!message.has_name());
```

- If the parser encounters multiple members of the same oneof on the wire, only the last member seen is used in the parsed message.
  > 如果解析器解析到同一个oneof的多个成员，则在解析后的消息中只使用最后一个看到的成员。

- A oneof cannot be repeated.
  > oneof类型字段不能作为重复字段使用。

- Reflection APIs work for oneof fields.
  > 反射API可以用于oneof字段。

- If you set a oneof field to the default value (such as setting an int32 oneof field to 0), the "case" of that oneof field will be set, and the value will be serialized on the wire.
  > 如果将一个oneof字段设置为默认值(例如将一个int32的oneof字段设置为0)，则该oneof字段的`case`将被设置，并且该值会被序列化。

- If you're using C++, make sure your code doesn't cause memory crashes. The following sample code will crash because sub_message was already deleted by calling the set_name() method.
  > 使用C++时，务必确保代码不会导致内存崩溃。下面的示例代码将会崩溃，因为通过调用`set_name()`方法已经删除了`sub_message`。

```cpp
SampleMessage message;
SubMessage* sub_message = message.mutable_sub_message();
message.set_name("name");      // Will delete sub_message
sub_message->set_...            // Crashes here
```

- Again in C++, if you Swap() two messages with oneofs, each message will end up with the other’s oneof case: in the example below, msg1 will have a sub_message and msg2 will have a name.
  > 同样在C++中，如果对两个oneof类型的消息执行`Swap()`操作，那么每个消息将以另一个消息的状态结束: 在下面的例子中，ms1最终具有成员`sub_message`，而ms2最终具有成员`name`。

```cpp
SampleMessage msg1;
msg1.set_name("name");
SampleMessage msg2;
msg2.mutable_sub_message();
msg1.swap(&msg2);
CHECK(msg1.has_sub_message());
CHECK(msg2.has_name());
```

## Backwards-compatibility issues
向后兼容性问题。

Be careful when adding or removing oneof fields. If checking the value of a oneof returns None/NOT_SET, it could mean that the oneof has not been set or it has been set to a field in a different version of the oneof. There is no way to tell the difference, since there's no way to know if an unknown field on the wire is a member of the oneof.
> 添加或者删除oneof字段一定要谨慎。如果检查oneof字段的值时返回的是`None/NOT_SET`，那么有可能是这个oneof字段并没有被设值，但也有可能是另一个版本的`.proto`中这个oneof有另一个成员且设置的是这个成员。没有办法辨别其中的区别，因为没有办法知道数据中的未知字段是否是oneof的成员。

Tag Reuse Issues
> 标签重用问题
> 
> 下面三个BUG应该是目前依然存在的问题，但我不明白的是，它们跟`Tag Reuse`是什么关系？跟`向后兼容性`又有什么关系？`Tag Reuse`是这些问题的共同特征吗？感觉这里语焉不详，也没有更详细的资料链接。。。

- Move fields into or out of a oneof: You may lose some of your information (some fields will be cleared) after the message is serialized and parsed. However, you can safely move a single field into a new oneof and may be able to move multiple fields if it is known that only one is ever set.
  > 如果对oneof字段做成员的添加或移除，那么在消息被序列化和解析之后，可能会丢失一些信息，一些字段会被清除。但是，可以安全地将一个单独的字段移动到一个新的oneof字段中。如果确切地知道某些字段始终只会设置其中一个字段的值，那么也可以安全地将这些字段加入一个新的oneof。
  > 
  > `and may be able to move multiple fields if it is known that only one is ever set`。。。到底应该怎么理解? 目前的翻译是个人理解，但感觉是废话，这就是什么时候可以创建一个新的oneof嘛。。。不知道特地在这里强调的意义何在，只是强调oneof的成员添加/移除的BUG不包括这两种新建oneof的情况？这有必要强调吗？
  > 
  > 另外，这个BUG跟`Tag Reuse`有啥关系？

- Delete a oneof field and add it back: This may clear your currently set oneof field after the message is serialized and parsed.
  > 删除一个oneof字段再将其添加回去: 这可能会在消息被序列化和解析后清除当前设置的oneof字段。
  > 
  > 所以这个BUG跟`Tag Reuse`有啥关系？再添加的时候，oneof字段名没变？字段编号呢？

- Split or merge oneof: This has similar issues to moving regular fields.
  > 拆分或合并oneof: 这与移动常规字段有类似的问题。
  >
  > 所以移动常规字段有什么问题呢？跟`Tag Reuse`有啥关系呢？

# Maps
映射集合。

If you want to create an associative map as part of your data definition, protocol buffers provides a handy shortcut syntax:
> 如果希望创建关联映射作为数据定义的一部分，protobuff提供了一种方便的快捷语法:

```protobuf
map<key_type, value_type> map_field = N;
```

...where the key_type can be any integral or string type (so, any scalar type except for floating point types and bytes). Note that enum is not a valid key_type. The value_type can be any type except another map.
> 其中键类型可以是任何整数类型或字符串类型(即，除了浮点类型和字节之外的任何标量类型)。注意，enum不是有效的键类型。值类型可以是除map之外的任何类型。

So, for example, if you wanted to create a map of projects where each Project message is associated with a string key, you could define it like this:
> 例如，创建一个项目映射，其中每个项目消息都与一个字符串键相关联，则可以像这样定义它:

```protobuf
map<string, Project> projects = 3;
```

- Map fields cannot be repeated.
  > map字段不能作为重复字段使用。

- Wire format ordering and map iteration ordering of map values are undefined, so you cannot rely on your map items being in a particular order.
  > map的values在二进制格式消息中的顺序，以及内存中的迭代顺序都没有定义，因此不能认为map中的元素有某种特定的顺序。

- When generating text format for a .proto, maps are sorted by key. Numeric keys are sorted numerically.
  > 当为`.proto`创建文本格式消息时，map会根据键值排序，对于数字类型的键值就按照数字顺序排序。
  > 
  > json格式消息是否认为是文本格式?

- When parsing from the wire or when merging, if there are duplicate map keys the last key seen is used. When parsing a map from text format, parsing may fail if there are duplicate keys.
  > map中的键值重复的话，在二进制消息和文本消息中的表现是不一样的。对于二进制消息的解析或合并，会使用重复键值中的后一个，覆盖前者。而对于文本格式消息的解析，键值重复会引起解析错误。

- If you provide a key but no value for a map field, the behavior when the field is serialized is language-dependent. In C++, Java, Kotlin, and Python the default value for the type is serialized, while in other languages nothing is serialized.
  > 如果向map字段提供了key却没有提供value，那么序列化该字段时的行为是依赖于具体语言的。在C++、Java、Kotlin和Python中，会根据value类型的默认值做序列化。而在其他语言中什么都没有序列化。
  >
  > `nothing is serialized`如何理解？是key和value都没有被序列化？还是只有value没有被序列化？讲道理，map里每个key就应该映射到一个非空的value，value为空的话就不应该有对应的key。所以我偏向于理解为连key都不存在，即key和value都没有序列化。

The generated map API is currently available for all proto3 supported languages. You can find out more about the map API for your chosen language in the relevant API reference.
> 目前map的API文档对所有proto3支持的语言都是有效的。你可以在对应语言的`API reference`里找到更多关于map的API信息。

## Backwards compatibility
向后兼容性。

The map syntax is equivalent to the following on the wire, so protocol buffers implementations that do not support maps can still handle your data:
> map在语法上等价于二进制消息上的下列写法，因此即使不支持map的protbuff也仍然能处理map数据:

```protobuf
message MapFieldEntry {
  key_type key = 1;
  value_type value = 2;
}

repeated MapFieldEntry map_field = N;
```

Any protocol buffers implementation that supports maps must both produce and accept data that can be accepted by the above definition.
> 任何支持map的protobuff实现都必须产生和接受上述定义可以接受的数据。


