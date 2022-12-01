// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: asset/basic_asset.proto

package asset

import (
	owner "github.com/zhaochuninhefei/myproto-go/owner"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BasicAsset 基础资产
type BasicAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId    int64  `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	AssetName  string `protobuf:"bytes,2,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"`
	AssetPrice int64  `protobuf:"varint,3,opt,name=asset_price,json=assetPrice,proto3" json:"asset_price,omitempty"`
	// 注意这里引用owner.proto中的Owner时，前面要加上它的包名owner
	AssetOwner *owner.Owner `protobuf:"bytes,4,opt,name=asset_owner,json=assetOwner,proto3" json:"asset_owner,omitempty"`
	AssetDesc  string       `protobuf:"bytes,16,opt,name=asset_desc,json=assetDesc,proto3" json:"asset_desc,omitempty"`
	// oneof,只有一个字段会被实际使用
	//
	// Types that are assignable to AssetNum:
	//	*BasicAsset_AssetNumStr
	//	*BasicAsset_AssetNumInt
	AssetNum isBasicAsset_AssetNum `protobuf_oneof:"asset_num"`
}

func (x *BasicAsset) Reset() {
	*x = BasicAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_basic_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAsset) ProtoMessage() {}

func (x *BasicAsset) ProtoReflect() protoreflect.Message {
	mi := &file_asset_basic_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicAsset.ProtoReflect.Descriptor instead.
func (*BasicAsset) Descriptor() ([]byte, []int) {
	return file_asset_basic_asset_proto_rawDescGZIP(), []int{0}
}

func (x *BasicAsset) GetAssetId() int64 {
	if x != nil {
		return x.AssetId
	}
	return 0
}

func (x *BasicAsset) GetAssetName() string {
	if x != nil {
		return x.AssetName
	}
	return ""
}

func (x *BasicAsset) GetAssetPrice() int64 {
	if x != nil {
		return x.AssetPrice
	}
	return 0
}

func (x *BasicAsset) GetAssetOwner() *owner.Owner {
	if x != nil {
		return x.AssetOwner
	}
	return nil
}

func (x *BasicAsset) GetAssetDesc() string {
	if x != nil {
		return x.AssetDesc
	}
	return ""
}

func (m *BasicAsset) GetAssetNum() isBasicAsset_AssetNum {
	if m != nil {
		return m.AssetNum
	}
	return nil
}

func (x *BasicAsset) GetAssetNumStr() string {
	if x, ok := x.GetAssetNum().(*BasicAsset_AssetNumStr); ok {
		return x.AssetNumStr
	}
	return ""
}

func (x *BasicAsset) GetAssetNumInt() int64 {
	if x, ok := x.GetAssetNum().(*BasicAsset_AssetNumInt); ok {
		return x.AssetNumInt
	}
	return 0
}

type isBasicAsset_AssetNum interface {
	isBasicAsset_AssetNum()
}

type BasicAsset_AssetNumStr struct {
	AssetNumStr string `protobuf:"bytes,17,opt,name=asset_num_str,json=assetNumStr,proto3,oneof"`
}

type BasicAsset_AssetNumInt struct {
	AssetNumInt int64 `protobuf:"varint,18,opt,name=asset_num_int,json=assetNumInt,proto3,oneof"`
}

func (*BasicAsset_AssetNumStr) isBasicAsset_AssetNum() {}

func (*BasicAsset_AssetNumInt) isBasicAsset_AssetNum() {}

var File_asset_basic_asset_proto protoreflect.FileDescriptor

var file_asset_basic_asset_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x1a, 0x11, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a,
	0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x52, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x24, 0x0a, 0x0d, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x53, 0x74,
	0x72, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x69,
	0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x61, 0x6f, 0x63, 0x68, 0x75, 0x6e, 0x69, 0x6e, 0x68, 0x65, 0x66,
	0x65, 0x69, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_asset_basic_asset_proto_rawDescOnce sync.Once
	file_asset_basic_asset_proto_rawDescData = file_asset_basic_asset_proto_rawDesc
)

func file_asset_basic_asset_proto_rawDescGZIP() []byte {
	file_asset_basic_asset_proto_rawDescOnce.Do(func() {
		file_asset_basic_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_asset_basic_asset_proto_rawDescData)
	})
	return file_asset_basic_asset_proto_rawDescData
}

var file_asset_basic_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_asset_basic_asset_proto_goTypes = []interface{}{
	(*BasicAsset)(nil),  // 0: asset.BasicAsset
	(*owner.Owner)(nil), // 1: owner.Owner
}
var file_asset_basic_asset_proto_depIdxs = []int32{
	1, // 0: asset.BasicAsset.asset_owner:type_name -> owner.Owner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_asset_basic_asset_proto_init() }
func file_asset_basic_asset_proto_init() {
	if File_asset_basic_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_asset_basic_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicAsset); i {
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
	file_asset_basic_asset_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BasicAsset_AssetNumStr)(nil),
		(*BasicAsset_AssetNumInt)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_asset_basic_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_asset_basic_asset_proto_goTypes,
		DependencyIndexes: file_asset_basic_asset_proto_depIdxs,
		MessageInfos:      file_asset_basic_asset_proto_msgTypes,
	}.Build()
	File_asset_basic_asset_proto = out.File
	file_asset_basic_asset_proto_rawDesc = nil
	file_asset_basic_asset_proto_goTypes = nil
	file_asset_basic_asset_proto_depIdxs = nil
}
