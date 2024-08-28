// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueLineupInfo.proto

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

type ChessRogueLineupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChessAvatarList []*ChessRogueLineupAvatarInfo `protobuf:"bytes,11,rep,name=chess_avatar_list,json=chessAvatarList,proto3" json:"chess_avatar_list,omitempty"`
	ReviveInfo      *RogueAvatarReviveCost        `protobuf:"bytes,10,opt,name=revive_info,json=reviveInfo,proto3" json:"revive_info,omitempty"`
}

func (x *ChessRogueLineupInfo) Reset() {
	*x = ChessRogueLineupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueLineupInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueLineupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueLineupInfo) ProtoMessage() {}

func (x *ChessRogueLineupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueLineupInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueLineupInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueLineupInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueLineupInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueLineupInfo) GetChessAvatarList() []*ChessRogueLineupAvatarInfo {
	if x != nil {
		return x.ChessAvatarList
	}
	return nil
}

func (x *ChessRogueLineupInfo) GetReviveInfo() *RogueAvatarReviveCost {
	if x != nil {
		return x.ReviveInfo
	}
	return nil
}

var File_ChessRogueLineupInfo_proto protoreflect.FileDescriptor

var file_ChessRogueLineupInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x76, 0x69, 0x76,
	0x65, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x14,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x11, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x63, 0x68,
	0x65, 0x73, 0x73, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x0b, 0x72, 0x65, 0x76, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x52, 0x65, 0x76, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69,
	0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueLineupInfo_proto_rawDescOnce sync.Once
	file_ChessRogueLineupInfo_proto_rawDescData = file_ChessRogueLineupInfo_proto_rawDesc
)

func file_ChessRogueLineupInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueLineupInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueLineupInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueLineupInfo_proto_rawDescData)
	})
	return file_ChessRogueLineupInfo_proto_rawDescData
}

var file_ChessRogueLineupInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueLineupInfo_proto_goTypes = []interface{}{
	(*ChessRogueLineupInfo)(nil),       // 0: ChessRogueLineupInfo
	(*ChessRogueLineupAvatarInfo)(nil), // 1: ChessRogueLineupAvatarInfo
	(*RogueAvatarReviveCost)(nil),      // 2: RogueAvatarReviveCost
}
var file_ChessRogueLineupInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueLineupInfo.chess_avatar_list:type_name -> ChessRogueLineupAvatarInfo
	2, // 1: ChessRogueLineupInfo.revive_info:type_name -> RogueAvatarReviveCost
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ChessRogueLineupInfo_proto_init() }
func file_ChessRogueLineupInfo_proto_init() {
	if File_ChessRogueLineupInfo_proto != nil {
		return
	}
	file_ChessRogueLineupAvatarInfo_proto_init()
	file_RogueAvatarReviveCost_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueLineupInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueLineupInfo); i {
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
			RawDescriptor: file_ChessRogueLineupInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueLineupInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueLineupInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueLineupInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueLineupInfo_proto = out.File
	file_ChessRogueLineupInfo_proto_rawDesc = nil
	file_ChessRogueLineupInfo_proto_goTypes = nil
	file_ChessRogueLineupInfo_proto_depIdxs = nil
}
