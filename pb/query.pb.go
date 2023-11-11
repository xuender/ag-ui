// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pb/query.proto

package pb

import (
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

type Searcher int32

const (
	Searcher_rg Searcher = 0
	Searcher_ug Searcher = 1
	Searcher_ag Searcher = 2
)

// Enum value maps for Searcher.
var (
	Searcher_name = map[int32]string{
		0: "rg",
		1: "ug",
		2: "ag",
	}
	Searcher_value = map[string]int32{
		"rg": 0,
		"ug": 1,
		"ag": 2,
	}
)

func (x Searcher) Enum() *Searcher {
	p := new(Searcher)
	*p = x
	return p
}

func (x Searcher) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Searcher) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_query_proto_enumTypes[0].Descriptor()
}

func (Searcher) Type() protoreflect.EnumType {
	return &file_pb_query_proto_enumTypes[0]
}

func (x Searcher) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Searcher.Descriptor instead.
func (Searcher) EnumDescriptor() ([]byte, []int) {
	return file_pb_query_proto_rawDescGZIP(), []int{0}
}

// Query .
type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern  string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	MaxCount uint32   `protobuf:"varint,2,opt,name=maxCount,proto3" json:"maxCount,omitempty"`
	Searcher Searcher `protobuf:"varint,3,opt,name=searcher,proto3,enum=pb.Searcher" json:"searcher,omitempty"`
	Paths    []string `protobuf:"bytes,4,rep,name=paths,proto3" json:"paths,omitempty"`
	RgTypes  []string `protobuf:"bytes,5,rep,name=rgTypes,proto3" json:"rgTypes,omitempty"`
	AgTypes  []string `protobuf:"bytes,6,rep,name=agTypes,proto3" json:"agTypes,omitempty"`
	UgTypes  []string `protobuf:"bytes,7,rep,name=ugTypes,proto3" json:"ugTypes,omitempty"`
	GrepType string   `protobuf:"bytes,8,opt,name=grepType,proto3" json:"grepType,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_pb_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_pb_query_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *Query) GetMaxCount() uint32 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

func (x *Query) GetSearcher() Searcher {
	if x != nil {
		return x.Searcher
	}
	return Searcher_rg
}

func (x *Query) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *Query) GetRgTypes() []string {
	if x != nil {
		return x.RgTypes
	}
	return nil
}

func (x *Query) GetAgTypes() []string {
	if x != nil {
		return x.AgTypes
	}
	return nil
}

func (x *Query) GetUgTypes() []string {
	if x != nil {
		return x.UgTypes
	}
	return nil
}

func (x *Query) GetGrepType() string {
	if x != nil {
		return x.GrepType
	}
	return ""
}

var File_pb_query_proto protoreflect.FileDescriptor

var file_pb_query_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0xe7, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x22,
	0x0a, 0x08, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x72, 0x67,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x75, 0x67, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x61, 0x67,
	0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_query_proto_rawDescOnce sync.Once
	file_pb_query_proto_rawDescData = file_pb_query_proto_rawDesc
)

func file_pb_query_proto_rawDescGZIP() []byte {
	file_pb_query_proto_rawDescOnce.Do(func() {
		file_pb_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_query_proto_rawDescData)
	})
	return file_pb_query_proto_rawDescData
}

var file_pb_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_query_proto_goTypes = []interface{}{
	(Searcher)(0), // 0: pb.Searcher
	(*Query)(nil), // 1: pb.Query
}
var file_pb_query_proto_depIdxs = []int32{
	0, // 0: pb.Query.searcher:type_name -> pb.Searcher
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_query_proto_init() }
func file_pb_query_proto_init() {
	if File_pb_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
			RawDescriptor: file_pb_query_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_query_proto_goTypes,
		DependencyIndexes: file_pb_query_proto_depIdxs,
		EnumInfos:         file_pb_query_proto_enumTypes,
		MessageInfos:      file_pb_query_proto_msgTypes,
	}.Build()
	File_pb_query_proto = out.File
	file_pb_query_proto_rawDesc = nil
	file_pb_query_proto_goTypes = nil
	file_pb_query_proto_depIdxs = nil
}
