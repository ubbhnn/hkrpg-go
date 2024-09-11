// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PVEBattleResultCsReq.proto

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

type PVEBattleResultCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId                uint32            `protobuf:"varint,5,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	OFBINFHOHKN            map[string]uint32 `protobuf:"bytes,7,rep,name=OFBINFHOHKN,proto3" json:"OFBINFHOHKN,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	CostTime               uint32            `protobuf:"varint,13,opt,name=cost_time,json=costTime,proto3" json:"cost_time,omitempty"`
	ClientVersion          uint32            `protobuf:"varint,14,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	IFJFPGFAHGO            uint32            `protobuf:"varint,4,opt,name=IFJFPGFAHGO,proto3" json:"IFJFPGFAHGO,omitempty"`
	IsAutoFight            bool              `protobuf:"varint,15,opt,name=is_auto_fight,json=isAutoFight,proto3" json:"is_auto_fight,omitempty"`
	EndStatus              BattleEndStatus   `protobuf:"varint,1,opt,name=end_status,json=endStatus,proto3,enum=BattleEndStatus" json:"end_status,omitempty"`
	TurnSnapshotHash       []byte            `protobuf:"bytes,8,opt,name=turn_snapshot_hash,json=turnSnapshotHash,proto3" json:"turn_snapshot_hash,omitempty"`
	Stt                    *BattleStatistics `protobuf:"bytes,12,opt,name=stt,proto3" json:"stt,omitempty"`
	IsAiConsiderUltraSkill bool              `protobuf:"varint,2,opt,name=is_ai_consider_ultra_skill,json=isAiConsiderUltraSkill,proto3" json:"is_ai_consider_ultra_skill,omitempty"`
	BattleId               uint32            `protobuf:"varint,9,opt,name=battle_id,json=battleId,proto3" json:"battle_id,omitempty"`
	ALLDMMPDLDL            bool              `protobuf:"varint,3,opt,name=ALLDMMPDLDL,proto3" json:"ALLDMMPDLDL,omitempty"`
	DebugExtraInfo         string            `protobuf:"bytes,11,opt,name=debug_extra_info,json=debugExtraInfo,proto3" json:"debug_extra_info,omitempty"`
	ResVersion             uint32            `protobuf:"varint,10,opt,name=res_version,json=resVersion,proto3" json:"res_version,omitempty"`
	OpList                 []*BattleOp       `protobuf:"bytes,6,rep,name=op_list,json=opList,proto3" json:"op_list,omitempty"`
}

func (x *PVEBattleResultCsReq) Reset() {
	*x = PVEBattleResultCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PVEBattleResultCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PVEBattleResultCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PVEBattleResultCsReq) ProtoMessage() {}

func (x *PVEBattleResultCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_PVEBattleResultCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PVEBattleResultCsReq.ProtoReflect.Descriptor instead.
func (*PVEBattleResultCsReq) Descriptor() ([]byte, []int) {
	return file_PVEBattleResultCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *PVEBattleResultCsReq) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *PVEBattleResultCsReq) GetOFBINFHOHKN() map[string]uint32 {
	if x != nil {
		return x.OFBINFHOHKN
	}
	return nil
}

func (x *PVEBattleResultCsReq) GetCostTime() uint32 {
	if x != nil {
		return x.CostTime
	}
	return 0
}

func (x *PVEBattleResultCsReq) GetClientVersion() uint32 {
	if x != nil {
		return x.ClientVersion
	}
	return 0
}

func (x *PVEBattleResultCsReq) GetIFJFPGFAHGO() uint32 {
	if x != nil {
		return x.IFJFPGFAHGO
	}
	return 0
}

func (x *PVEBattleResultCsReq) GetIsAutoFight() bool {
	if x != nil {
		return x.IsAutoFight
	}
	return false
}

func (x *PVEBattleResultCsReq) GetEndStatus() BattleEndStatus {
	if x != nil {
		return x.EndStatus
	}
	return BattleEndStatus_BATTLE_END_NONE
}

func (x *PVEBattleResultCsReq) GetTurnSnapshotHash() []byte {
	if x != nil {
		return x.TurnSnapshotHash
	}
	return nil
}

func (x *PVEBattleResultCsReq) GetStt() *BattleStatistics {
	if x != nil {
		return x.Stt
	}
	return nil
}

func (x *PVEBattleResultCsReq) GetIsAiConsiderUltraSkill() bool {
	if x != nil {
		return x.IsAiConsiderUltraSkill
	}
	return false
}

func (x *PVEBattleResultCsReq) GetBattleId() uint32 {
	if x != nil {
		return x.BattleId
	}
	return 0
}

func (x *PVEBattleResultCsReq) GetALLDMMPDLDL() bool {
	if x != nil {
		return x.ALLDMMPDLDL
	}
	return false
}

func (x *PVEBattleResultCsReq) GetDebugExtraInfo() string {
	if x != nil {
		return x.DebugExtraInfo
	}
	return ""
}

func (x *PVEBattleResultCsReq) GetResVersion() uint32 {
	if x != nil {
		return x.ResVersion
	}
	return 0
}

func (x *PVEBattleResultCsReq) GetOpList() []*BattleOp {
	if x != nil {
		return x.OpList
	}
	return nil
}

var File_PVEBattleResultCsReq_proto protoreflect.FileDescriptor

var file_PVEBattleResultCsReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x50, 0x56, 0x45, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x4f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x05, 0x0a, 0x14,
	0x50, 0x56, 0x45, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x48, 0x0a, 0x0b, 0x4f, 0x46, 0x42, 0x49, 0x4e, 0x46, 0x48, 0x4f, 0x48, 0x4b, 0x4e, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x50, 0x56, 0x45, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x4f, 0x46, 0x42, 0x49,
	0x4e, 0x46, 0x48, 0x4f, 0x48, 0x4b, 0x4e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4f, 0x46,
	0x42, 0x49, 0x4e, 0x46, 0x48, 0x4f, 0x48, 0x4b, 0x4e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x49, 0x46, 0x4a, 0x46, 0x50, 0x47, 0x46, 0x41, 0x48, 0x47, 0x4f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x49, 0x46, 0x4a, 0x46, 0x50, 0x47, 0x46, 0x41, 0x48, 0x47, 0x4f, 0x12,
	0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x46, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x45, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x10, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x23, 0x0a, 0x03, 0x73, 0x74, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x52, 0x03, 0x73, 0x74, 0x74, 0x12, 0x3a, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x61, 0x69,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x6c, 0x74, 0x72, 0x61, 0x5f,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x73, 0x41,
	0x69, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x72, 0x55, 0x6c, 0x74, 0x72, 0x61, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x41, 0x4c, 0x4c, 0x44, 0x4d, 0x4d, 0x50, 0x44, 0x4c, 0x44, 0x4c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x4c, 0x4c, 0x44, 0x4d, 0x4d, 0x50, 0x44, 0x4c,
	0x44, 0x4c, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a,
	0x07, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x70, 0x52, 0x06, 0x6f, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x3e, 0x0a, 0x10, 0x4f, 0x46, 0x42, 0x49, 0x4e, 0x46, 0x48, 0x4f, 0x48, 0x4b, 0x4e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PVEBattleResultCsReq_proto_rawDescOnce sync.Once
	file_PVEBattleResultCsReq_proto_rawDescData = file_PVEBattleResultCsReq_proto_rawDesc
)

func file_PVEBattleResultCsReq_proto_rawDescGZIP() []byte {
	file_PVEBattleResultCsReq_proto_rawDescOnce.Do(func() {
		file_PVEBattleResultCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PVEBattleResultCsReq_proto_rawDescData)
	})
	return file_PVEBattleResultCsReq_proto_rawDescData
}

var file_PVEBattleResultCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_PVEBattleResultCsReq_proto_goTypes = []interface{}{
	(*PVEBattleResultCsReq)(nil), // 0: PVEBattleResultCsReq
	nil,                          // 1: PVEBattleResultCsReq.OFBINFHOHKNEntry
	(BattleEndStatus)(0),         // 2: BattleEndStatus
	(*BattleStatistics)(nil),     // 3: BattleStatistics
	(*BattleOp)(nil),             // 4: BattleOp
}
var file_PVEBattleResultCsReq_proto_depIdxs = []int32{
	1, // 0: PVEBattleResultCsReq.OFBINFHOHKN:type_name -> PVEBattleResultCsReq.OFBINFHOHKNEntry
	2, // 1: PVEBattleResultCsReq.end_status:type_name -> BattleEndStatus
	3, // 2: PVEBattleResultCsReq.stt:type_name -> BattleStatistics
	4, // 3: PVEBattleResultCsReq.op_list:type_name -> BattleOp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_PVEBattleResultCsReq_proto_init() }
func file_PVEBattleResultCsReq_proto_init() {
	if File_PVEBattleResultCsReq_proto != nil {
		return
	}
	file_BattleOp_proto_init()
	file_BattleEndStatus_proto_init()
	file_BattleStatistics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PVEBattleResultCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PVEBattleResultCsReq); i {
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
			RawDescriptor: file_PVEBattleResultCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PVEBattleResultCsReq_proto_goTypes,
		DependencyIndexes: file_PVEBattleResultCsReq_proto_depIdxs,
		MessageInfos:      file_PVEBattleResultCsReq_proto_msgTypes,
	}.Build()
	File_PVEBattleResultCsReq_proto = out.File
	file_PVEBattleResultCsReq_proto_rawDesc = nil
	file_PVEBattleResultCsReq_proto_goTypes = nil
	file_PVEBattleResultCsReq_proto_depIdxs = nil
}
