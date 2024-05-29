// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MissionSyncRecord.proto

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

type MissionSyncRecord int32

const (
	MissionSyncRecord_MISSION_SYNC_RECORD_NONE                MissionSyncRecord = 0
	MissionSyncRecord_MISSION_SYNC_RECORD_MAIN_MISSION_ACCEPT MissionSyncRecord = 1
	MissionSyncRecord_MISSION_SYNC_RECORD_MAIN_MISSION_START  MissionSyncRecord = 2
	MissionSyncRecord_MISSION_SYNC_RECORD_MAIN_MISSION_FINISH MissionSyncRecord = 3
	MissionSyncRecord_MISSION_SYNC_RECORD_MAIN_MISSION_DELETE MissionSyncRecord = 4
	MissionSyncRecord_MISSION_SYNC_RECORD_MISSION_ACCEPT      MissionSyncRecord = 11
	MissionSyncRecord_MISSION_SYNC_RECORD_MISSION_START       MissionSyncRecord = 12
	MissionSyncRecord_MISSION_SYNC_RECORD_MISSION_FINISH      MissionSyncRecord = 13
	MissionSyncRecord_MISSION_SYNC_RECORD_MISSION_DELETE      MissionSyncRecord = 14
	MissionSyncRecord_MISSION_SYNC_RECORD_MISSION_PROGRESS    MissionSyncRecord = 15
)

// Enum value maps for MissionSyncRecord.
var (
	MissionSyncRecord_name = map[int32]string{
		0:  "MISSION_SYNC_RECORD_NONE",
		1:  "MISSION_SYNC_RECORD_MAIN_MISSION_ACCEPT",
		2:  "MISSION_SYNC_RECORD_MAIN_MISSION_START",
		3:  "MISSION_SYNC_RECORD_MAIN_MISSION_FINISH",
		4:  "MISSION_SYNC_RECORD_MAIN_MISSION_DELETE",
		11: "MISSION_SYNC_RECORD_MISSION_ACCEPT",
		12: "MISSION_SYNC_RECORD_MISSION_START",
		13: "MISSION_SYNC_RECORD_MISSION_FINISH",
		14: "MISSION_SYNC_RECORD_MISSION_DELETE",
		15: "MISSION_SYNC_RECORD_MISSION_PROGRESS",
	}
	MissionSyncRecord_value = map[string]int32{
		"MISSION_SYNC_RECORD_NONE":                0,
		"MISSION_SYNC_RECORD_MAIN_MISSION_ACCEPT": 1,
		"MISSION_SYNC_RECORD_MAIN_MISSION_START":  2,
		"MISSION_SYNC_RECORD_MAIN_MISSION_FINISH": 3,
		"MISSION_SYNC_RECORD_MAIN_MISSION_DELETE": 4,
		"MISSION_SYNC_RECORD_MISSION_ACCEPT":      11,
		"MISSION_SYNC_RECORD_MISSION_START":       12,
		"MISSION_SYNC_RECORD_MISSION_FINISH":      13,
		"MISSION_SYNC_RECORD_MISSION_DELETE":      14,
		"MISSION_SYNC_RECORD_MISSION_PROGRESS":    15,
	}
)

func (x MissionSyncRecord) Enum() *MissionSyncRecord {
	p := new(MissionSyncRecord)
	*p = x
	return p
}

func (x MissionSyncRecord) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MissionSyncRecord) Descriptor() protoreflect.EnumDescriptor {
	return file_MissionSyncRecord_proto_enumTypes[0].Descriptor()
}

func (MissionSyncRecord) Type() protoreflect.EnumType {
	return &file_MissionSyncRecord_proto_enumTypes[0]
}

func (x MissionSyncRecord) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MissionSyncRecord.Descriptor instead.
func (MissionSyncRecord) EnumDescriptor() ([]byte, []int) {
	return file_MissionSyncRecord_proto_rawDescGZIP(), []int{0}
}

var File_MissionSyncRecord_proto protoreflect.FileDescriptor

var file_MissionSyncRecord_proto_rawDesc = []byte{
	0x0a, 0x17, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xad, 0x03, 0x0a, 0x11, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x1c, 0x0a, 0x18, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x2b, 0x0a,
	0x27, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x52, 0x45,
	0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x01, 0x12, 0x2a, 0x0a, 0x26, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52,
	0x44, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x10, 0x02, 0x12, 0x2b, 0x0a, 0x27, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x41,
	0x49, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53,
	0x48, 0x10, 0x03, 0x12, 0x2b, 0x0a, 0x27, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x59, 0x4e, 0x43, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04,
	0x12, 0x26, 0x0a, 0x22, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43,
	0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x0b, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x0c, 0x12,
	0x26, 0x0a, 0x22, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x46,
	0x49, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x0d, 0x12, 0x26, 0x0a, 0x22, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x0e, 0x12,
	0x28, 0x0a, 0x24, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x0f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MissionSyncRecord_proto_rawDescOnce sync.Once
	file_MissionSyncRecord_proto_rawDescData = file_MissionSyncRecord_proto_rawDesc
)

func file_MissionSyncRecord_proto_rawDescGZIP() []byte {
	file_MissionSyncRecord_proto_rawDescOnce.Do(func() {
		file_MissionSyncRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_MissionSyncRecord_proto_rawDescData)
	})
	return file_MissionSyncRecord_proto_rawDescData
}

var file_MissionSyncRecord_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MissionSyncRecord_proto_goTypes = []interface{}{
	(MissionSyncRecord)(0), // 0: MissionSyncRecord
}
var file_MissionSyncRecord_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MissionSyncRecord_proto_init() }
func file_MissionSyncRecord_proto_init() {
	if File_MissionSyncRecord_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MissionSyncRecord_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MissionSyncRecord_proto_goTypes,
		DependencyIndexes: file_MissionSyncRecord_proto_depIdxs,
		EnumInfos:         file_MissionSyncRecord_proto_enumTypes,
	}.Build()
	File_MissionSyncRecord_proto = out.File
	file_MissionSyncRecord_proto_rawDesc = nil
	file_MissionSyncRecord_proto_goTypes = nil
	file_MissionSyncRecord_proto_depIdxs = nil
}