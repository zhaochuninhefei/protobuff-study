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




