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


