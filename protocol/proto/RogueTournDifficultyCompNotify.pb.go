// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournDifficultyCompNotify.proto

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

type RogueTournDifficultyCompNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JFEMDFAMLHL []uint32 `protobuf:"varint,6,rep,packed,name=JFEMDFAMLHL,proto3" json:"JFEMDFAMLHL,omitempty"`
}

func (x *RogueTournDifficultyCompNotify) Reset() {
	*x = RogueTournDifficultyCompNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournDifficultyCompNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournDifficultyCompNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournDifficultyCompNotify) ProtoMessage() {}

func (x *RogueTournDifficultyCompNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournDifficultyCompNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournDifficultyCompNotify.ProtoReflect.Descriptor instead.
func (*RogueTournDifficultyCompNotify) Descriptor() ([]byte, []int) {
	return file_RogueTournDifficultyCompNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournDifficultyCompNotify) GetJFEMDFAMLHL() []uint32 {
	if x != nil {
		return x.JFEMDFAMLHL
	}
	return nil
}

var File_RogueTournDifficultyCompNotify_proto protoreflect.FileDescriptor

var file_RogueTournDifficultyCompNotify_proto_rawDesc = []byte{
	0x0a, 0x24, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x44, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x43, 0x6f,
	0x6d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x46, 0x45, 0x4d,
	0x44, 0x46, 0x41, 0x4d, 0x4c, 0x48, 0x4c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4a,
	0x46, 0x45, 0x4d, 0x44, 0x46, 0x41, 0x4d, 0x4c, 0x48, 0x4c, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournDifficultyCompNotify_proto_rawDescOnce sync.Once
	file_RogueTournDifficultyCompNotify_proto_rawDescData = file_RogueTournDifficultyCompNotify_proto_rawDesc
)

func file_RogueTournDifficultyCompNotify_proto_rawDescGZIP() []byte {
	file_RogueTournDifficultyCompNotify_proto_rawDescOnce.Do(func() {
		file_RogueTournDifficultyCompNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournDifficultyCompNotify_proto_rawDescData)
	})
	return file_RogueTournDifficultyCompNotify_proto_rawDescData
}

var file_RogueTournDifficultyCompNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournDifficultyCompNotify_proto_goTypes = []interface{}{
	(*RogueTournDifficultyCompNotify)(nil), // 0: RogueTournDifficultyCompNotify
}
var file_RogueTournDifficultyCompNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueTournDifficultyCompNotify_proto_init() }
func file_RogueTournDifficultyCompNotify_proto_init() {
	if File_RogueTournDifficultyCompNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueTournDifficultyCompNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournDifficultyCompNotify); i {
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
			RawDescriptor: file_RogueTournDifficultyCompNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournDifficultyCompNotify_proto_goTypes,
		DependencyIndexes: file_RogueTournDifficultyCompNotify_proto_depIdxs,
		MessageInfos:      file_RogueTournDifficultyCompNotify_proto_msgTypes,
	}.Build()
	File_RogueTournDifficultyCompNotify_proto = out.File
	file_RogueTournDifficultyCompNotify_proto_rawDesc = nil
	file_RogueTournDifficultyCompNotify_proto_goTypes = nil
	file_RogueTournDifficultyCompNotify_proto_depIdxs = nil
}