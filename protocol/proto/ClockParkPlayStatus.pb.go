// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ClockParkPlayStatus.proto

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

type ClockParkPlayStatus int32

const (
	ClockParkPlayStatus_CLOCK_PARK_PLAY_NONE          ClockParkPlayStatus = 0
	ClockParkPlayStatus_CLOCK_PARK_PLAY_NORMAL_DEATH  ClockParkPlayStatus = 1
	ClockParkPlayStatus_CLOCK_PARK_PLAY_NORMAL_PASS   ClockParkPlayStatus = 2
	ClockParkPlayStatus_CLOCK_PARK_PLAY_FINISH_SCRIPT ClockParkPlayStatus = 5
)

// Enum value maps for ClockParkPlayStatus.
var (
	ClockParkPlayStatus_name = map[int32]string{
		0: "CLOCK_PARK_PLAY_NONE",
		1: "CLOCK_PARK_PLAY_NORMAL_DEATH",
		2: "CLOCK_PARK_PLAY_NORMAL_PASS",
		5: "CLOCK_PARK_PLAY_FINISH_SCRIPT",
	}
	ClockParkPlayStatus_value = map[string]int32{
		"CLOCK_PARK_PLAY_NONE":          0,
		"CLOCK_PARK_PLAY_NORMAL_DEATH":  1,
		"CLOCK_PARK_PLAY_NORMAL_PASS":   2,
		"CLOCK_PARK_PLAY_FINISH_SCRIPT": 5,
	}
)

func (x ClockParkPlayStatus) Enum() *ClockParkPlayStatus {
	p := new(ClockParkPlayStatus)
	*p = x
	return p
}

func (x ClockParkPlayStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClockParkPlayStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_ClockParkPlayStatus_proto_enumTypes[0].Descriptor()
}

func (ClockParkPlayStatus) Type() protoreflect.EnumType {
	return &file_ClockParkPlayStatus_proto_enumTypes[0]
}

func (x ClockParkPlayStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClockParkPlayStatus.Descriptor instead.
func (ClockParkPlayStatus) EnumDescriptor() ([]byte, []int) {
	return file_ClockParkPlayStatus_proto_rawDescGZIP(), []int{0}
}

var File_ClockParkPlayStatus_proto protoreflect.FileDescriptor

var file_ClockParkPlayStatus_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x95, 0x01, 0x0a, 0x13,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x50, 0x41, 0x52,
	0x4b, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x20, 0x0a,
	0x1c, 0x43, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x50, 0x41, 0x52, 0x4b, 0x5f, 0x50, 0x4c, 0x41, 0x59,
	0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x44, 0x45, 0x41, 0x54, 0x48, 0x10, 0x01, 0x12,
	0x1f, 0x0a, 0x1b, 0x43, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x50, 0x41, 0x52, 0x4b, 0x5f, 0x50, 0x4c,
	0x41, 0x59, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x02,
	0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x50, 0x41, 0x52, 0x4b, 0x5f, 0x50,
	0x4c, 0x41, 0x59, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x53, 0x43, 0x52, 0x49, 0x50,
	0x54, 0x10, 0x05, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClockParkPlayStatus_proto_rawDescOnce sync.Once
	file_ClockParkPlayStatus_proto_rawDescData = file_ClockParkPlayStatus_proto_rawDesc
)

func file_ClockParkPlayStatus_proto_rawDescGZIP() []byte {
	file_ClockParkPlayStatus_proto_rawDescOnce.Do(func() {
		file_ClockParkPlayStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClockParkPlayStatus_proto_rawDescData)
	})
	return file_ClockParkPlayStatus_proto_rawDescData
}

var file_ClockParkPlayStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ClockParkPlayStatus_proto_goTypes = []interface{}{
	(ClockParkPlayStatus)(0), // 0: ClockParkPlayStatus
}
var file_ClockParkPlayStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ClockParkPlayStatus_proto_init() }
func file_ClockParkPlayStatus_proto_init() {
	if File_ClockParkPlayStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ClockParkPlayStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClockParkPlayStatus_proto_goTypes,
		DependencyIndexes: file_ClockParkPlayStatus_proto_depIdxs,
		EnumInfos:         file_ClockParkPlayStatus_proto_enumTypes,
	}.Build()
	File_ClockParkPlayStatus_proto = out.File
	file_ClockParkPlayStatus_proto_rawDesc = nil
	file_ClockParkPlayStatus_proto_goTypes = nil
	file_ClockParkPlayStatus_proto_depIdxs = nil
}