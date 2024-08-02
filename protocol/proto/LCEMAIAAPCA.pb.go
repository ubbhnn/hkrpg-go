// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LCEMAIAAPCA.proto

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

type LCEMAIAAPCA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogueId uint32 `protobuf:"varint,7,opt,name=dialogue_id,json=dialogueId,proto3" json:"dialogue_id,omitempty"`
}

func (x *LCEMAIAAPCA) Reset() {
	*x = LCEMAIAAPCA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LCEMAIAAPCA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LCEMAIAAPCA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LCEMAIAAPCA) ProtoMessage() {}

func (x *LCEMAIAAPCA) ProtoReflect() protoreflect.Message {
	mi := &file_LCEMAIAAPCA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LCEMAIAAPCA.ProtoReflect.Descriptor instead.
func (*LCEMAIAAPCA) Descriptor() ([]byte, []int) {
	return file_LCEMAIAAPCA_proto_rawDescGZIP(), []int{0}
}

func (x *LCEMAIAAPCA) GetDialogueId() uint32 {
	if x != nil {
		return x.DialogueId
	}
	return 0
}

var File_LCEMAIAAPCA_proto protoreflect.FileDescriptor

var file_LCEMAIAAPCA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x43, 0x45, 0x4d, 0x41, 0x49, 0x41, 0x41, 0x50, 0x43, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0b, 0x4c, 0x43, 0x45, 0x4d, 0x41, 0x49, 0x41, 0x41, 0x50,
	0x43, 0x41, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LCEMAIAAPCA_proto_rawDescOnce sync.Once
	file_LCEMAIAAPCA_proto_rawDescData = file_LCEMAIAAPCA_proto_rawDesc
)

func file_LCEMAIAAPCA_proto_rawDescGZIP() []byte {
	file_LCEMAIAAPCA_proto_rawDescOnce.Do(func() {
		file_LCEMAIAAPCA_proto_rawDescData = protoimpl.X.CompressGZIP(file_LCEMAIAAPCA_proto_rawDescData)
	})
	return file_LCEMAIAAPCA_proto_rawDescData
}

var file_LCEMAIAAPCA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LCEMAIAAPCA_proto_goTypes = []interface{}{
	(*LCEMAIAAPCA)(nil), // 0: LCEMAIAAPCA
}
var file_LCEMAIAAPCA_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LCEMAIAAPCA_proto_init() }
func file_LCEMAIAAPCA_proto_init() {
	if File_LCEMAIAAPCA_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LCEMAIAAPCA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LCEMAIAAPCA); i {
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
			RawDescriptor: file_LCEMAIAAPCA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LCEMAIAAPCA_proto_goTypes,
		DependencyIndexes: file_LCEMAIAAPCA_proto_depIdxs,
		MessageInfos:      file_LCEMAIAAPCA_proto_msgTypes,
	}.Build()
	File_LCEMAIAAPCA_proto = out.File
	file_LCEMAIAAPCA_proto_rawDesc = nil
	file_LCEMAIAAPCA_proto_goTypes = nil
	file_LCEMAIAAPCA_proto_depIdxs = nil
}