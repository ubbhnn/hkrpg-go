// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetRogueExhibitionScRsp.proto

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

type SetRogueExhibitionScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JDBJENNKOBK []*EMIECBDCBNM `protobuf:"bytes,14,rep,name=JDBJENNKOBK,proto3" json:"JDBJENNKOBK,omitempty"`
	FMKFFPMHCPE []*GGPEHIBFFCB `protobuf:"bytes,13,rep,name=FMKFFPMHCPE,proto3" json:"FMKFFPMHCPE,omitempty"`
	Retcode     uint32         `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SetRogueExhibitionScRsp) Reset() {
	*x = SetRogueExhibitionScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetRogueExhibitionScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRogueExhibitionScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRogueExhibitionScRsp) ProtoMessage() {}

func (x *SetRogueExhibitionScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetRogueExhibitionScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRogueExhibitionScRsp.ProtoReflect.Descriptor instead.
func (*SetRogueExhibitionScRsp) Descriptor() ([]byte, []int) {
	return file_SetRogueExhibitionScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetRogueExhibitionScRsp) GetJDBJENNKOBK() []*EMIECBDCBNM {
	if x != nil {
		return x.JDBJENNKOBK
	}
	return nil
}

func (x *SetRogueExhibitionScRsp) GetFMKFFPMHCPE() []*GGPEHIBFFCB {
	if x != nil {
		return x.FMKFFPMHCPE
	}
	return nil
}

func (x *SetRogueExhibitionScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SetRogueExhibitionScRsp_proto protoreflect.FileDescriptor

var file_SetRogueExhibitionScRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x47, 0x47, 0x50, 0x45, 0x48, 0x49, 0x42, 0x46, 0x46, 0x43, 0x42, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4d, 0x49, 0x45, 0x43, 0x42, 0x44, 0x43, 0x42, 0x4e, 0x4d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x52, 0x73,
	0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x44, 0x42, 0x4a, 0x45, 0x4e, 0x4e, 0x4b, 0x4f, 0x42, 0x4b,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x4d, 0x49, 0x45, 0x43, 0x42, 0x44,
	0x43, 0x42, 0x4e, 0x4d, 0x52, 0x0b, 0x4a, 0x44, 0x42, 0x4a, 0x45, 0x4e, 0x4e, 0x4b, 0x4f, 0x42,
	0x4b, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x4d, 0x4b, 0x46, 0x46, 0x50, 0x4d, 0x48, 0x43, 0x50, 0x45,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x47, 0x50, 0x45, 0x48, 0x49, 0x42,
	0x46, 0x46, 0x43, 0x42, 0x52, 0x0b, 0x46, 0x4d, 0x4b, 0x46, 0x46, 0x50, 0x4d, 0x48, 0x43, 0x50,
	0x45, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_SetRogueExhibitionScRsp_proto_rawDescOnce sync.Once
	file_SetRogueExhibitionScRsp_proto_rawDescData = file_SetRogueExhibitionScRsp_proto_rawDesc
)

func file_SetRogueExhibitionScRsp_proto_rawDescGZIP() []byte {
	file_SetRogueExhibitionScRsp_proto_rawDescOnce.Do(func() {
		file_SetRogueExhibitionScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetRogueExhibitionScRsp_proto_rawDescData)
	})
	return file_SetRogueExhibitionScRsp_proto_rawDescData
}

var file_SetRogueExhibitionScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetRogueExhibitionScRsp_proto_goTypes = []interface{}{
	(*SetRogueExhibitionScRsp)(nil), // 0: SetRogueExhibitionScRsp
	(*EMIECBDCBNM)(nil),             // 1: EMIECBDCBNM
	(*GGPEHIBFFCB)(nil),             // 2: GGPEHIBFFCB
}
var file_SetRogueExhibitionScRsp_proto_depIdxs = []int32{
	1, // 0: SetRogueExhibitionScRsp.JDBJENNKOBK:type_name -> EMIECBDCBNM
	2, // 1: SetRogueExhibitionScRsp.FMKFFPMHCPE:type_name -> GGPEHIBFFCB
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SetRogueExhibitionScRsp_proto_init() }
func file_SetRogueExhibitionScRsp_proto_init() {
	if File_SetRogueExhibitionScRsp_proto != nil {
		return
	}
	file_GGPEHIBFFCB_proto_init()
	file_EMIECBDCBNM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SetRogueExhibitionScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRogueExhibitionScRsp); i {
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
			RawDescriptor: file_SetRogueExhibitionScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetRogueExhibitionScRsp_proto_goTypes,
		DependencyIndexes: file_SetRogueExhibitionScRsp_proto_depIdxs,
		MessageInfos:      file_SetRogueExhibitionScRsp_proto_msgTypes,
	}.Build()
	File_SetRogueExhibitionScRsp_proto = out.File
	file_SetRogueExhibitionScRsp_proto_rawDesc = nil
	file_SetRogueExhibitionScRsp_proto_goTypes = nil
	file_SetRogueExhibitionScRsp_proto_depIdxs = nil
}
