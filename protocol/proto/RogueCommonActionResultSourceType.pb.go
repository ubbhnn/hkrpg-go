// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueCommonActionResultSourceType.proto

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

type RogueCommonActionResultSourceType int32

const (
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_NONE            RogueCommonActionResultSourceType = 0
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_SELECT          RogueCommonActionResultSourceType = 1
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_ENHANCE         RogueCommonActionResultSourceType = 2
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_MIRACLE         RogueCommonActionResultSourceType = 3
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_DIALOGUE        RogueCommonActionResultSourceType = 4
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_BONUS           RogueCommonActionResultSourceType = 5
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_SHOP            RogueCommonActionResultSourceType = 6
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_DICE            RogueCommonActionResultSourceType = 7
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_AEON            RogueCommonActionResultSourceType = 8
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_BOARD_EVENT     RogueCommonActionResultSourceType = 9
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_MAZE_SKILL      RogueCommonActionResultSourceType = 10
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_LEVEL_MECHANISM RogueCommonActionResultSourceType = 11
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_BUFF            RogueCommonActionResultSourceType = 12
	RogueCommonActionResultSourceType_ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_REFORGE         RogueCommonActionResultSourceType = 13
)

// Enum value maps for RogueCommonActionResultSourceType.
var (
	RogueCommonActionResultSourceType_name = map[int32]string{
		0:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_NONE",
		1:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_SELECT",
		2:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_ENHANCE",
		3:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_MIRACLE",
		4:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_DIALOGUE",
		5:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_BONUS",
		6:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_SHOP",
		7:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_DICE",
		8:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_AEON",
		9:  "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_BOARD_EVENT",
		10: "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_MAZE_SKILL",
		11: "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_LEVEL_MECHANISM",
		12: "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_BUFF",
		13: "ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_REFORGE",
	}
	RogueCommonActionResultSourceType_value = map[string]int32{
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_NONE":            0,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_SELECT":          1,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_ENHANCE":         2,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_MIRACLE":         3,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_DIALOGUE":        4,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_BONUS":           5,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_SHOP":            6,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_DICE":            7,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_AEON":            8,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_BOARD_EVENT":     9,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_MAZE_SKILL":      10,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_LEVEL_MECHANISM": 11,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_BUFF":            12,
		"ROGUE_COMMON_ACTION_RESULT_SOURCE_TYPE_REFORGE":         13,
	}
)

func (x RogueCommonActionResultSourceType) Enum() *RogueCommonActionResultSourceType {
	p := new(RogueCommonActionResultSourceType)
	*p = x
	return p
}

func (x RogueCommonActionResultSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RogueCommonActionResultSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_RogueCommonActionResultSourceType_proto_enumTypes[0].Descriptor()
}

func (RogueCommonActionResultSourceType) Type() protoreflect.EnumType {
	return &file_RogueCommonActionResultSourceType_proto_enumTypes[0]
}

func (x RogueCommonActionResultSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RogueCommonActionResultSourceType.Descriptor instead.
func (RogueCommonActionResultSourceType) EnumDescriptor() ([]byte, []int) {
	return file_RogueCommonActionResultSourceType_proto_rawDescGZIP(), []int{0}
}

var File_RogueCommonActionResultSourceType_proto protoreflect.FileDescriptor

var file_RogueCommonActionResultSourceType_proto_rawDesc = []byte{
	0x0a, 0x27, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf9, 0x05, 0x0a, 0x21, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2f, 0x0a, 0x2b, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x31, 0x0a, 0x2d, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43,
	0x54, 0x10, 0x01, 0x12, 0x32, 0x0a, 0x2e, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d,
	0x4d, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c,
	0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e,
	0x48, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x32, 0x0a, 0x2e, 0x52, 0x4f, 0x47, 0x55, 0x45,
	0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x49, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x33, 0x0a, 0x2f, 0x52,
	0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x4f, 0x47, 0x55, 0x45, 0x10, 0x04,
	0x12, 0x30, 0x0a, 0x2c, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4e, 0x55, 0x53,
	0x10, 0x05, 0x12, 0x2f, 0x0a, 0x2b, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x48, 0x4f,
	0x50, 0x10, 0x06, 0x12, 0x2f, 0x0a, 0x2b, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d,
	0x4d, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c,
	0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49,
	0x43, 0x45, 0x10, 0x07, 0x12, 0x2f, 0x0a, 0x2b, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x45, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x36, 0x0a, 0x32, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x42, 0x4f, 0x41, 0x52, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x09, 0x12, 0x35, 0x0a,
	0x31, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x5a, 0x45, 0x5f, 0x53, 0x4b, 0x49,
	0x4c, 0x4c, 0x10, 0x0a, 0x12, 0x3a, 0x0a, 0x36, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x53, 0x4d, 0x10, 0x0b,
	0x12, 0x2f, 0x0a, 0x2b, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x10,
	0x0c, 0x12, 0x32, 0x0a, 0x2e, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f,
	0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x46, 0x4f,
	0x52, 0x47, 0x45, 0x10, 0x0d, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueCommonActionResultSourceType_proto_rawDescOnce sync.Once
	file_RogueCommonActionResultSourceType_proto_rawDescData = file_RogueCommonActionResultSourceType_proto_rawDesc
)

func file_RogueCommonActionResultSourceType_proto_rawDescGZIP() []byte {
	file_RogueCommonActionResultSourceType_proto_rawDescOnce.Do(func() {
		file_RogueCommonActionResultSourceType_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueCommonActionResultSourceType_proto_rawDescData)
	})
	return file_RogueCommonActionResultSourceType_proto_rawDescData
}

var file_RogueCommonActionResultSourceType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RogueCommonActionResultSourceType_proto_goTypes = []interface{}{
	(RogueCommonActionResultSourceType)(0), // 0: RogueCommonActionResultSourceType
}
var file_RogueCommonActionResultSourceType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueCommonActionResultSourceType_proto_init() }
func file_RogueCommonActionResultSourceType_proto_init() {
	if File_RogueCommonActionResultSourceType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueCommonActionResultSourceType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueCommonActionResultSourceType_proto_goTypes,
		DependencyIndexes: file_RogueCommonActionResultSourceType_proto_depIdxs,
		EnumInfos:         file_RogueCommonActionResultSourceType_proto_enumTypes,
	}.Build()
	File_RogueCommonActionResultSourceType_proto = out.File
	file_RogueCommonActionResultSourceType_proto_rawDesc = nil
	file_RogueCommonActionResultSourceType_proto_goTypes = nil
	file_RogueCommonActionResultSourceType_proto_depIdxs = nil
}