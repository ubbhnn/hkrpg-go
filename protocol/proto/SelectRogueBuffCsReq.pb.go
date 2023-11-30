// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SelectRogueBuffCsReq.proto

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

type SelectRogueBuffCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MazeBuffId uint32 `protobuf:"varint,13,opt,name=maze_buff_id,json=mazeBuffId,proto3" json:"maze_buff_id,omitempty"`
}

func (x *SelectRogueBuffCsReq) Reset() {
	*x = SelectRogueBuffCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SelectRogueBuffCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectRogueBuffCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectRogueBuffCsReq) ProtoMessage() {}

func (x *SelectRogueBuffCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SelectRogueBuffCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectRogueBuffCsReq.ProtoReflect.Descriptor instead.
func (*SelectRogueBuffCsReq) Descriptor() ([]byte, []int) {
	return file_SelectRogueBuffCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SelectRogueBuffCsReq) GetMazeBuffId() uint32 {
	if x != nil {
		return x.MazeBuffId
	}
	return 0
}

var File_SelectRogueBuffCsReq_proto protoreflect.FileDescriptor

var file_SelectRogueBuffCsReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66,
	0x66, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x14,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x7a, 0x65, 0x5f, 0x62, 0x75, 0x66,
	0x66, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x7a, 0x65,
	0x42, 0x75, 0x66, 0x66, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SelectRogueBuffCsReq_proto_rawDescOnce sync.Once
	file_SelectRogueBuffCsReq_proto_rawDescData = file_SelectRogueBuffCsReq_proto_rawDesc
)

func file_SelectRogueBuffCsReq_proto_rawDescGZIP() []byte {
	file_SelectRogueBuffCsReq_proto_rawDescOnce.Do(func() {
		file_SelectRogueBuffCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SelectRogueBuffCsReq_proto_rawDescData)
	})
	return file_SelectRogueBuffCsReq_proto_rawDescData
}

var file_SelectRogueBuffCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SelectRogueBuffCsReq_proto_goTypes = []interface{}{
	(*SelectRogueBuffCsReq)(nil), // 0: SelectRogueBuffCsReq
}
var file_SelectRogueBuffCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SelectRogueBuffCsReq_proto_init() }
func file_SelectRogueBuffCsReq_proto_init() {
	if File_SelectRogueBuffCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SelectRogueBuffCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectRogueBuffCsReq); i {
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
			RawDescriptor: file_SelectRogueBuffCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SelectRogueBuffCsReq_proto_goTypes,
		DependencyIndexes: file_SelectRogueBuffCsReq_proto_depIdxs,
		MessageInfos:      file_SelectRogueBuffCsReq_proto_msgTypes,
	}.Build()
	File_SelectRogueBuffCsReq_proto = out.File
	file_SelectRogueBuffCsReq_proto_rawDesc = nil
	file_SelectRogueBuffCsReq_proto_goTypes = nil
	file_SelectRogueBuffCsReq_proto_depIdxs = nil
}
