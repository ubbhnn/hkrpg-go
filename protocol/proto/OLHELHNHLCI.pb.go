// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: OLHELHNHLCI.proto

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

type OLHELHNHLCI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BKMODLIJKGO DevelopmentType `protobuf:"varint,8,opt,name=BKMODLIJKGO,proto3,enum=DevelopmentType" json:"BKMODLIJKGO,omitempty"`
	Time        int64           `protobuf:"varint,7,opt,name=time,proto3" json:"time,omitempty"`
	// Types that are assignable to BCPIDNIOFPM:
	//
	//	*OLHELHNHLCI_GIJPNCCEGEK
	//	*OLHELHNHLCI_AIIEOGKGBIJ
	//	*OLHELHNHLCI_AvatarId
	//	*OLHELHNHLCI_PDIGONIOFGC
	//	*OLHELHNHLCI_LLJLJLHAKKB
	//	*OLHELHNHLCI_CEFPILPICIN
	BCPIDNIOFPM isOLHELHNHLCI_BCPIDNIOFPM `protobuf_oneof:"BCPIDNIOFPM"`
}

func (x *OLHELHNHLCI) Reset() {
	*x = OLHELHNHLCI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OLHELHNHLCI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OLHELHNHLCI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OLHELHNHLCI) ProtoMessage() {}

func (x *OLHELHNHLCI) ProtoReflect() protoreflect.Message {
	mi := &file_OLHELHNHLCI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OLHELHNHLCI.ProtoReflect.Descriptor instead.
func (*OLHELHNHLCI) Descriptor() ([]byte, []int) {
	return file_OLHELHNHLCI_proto_rawDescGZIP(), []int{0}
}

func (x *OLHELHNHLCI) GetBKMODLIJKGO() DevelopmentType {
	if x != nil {
		return x.BKMODLIJKGO
	}
	return DevelopmentType_DEVELOPMENT_NONE
}

func (x *OLHELHNHLCI) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (m *OLHELHNHLCI) GetBCPIDNIOFPM() isOLHELHNHLCI_BCPIDNIOFPM {
	if m != nil {
		return m.BCPIDNIOFPM
	}
	return nil
}

func (x *OLHELHNHLCI) GetGIJPNCCEGEK() *HLNCNPHPNJB {
	if x, ok := x.GetBCPIDNIOFPM().(*OLHELHNHLCI_GIJPNCCEGEK); ok {
		return x.GIJPNCCEGEK
	}
	return nil
}

func (x *OLHELHNHLCI) GetAIIEOGKGBIJ() *NJDHPPFMDJO {
	if x, ok := x.GetBCPIDNIOFPM().(*OLHELHNHLCI_AIIEOGKGBIJ); ok {
		return x.AIIEOGKGBIJ
	}
	return nil
}

func (x *OLHELHNHLCI) GetAvatarId() uint32 {
	if x, ok := x.GetBCPIDNIOFPM().(*OLHELHNHLCI_AvatarId); ok {
		return x.AvatarId
	}
	return 0
}

func (x *OLHELHNHLCI) GetPDIGONIOFGC() uint32 {
	if x, ok := x.GetBCPIDNIOFPM().(*OLHELHNHLCI_PDIGONIOFGC); ok {
		return x.PDIGONIOFGC
	}
	return 0
}

func (x *OLHELHNHLCI) GetLLJLJLHAKKB() uint32 {
	if x, ok := x.GetBCPIDNIOFPM().(*OLHELHNHLCI_LLJLJLHAKKB); ok {
		return x.LLJLJLHAKKB
	}
	return 0
}

func (x *OLHELHNHLCI) GetCEFPILPICIN() *CJHMPNNALNL {
	if x, ok := x.GetBCPIDNIOFPM().(*OLHELHNHLCI_CEFPILPICIN); ok {
		return x.CEFPILPICIN
	}
	return nil
}

type isOLHELHNHLCI_BCPIDNIOFPM interface {
	isOLHELHNHLCI_BCPIDNIOFPM()
}

type OLHELHNHLCI_GIJPNCCEGEK struct {
	GIJPNCCEGEK *HLNCNPHPNJB `protobuf:"bytes,1867,opt,name=GIJPNCCEGEK,proto3,oneof"`
}

type OLHELHNHLCI_AIIEOGKGBIJ struct {
	AIIEOGKGBIJ *NJDHPPFMDJO `protobuf:"bytes,1350,opt,name=AIIEOGKGBIJ,proto3,oneof"`
}

type OLHELHNHLCI_AvatarId struct {
	AvatarId uint32 `protobuf:"varint,1495,opt,name=avatar_id,json=avatarId,proto3,oneof"`
}

type OLHELHNHLCI_PDIGONIOFGC struct {
	PDIGONIOFGC uint32 `protobuf:"varint,936,opt,name=PDIGONIOFGC,proto3,oneof"`
}

type OLHELHNHLCI_LLJLJLHAKKB struct {
	LLJLJLHAKKB uint32 `protobuf:"varint,955,opt,name=LLJLJLHAKKB,proto3,oneof"`
}

type OLHELHNHLCI_CEFPILPICIN struct {
	CEFPILPICIN *CJHMPNNALNL `protobuf:"bytes,606,opt,name=CEFPILPICIN,proto3,oneof"`
}

func (*OLHELHNHLCI_GIJPNCCEGEK) isOLHELHNHLCI_BCPIDNIOFPM() {}

func (*OLHELHNHLCI_AIIEOGKGBIJ) isOLHELHNHLCI_BCPIDNIOFPM() {}

func (*OLHELHNHLCI_AvatarId) isOLHELHNHLCI_BCPIDNIOFPM() {}

func (*OLHELHNHLCI_PDIGONIOFGC) isOLHELHNHLCI_BCPIDNIOFPM() {}

func (*OLHELHNHLCI_LLJLJLHAKKB) isOLHELHNHLCI_BCPIDNIOFPM() {}

func (*OLHELHNHLCI_CEFPILPICIN) isOLHELHNHLCI_BCPIDNIOFPM() {}

var File_OLHELHNHLCI_proto protoreflect.FileDescriptor

var file_OLHELHNHLCI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4f, 0x4c, 0x48, 0x45, 0x4c, 0x48, 0x4e, 0x48, 0x4c, 0x43, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4a, 0x48, 0x4d, 0x50, 0x4e, 0x4e, 0x41, 0x4c, 0x4e, 0x4c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x4c, 0x4e, 0x43, 0x4e, 0x50, 0x48, 0x50,
	0x4e, 0x4a, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x4a, 0x44, 0x48, 0x50,
	0x50, 0x46, 0x4d, 0x44, 0x4a, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x44, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x0b, 0x4f, 0x4c, 0x48, 0x45, 0x4c, 0x48, 0x4e, 0x48,
	0x4c, 0x43, 0x49, 0x12, 0x32, 0x0a, 0x0b, 0x42, 0x4b, 0x4d, 0x4f, 0x44, 0x4c, 0x49, 0x4a, 0x4b,
	0x47, 0x4f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x42, 0x4b, 0x4d, 0x4f,
	0x44, 0x4c, 0x49, 0x4a, 0x4b, 0x47, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x47,
	0x49, 0x4a, 0x50, 0x4e, 0x43, 0x43, 0x45, 0x47, 0x45, 0x4b, 0x18, 0xcb, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x4c, 0x4e, 0x43, 0x4e, 0x50, 0x48, 0x50, 0x4e, 0x4a, 0x42, 0x48,
	0x00, 0x52, 0x0b, 0x47, 0x49, 0x4a, 0x50, 0x4e, 0x43, 0x43, 0x45, 0x47, 0x45, 0x4b, 0x12, 0x31,
	0x0a, 0x0b, 0x41, 0x49, 0x49, 0x45, 0x4f, 0x47, 0x4b, 0x47, 0x42, 0x49, 0x4a, 0x18, 0xc6, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x4a, 0x44, 0x48, 0x50, 0x50, 0x46, 0x4d, 0x44,
	0x4a, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x41, 0x49, 0x49, 0x45, 0x4f, 0x47, 0x4b, 0x47, 0x42, 0x49,
	0x4a, 0x12, 0x1e, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0xd7,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0b, 0x50, 0x44, 0x49, 0x47, 0x4f, 0x4e, 0x49, 0x4f, 0x46, 0x47, 0x43,
	0x18, 0xa8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0b, 0x50, 0x44, 0x49, 0x47, 0x4f,
	0x4e, 0x49, 0x4f, 0x46, 0x47, 0x43, 0x12, 0x23, 0x0a, 0x0b, 0x4c, 0x4c, 0x4a, 0x4c, 0x4a, 0x4c,
	0x48, 0x41, 0x4b, 0x4b, 0x42, 0x18, 0xbb, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0b,
	0x4c, 0x4c, 0x4a, 0x4c, 0x4a, 0x4c, 0x48, 0x41, 0x4b, 0x4b, 0x42, 0x12, 0x31, 0x0a, 0x0b, 0x43,
	0x45, 0x46, 0x50, 0x49, 0x4c, 0x50, 0x49, 0x43, 0x49, 0x4e, 0x18, 0xde, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4a, 0x48, 0x4d, 0x50, 0x4e, 0x4e, 0x41, 0x4c, 0x4e, 0x4c, 0x48,
	0x00, 0x52, 0x0b, 0x43, 0x45, 0x46, 0x50, 0x49, 0x4c, 0x50, 0x49, 0x43, 0x49, 0x4e, 0x42, 0x0d,
	0x0a, 0x0b, 0x42, 0x43, 0x50, 0x49, 0x44, 0x4e, 0x49, 0x4f, 0x46, 0x50, 0x4d, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OLHELHNHLCI_proto_rawDescOnce sync.Once
	file_OLHELHNHLCI_proto_rawDescData = file_OLHELHNHLCI_proto_rawDesc
)

func file_OLHELHNHLCI_proto_rawDescGZIP() []byte {
	file_OLHELHNHLCI_proto_rawDescOnce.Do(func() {
		file_OLHELHNHLCI_proto_rawDescData = protoimpl.X.CompressGZIP(file_OLHELHNHLCI_proto_rawDescData)
	})
	return file_OLHELHNHLCI_proto_rawDescData
}

var file_OLHELHNHLCI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OLHELHNHLCI_proto_goTypes = []interface{}{
	(*OLHELHNHLCI)(nil),  // 0: OLHELHNHLCI
	(DevelopmentType)(0), // 1: DevelopmentType
	(*HLNCNPHPNJB)(nil),  // 2: HLNCNPHPNJB
	(*NJDHPPFMDJO)(nil),  // 3: NJDHPPFMDJO
	(*CJHMPNNALNL)(nil),  // 4: CJHMPNNALNL
}
var file_OLHELHNHLCI_proto_depIdxs = []int32{
	1, // 0: OLHELHNHLCI.BKMODLIJKGO:type_name -> DevelopmentType
	2, // 1: OLHELHNHLCI.GIJPNCCEGEK:type_name -> HLNCNPHPNJB
	3, // 2: OLHELHNHLCI.AIIEOGKGBIJ:type_name -> NJDHPPFMDJO
	4, // 3: OLHELHNHLCI.CEFPILPICIN:type_name -> CJHMPNNALNL
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_OLHELHNHLCI_proto_init() }
func file_OLHELHNHLCI_proto_init() {
	if File_OLHELHNHLCI_proto != nil {
		return
	}
	file_CJHMPNNALNL_proto_init()
	file_HLNCNPHPNJB_proto_init()
	file_NJDHPPFMDJO_proto_init()
	file_DevelopmentType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OLHELHNHLCI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OLHELHNHLCI); i {
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
	file_OLHELHNHLCI_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OLHELHNHLCI_GIJPNCCEGEK)(nil),
		(*OLHELHNHLCI_AIIEOGKGBIJ)(nil),
		(*OLHELHNHLCI_AvatarId)(nil),
		(*OLHELHNHLCI_PDIGONIOFGC)(nil),
		(*OLHELHNHLCI_LLJLJLHAKKB)(nil),
		(*OLHELHNHLCI_CEFPILPICIN)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_OLHELHNHLCI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OLHELHNHLCI_proto_goTypes,
		DependencyIndexes: file_OLHELHNHLCI_proto_depIdxs,
		MessageInfos:      file_OLHELHNHLCI_proto_msgTypes,
	}.Build()
	File_OLHELHNHLCI_proto = out.File
	file_OLHELHNHLCI_proto_rawDesc = nil
	file_OLHELHNHLCI_proto_goTypes = nil
	file_OLHELHNHLCI_proto_depIdxs = nil
}
