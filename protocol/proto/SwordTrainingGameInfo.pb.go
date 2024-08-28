// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SwordTrainingGameInfo.proto

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

type SwordTrainingGameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillInfo               *SwordTrainingSkillInfo  `protobuf:"bytes,15,opt,name=skill_info,json=skillInfo,proto3" json:"skill_info,omitempty"`
	MCKIEJODKGE             *ECIGNEGEAIH             `protobuf:"bytes,3,opt,name=MCKIEJODKGE,proto3" json:"MCKIEJODKGE,omitempty"`
	PowerInfo               *SwordTrainingPowerInfo  `protobuf:"bytes,12,opt,name=power_info,json=powerInfo,proto3" json:"power_info,omitempty"`
	IIHDBEFHEOC             *KCFNOAMOLFC             `protobuf:"bytes,14,opt,name=IIHDBEFHEOC,proto3" json:"IIHDBEFHEOC,omitempty"`
	INJNGCMDBGL             *KKIPILADIGB             `protobuf:"bytes,2,opt,name=INJNGCMDBGL,proto3" json:"INJNGCMDBGL,omitempty"`
	PendingAction           *KKDCJKFPLMF             `protobuf:"bytes,1,opt,name=pending_action,json=pendingAction,proto3" json:"pending_action,omitempty"`
	SkillPower              uint32                   `protobuf:"varint,5,opt,name=skill_power,json=skillPower,proto3" json:"skill_power,omitempty"`
	OptionResultInfo        []*DLBAMGFIACN           `protobuf:"bytes,8,rep,name=option_result_info,json=optionResultInfo,proto3" json:"option_result_info,omitempty"`
	DPHNBKLJEHM             []uint32                 `protobuf:"varint,6,rep,packed,name=DPHNBKLJEHM,proto3" json:"DPHNBKLJEHM,omitempty"`
	SwordTrainingActionInfo *SwordTrainingActionInfo `protobuf:"bytes,11,opt,name=sword_training_action_info,json=swordTrainingActionInfo,proto3" json:"sword_training_action_info,omitempty"`
}

func (x *SwordTrainingGameInfo) Reset() {
	*x = SwordTrainingGameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SwordTrainingGameInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwordTrainingGameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwordTrainingGameInfo) ProtoMessage() {}

func (x *SwordTrainingGameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SwordTrainingGameInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwordTrainingGameInfo.ProtoReflect.Descriptor instead.
func (*SwordTrainingGameInfo) Descriptor() ([]byte, []int) {
	return file_SwordTrainingGameInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SwordTrainingGameInfo) GetSkillInfo() *SwordTrainingSkillInfo {
	if x != nil {
		return x.SkillInfo
	}
	return nil
}

func (x *SwordTrainingGameInfo) GetMCKIEJODKGE() *ECIGNEGEAIH {
	if x != nil {
		return x.MCKIEJODKGE
	}
	return nil
}

func (x *SwordTrainingGameInfo) GetPowerInfo() *SwordTrainingPowerInfo {
	if x != nil {
		return x.PowerInfo
	}
	return nil
}

func (x *SwordTrainingGameInfo) GetIIHDBEFHEOC() *KCFNOAMOLFC {
	if x != nil {
		return x.IIHDBEFHEOC
	}
	return nil
}

func (x *SwordTrainingGameInfo) GetINJNGCMDBGL() *KKIPILADIGB {
	if x != nil {
		return x.INJNGCMDBGL
	}
	return nil
}

func (x *SwordTrainingGameInfo) GetPendingAction() *KKDCJKFPLMF {
	if x != nil {
		return x.PendingAction
	}
	return nil
}

func (x *SwordTrainingGameInfo) GetSkillPower() uint32 {
	if x != nil {
		return x.SkillPower
	}
	return 0
}

func (x *SwordTrainingGameInfo) GetOptionResultInfo() []*DLBAMGFIACN {
	if x != nil {
		return x.OptionResultInfo
	}
	return nil
}

func (x *SwordTrainingGameInfo) GetDPHNBKLJEHM() []uint32 {
	if x != nil {
		return x.DPHNBKLJEHM
	}
	return nil
}

func (x *SwordTrainingGameInfo) GetSwordTrainingActionInfo() *SwordTrainingActionInfo {
	if x != nil {
		return x.SwordTrainingActionInfo
	}
	return nil
}

var File_SwordTrainingGameInfo_proto protoreflect.FileDescriptor

var file_SwordTrainingGameInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45,
	0x43, 0x49, 0x47, 0x4e, 0x45, 0x47, 0x45, 0x41, 0x49, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x4b, 0x43, 0x46, 0x4e, 0x4f, 0x41, 0x4d, 0x4f, 0x4c, 0x46, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x44, 0x4c, 0x42, 0x41, 0x4d, 0x47, 0x46, 0x49, 0x41, 0x43, 0x4e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4b, 0x49, 0x50, 0x49, 0x4c, 0x41, 0x44, 0x49, 0x47, 0x42, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4b, 0x44, 0x43, 0x4a, 0x4b, 0x46, 0x50, 0x4c, 0x4d,
	0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x04, 0x0a, 0x15, 0x53, 0x77, 0x6f, 0x72,
	0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x4d, 0x43, 0x4b,
	0x49, 0x45, 0x4a, 0x4f, 0x44, 0x4b, 0x47, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x45, 0x43, 0x49, 0x47, 0x4e, 0x45, 0x47, 0x45, 0x41, 0x49, 0x48, 0x52, 0x0b, 0x4d, 0x43,
	0x4b, 0x49, 0x45, 0x4a, 0x4f, 0x44, 0x4b, 0x47, 0x45, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x49, 0x48, 0x44, 0x42, 0x45, 0x46, 0x48, 0x45, 0x4f, 0x43,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x43, 0x46, 0x4e, 0x4f, 0x41, 0x4d,
	0x4f, 0x4c, 0x46, 0x43, 0x52, 0x0b, 0x49, 0x49, 0x48, 0x44, 0x42, 0x45, 0x46, 0x48, 0x45, 0x4f,
	0x43, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x4e, 0x4a, 0x4e, 0x47, 0x43, 0x4d, 0x44, 0x42, 0x47, 0x4c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4b, 0x49, 0x50, 0x49, 0x4c, 0x41,
	0x44, 0x49, 0x47, 0x42, 0x52, 0x0b, 0x49, 0x4e, 0x4a, 0x4e, 0x47, 0x43, 0x4d, 0x44, 0x42, 0x47,
	0x4c, 0x12, 0x33, 0x0a, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4b, 0x44, 0x43,
	0x4a, 0x4b, 0x46, 0x50, 0x4c, 0x4d, 0x46, 0x52, 0x0d, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4c, 0x42, 0x41, 0x4d, 0x47, 0x46, 0x49, 0x41, 0x43,
	0x4e, 0x52, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x50, 0x48, 0x4e, 0x42, 0x4b, 0x4c, 0x4a, 0x45,
	0x48, 0x4d, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x50, 0x48, 0x4e, 0x42, 0x4b,
	0x4c, 0x4a, 0x45, 0x48, 0x4d, 0x12, 0x55, 0x0a, 0x1a, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x53, 0x77, 0x6f, 0x72,
	0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x17, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SwordTrainingGameInfo_proto_rawDescOnce sync.Once
	file_SwordTrainingGameInfo_proto_rawDescData = file_SwordTrainingGameInfo_proto_rawDesc
)

func file_SwordTrainingGameInfo_proto_rawDescGZIP() []byte {
	file_SwordTrainingGameInfo_proto_rawDescOnce.Do(func() {
		file_SwordTrainingGameInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SwordTrainingGameInfo_proto_rawDescData)
	})
	return file_SwordTrainingGameInfo_proto_rawDescData
}

var file_SwordTrainingGameInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SwordTrainingGameInfo_proto_goTypes = []interface{}{
	(*SwordTrainingGameInfo)(nil),   // 0: SwordTrainingGameInfo
	(*SwordTrainingSkillInfo)(nil),  // 1: SwordTrainingSkillInfo
	(*ECIGNEGEAIH)(nil),             // 2: ECIGNEGEAIH
	(*SwordTrainingPowerInfo)(nil),  // 3: SwordTrainingPowerInfo
	(*KCFNOAMOLFC)(nil),             // 4: KCFNOAMOLFC
	(*KKIPILADIGB)(nil),             // 5: KKIPILADIGB
	(*KKDCJKFPLMF)(nil),             // 6: KKDCJKFPLMF
	(*DLBAMGFIACN)(nil),             // 7: DLBAMGFIACN
	(*SwordTrainingActionInfo)(nil), // 8: SwordTrainingActionInfo
}
var file_SwordTrainingGameInfo_proto_depIdxs = []int32{
	1, // 0: SwordTrainingGameInfo.skill_info:type_name -> SwordTrainingSkillInfo
	2, // 1: SwordTrainingGameInfo.MCKIEJODKGE:type_name -> ECIGNEGEAIH
	3, // 2: SwordTrainingGameInfo.power_info:type_name -> SwordTrainingPowerInfo
	4, // 3: SwordTrainingGameInfo.IIHDBEFHEOC:type_name -> KCFNOAMOLFC
	5, // 4: SwordTrainingGameInfo.INJNGCMDBGL:type_name -> KKIPILADIGB
	6, // 5: SwordTrainingGameInfo.pending_action:type_name -> KKDCJKFPLMF
	7, // 6: SwordTrainingGameInfo.option_result_info:type_name -> DLBAMGFIACN
	8, // 7: SwordTrainingGameInfo.sword_training_action_info:type_name -> SwordTrainingActionInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_SwordTrainingGameInfo_proto_init() }
func file_SwordTrainingGameInfo_proto_init() {
	if File_SwordTrainingGameInfo_proto != nil {
		return
	}
	file_ECIGNEGEAIH_proto_init()
	file_SwordTrainingSkillInfo_proto_init()
	file_KCFNOAMOLFC_proto_init()
	file_SwordTrainingPowerInfo_proto_init()
	file_DLBAMGFIACN_proto_init()
	file_KKIPILADIGB_proto_init()
	file_SwordTrainingActionInfo_proto_init()
	file_KKDCJKFPLMF_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SwordTrainingGameInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwordTrainingGameInfo); i {
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
			RawDescriptor: file_SwordTrainingGameInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SwordTrainingGameInfo_proto_goTypes,
		DependencyIndexes: file_SwordTrainingGameInfo_proto_depIdxs,
		MessageInfos:      file_SwordTrainingGameInfo_proto_msgTypes,
	}.Build()
	File_SwordTrainingGameInfo_proto = out.File
	file_SwordTrainingGameInfo_proto_rawDesc = nil
	file_SwordTrainingGameInfo_proto_goTypes = nil
	file_SwordTrainingGameInfo_proto_depIdxs = nil
}
