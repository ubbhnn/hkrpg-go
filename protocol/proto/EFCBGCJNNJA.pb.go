// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EFCBGCJNNJA.proto

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

type EFCBGCJNNJA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to NBEIPGABFCF:
	//
	//	*EFCBGCJNNJA_PPNAPEDNKEF
	//	*EFCBGCJNNJA_ACHLGHPPMNO
	NBEIPGABFCF isEFCBGCJNNJA_NBEIPGABFCF `protobuf_oneof:"NBEIPGABFCF"`
}

func (x *EFCBGCJNNJA) Reset() {
	*x = EFCBGCJNNJA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EFCBGCJNNJA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EFCBGCJNNJA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EFCBGCJNNJA) ProtoMessage() {}

func (x *EFCBGCJNNJA) ProtoReflect() protoreflect.Message {
	mi := &file_EFCBGCJNNJA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EFCBGCJNNJA.ProtoReflect.Descriptor instead.
func (*EFCBGCJNNJA) Descriptor() ([]byte, []int) {
	return file_EFCBGCJNNJA_proto_rawDescGZIP(), []int{0}
}

func (m *EFCBGCJNNJA) GetNBEIPGABFCF() isEFCBGCJNNJA_NBEIPGABFCF {
	if m != nil {
		return m.NBEIPGABFCF
	}
	return nil
}

func (x *EFCBGCJNNJA) GetPPNAPEDNKEF() *AGAAGPJDOEA {
	if x, ok := x.GetNBEIPGABFCF().(*EFCBGCJNNJA_PPNAPEDNKEF); ok {
		return x.PPNAPEDNKEF
	}
	return nil
}

func (x *EFCBGCJNNJA) GetACHLGHPPMNO() *OJAHMDLNMAK {
	if x, ok := x.GetNBEIPGABFCF().(*EFCBGCJNNJA_ACHLGHPPMNO); ok {
		return x.ACHLGHPPMNO
	}
	return nil
}

type isEFCBGCJNNJA_NBEIPGABFCF interface {
	isEFCBGCJNNJA_NBEIPGABFCF()
}

type EFCBGCJNNJA_PPNAPEDNKEF struct {
	PPNAPEDNKEF *AGAAGPJDOEA `protobuf:"bytes,11,opt,name=PPNAPEDNKEF,proto3,oneof"`
}

type EFCBGCJNNJA_ACHLGHPPMNO struct {
	ACHLGHPPMNO *OJAHMDLNMAK `protobuf:"bytes,3,opt,name=ACHLGHPPMNO,proto3,oneof"`
}

func (*EFCBGCJNNJA_PPNAPEDNKEF) isEFCBGCJNNJA_NBEIPGABFCF() {}

func (*EFCBGCJNNJA_ACHLGHPPMNO) isEFCBGCJNNJA_NBEIPGABFCF() {}

var File_EFCBGCJNNJA_proto protoreflect.FileDescriptor

var file_EFCBGCJNNJA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x45, 0x46, 0x43, 0x42, 0x47, 0x43, 0x4a, 0x4e, 0x4e, 0x4a, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4a, 0x41, 0x48, 0x4d, 0x44, 0x4c, 0x4e, 0x4d, 0x41, 0x4b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x47, 0x41, 0x41, 0x47, 0x50, 0x4a, 0x44,
	0x4f, 0x45, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x45, 0x46,
	0x43, 0x42, 0x47, 0x43, 0x4a, 0x4e, 0x4e, 0x4a, 0x41, 0x12, 0x30, 0x0a, 0x0b, 0x50, 0x50, 0x4e,
	0x41, 0x50, 0x45, 0x44, 0x4e, 0x4b, 0x45, 0x46, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x41, 0x47, 0x41, 0x41, 0x47, 0x50, 0x4a, 0x44, 0x4f, 0x45, 0x41, 0x48, 0x00, 0x52, 0x0b,
	0x50, 0x50, 0x4e, 0x41, 0x50, 0x45, 0x44, 0x4e, 0x4b, 0x45, 0x46, 0x12, 0x30, 0x0a, 0x0b, 0x41,
	0x43, 0x48, 0x4c, 0x47, 0x48, 0x50, 0x50, 0x4d, 0x4e, 0x4f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4f, 0x4a, 0x41, 0x48, 0x4d, 0x44, 0x4c, 0x4e, 0x4d, 0x41, 0x4b, 0x48, 0x00,
	0x52, 0x0b, 0x41, 0x43, 0x48, 0x4c, 0x47, 0x48, 0x50, 0x50, 0x4d, 0x4e, 0x4f, 0x42, 0x0d, 0x0a,
	0x0b, 0x4e, 0x42, 0x45, 0x49, 0x50, 0x47, 0x41, 0x42, 0x46, 0x43, 0x46, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EFCBGCJNNJA_proto_rawDescOnce sync.Once
	file_EFCBGCJNNJA_proto_rawDescData = file_EFCBGCJNNJA_proto_rawDesc
)

func file_EFCBGCJNNJA_proto_rawDescGZIP() []byte {
	file_EFCBGCJNNJA_proto_rawDescOnce.Do(func() {
		file_EFCBGCJNNJA_proto_rawDescData = protoimpl.X.CompressGZIP(file_EFCBGCJNNJA_proto_rawDescData)
	})
	return file_EFCBGCJNNJA_proto_rawDescData
}

var file_EFCBGCJNNJA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EFCBGCJNNJA_proto_goTypes = []interface{}{
	(*EFCBGCJNNJA)(nil), // 0: EFCBGCJNNJA
	(*AGAAGPJDOEA)(nil), // 1: AGAAGPJDOEA
	(*OJAHMDLNMAK)(nil), // 2: OJAHMDLNMAK
}
var file_EFCBGCJNNJA_proto_depIdxs = []int32{
	1, // 0: EFCBGCJNNJA.PPNAPEDNKEF:type_name -> AGAAGPJDOEA
	2, // 1: EFCBGCJNNJA.ACHLGHPPMNO:type_name -> OJAHMDLNMAK
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EFCBGCJNNJA_proto_init() }
func file_EFCBGCJNNJA_proto_init() {
	if File_EFCBGCJNNJA_proto != nil {
		return
	}
	file_OJAHMDLNMAK_proto_init()
	file_AGAAGPJDOEA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EFCBGCJNNJA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EFCBGCJNNJA); i {
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
	file_EFCBGCJNNJA_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EFCBGCJNNJA_PPNAPEDNKEF)(nil),
		(*EFCBGCJNNJA_ACHLGHPPMNO)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_EFCBGCJNNJA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EFCBGCJNNJA_proto_goTypes,
		DependencyIndexes: file_EFCBGCJNNJA_proto_depIdxs,
		MessageInfos:      file_EFCBGCJNNJA_proto_msgTypes,
	}.Build()
	File_EFCBGCJNNJA_proto = out.File
	file_EFCBGCJNNJA_proto_rawDesc = nil
	file_EFCBGCJNNJA_proto_goTypes = nil
	file_EFCBGCJNNJA_proto_depIdxs = nil
}