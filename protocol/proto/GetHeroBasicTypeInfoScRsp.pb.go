// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetHeroBasicTypeInfoScRsp.proto

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

type GetHeroBasicTypeInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurBasicType      HeroBasicType        `protobuf:"varint,11,opt,name=cur_basic_type,json=curBasicType,proto3,enum=HeroBasicType" json:"cur_basic_type,omitempty"`
	BasicTypeInfoList []*HeroBasicTypeInfo `protobuf:"bytes,6,rep,name=basic_type_info_list,json=basicTypeInfoList,proto3" json:"basic_type_info_list,omitempty"`
	Gender            Gender               `protobuf:"varint,5,opt,name=gender,proto3,enum=Gender" json:"gender,omitempty"`
	Retcode           uint32               `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetHeroBasicTypeInfoScRsp) Reset() {
	*x = GetHeroBasicTypeInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetHeroBasicTypeInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeroBasicTypeInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeroBasicTypeInfoScRsp) ProtoMessage() {}

func (x *GetHeroBasicTypeInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetHeroBasicTypeInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeroBasicTypeInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetHeroBasicTypeInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetHeroBasicTypeInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetHeroBasicTypeInfoScRsp) GetCurBasicType() HeroBasicType {
	if x != nil {
		return x.CurBasicType
	}
	return HeroBasicType_None
}

func (x *GetHeroBasicTypeInfoScRsp) GetBasicTypeInfoList() []*HeroBasicTypeInfo {
	if x != nil {
		return x.BasicTypeInfoList
	}
	return nil
}

func (x *GetHeroBasicTypeInfoScRsp) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_GenderNone
}

func (x *GetHeroBasicTypeInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetHeroBasicTypeInfoScRsp_proto protoreflect.FileDescriptor

var file_GetHeroBasicTypeInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x0e, 0x63,
	0x75, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x43, 0x0a, 0x14, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetHeroBasicTypeInfoScRsp_proto_rawDescOnce sync.Once
	file_GetHeroBasicTypeInfoScRsp_proto_rawDescData = file_GetHeroBasicTypeInfoScRsp_proto_rawDesc
)

func file_GetHeroBasicTypeInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetHeroBasicTypeInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetHeroBasicTypeInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetHeroBasicTypeInfoScRsp_proto_rawDescData)
	})
	return file_GetHeroBasicTypeInfoScRsp_proto_rawDescData
}

var file_GetHeroBasicTypeInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetHeroBasicTypeInfoScRsp_proto_goTypes = []interface{}{
	(*GetHeroBasicTypeInfoScRsp)(nil), // 0: GetHeroBasicTypeInfoScRsp
	(HeroBasicType)(0),                // 1: HeroBasicType
	(*HeroBasicTypeInfo)(nil),         // 2: HeroBasicTypeInfo
	(Gender)(0),                       // 3: Gender
}
var file_GetHeroBasicTypeInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetHeroBasicTypeInfoScRsp.cur_basic_type:type_name -> HeroBasicType
	2, // 1: GetHeroBasicTypeInfoScRsp.basic_type_info_list:type_name -> HeroBasicTypeInfo
	3, // 2: GetHeroBasicTypeInfoScRsp.gender:type_name -> Gender
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetHeroBasicTypeInfoScRsp_proto_init() }
func file_GetHeroBasicTypeInfoScRsp_proto_init() {
	if File_GetHeroBasicTypeInfoScRsp_proto != nil {
		return
	}
	file_HeroBasicType_proto_init()
	file_Gender_proto_init()
	file_HeroBasicTypeInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetHeroBasicTypeInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHeroBasicTypeInfoScRsp); i {
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
			RawDescriptor: file_GetHeroBasicTypeInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetHeroBasicTypeInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetHeroBasicTypeInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetHeroBasicTypeInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetHeroBasicTypeInfoScRsp_proto = out.File
	file_GetHeroBasicTypeInfoScRsp_proto_rawDesc = nil
	file_GetHeroBasicTypeInfoScRsp_proto_goTypes = nil
	file_GetHeroBasicTypeInfoScRsp_proto_depIdxs = nil
}
