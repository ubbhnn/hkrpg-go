// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MANNPANJCLL.proto

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

type MANNPANJCLL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BDDHIMGOGLC uint32 `protobuf:"varint,12,opt,name=BDDHIMGOGLC,proto3" json:"BDDHIMGOGLC,omitempty"`
	// Types that are assignable to AOPJLPPLLBP:
	//
	//	*MANNPANJCLL_JAEAHBGOMJG
	//	*MANNPANJCLL_ELIFOGEMLDE
	//	*MANNPANJCLL_IDEGMFPEMLN
	//	*MANNPANJCLL_CNDNDNNKLAB
	//	*MANNPANJCLL_LLNONJCICAA
	//	*MANNPANJCLL_GNBKIDONJPH
	//	*MANNPANJCLL_BHJMMAGEFIJ
	//	*MANNPANJCLL_FIIHOANMOKE
	AOPJLPPLLBP isMANNPANJCLL_AOPJLPPLLBP `protobuf_oneof:"AOPJLPPLLBP"`
}

func (x *MANNPANJCLL) Reset() {
	*x = MANNPANJCLL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MANNPANJCLL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MANNPANJCLL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MANNPANJCLL) ProtoMessage() {}

func (x *MANNPANJCLL) ProtoReflect() protoreflect.Message {
	mi := &file_MANNPANJCLL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MANNPANJCLL.ProtoReflect.Descriptor instead.
func (*MANNPANJCLL) Descriptor() ([]byte, []int) {
	return file_MANNPANJCLL_proto_rawDescGZIP(), []int{0}
}

func (x *MANNPANJCLL) GetBDDHIMGOGLC() uint32 {
	if x != nil {
		return x.BDDHIMGOGLC
	}
	return 0
}

func (m *MANNPANJCLL) GetAOPJLPPLLBP() isMANNPANJCLL_AOPJLPPLLBP {
	if m != nil {
		return m.AOPJLPPLLBP
	}
	return nil
}

func (x *MANNPANJCLL) GetJAEAHBGOMJG() *AKJLICDOOND {
	if x, ok := x.GetAOPJLPPLLBP().(*MANNPANJCLL_JAEAHBGOMJG); ok {
		return x.JAEAHBGOMJG
	}
	return nil
}

func (x *MANNPANJCLL) GetELIFOGEMLDE() *ONDNCDGJABI {
	if x, ok := x.GetAOPJLPPLLBP().(*MANNPANJCLL_ELIFOGEMLDE); ok {
		return x.ELIFOGEMLDE
	}
	return nil
}

func (x *MANNPANJCLL) GetIDEGMFPEMLN() *ONAKBKCHNFK {
	if x, ok := x.GetAOPJLPPLLBP().(*MANNPANJCLL_IDEGMFPEMLN); ok {
		return x.IDEGMFPEMLN
	}
	return nil
}

func (x *MANNPANJCLL) GetCNDNDNNKLAB() *OIMAHOEJEKN {
	if x, ok := x.GetAOPJLPPLLBP().(*MANNPANJCLL_CNDNDNNKLAB); ok {
		return x.CNDNDNNKLAB
	}
	return nil
}

func (x *MANNPANJCLL) GetLLNONJCICAA() *MOPCPNPJCJM {
	if x, ok := x.GetAOPJLPPLLBP().(*MANNPANJCLL_LLNONJCICAA); ok {
		return x.LLNONJCICAA
	}
	return nil
}

func (x *MANNPANJCLL) GetGNBKIDONJPH() *GGMKAOBPJIB {
	if x, ok := x.GetAOPJLPPLLBP().(*MANNPANJCLL_GNBKIDONJPH); ok {
		return x.GNBKIDONJPH
	}
	return nil
}

func (x *MANNPANJCLL) GetBHJMMAGEFIJ() *JBBLICFMJAO {
	if x, ok := x.GetAOPJLPPLLBP().(*MANNPANJCLL_BHJMMAGEFIJ); ok {
		return x.BHJMMAGEFIJ
	}
	return nil
}

func (x *MANNPANJCLL) GetFIIHOANMOKE() bool {
	if x, ok := x.GetAOPJLPPLLBP().(*MANNPANJCLL_FIIHOANMOKE); ok {
		return x.FIIHOANMOKE
	}
	return false
}

type isMANNPANJCLL_AOPJLPPLLBP interface {
	isMANNPANJCLL_AOPJLPPLLBP()
}

type MANNPANJCLL_JAEAHBGOMJG struct {
	JAEAHBGOMJG *AKJLICDOOND `protobuf:"bytes,1,opt,name=JAEAHBGOMJG,proto3,oneof"`
}

type MANNPANJCLL_ELIFOGEMLDE struct {
	ELIFOGEMLDE *ONDNCDGJABI `protobuf:"bytes,3,opt,name=ELIFOGEMLDE,proto3,oneof"`
}

type MANNPANJCLL_IDEGMFPEMLN struct {
	IDEGMFPEMLN *ONAKBKCHNFK `protobuf:"bytes,13,opt,name=IDEGMFPEMLN,proto3,oneof"`
}

type MANNPANJCLL_CNDNDNNKLAB struct {
	CNDNDNNKLAB *OIMAHOEJEKN `protobuf:"bytes,10,opt,name=CNDNDNNKLAB,proto3,oneof"`
}

type MANNPANJCLL_LLNONJCICAA struct {
	LLNONJCICAA *MOPCPNPJCJM `protobuf:"bytes,5,opt,name=LLNONJCICAA,proto3,oneof"`
}

type MANNPANJCLL_GNBKIDONJPH struct {
	GNBKIDONJPH *GGMKAOBPJIB `protobuf:"bytes,11,opt,name=GNBKIDONJPH,proto3,oneof"`
}

type MANNPANJCLL_BHJMMAGEFIJ struct {
	BHJMMAGEFIJ *JBBLICFMJAO `protobuf:"bytes,2,opt,name=BHJMMAGEFIJ,proto3,oneof"`
}

type MANNPANJCLL_FIIHOANMOKE struct {
	FIIHOANMOKE bool `protobuf:"varint,14,opt,name=FIIHOANMOKE,proto3,oneof"`
}

func (*MANNPANJCLL_JAEAHBGOMJG) isMANNPANJCLL_AOPJLPPLLBP() {}

func (*MANNPANJCLL_ELIFOGEMLDE) isMANNPANJCLL_AOPJLPPLLBP() {}

func (*MANNPANJCLL_IDEGMFPEMLN) isMANNPANJCLL_AOPJLPPLLBP() {}

func (*MANNPANJCLL_CNDNDNNKLAB) isMANNPANJCLL_AOPJLPPLLBP() {}

func (*MANNPANJCLL_LLNONJCICAA) isMANNPANJCLL_AOPJLPPLLBP() {}

func (*MANNPANJCLL_GNBKIDONJPH) isMANNPANJCLL_AOPJLPPLLBP() {}

func (*MANNPANJCLL_BHJMMAGEFIJ) isMANNPANJCLL_AOPJLPPLLBP() {}

func (*MANNPANJCLL_FIIHOANMOKE) isMANNPANJCLL_AOPJLPPLLBP() {}

var File_MANNPANJCLL_proto protoreflect.FileDescriptor

var file_MANNPANJCLL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x41, 0x4e, 0x4e, 0x50, 0x41, 0x4e, 0x4a, 0x43, 0x4c, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4e, 0x44, 0x4e, 0x43, 0x44, 0x47, 0x4a, 0x41, 0x42, 0x49,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x4f, 0x50, 0x43, 0x50, 0x4e, 0x50, 0x4a,
	0x43, 0x4a, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x47, 0x4d, 0x4b, 0x41,
	0x4f, 0x42, 0x50, 0x4a, 0x49, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x49,
	0x4d, 0x41, 0x48, 0x4f, 0x45, 0x4a, 0x45, 0x4b, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4f, 0x4e, 0x41, 0x4b, 0x42, 0x4b, 0x43, 0x48, 0x4e, 0x46, 0x4b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x41, 0x4b, 0x4a, 0x4c, 0x49, 0x43, 0x44, 0x4f, 0x4f, 0x4e, 0x44, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x42, 0x42, 0x4c, 0x49, 0x43, 0x46, 0x4d, 0x4a,
	0x41, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x03, 0x0a, 0x0b, 0x4d, 0x41, 0x4e,
	0x4e, 0x50, 0x41, 0x4e, 0x4a, 0x43, 0x4c, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x44, 0x44, 0x48,
	0x49, 0x4d, 0x47, 0x4f, 0x47, 0x4c, 0x43, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42,
	0x44, 0x44, 0x48, 0x49, 0x4d, 0x47, 0x4f, 0x47, 0x4c, 0x43, 0x12, 0x30, 0x0a, 0x0b, 0x4a, 0x41,
	0x45, 0x41, 0x48, 0x42, 0x47, 0x4f, 0x4d, 0x4a, 0x47, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x41, 0x4b, 0x4a, 0x4c, 0x49, 0x43, 0x44, 0x4f, 0x4f, 0x4e, 0x44, 0x48, 0x00, 0x52,
	0x0b, 0x4a, 0x41, 0x45, 0x41, 0x48, 0x42, 0x47, 0x4f, 0x4d, 0x4a, 0x47, 0x12, 0x30, 0x0a, 0x0b,
	0x45, 0x4c, 0x49, 0x46, 0x4f, 0x47, 0x45, 0x4d, 0x4c, 0x44, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4e, 0x44, 0x4e, 0x43, 0x44, 0x47, 0x4a, 0x41, 0x42, 0x49, 0x48,
	0x00, 0x52, 0x0b, 0x45, 0x4c, 0x49, 0x46, 0x4f, 0x47, 0x45, 0x4d, 0x4c, 0x44, 0x45, 0x12, 0x30,
	0x0a, 0x0b, 0x49, 0x44, 0x45, 0x47, 0x4d, 0x46, 0x50, 0x45, 0x4d, 0x4c, 0x4e, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4e, 0x41, 0x4b, 0x42, 0x4b, 0x43, 0x48, 0x4e, 0x46,
	0x4b, 0x48, 0x00, 0x52, 0x0b, 0x49, 0x44, 0x45, 0x47, 0x4d, 0x46, 0x50, 0x45, 0x4d, 0x4c, 0x4e,
	0x12, 0x30, 0x0a, 0x0b, 0x43, 0x4e, 0x44, 0x4e, 0x44, 0x4e, 0x4e, 0x4b, 0x4c, 0x41, 0x42, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x49, 0x4d, 0x41, 0x48, 0x4f, 0x45, 0x4a,
	0x45, 0x4b, 0x4e, 0x48, 0x00, 0x52, 0x0b, 0x43, 0x4e, 0x44, 0x4e, 0x44, 0x4e, 0x4e, 0x4b, 0x4c,
	0x41, 0x42, 0x12, 0x30, 0x0a, 0x0b, 0x4c, 0x4c, 0x4e, 0x4f, 0x4e, 0x4a, 0x43, 0x49, 0x43, 0x41,
	0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x4f, 0x50, 0x43, 0x50, 0x4e,
	0x50, 0x4a, 0x43, 0x4a, 0x4d, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x4c, 0x4e, 0x4f, 0x4e, 0x4a, 0x43,
	0x49, 0x43, 0x41, 0x41, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x4e, 0x42, 0x4b, 0x49, 0x44, 0x4f, 0x4e,
	0x4a, 0x50, 0x48, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x47, 0x4d, 0x4b,
	0x41, 0x4f, 0x42, 0x50, 0x4a, 0x49, 0x42, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x4e, 0x42, 0x4b, 0x49,
	0x44, 0x4f, 0x4e, 0x4a, 0x50, 0x48, 0x12, 0x30, 0x0a, 0x0b, 0x42, 0x48, 0x4a, 0x4d, 0x4d, 0x41,
	0x47, 0x45, 0x46, 0x49, 0x4a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x42,
	0x42, 0x4c, 0x49, 0x43, 0x46, 0x4d, 0x4a, 0x41, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x42, 0x48, 0x4a,
	0x4d, 0x4d, 0x41, 0x47, 0x45, 0x46, 0x49, 0x4a, 0x12, 0x22, 0x0a, 0x0b, 0x46, 0x49, 0x49, 0x48,
	0x4f, 0x41, 0x4e, 0x4d, 0x4f, 0x4b, 0x45, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x0b, 0x46, 0x49, 0x49, 0x48, 0x4f, 0x41, 0x4e, 0x4d, 0x4f, 0x4b, 0x45, 0x42, 0x0d, 0x0a, 0x0b,
	0x41, 0x4f, 0x50, 0x4a, 0x4c, 0x50, 0x50, 0x4c, 0x4c, 0x42, 0x50, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_MANNPANJCLL_proto_rawDescOnce sync.Once
	file_MANNPANJCLL_proto_rawDescData = file_MANNPANJCLL_proto_rawDesc
)

func file_MANNPANJCLL_proto_rawDescGZIP() []byte {
	file_MANNPANJCLL_proto_rawDescOnce.Do(func() {
		file_MANNPANJCLL_proto_rawDescData = protoimpl.X.CompressGZIP(file_MANNPANJCLL_proto_rawDescData)
	})
	return file_MANNPANJCLL_proto_rawDescData
}

var file_MANNPANJCLL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MANNPANJCLL_proto_goTypes = []interface{}{
	(*MANNPANJCLL)(nil), // 0: MANNPANJCLL
	(*AKJLICDOOND)(nil), // 1: AKJLICDOOND
	(*ONDNCDGJABI)(nil), // 2: ONDNCDGJABI
	(*ONAKBKCHNFK)(nil), // 3: ONAKBKCHNFK
	(*OIMAHOEJEKN)(nil), // 4: OIMAHOEJEKN
	(*MOPCPNPJCJM)(nil), // 5: MOPCPNPJCJM
	(*GGMKAOBPJIB)(nil), // 6: GGMKAOBPJIB
	(*JBBLICFMJAO)(nil), // 7: JBBLICFMJAO
}
var file_MANNPANJCLL_proto_depIdxs = []int32{
	1, // 0: MANNPANJCLL.JAEAHBGOMJG:type_name -> AKJLICDOOND
	2, // 1: MANNPANJCLL.ELIFOGEMLDE:type_name -> ONDNCDGJABI
	3, // 2: MANNPANJCLL.IDEGMFPEMLN:type_name -> ONAKBKCHNFK
	4, // 3: MANNPANJCLL.CNDNDNNKLAB:type_name -> OIMAHOEJEKN
	5, // 4: MANNPANJCLL.LLNONJCICAA:type_name -> MOPCPNPJCJM
	6, // 5: MANNPANJCLL.GNBKIDONJPH:type_name -> GGMKAOBPJIB
	7, // 6: MANNPANJCLL.BHJMMAGEFIJ:type_name -> JBBLICFMJAO
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_MANNPANJCLL_proto_init() }
func file_MANNPANJCLL_proto_init() {
	if File_MANNPANJCLL_proto != nil {
		return
	}
	file_ONDNCDGJABI_proto_init()
	file_MOPCPNPJCJM_proto_init()
	file_GGMKAOBPJIB_proto_init()
	file_OIMAHOEJEKN_proto_init()
	file_ONAKBKCHNFK_proto_init()
	file_AKJLICDOOND_proto_init()
	file_JBBLICFMJAO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MANNPANJCLL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MANNPANJCLL); i {
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
	file_MANNPANJCLL_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MANNPANJCLL_JAEAHBGOMJG)(nil),
		(*MANNPANJCLL_ELIFOGEMLDE)(nil),
		(*MANNPANJCLL_IDEGMFPEMLN)(nil),
		(*MANNPANJCLL_CNDNDNNKLAB)(nil),
		(*MANNPANJCLL_LLNONJCICAA)(nil),
		(*MANNPANJCLL_GNBKIDONJPH)(nil),
		(*MANNPANJCLL_BHJMMAGEFIJ)(nil),
		(*MANNPANJCLL_FIIHOANMOKE)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MANNPANJCLL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MANNPANJCLL_proto_goTypes,
		DependencyIndexes: file_MANNPANJCLL_proto_depIdxs,
		MessageInfos:      file_MANNPANJCLL_proto_msgTypes,
	}.Build()
	File_MANNPANJCLL_proto = out.File
	file_MANNPANJCLL_proto_rawDesc = nil
	file_MANNPANJCLL_proto_goTypes = nil
	file_MANNPANJCLL_proto_depIdxs = nil
}
