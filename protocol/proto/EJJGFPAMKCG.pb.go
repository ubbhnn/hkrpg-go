// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EJJGFPAMKCG.proto

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

type EJJGFPAMKCG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MMODGAKDNCJ *MJHCOCMPAGC `protobuf:"bytes,2,opt,name=MMODGAKDNCJ,proto3" json:"MMODGAKDNCJ,omitempty"`
	OFMEFDIKHPH *CLCICIDMHKE `protobuf:"bytes,10,opt,name=OFMEFDIKHPH,proto3" json:"OFMEFDIKHPH,omitempty"`
	PPOBJLPAANM *JFCMKBOKFOF `protobuf:"bytes,15,opt,name=PPOBJLPAANM,proto3" json:"PPOBJLPAANM,omitempty"`
	LJEBNEBDLOB *FFLOEAGFOIA `protobuf:"bytes,12,opt,name=LJEBNEBDLOB,proto3" json:"LJEBNEBDLOB,omitempty"`
	MIDGOFHDBMF *GPIANNCLDCP `protobuf:"bytes,9,opt,name=MIDGOFHDBMF,proto3" json:"MIDGOFHDBMF,omitempty"`
}

func (x *EJJGFPAMKCG) Reset() {
	*x = EJJGFPAMKCG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EJJGFPAMKCG_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EJJGFPAMKCG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EJJGFPAMKCG) ProtoMessage() {}

func (x *EJJGFPAMKCG) ProtoReflect() protoreflect.Message {
	mi := &file_EJJGFPAMKCG_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EJJGFPAMKCG.ProtoReflect.Descriptor instead.
func (*EJJGFPAMKCG) Descriptor() ([]byte, []int) {
	return file_EJJGFPAMKCG_proto_rawDescGZIP(), []int{0}
}

func (x *EJJGFPAMKCG) GetMMODGAKDNCJ() *MJHCOCMPAGC {
	if x != nil {
		return x.MMODGAKDNCJ
	}
	return nil
}

func (x *EJJGFPAMKCG) GetOFMEFDIKHPH() *CLCICIDMHKE {
	if x != nil {
		return x.OFMEFDIKHPH
	}
	return nil
}

func (x *EJJGFPAMKCG) GetPPOBJLPAANM() *JFCMKBOKFOF {
	if x != nil {
		return x.PPOBJLPAANM
	}
	return nil
}

func (x *EJJGFPAMKCG) GetLJEBNEBDLOB() *FFLOEAGFOIA {
	if x != nil {
		return x.LJEBNEBDLOB
	}
	return nil
}

func (x *EJJGFPAMKCG) GetMIDGOFHDBMF() *GPIANNCLDCP {
	if x != nil {
		return x.MIDGOFHDBMF
	}
	return nil
}

var File_EJJGFPAMKCG_proto protoreflect.FileDescriptor

var file_EJJGFPAMKCG_proto_rawDesc = []byte{
	0x0a, 0x11, 0x45, 0x4a, 0x4a, 0x47, 0x46, 0x50, 0x41, 0x4d, 0x4b, 0x43, 0x47, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4c, 0x43, 0x49, 0x43, 0x49, 0x44, 0x4d, 0x48, 0x4b, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x4a, 0x48, 0x43, 0x4f, 0x43, 0x4d, 0x50,
	0x41, 0x47, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x50, 0x49, 0x41, 0x4e,
	0x4e, 0x43, 0x4c, 0x44, 0x43, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x46,
	0x4c, 0x4f, 0x45, 0x41, 0x47, 0x46, 0x4f, 0x49, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4a, 0x46, 0x43, 0x4d, 0x4b, 0x42, 0x4f, 0x4b, 0x46, 0x4f, 0x46, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x0b, 0x45, 0x4a, 0x4a, 0x47, 0x46, 0x50, 0x41, 0x4d, 0x4b,
	0x43, 0x47, 0x12, 0x2e, 0x0a, 0x0b, 0x4d, 0x4d, 0x4f, 0x44, 0x47, 0x41, 0x4b, 0x44, 0x4e, 0x43,
	0x4a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x4a, 0x48, 0x43, 0x4f, 0x43,
	0x4d, 0x50, 0x41, 0x47, 0x43, 0x52, 0x0b, 0x4d, 0x4d, 0x4f, 0x44, 0x47, 0x41, 0x4b, 0x44, 0x4e,
	0x43, 0x4a, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x46, 0x4d, 0x45, 0x46, 0x44, 0x49, 0x4b, 0x48, 0x50,
	0x48, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4c, 0x43, 0x49, 0x43, 0x49,
	0x44, 0x4d, 0x48, 0x4b, 0x45, 0x52, 0x0b, 0x4f, 0x46, 0x4d, 0x45, 0x46, 0x44, 0x49, 0x4b, 0x48,
	0x50, 0x48, 0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x50, 0x4f, 0x42, 0x4a, 0x4c, 0x50, 0x41, 0x41, 0x4e,
	0x4d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x46, 0x43, 0x4d, 0x4b, 0x42,
	0x4f, 0x4b, 0x46, 0x4f, 0x46, 0x52, 0x0b, 0x50, 0x50, 0x4f, 0x42, 0x4a, 0x4c, 0x50, 0x41, 0x41,
	0x4e, 0x4d, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x4a, 0x45, 0x42, 0x4e, 0x45, 0x42, 0x44, 0x4c, 0x4f,
	0x42, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x46, 0x4c, 0x4f, 0x45, 0x41,
	0x47, 0x46, 0x4f, 0x49, 0x41, 0x52, 0x0b, 0x4c, 0x4a, 0x45, 0x42, 0x4e, 0x45, 0x42, 0x44, 0x4c,
	0x4f, 0x42, 0x12, 0x2e, 0x0a, 0x0b, 0x4d, 0x49, 0x44, 0x47, 0x4f, 0x46, 0x48, 0x44, 0x42, 0x4d,
	0x46, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x50, 0x49, 0x41, 0x4e, 0x4e,
	0x43, 0x4c, 0x44, 0x43, 0x50, 0x52, 0x0b, 0x4d, 0x49, 0x44, 0x47, 0x4f, 0x46, 0x48, 0x44, 0x42,
	0x4d, 0x46, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EJJGFPAMKCG_proto_rawDescOnce sync.Once
	file_EJJGFPAMKCG_proto_rawDescData = file_EJJGFPAMKCG_proto_rawDesc
)

func file_EJJGFPAMKCG_proto_rawDescGZIP() []byte {
	file_EJJGFPAMKCG_proto_rawDescOnce.Do(func() {
		file_EJJGFPAMKCG_proto_rawDescData = protoimpl.X.CompressGZIP(file_EJJGFPAMKCG_proto_rawDescData)
	})
	return file_EJJGFPAMKCG_proto_rawDescData
}

var file_EJJGFPAMKCG_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EJJGFPAMKCG_proto_goTypes = []interface{}{
	(*EJJGFPAMKCG)(nil), // 0: EJJGFPAMKCG
	(*MJHCOCMPAGC)(nil), // 1: MJHCOCMPAGC
	(*CLCICIDMHKE)(nil), // 2: CLCICIDMHKE
	(*JFCMKBOKFOF)(nil), // 3: JFCMKBOKFOF
	(*FFLOEAGFOIA)(nil), // 4: FFLOEAGFOIA
	(*GPIANNCLDCP)(nil), // 5: GPIANNCLDCP
}
var file_EJJGFPAMKCG_proto_depIdxs = []int32{
	1, // 0: EJJGFPAMKCG.MMODGAKDNCJ:type_name -> MJHCOCMPAGC
	2, // 1: EJJGFPAMKCG.OFMEFDIKHPH:type_name -> CLCICIDMHKE
	3, // 2: EJJGFPAMKCG.PPOBJLPAANM:type_name -> JFCMKBOKFOF
	4, // 3: EJJGFPAMKCG.LJEBNEBDLOB:type_name -> FFLOEAGFOIA
	5, // 4: EJJGFPAMKCG.MIDGOFHDBMF:type_name -> GPIANNCLDCP
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_EJJGFPAMKCG_proto_init() }
func file_EJJGFPAMKCG_proto_init() {
	if File_EJJGFPAMKCG_proto != nil {
		return
	}
	file_CLCICIDMHKE_proto_init()
	file_MJHCOCMPAGC_proto_init()
	file_GPIANNCLDCP_proto_init()
	file_FFLOEAGFOIA_proto_init()
	file_JFCMKBOKFOF_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EJJGFPAMKCG_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EJJGFPAMKCG); i {
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
			RawDescriptor: file_EJJGFPAMKCG_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EJJGFPAMKCG_proto_goTypes,
		DependencyIndexes: file_EJJGFPAMKCG_proto_depIdxs,
		MessageInfos:      file_EJJGFPAMKCG_proto_msgTypes,
	}.Build()
	File_EJJGFPAMKCG_proto = out.File
	file_EJJGFPAMKCG_proto_rawDesc = nil
	file_EJJGFPAMKCG_proto_goTypes = nil
	file_EJJGFPAMKCG_proto_depIdxs = nil
}