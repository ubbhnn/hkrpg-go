// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerKickOutScNotify.proto

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

type PlayerKickOutScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlackInfo *BlackInfo `protobuf:"bytes,10,opt,name=black_info,json=blackInfo,proto3" json:"black_info,omitempty"`
	KickType  KickType   `protobuf:"varint,12,opt,name=kick_type,json=kickType,proto3,enum=KickType" json:"kick_type,omitempty"`
}

func (x *PlayerKickOutScNotify) Reset() {
	*x = PlayerKickOutScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerKickOutScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerKickOutScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerKickOutScNotify) ProtoMessage() {}

func (x *PlayerKickOutScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerKickOutScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerKickOutScNotify.ProtoReflect.Descriptor instead.
func (*PlayerKickOutScNotify) Descriptor() ([]byte, []int) {
	return file_PlayerKickOutScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerKickOutScNotify) GetBlackInfo() *BlackInfo {
	if x != nil {
		return x.BlackInfo
	}
	return nil
}

func (x *PlayerKickOutScNotify) GetKickType() KickType {
	if x != nil {
		return x.KickType
	}
	return KickType_KICK_SQUEEZED
}

var File_PlayerKickOutScNotify_proto protoreflect.FileDescriptor

var file_PlayerKickOutScNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x53,
	0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x42,
	0x6c, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x4b, 0x69, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a,
	0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x53,
	0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x29, 0x0a, 0x0a, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x26, 0x0a, 0x09, 0x6b, 0x69, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x6b, 0x69, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_PlayerKickOutScNotify_proto_rawDescOnce sync.Once
	file_PlayerKickOutScNotify_proto_rawDescData = file_PlayerKickOutScNotify_proto_rawDesc
)

func file_PlayerKickOutScNotify_proto_rawDescGZIP() []byte {
	file_PlayerKickOutScNotify_proto_rawDescOnce.Do(func() {
		file_PlayerKickOutScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerKickOutScNotify_proto_rawDescData)
	})
	return file_PlayerKickOutScNotify_proto_rawDescData
}

var file_PlayerKickOutScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerKickOutScNotify_proto_goTypes = []interface{}{
	(*PlayerKickOutScNotify)(nil), // 0: PlayerKickOutScNotify
	(*BlackInfo)(nil),             // 1: BlackInfo
	(KickType)(0),                 // 2: KickType
}
var file_PlayerKickOutScNotify_proto_depIdxs = []int32{
	1, // 0: PlayerKickOutScNotify.black_info:type_name -> BlackInfo
	2, // 1: PlayerKickOutScNotify.kick_type:type_name -> KickType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerKickOutScNotify_proto_init() }
func file_PlayerKickOutScNotify_proto_init() {
	if File_PlayerKickOutScNotify_proto != nil {
		return
	}
	file_BlackInfo_proto_init()
	file_KickType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerKickOutScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerKickOutScNotify); i {
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
			RawDescriptor: file_PlayerKickOutScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerKickOutScNotify_proto_goTypes,
		DependencyIndexes: file_PlayerKickOutScNotify_proto_depIdxs,
		MessageInfos:      file_PlayerKickOutScNotify_proto_msgTypes,
	}.Build()
	File_PlayerKickOutScNotify_proto = out.File
	file_PlayerKickOutScNotify_proto_rawDesc = nil
	file_PlayerKickOutScNotify_proto_goTypes = nil
	file_PlayerKickOutScNotify_proto_depIdxs = nil
}
