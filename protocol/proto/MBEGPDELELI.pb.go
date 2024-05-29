// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MBEGPDELELI.proto

package proto

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

type MBEGPDELELI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiracleInfo *GameRogueMiracle `protobuf:"bytes,9,opt,name=miracle_info,json=miracleInfo,proto3" json:"miracle_info,omitempty"`
}

func (x *MBEGPDELELI) Reset() {
	*x = MBEGPDELELI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MBEGPDELELI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MBEGPDELELI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MBEGPDELELI) ProtoMessage() {}

func (x *MBEGPDELELI) ProtoReflect() protoreflect.Message {
	mi := &file_MBEGPDELELI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MBEGPDELELI.ProtoReflect.Descriptor instead.
func (*MBEGPDELELI) Descriptor() ([]byte, []int) {
	return file_MBEGPDELELI_proto_rawDescGZIP(), []int{0}
}

func (x *MBEGPDELELI) GetMiracleInfo() *GameRogueMiracle {
	if x != nil {
		return x.MiracleInfo
	}
	return nil
}

var File_MBEGPDELELI_proto protoreflect.FileDescriptor

var file_MBEGPDELELI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x42, 0x45, 0x47, 0x50, 0x44, 0x45, 0x4c, 0x45, 0x4c, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0b, 0x4d,
	0x42, 0x45, 0x47, 0x50, 0x44, 0x45, 0x4c, 0x45, 0x4c, 0x49, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x52, 0x0b, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_MBEGPDELELI_proto_rawDescOnce sync.Once
	file_MBEGPDELELI_proto_rawDescData = file_MBEGPDELELI_proto_rawDesc
)

func file_MBEGPDELELI_proto_rawDescGZIP() []byte {
	file_MBEGPDELELI_proto_rawDescOnce.Do(func() {
		file_MBEGPDELELI_proto_rawDescData = protoimpl.X.CompressGZIP(file_MBEGPDELELI_proto_rawDescData)
	})
	return file_MBEGPDELELI_proto_rawDescData
}

var file_MBEGPDELELI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MBEGPDELELI_proto_goTypes = []interface{}{
	(*MBEGPDELELI)(nil),      // 0: MBEGPDELELI
	(*GameRogueMiracle)(nil), // 1: GameRogueMiracle
}
var file_MBEGPDELELI_proto_depIdxs = []int32{
	1, // 0: MBEGPDELELI.miracle_info:type_name -> GameRogueMiracle
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MBEGPDELELI_proto_init() }
func file_MBEGPDELELI_proto_init() {
	if File_MBEGPDELELI_proto != nil {
		return
	}
	file_GameRogueMiracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MBEGPDELELI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MBEGPDELELI); i {
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
			RawDescriptor: file_MBEGPDELELI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MBEGPDELELI_proto_goTypes,
		DependencyIndexes: file_MBEGPDELELI_proto_depIdxs,
		MessageInfos:      file_MBEGPDELELI_proto_msgTypes,
	}.Build()
	File_MBEGPDELELI_proto = out.File
	file_MBEGPDELELI_proto_rawDesc = nil
	file_MBEGPDELELI_proto_goTypes = nil
	file_MBEGPDELELI_proto_depIdxs = nil
}