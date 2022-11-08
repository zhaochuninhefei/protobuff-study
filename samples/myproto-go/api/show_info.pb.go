// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: api/show_info.proto

package api

import (
	asset "github.com/zhaochuninhefei/myproto-go/asset"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShowInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoId   int64                  `protobuf:"varint,1,opt,name=info_id,json=infoId,proto3" json:"info_id,omitempty"`
	Assets   []*asset.BasicAsset    `protobuf:"bytes,2,rep,name=assets,proto3" json:"assets,omitempty"`
	ShowTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=show_time,json=showTime,proto3" json:"show_time,omitempty"`
}

func (x *ShowInfo) Reset() {
	*x = ShowInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_show_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowInfo) ProtoMessage() {}

func (x *ShowInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_show_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowInfo.ProtoReflect.Descriptor instead.
func (*ShowInfo) Descriptor() ([]byte, []int) {
	return file_api_show_info_proto_rawDescGZIP(), []int{0}
}

func (x *ShowInfo) GetInfoId() int64 {
	if x != nil {
		return x.InfoId
	}
	return 0
}

func (x *ShowInfo) GetAssets() []*asset.BasicAsset {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *ShowInfo) GetShowTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ShowTime
	}
	return nil
}

type ShowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *ShowRequest) Reset() {
	*x = ShowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_show_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowRequest) ProtoMessage() {}

func (x *ShowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_show_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowRequest.ProtoReflect.Descriptor instead.
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return file_api_show_info_proto_rawDescGZIP(), []int{1}
}

func (x *ShowRequest) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

var File_api_show_info_proto protoreflect.FileDescriptor

var file_api_show_info_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x69, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x06, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2c,
	0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x32, 0x36, 0x0a, 0x0b,
	0x53, 0x68, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x53,
	0x68, 0x6f, 0x77, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x77,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x61, 0x6f, 0x63, 0x68, 0x75, 0x6e, 0x69, 0x6e, 0x68, 0x65, 0x66,
	0x65, 0x69, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_show_info_proto_rawDescOnce sync.Once
	file_api_show_info_proto_rawDescData = file_api_show_info_proto_rawDesc
)

func file_api_show_info_proto_rawDescGZIP() []byte {
	file_api_show_info_proto_rawDescOnce.Do(func() {
		file_api_show_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_show_info_proto_rawDescData)
	})
	return file_api_show_info_proto_rawDescData
}

var file_api_show_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_show_info_proto_goTypes = []interface{}{
	(*ShowInfo)(nil),              // 0: api.ShowInfo
	(*ShowRequest)(nil),           // 1: api.ShowRequest
	(*asset.BasicAsset)(nil),      // 2: asset.BasicAsset
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_api_show_info_proto_depIdxs = []int32{
	2, // 0: api.ShowInfo.assets:type_name -> asset.BasicAsset
	3, // 1: api.ShowInfo.show_time:type_name -> google.protobuf.Timestamp
	1, // 2: api.ShowService.Show:input_type -> api.ShowRequest
	0, // 3: api.ShowService.Show:output_type -> api.ShowInfo
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_show_info_proto_init() }
func file_api_show_info_proto_init() {
	if File_api_show_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_show_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_show_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_show_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_show_info_proto_goTypes,
		DependencyIndexes: file_api_show_info_proto_depIdxs,
		MessageInfos:      file_api_show_info_proto_msgTypes,
	}.Build()
	File_api_show_info_proto = out.File
	file_api_show_info_proto_rawDesc = nil
	file_api_show_info_proto_goTypes = nil
	file_api_show_info_proto_depIdxs = nil
}
