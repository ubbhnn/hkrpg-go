// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LEKCLMAMABB.proto

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

type LEKCLMAMABB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PendingAction *BLBPCMEMJNC `protobuf:"bytes,11,opt,name=pending_action,json=pendingAction,proto3" json:"pending_action,omitempty"`
	KDFJLNHLHEO   *OFKENDBKCBL `protobuf:"bytes,7,opt,name=KDFJLNHLHEO,proto3" json:"KDFJLNHLHEO,omitempty"`
	OMCHNPNOMDC   *FHMHHEFPGIN `protobuf:"bytes,1,opt,name=OMCHNPNOMDC,proto3" json:"OMCHNPNOMDC,omitempty"`
	FIPBJGEADMI   *CNIFAGKECLJ `protobuf:"bytes,10,opt,name=FIPBJGEADMI,proto3" json:"FIPBJGEADMI,omitempty"`
	HOJIDCOKIDA   []uint32     `protobuf:"varint,14,rep,packed,name=HOJIDCOKIDA,proto3" json:"HOJIDCOKIDA,omitempty"`
	NMKBNCBHJKG   uint32       `protobuf:"varint,8,opt,name=NMKBNCBHJKG,proto3" json:"NMKBNCBHJKG,omitempty"`
	KKONPFDGJIE   *GKJMKHADBHM `protobuf:"bytes,4,opt,name=KKONPFDGJIE,proto3" json:"KKONPFDGJIE,omitempty"`
	SkillInfo     *ANMOHKLGHLA `protobuf:"bytes,2,opt,name=skill_info,json=skillInfo,proto3" json:"skill_info,omitempty"`
}

func (x *LEKCLMAMABB) Reset() {
	*x = LEKCLMAMABB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LEKCLMAMABB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LEKCLMAMABB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LEKCLMAMABB) ProtoMessage() {}

func (x *LEKCLMAMABB) ProtoReflect() protoreflect.Message {
	mi := &file_LEKCLMAMABB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LEKCLMAMABB.ProtoReflect.Descriptor instead.
func (*LEKCLMAMABB) Descriptor() ([]byte, []int) {
	return file_LEKCLMAMABB_proto_rawDescGZIP(), []int{0}
}

func (x *LEKCLMAMABB) GetPendingAction() *BLBPCMEMJNC {
	if x != nil {
		return x.PendingAction
	}
	return nil
}

func (x *LEKCLMAMABB) GetKDFJLNHLHEO() *OFKENDBKCBL {
	if x != nil {
		return x.KDFJLNHLHEO
	}
	return nil
}

func (x *LEKCLMAMABB) GetOMCHNPNOMDC() *FHMHHEFPGIN {
	if x != nil {
		return x.OMCHNPNOMDC
	}
	return nil
}

func (x *LEKCLMAMABB) GetFIPBJGEADMI() *CNIFAGKECLJ {
	if x != nil {
		return x.FIPBJGEADMI
	}
	return nil
}

func (x *LEKCLMAMABB) GetHOJIDCOKIDA() []uint32 {
	if x != nil {
		return x.HOJIDCOKIDA
	}
	return nil
}

func (x *LEKCLMAMABB) GetNMKBNCBHJKG() uint32 {
	if x != nil {
		return x.NMKBNCBHJKG
	}
	return 0
}

func (x *LEKCLMAMABB) GetKKONPFDGJIE() *GKJMKHADBHM {
	if x != nil {
		return x.KKONPFDGJIE
	}
	return nil
}

func (x *LEKCLMAMABB) GetSkillInfo() *ANMOHKLGHLA {
	if x != nil {
		return x.SkillInfo
	}
	return nil
}

var File_LEKCLMAMABB_proto protoreflect.FileDescriptor

var file_LEKCLMAMABB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x45, 0x4b, 0x43, 0x4c, 0x4d, 0x41, 0x4d, 0x41, 0x42, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x46, 0x4b, 0x45, 0x4e, 0x44, 0x42, 0x4b, 0x43, 0x42, 0x4c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x48, 0x4d, 0x48, 0x48, 0x45, 0x46, 0x50,
	0x47, 0x49, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x4c, 0x42, 0x50, 0x43,
	0x4d, 0x45, 0x4d, 0x4a, 0x4e, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x4e,
	0x4d, 0x4f, 0x48, 0x4b, 0x4c, 0x47, 0x48, 0x4c, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x43, 0x4e, 0x49, 0x46, 0x41, 0x47, 0x4b, 0x45, 0x43, 0x4c, 0x4a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x47, 0x4b, 0x4a, 0x4d, 0x4b, 0x48, 0x41, 0x44, 0x42, 0x48, 0x4d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x02, 0x0a, 0x0b, 0x4c, 0x45, 0x4b, 0x43, 0x4c, 0x4d,
	0x41, 0x4d, 0x41, 0x42, 0x42, 0x12, 0x33, 0x0a, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x42, 0x4c, 0x42, 0x50, 0x43, 0x4d, 0x45, 0x4d, 0x4a, 0x4e, 0x43, 0x52, 0x0d, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x0b, 0x4b, 0x44,
	0x46, 0x4a, 0x4c, 0x4e, 0x48, 0x4c, 0x48, 0x45, 0x4f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4f, 0x46, 0x4b, 0x45, 0x4e, 0x44, 0x42, 0x4b, 0x43, 0x42, 0x4c, 0x52, 0x0b, 0x4b,
	0x44, 0x46, 0x4a, 0x4c, 0x4e, 0x48, 0x4c, 0x48, 0x45, 0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x4d,
	0x43, 0x48, 0x4e, 0x50, 0x4e, 0x4f, 0x4d, 0x44, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x46, 0x48, 0x4d, 0x48, 0x48, 0x45, 0x46, 0x50, 0x47, 0x49, 0x4e, 0x52, 0x0b, 0x4f,
	0x4d, 0x43, 0x48, 0x4e, 0x50, 0x4e, 0x4f, 0x4d, 0x44, 0x43, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x49,
	0x50, 0x42, 0x4a, 0x47, 0x45, 0x41, 0x44, 0x4d, 0x49, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x43, 0x4e, 0x49, 0x46, 0x41, 0x47, 0x4b, 0x45, 0x43, 0x4c, 0x4a, 0x52, 0x0b, 0x46,
	0x49, 0x50, 0x42, 0x4a, 0x47, 0x45, 0x41, 0x44, 0x4d, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4f,
	0x4a, 0x49, 0x44, 0x43, 0x4f, 0x4b, 0x49, 0x44, 0x41, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0b, 0x48, 0x4f, 0x4a, 0x49, 0x44, 0x43, 0x4f, 0x4b, 0x49, 0x44, 0x41, 0x12, 0x20, 0x0a, 0x0b,
	0x4e, 0x4d, 0x4b, 0x42, 0x4e, 0x43, 0x42, 0x48, 0x4a, 0x4b, 0x47, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4e, 0x4d, 0x4b, 0x42, 0x4e, 0x43, 0x42, 0x48, 0x4a, 0x4b, 0x47, 0x12, 0x2e,
	0x0a, 0x0b, 0x4b, 0x4b, 0x4f, 0x4e, 0x50, 0x46, 0x44, 0x47, 0x4a, 0x49, 0x45, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x4b, 0x4a, 0x4d, 0x4b, 0x48, 0x41, 0x44, 0x42, 0x48,
	0x4d, 0x52, 0x0b, 0x4b, 0x4b, 0x4f, 0x4e, 0x50, 0x46, 0x44, 0x47, 0x4a, 0x49, 0x45, 0x12, 0x2b,
	0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x4e, 0x4d, 0x4f, 0x48, 0x4b, 0x4c, 0x47, 0x48, 0x4c, 0x41,
	0x52, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_LEKCLMAMABB_proto_rawDescOnce sync.Once
	file_LEKCLMAMABB_proto_rawDescData = file_LEKCLMAMABB_proto_rawDesc
)

func file_LEKCLMAMABB_proto_rawDescGZIP() []byte {
	file_LEKCLMAMABB_proto_rawDescOnce.Do(func() {
		file_LEKCLMAMABB_proto_rawDescData = protoimpl.X.CompressGZIP(file_LEKCLMAMABB_proto_rawDescData)
	})
	return file_LEKCLMAMABB_proto_rawDescData
}

var file_LEKCLMAMABB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LEKCLMAMABB_proto_goTypes = []interface{}{
	(*LEKCLMAMABB)(nil), // 0: LEKCLMAMABB
	(*BLBPCMEMJNC)(nil), // 1: BLBPCMEMJNC
	(*OFKENDBKCBL)(nil), // 2: OFKENDBKCBL
	(*FHMHHEFPGIN)(nil), // 3: FHMHHEFPGIN
	(*CNIFAGKECLJ)(nil), // 4: CNIFAGKECLJ
	(*GKJMKHADBHM)(nil), // 5: GKJMKHADBHM
	(*ANMOHKLGHLA)(nil), // 6: ANMOHKLGHLA
}
var file_LEKCLMAMABB_proto_depIdxs = []int32{
	1, // 0: LEKCLMAMABB.pending_action:type_name -> BLBPCMEMJNC
	2, // 1: LEKCLMAMABB.KDFJLNHLHEO:type_name -> OFKENDBKCBL
	3, // 2: LEKCLMAMABB.OMCHNPNOMDC:type_name -> FHMHHEFPGIN
	4, // 3: LEKCLMAMABB.FIPBJGEADMI:type_name -> CNIFAGKECLJ
	5, // 4: LEKCLMAMABB.KKONPFDGJIE:type_name -> GKJMKHADBHM
	6, // 5: LEKCLMAMABB.skill_info:type_name -> ANMOHKLGHLA
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_LEKCLMAMABB_proto_init() }
func file_LEKCLMAMABB_proto_init() {
	if File_LEKCLMAMABB_proto != nil {
		return
	}
	file_OFKENDBKCBL_proto_init()
	file_FHMHHEFPGIN_proto_init()
	file_BLBPCMEMJNC_proto_init()
	file_ANMOHKLGHLA_proto_init()
	file_CNIFAGKECLJ_proto_init()
	file_GKJMKHADBHM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LEKCLMAMABB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LEKCLMAMABB); i {
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
			RawDescriptor: file_LEKCLMAMABB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LEKCLMAMABB_proto_goTypes,
		DependencyIndexes: file_LEKCLMAMABB_proto_depIdxs,
		MessageInfos:      file_LEKCLMAMABB_proto_msgTypes,
	}.Build()
	File_LEKCLMAMABB_proto = out.File
	file_LEKCLMAMABB_proto_rawDesc = nil
	file_LEKCLMAMABB_proto_goTypes = nil
	file_LEKCLMAMABB_proto_depIdxs = nil
}
