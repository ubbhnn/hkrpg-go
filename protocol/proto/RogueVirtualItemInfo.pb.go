// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: RogueVirtualItemInfo.proto

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

type RogueVirtualItemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoguePumanCoupon  uint32 `protobuf:"varint,8,opt,name=rogue_puman_coupon,json=roguePumanCoupon,proto3" json:"rogue_puman_coupon,omitempty"`
	RogueCoin         uint32 `protobuf:"varint,10,opt,name=rogue_coin,json=rogueCoin,proto3" json:"rogue_coin,omitempty"`
	RogueImmersifier  uint32 `protobuf:"varint,13,opt,name=rogue_immersifier,json=rogueImmersifier,proto3" json:"rogue_immersifier,omitempty"`
	RogueAbilityPoint uint32 `protobuf:"varint,7,opt,name=rogue_ability_point,json=rogueAbilityPoint,proto3" json:"rogue_ability_point,omitempty"`
}

func (x *RogueVirtualItemInfo) Reset() {
	*x = RogueVirtualItemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueVirtualItemInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueVirtualItemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueVirtualItemInfo) ProtoMessage() {}

func (x *RogueVirtualItemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueVirtualItemInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueVirtualItemInfo.ProtoReflect.Descriptor instead.
func (*RogueVirtualItemInfo) Descriptor() ([]byte, []int) {
	return file_RogueVirtualItemInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueVirtualItemInfo) GetRoguePumanCoupon() uint32 {
	if x != nil {
		return x.RoguePumanCoupon
	}
	return 0
}

func (x *RogueVirtualItemInfo) GetRogueCoin() uint32 {
	if x != nil {
		return x.RogueCoin
	}
	return 0
}

func (x *RogueVirtualItemInfo) GetRogueImmersifier() uint32 {
	if x != nil {
		return x.RogueImmersifier
	}
	return 0
}

func (x *RogueVirtualItemInfo) GetRogueAbilityPoint() uint32 {
	if x != nil {
		return x.RogueAbilityPoint
	}
	return 0
}

var File_RogueVirtualItemInfo_proto protoreflect.FileDescriptor

var file_RogueVirtualItemInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a,
	0x14, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x70,
	0x75, 0x6d, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x50, 0x75, 0x6d, 0x61, 0x6e, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x69,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f,
	0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x69, 0x6d, 0x6d, 0x65,
	0x72, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x49, 0x6d, 0x6d, 0x65, 0x72, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x2e, 0x0a, 0x13, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x72, 0x6f,
	0x67, 0x75, 0x65, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_RogueVirtualItemInfo_proto_rawDescOnce sync.Once
	file_RogueVirtualItemInfo_proto_rawDescData = file_RogueVirtualItemInfo_proto_rawDesc
)

func file_RogueVirtualItemInfo_proto_rawDescGZIP() []byte {
	file_RogueVirtualItemInfo_proto_rawDescOnce.Do(func() {
		file_RogueVirtualItemInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueVirtualItemInfo_proto_rawDescData)
	})
	return file_RogueVirtualItemInfo_proto_rawDescData
}

var file_RogueVirtualItemInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueVirtualItemInfo_proto_goTypes = []interface{}{
	(*RogueVirtualItemInfo)(nil), // 0: RogueVirtualItemInfo
}
var file_RogueVirtualItemInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueVirtualItemInfo_proto_init() }
func file_RogueVirtualItemInfo_proto_init() {
	if File_RogueVirtualItemInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueVirtualItemInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueVirtualItemInfo); i {
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
			RawDescriptor: file_RogueVirtualItemInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueVirtualItemInfo_proto_goTypes,
		DependencyIndexes: file_RogueVirtualItemInfo_proto_depIdxs,
		MessageInfos:      file_RogueVirtualItemInfo_proto_msgTypes,
	}.Build()
	File_RogueVirtualItemInfo_proto = out.File
	file_RogueVirtualItemInfo_proto_rawDesc = nil
	file_RogueVirtualItemInfo_proto_goTypes = nil
	file_RogueVirtualItemInfo_proto_depIdxs = nil
}
