// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PunkLordBattleResultScNotify.proto

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

type PunkLordBattleResultScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BPGFNCNJKHH *PunkLordMonsterBasicInfo `protobuf:"bytes,3,opt,name=BPGFNCNJKHH,proto3" json:"BPGFNCNJKHH,omitempty"`
	HLLELCGOCNK *PunkLordBattleRecord     `protobuf:"bytes,12,opt,name=HLLELCGOCNK,proto3" json:"HLLELCGOCNK,omitempty"`
	NIHCBBOJAEC uint32                    `protobuf:"varint,10,opt,name=NIHCBBOJAEC,proto3" json:"NIHCBBOJAEC,omitempty"`
	BPAFMEHEPKF uint32                    `protobuf:"varint,13,opt,name=BPAFMEHEPKF,proto3" json:"BPAFMEHEPKF,omitempty"`
	MGMEANGENHK uint32                    `protobuf:"varint,4,opt,name=MGMEANGENHK,proto3" json:"MGMEANGENHK,omitempty"`
}

func (x *PunkLordBattleResultScNotify) Reset() {
	*x = PunkLordBattleResultScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PunkLordBattleResultScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PunkLordBattleResultScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PunkLordBattleResultScNotify) ProtoMessage() {}

func (x *PunkLordBattleResultScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PunkLordBattleResultScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PunkLordBattleResultScNotify.ProtoReflect.Descriptor instead.
func (*PunkLordBattleResultScNotify) Descriptor() ([]byte, []int) {
	return file_PunkLordBattleResultScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PunkLordBattleResultScNotify) GetBPGFNCNJKHH() *PunkLordMonsterBasicInfo {
	if x != nil {
		return x.BPGFNCNJKHH
	}
	return nil
}

func (x *PunkLordBattleResultScNotify) GetHLLELCGOCNK() *PunkLordBattleRecord {
	if x != nil {
		return x.HLLELCGOCNK
	}
	return nil
}

func (x *PunkLordBattleResultScNotify) GetNIHCBBOJAEC() uint32 {
	if x != nil {
		return x.NIHCBBOJAEC
	}
	return 0
}

func (x *PunkLordBattleResultScNotify) GetBPAFMEHEPKF() uint32 {
	if x != nil {
		return x.BPAFMEHEPKF
	}
	return 0
}

func (x *PunkLordBattleResultScNotify) GetMGMEANGENHK() uint32 {
	if x != nil {
		return x.MGMEANGENHK
	}
	return 0
}

var File_PunkLordBattleResultScNotify_proto protoreflect.FileDescriptor

var file_PunkLordBattleResultScNotify_proto_rawDesc = []byte{
	0x0a, 0x22, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfa, 0x01, 0x0a, 0x1c, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x42, 0x50, 0x47, 0x46, 0x4e, 0x43, 0x4e, 0x4a, 0x4b, 0x48, 0x48,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72,
	0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0b, 0x42, 0x50, 0x47, 0x46, 0x4e, 0x43, 0x4e, 0x4a, 0x4b, 0x48, 0x48, 0x12, 0x37,
	0x0a, 0x0b, 0x48, 0x4c, 0x4c, 0x45, 0x4c, 0x43, 0x47, 0x4f, 0x43, 0x4e, 0x4b, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x48, 0x4c, 0x4c, 0x45,
	0x4c, 0x43, 0x47, 0x4f, 0x43, 0x4e, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x49, 0x48, 0x43, 0x42,
	0x42, 0x4f, 0x4a, 0x41, 0x45, 0x43, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x49,
	0x48, 0x43, 0x42, 0x42, 0x4f, 0x4a, 0x41, 0x45, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x50, 0x41,
	0x46, 0x4d, 0x45, 0x48, 0x45, 0x50, 0x4b, 0x46, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x42, 0x50, 0x41, 0x46, 0x4d, 0x45, 0x48, 0x45, 0x50, 0x4b, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x4d,
	0x47, 0x4d, 0x45, 0x41, 0x4e, 0x47, 0x45, 0x4e, 0x48, 0x4b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4d, 0x47, 0x4d, 0x45, 0x41, 0x4e, 0x47, 0x45, 0x4e, 0x48, 0x4b, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PunkLordBattleResultScNotify_proto_rawDescOnce sync.Once
	file_PunkLordBattleResultScNotify_proto_rawDescData = file_PunkLordBattleResultScNotify_proto_rawDesc
)

func file_PunkLordBattleResultScNotify_proto_rawDescGZIP() []byte {
	file_PunkLordBattleResultScNotify_proto_rawDescOnce.Do(func() {
		file_PunkLordBattleResultScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PunkLordBattleResultScNotify_proto_rawDescData)
	})
	return file_PunkLordBattleResultScNotify_proto_rawDescData
}

var file_PunkLordBattleResultScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PunkLordBattleResultScNotify_proto_goTypes = []interface{}{
	(*PunkLordBattleResultScNotify)(nil), // 0: PunkLordBattleResultScNotify
	(*PunkLordMonsterBasicInfo)(nil),     // 1: PunkLordMonsterBasicInfo
	(*PunkLordBattleRecord)(nil),         // 2: PunkLordBattleRecord
}
var file_PunkLordBattleResultScNotify_proto_depIdxs = []int32{
	1, // 0: PunkLordBattleResultScNotify.BPGFNCNJKHH:type_name -> PunkLordMonsterBasicInfo
	2, // 1: PunkLordBattleResultScNotify.HLLELCGOCNK:type_name -> PunkLordBattleRecord
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PunkLordBattleResultScNotify_proto_init() }
func file_PunkLordBattleResultScNotify_proto_init() {
	if File_PunkLordBattleResultScNotify_proto != nil {
		return
	}
	file_PunkLordBattleRecord_proto_init()
	file_PunkLordMonsterBasicInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PunkLordBattleResultScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PunkLordBattleResultScNotify); i {
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
			RawDescriptor: file_PunkLordBattleResultScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PunkLordBattleResultScNotify_proto_goTypes,
		DependencyIndexes: file_PunkLordBattleResultScNotify_proto_depIdxs,
		MessageInfos:      file_PunkLordBattleResultScNotify_proto_msgTypes,
	}.Build()
	File_PunkLordBattleResultScNotify_proto = out.File
	file_PunkLordBattleResultScNotify_proto_rawDesc = nil
	file_PunkLordBattleResultScNotify_proto_goTypes = nil
	file_PunkLordBattleResultScNotify_proto_depIdxs = nil
}
