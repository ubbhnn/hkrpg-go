// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RelicFilterPlan.proto

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

type RelicFilterPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IconFieldNumber            *RelicFilterPlanIcon     `protobuf:"bytes,3,opt,name=IconFieldNumber,proto3" json:"IconFieldNumber,omitempty"`
	UpdateTimestampFieldNumber int64                    `protobuf:"varint,5,opt,name=UpdateTimestampFieldNumber,proto3" json:"UpdateTimestampFieldNumber,omitempty"`
	SettingsFieldNumber        *RelicFilterPlanSettings `protobuf:"bytes,8,opt,name=SettingsFieldNumber,proto3" json:"SettingsFieldNumber,omitempty"`
	NameFieldNumber            string                   `protobuf:"bytes,12,opt,name=NameFieldNumber,proto3" json:"NameFieldNumber,omitempty"`
	SlotIndexFieldNumber       uint32                   `protobuf:"varint,15,opt,name=SlotIndexFieldNumber,proto3" json:"SlotIndexFieldNumber,omitempty"`
	IsMarkedFieldNumber        bool                     `protobuf:"varint,11,opt,name=IsMarkedFieldNumber,proto3" json:"IsMarkedFieldNumber,omitempty"`
}

func (x *RelicFilterPlan) Reset() {
	*x = RelicFilterPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RelicFilterPlan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelicFilterPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelicFilterPlan) ProtoMessage() {}

func (x *RelicFilterPlan) ProtoReflect() protoreflect.Message {
	mi := &file_RelicFilterPlan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelicFilterPlan.ProtoReflect.Descriptor instead.
func (*RelicFilterPlan) Descriptor() ([]byte, []int) {
	return file_RelicFilterPlan_proto_rawDescGZIP(), []int{0}
}

func (x *RelicFilterPlan) GetIconFieldNumber() *RelicFilterPlanIcon {
	if x != nil {
		return x.IconFieldNumber
	}
	return nil
}

func (x *RelicFilterPlan) GetUpdateTimestampFieldNumber() int64 {
	if x != nil {
		return x.UpdateTimestampFieldNumber
	}
	return 0
}

func (x *RelicFilterPlan) GetSettingsFieldNumber() *RelicFilterPlanSettings {
	if x != nil {
		return x.SettingsFieldNumber
	}
	return nil
}

func (x *RelicFilterPlan) GetNameFieldNumber() string {
	if x != nil {
		return x.NameFieldNumber
	}
	return ""
}

func (x *RelicFilterPlan) GetSlotIndexFieldNumber() uint32 {
	if x != nil {
		return x.SlotIndexFieldNumber
	}
	return 0
}

func (x *RelicFilterPlan) GetIsMarkedFieldNumber() bool {
	if x != nil {
		return x.IsMarkedFieldNumber
	}
	return false
}

var File_RelicFilterPlan_proto protoreflect.FileDescriptor

var file_RelicFilterPlan_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50,
	0x6c, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xed, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x3e, 0x0a, 0x0f, 0x49, 0x63, 0x6f, 0x6e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e,
	0x49, 0x63, 0x6f, 0x6e, 0x52, 0x0f, 0x49, 0x63, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x50, 0x6c, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x13, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x28, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4e, 0x61, 0x6d, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x14, 0x53,
	0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x53, 0x6c, 0x6f, 0x74, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x13, 0x49, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x49, 0x73,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RelicFilterPlan_proto_rawDescOnce sync.Once
	file_RelicFilterPlan_proto_rawDescData = file_RelicFilterPlan_proto_rawDesc
)

func file_RelicFilterPlan_proto_rawDescGZIP() []byte {
	file_RelicFilterPlan_proto_rawDescOnce.Do(func() {
		file_RelicFilterPlan_proto_rawDescData = protoimpl.X.CompressGZIP(file_RelicFilterPlan_proto_rawDescData)
	})
	return file_RelicFilterPlan_proto_rawDescData
}

var file_RelicFilterPlan_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RelicFilterPlan_proto_goTypes = []interface{}{
	(*RelicFilterPlan)(nil),         // 0: RelicFilterPlan
	(*RelicFilterPlanIcon)(nil),     // 1: RelicFilterPlanIcon
	(*RelicFilterPlanSettings)(nil), // 2: RelicFilterPlanSettings
}
var file_RelicFilterPlan_proto_depIdxs = []int32{
	1, // 0: RelicFilterPlan.IconFieldNumber:type_name -> RelicFilterPlanIcon
	2, // 1: RelicFilterPlan.SettingsFieldNumber:type_name -> RelicFilterPlanSettings
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RelicFilterPlan_proto_init() }
func file_RelicFilterPlan_proto_init() {
	if File_RelicFilterPlan_proto != nil {
		return
	}
	file_RelicFilterPlanIcon_proto_init()
	file_RelicFilterPlanSettings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RelicFilterPlan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelicFilterPlan); i {
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
			RawDescriptor: file_RelicFilterPlan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RelicFilterPlan_proto_goTypes,
		DependencyIndexes: file_RelicFilterPlan_proto_depIdxs,
		MessageInfos:      file_RelicFilterPlan_proto_msgTypes,
	}.Build()
	File_RelicFilterPlan_proto = out.File
	file_RelicFilterPlan_proto_rawDesc = nil
	file_RelicFilterPlan_proto_goTypes = nil
	file_RelicFilterPlan_proto_depIdxs = nil
}
