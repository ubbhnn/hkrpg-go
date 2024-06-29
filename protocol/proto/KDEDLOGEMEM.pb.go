// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: KDEDLOGEMEM.proto

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

type KDEDLOGEMEM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	POCPHLJELHH []*FPONKPLDELJ `protobuf:"bytes,1,rep,name=POCPHLJELHH,proto3" json:"POCPHLJELHH,omitempty"`
	ReviveInfo  *IBHFIGDHELO   `protobuf:"bytes,6,opt,name=revive_info,json=reviveInfo,proto3" json:"revive_info,omitempty"`
}

func (x *KDEDLOGEMEM) Reset() {
	*x = KDEDLOGEMEM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KDEDLOGEMEM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KDEDLOGEMEM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KDEDLOGEMEM) ProtoMessage() {}

func (x *KDEDLOGEMEM) ProtoReflect() protoreflect.Message {
	mi := &file_KDEDLOGEMEM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KDEDLOGEMEM.ProtoReflect.Descriptor instead.
func (*KDEDLOGEMEM) Descriptor() ([]byte, []int) {
	return file_KDEDLOGEMEM_proto_rawDescGZIP(), []int{0}
}

func (x *KDEDLOGEMEM) GetPOCPHLJELHH() []*FPONKPLDELJ {
	if x != nil {
		return x.POCPHLJELHH
	}
	return nil
}

func (x *KDEDLOGEMEM) GetReviveInfo() *IBHFIGDHELO {
	if x != nil {
		return x.ReviveInfo
	}
	return nil
}

var File_KDEDLOGEMEM_proto protoreflect.FileDescriptor

var file_KDEDLOGEMEM_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4b, 0x44, 0x45, 0x44, 0x4c, 0x4f, 0x47, 0x45, 0x4d, 0x45, 0x4d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x42, 0x48, 0x46, 0x49, 0x47, 0x44, 0x48, 0x45, 0x4c, 0x4f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x50, 0x4f, 0x4e, 0x4b, 0x50, 0x4c, 0x44,
	0x45, 0x4c, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0b, 0x4b, 0x44, 0x45,
	0x44, 0x4c, 0x4f, 0x47, 0x45, 0x4d, 0x45, 0x4d, 0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x4f, 0x43, 0x50,
	0x48, 0x4c, 0x4a, 0x45, 0x4c, 0x48, 0x48, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x46, 0x50, 0x4f, 0x4e, 0x4b, 0x50, 0x4c, 0x44, 0x45, 0x4c, 0x4a, 0x52, 0x0b, 0x50, 0x4f, 0x43,
	0x50, 0x48, 0x4c, 0x4a, 0x45, 0x4c, 0x48, 0x48, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69,
	0x76, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x49, 0x42, 0x48, 0x46, 0x49, 0x47, 0x44, 0x48, 0x45, 0x4c, 0x4f, 0x52, 0x0a, 0x72, 0x65, 0x76,
	0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KDEDLOGEMEM_proto_rawDescOnce sync.Once
	file_KDEDLOGEMEM_proto_rawDescData = file_KDEDLOGEMEM_proto_rawDesc
)

func file_KDEDLOGEMEM_proto_rawDescGZIP() []byte {
	file_KDEDLOGEMEM_proto_rawDescOnce.Do(func() {
		file_KDEDLOGEMEM_proto_rawDescData = protoimpl.X.CompressGZIP(file_KDEDLOGEMEM_proto_rawDescData)
	})
	return file_KDEDLOGEMEM_proto_rawDescData
}

var file_KDEDLOGEMEM_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KDEDLOGEMEM_proto_goTypes = []interface{}{
	(*KDEDLOGEMEM)(nil), // 0: KDEDLOGEMEM
	(*FPONKPLDELJ)(nil), // 1: FPONKPLDELJ
	(*IBHFIGDHELO)(nil), // 2: IBHFIGDHELO
}
var file_KDEDLOGEMEM_proto_depIdxs = []int32{
	1, // 0: KDEDLOGEMEM.POCPHLJELHH:type_name -> FPONKPLDELJ
	2, // 1: KDEDLOGEMEM.revive_info:type_name -> IBHFIGDHELO
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_KDEDLOGEMEM_proto_init() }
func file_KDEDLOGEMEM_proto_init() {
	if File_KDEDLOGEMEM_proto != nil {
		return
	}
	file_IBHFIGDHELO_proto_init()
	file_FPONKPLDELJ_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_KDEDLOGEMEM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KDEDLOGEMEM); i {
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
			RawDescriptor: file_KDEDLOGEMEM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KDEDLOGEMEM_proto_goTypes,
		DependencyIndexes: file_KDEDLOGEMEM_proto_depIdxs,
		MessageInfos:      file_KDEDLOGEMEM_proto_msgTypes,
	}.Build()
	File_KDEDLOGEMEM_proto = out.File
	file_KDEDLOGEMEM_proto_rawDesc = nil
	file_KDEDLOGEMEM_proto_goTypes = nil
	file_KDEDLOGEMEM_proto_depIdxs = nil
}