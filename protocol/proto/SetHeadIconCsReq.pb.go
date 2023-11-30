// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SetHeadIconCsReq.proto

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

type SetHeadIconCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SetHeadIconCsReq) Reset() {
	*x = SetHeadIconCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetHeadIconCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHeadIconCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHeadIconCsReq) ProtoMessage() {}

func (x *SetHeadIconCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SetHeadIconCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHeadIconCsReq.ProtoReflect.Descriptor instead.
func (*SetHeadIconCsReq) Descriptor() ([]byte, []int) {
	return file_SetHeadIconCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SetHeadIconCsReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_SetHeadIconCsReq_proto protoreflect.FileDescriptor

var file_SetHeadIconCsReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x53, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetHeadIconCsReq_proto_rawDescOnce sync.Once
	file_SetHeadIconCsReq_proto_rawDescData = file_SetHeadIconCsReq_proto_rawDesc
)

func file_SetHeadIconCsReq_proto_rawDescGZIP() []byte {
	file_SetHeadIconCsReq_proto_rawDescOnce.Do(func() {
		file_SetHeadIconCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetHeadIconCsReq_proto_rawDescData)
	})
	return file_SetHeadIconCsReq_proto_rawDescData
}

var file_SetHeadIconCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetHeadIconCsReq_proto_goTypes = []interface{}{
	(*SetHeadIconCsReq)(nil), // 0: SetHeadIconCsReq
}
var file_SetHeadIconCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetHeadIconCsReq_proto_init() }
func file_SetHeadIconCsReq_proto_init() {
	if File_SetHeadIconCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetHeadIconCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHeadIconCsReq); i {
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
			RawDescriptor: file_SetHeadIconCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetHeadIconCsReq_proto_goTypes,
		DependencyIndexes: file_SetHeadIconCsReq_proto_depIdxs,
		MessageInfos:      file_SetHeadIconCsReq_proto_msgTypes,
	}.Build()
	File_SetHeadIconCsReq_proto = out.File
	file_SetHeadIconCsReq_proto_rawDesc = nil
	file_SetHeadIconCsReq_proto_goTypes = nil
	file_SetHeadIconCsReq_proto_depIdxs = nil
}
