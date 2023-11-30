// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BoardDataSync.proto

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

type BoardDataSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnlockedHeadIconList []*HeadIcon `protobuf:"bytes,4,rep,name=unlocked_head_icon_list,json=unlockedHeadIconList,proto3" json:"unlocked_head_icon_list,omitempty"`
	Signature            string      `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *BoardDataSync) Reset() {
	*x = BoardDataSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BoardDataSync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardDataSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardDataSync) ProtoMessage() {}

func (x *BoardDataSync) ProtoReflect() protoreflect.Message {
	mi := &file_BoardDataSync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardDataSync.ProtoReflect.Descriptor instead.
func (*BoardDataSync) Descriptor() ([]byte, []int) {
	return file_BoardDataSync_proto_rawDescGZIP(), []int{0}
}

func (x *BoardDataSync) GetUnlockedHeadIconList() []*HeadIcon {
	if x != nil {
		return x.UnlockedHeadIconList
	}
	return nil
}

func (x *BoardDataSync) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_BoardDataSync_proto protoreflect.FileDescriptor

var file_BoardDataSync_proto_rawDesc = []byte{
	0x0a, 0x13, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x79, 0x6e, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x48, 0x65, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0d, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x40, 0x0a, 0x17, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x49, 0x63,
	0x6f, 0x6e, 0x52, 0x14, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64,
	0x49, 0x63, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BoardDataSync_proto_rawDescOnce sync.Once
	file_BoardDataSync_proto_rawDescData = file_BoardDataSync_proto_rawDesc
)

func file_BoardDataSync_proto_rawDescGZIP() []byte {
	file_BoardDataSync_proto_rawDescOnce.Do(func() {
		file_BoardDataSync_proto_rawDescData = protoimpl.X.CompressGZIP(file_BoardDataSync_proto_rawDescData)
	})
	return file_BoardDataSync_proto_rawDescData
}

var file_BoardDataSync_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BoardDataSync_proto_goTypes = []interface{}{
	(*BoardDataSync)(nil), // 0: BoardDataSync
	(*HeadIcon)(nil),      // 1: HeadIcon
}
var file_BoardDataSync_proto_depIdxs = []int32{
	1, // 0: BoardDataSync.unlocked_head_icon_list:type_name -> HeadIcon
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BoardDataSync_proto_init() }
func file_BoardDataSync_proto_init() {
	if File_BoardDataSync_proto != nil {
		return
	}
	file_HeadIcon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BoardDataSync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardDataSync); i {
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
			RawDescriptor: file_BoardDataSync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BoardDataSync_proto_goTypes,
		DependencyIndexes: file_BoardDataSync_proto_depIdxs,
		MessageInfos:      file_BoardDataSync_proto_msgTypes,
	}.Build()
	File_BoardDataSync_proto = out.File
	file_BoardDataSync_proto_rawDesc = nil
	file_BoardDataSync_proto_goTypes = nil
	file_BoardDataSync_proto_depIdxs = nil
}
