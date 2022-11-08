Protocol Buffer Basics: Go
=====
This tutorial provides a basic Go programmer's introduction to working with protocol buffers, using the proto3 version of the protocol buffers language. By walking through creating a simple example application, it shows you how to
> 本教程为Go程序员提供了使用protobuf的proto3版本的基本介绍。通过创建一个简单的示例应用程序，本文向您展示了如何做到这一点:

- Define message formats in a .proto file.
    > 在`.proto`文件中定义消息格式。

- Use the protocol buffer compiler.
    > 使用protobuf编译器。

- Use the Go protocol buffer API to write and read messages.
    > 使用Go的protobuf的API读写消息。

This isn't a comprehensive guide to using protocol buffers in Go. For more detailed reference information, see the Protocol Buffer Language Guide, the Go API Reference, the Go Generated Code Guide, and the Encoding Reference.
> 这不是一个关于在Go中使用protobuf的全面指南。要了解更多详细的参考信息，请参见`Protocol Buffer Language Guide`、`Go API Reference`、`Go Generated Code Guide`、以及`Encoding Reference`。

# Why use protocol buffers?
为什么要使用protobuf?

The example we're going to use is a very simple "address book" application that can read and write people's contact details to and from a file. Each person in the address book has a name, an ID, an email address, and a contact phone number.
> 我们将要使用的示例是一个非常简单的“地址簿”应用程序，它可以从文件中读写人们的联系方式。地址簿中的每个人都有姓名、ID、电子邮件地址和联系电话号码。

How do you serialize and retrieve structured data like this? There are a few ways to solve this problem:
> 如何序列化和检索这样的结构化数据?有几种方法可以解决这个问题?

- Use gobs to serialize Go data structures. This is a good solution in a Go-specific environment, but it doesn't work well if you need to share data with applications written for other platforms.
    > 使用gobs序列化Go数据结构。在特定于go的环境中，这是一个很好的解决方案，但如果需要与为其他平台编写的应用程序共享数据，则不能很好地工作。

- You can invent an ad-hoc way to encode the data items into a single string – such as encoding 4 ints as "12:3:-23:67". This is a simple and flexible approach, although it does require writing one-off encoding and parsing code, and the parsing imposes a small run-time cost. This works best for encoding very simple data.
    > 您可以发明一种特别的方法将数据项编码为单个字符串，例如将4个整数编码为`12:3:-23:67`。这是一种简单而灵活的方法，但它需要设计并实现一次性的编码和解析代码，而且解析会带来一些的运行时成本。这种方法最好在编码非常简单的数据时使用。
    > 
    > 我可不想自己设计实现一套编码规则。。。

- Serialize the data to XML. This approach can be very attractive since XML is (sort of) human readable and there are binding libraries for lots of languages. This can be a good choice if you want to share data with other applications/projects. However, XML is notoriously space intensive, and encoding/decoding it can impose a huge performance penalty on applications. Also, navigating an XML DOM tree is considerably more complicated than navigating simple fields in a class normally would be.
    > 把数据序列化为XML。这种方法非常有吸引力，因为XML(在某种程度上)是人类可读的，而且许多语言都提供了相关库。如果需要与其他应用/项目共享数据，这可能是一个不错的方案。但是，众所周知，XML非常消耗存储空间的，并且其编码/解码会给应用程序带来巨大的性能消耗。此外，在XML的DOM树实现导航功能(即根据字段读写其值)通常要比在类中访问字段要复杂得多。

Protocol buffers are the flexible, efficient, automated solution to solve exactly this problem. With protocol buffers, you write a .proto description of the data structure you wish to store. From that, the protocol buffer compiler creates a class that implements automatic encoding and parsing of the protocol buffer data with an efficient binary format. The generated class provides getters and setters for the fields that make up a protocol buffer and takes care of the details of reading and writing the protocol buffer as a unit. Importantly, the protocol buffer format supports the idea of extending the format over time in such a way that the code can still read data encoded with the old format.
> protobuf恰好可以灵活、高效、自动化地解决这些问题。使用protobuf，可以编写一个`.proto`用来描述需要存储的数据结构。基于这个`.proto`文件，protobuf编译器会生成对应的类，这些类会使用高效的二进制格式实现protobuf数据的编码和解析。生成的类为对应的protobuf数据中的字段提供getter和setter操作，并负责将protobuf数据作为一个单元进行读写。重要的是，protobuf的格式支持扩展，扩展以后，那些用旧格式编码的数据仍然能够被读取。

# Where to find the example code
去哪找示例代码?

Our example is a set of command-line applications for managing an address book data file, encoded using protocol buffers. The command add_person_go adds a new entry to the data file. The command list_people_go parses the data file and prints the data to the console.
> 我们的示例是一组命令行应用程序，用于管理使用protobuf编码的地址簿数据文件。命令`add_person_go`向数据文件添加一个新条目。命令`list_people_go`解析数据文件并将数据打印到控制台。

You can find the complete example in the examples directory of the GitHub repository.
> 您可以在GitHub存储库的examples目录中找到完整的示例。

# Defining your protocol format
定义proto格式。

To create your address book application, you'll need to start with a .proto file. The definitions in a .proto file are simple: you add a message for each data structure you want to serialize, then specify a name and a type for each field in the message. In our example, the .proto file that defines the messages is addressbook.proto.
> 要创建地址簿应用程序，需要从`.proto`文件开始。proto文件中的定义很简单: 为想要序列化的每个数据结构添加一条消息，然后为消息中的每个字段指定名称和类型。在我们的示例中，定义消息的`.proto`文件是`addressbook.proto`。

The .proto file starts with a package declaration, which helps to prevent naming conflicts between different projects.
> proto文件以一个包声明开始，这有助于防止不同项目之间的命名冲突。

```protobuf
syntax = "proto3";
package tutorial;

import "google/protobuf/timestamp.proto";
```

The go_package option defines the import path of the package which will contain all the generated code for this file. The Go package name will be the last path component of the import path. For example, our example will use a package name of "tutorialpb".
> `go_package`选项定义了包的导入路径，该路径会包含所有为该文件生成的代码。Go的包名将是导入路径的最后一个路径组件。例如，我们的示例将使用名为“tutorialpb”的包。

```protobuf
option go_package = "github.com/protocolbuffers/protobuf/examples/go/tutorialpb";
```

Next, you have your message definitions. A message is just an aggregate containing a set of typed fields. Many standard simple data types are available as field types, including bool, int32, float, double, and string. You can also add further structure to your messages by using other message types as field types.
> 接下来添加消息定义。消息只是包含一组类型化字段的聚合。许多标准的简单数据类型都可以作为字段类型使用，包括bool、int32、float、double和string。还可以使用其他消息类型作为字段类型，从而向消息添加进一步的结构。

```protobuf
message Person {
  string name = 1;
  int32 id = 2;  // Unique ID number for this person.
  string email = 3;

  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }

  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }

  repeated PhoneNumber phones = 4;

  google.protobuf.Timestamp last_updated = 5;
}

// Our address book file is just one of these.
message AddressBook {
  repeated Person people = 1;
}
```

In the above example, the Person message contains PhoneNumber messages, while the AddressBook message contains Person messages. You can even define message types nested inside other messages – as you can see, the PhoneNumber type is defined inside Person. You can also define enum types if you want one of your fields to have one of a predefined list of values – here you want to specify that a phone number can be one of MOBILE, HOME, or WORK.
> 在上面的示例中，Person消息包含PhoneNumber消息，而AddressBook消息包含Person消息。您甚至可以定义嵌套在其他消息中的消息类型，如您所见，PhoneNumber类型是在Person中定义的。如果您希望其中一个字段具有预定义的值列表中的一个，您还可以定义枚举类型。比如上面的枚举`PhoneType`，指定电话号码的类型有`MOBILE`、`HOME`或`WORK`。

The " = 1", " = 2" markers on each element identify the unique "tag" that field uses in the binary encoding. Tag numbers 1-15 require one less byte to encode than higher numbers, so as an optimization you can decide to use those tags for the commonly used or repeated elements, leaving tags 16 and higher for less-commonly used optional elements. Each element in a repeated field requires re-encoding the tag number, so repeated fields are particularly good candidates for this optimization.
> 每个元素上的`= 1`，`= 2`是字段编号，是二进制编码中每个字段的唯一标记。字段编号`1-15`区间的数字需要的编码字节少一个，因此作为一种优化，一般建议将这些编号用于常用的或重复的元素，而将`16`以后的编号用于不常用的可选元素。重复字段中的每个元素都需要对编号进行重新编码，因此重复字段特别适合这种优化。

If a field value isn't set, a default value is used: zero for numeric types, the empty string for strings, false for bools. For embedded messages, the default value is always the "default instance" or "prototype" of the message, which has none of its fields set. Calling the accessor to get the value of a field which has not been explicitly set always returns that field's default value.
> 如果未设置字段值，则使用默认值: 数值类型为0，字符串为空字符串，boolean为false。对于嵌入式消息，默认值总是消息的“默认实例”或“原型”，其中没有设置任何字段。调用访问器来获取未显式设置的字段的值总是返回该字段的默认值。

If a field is repeated, the field may be repeated any number of times (including zero). The order of the repeated values will be preserved in the protocol buffer. Think of repeated fields as dynamically sized arrays.
> 对于一个重复字段，该字段可以重复任何次数(包括0)。重复值的顺序将保存在protobuf中。可以将重复字段看作动态大小的数组。

You'll find a complete guide to writing .proto files – including all the possible field types – in the Protocol Buffer Language Guide. Don't go looking for facilities similar to class inheritance, though – protocol buffers don't do that.
> 您将在`Protocol Buffer Language Guide`中找到一个完整的编写`.proto`文件的指南，其中包括所有可能的字段类型。不要去寻找类似于类继承的工具，protobuf不这么做。


# Compiling your protocol buffers
编译protobuf。

Now that you have a .proto, the next thing you need to do is generate the classes you'll need to read and write AddressBook (and hence Person and PhoneNumber) messages. To do this, you need to run the protocol buffer compiler protoc on your .proto:
> 现在您有了`.proto`，接下来需要做的是生成相关的类，来读取和写入AddressBook(以及Person和PhoneNumber)消息。为此，需要在`.proto`上运行protobuf编译器`protoc`:

1.If you haven't installed the compiler, download the package and follow the instructions in the README.
> 如果您还没有安装编译器，请下载该包并按照README中的说明操作。

2.Run the following command to install the Go protocol buffers plugin:
> 运行以下命令安装Go的protobuf插件:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

The compiler plugin protoc-gen-go will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler protoc to find it.
> 编译器插件`protoc-gen-go`会被安装在环境变量`$GOBIN`的目录下，默认是`$GOPATH/bin`。注意，这个路径需要加入你的本地环境变量`$PATH`，确保protobuf编译器protoc能够找到它。

3.Now run the compiler, specifying the source directory (where your application's source code lives – the current directory is used if you don't provide a value), the destination directory (where you want the generated code to go; often the same as $SRC_DIR), and the path to your .proto. In this case, you would invoke:
> 现在运行编译器，指定源目录(应用程序源代码所在的位置，如果不提供值则使用当前目录)、目标目录(您希望生成的代码存放的位置;通常与`$SRC_DIR`相同)，以及`.proto`的路径。在本例中，您将调用:

```bash
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```

Because you want Go code, you use the --go_out option – similar options are provided for other supported languages.
> 因为这里是要生成Go的代码，因此使用`--go_out`选项，类似选项用来支持其他语言。

This generates github.com/protocolbuffers/protobuf/examples/go/tutorialpb/addressbook.pb.go in your specified destination directory.
> 这里会生成`github.com/protocolbuffers/protobuf/examples/go/tutorialpb/addressbook.pb.go`代码文件到指定的目标目录。





