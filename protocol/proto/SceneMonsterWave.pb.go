// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneMonsterWave.proto

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

type SceneMonsterWave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DropList      []*ItemList            `protobuf:"bytes,5,rep,name=drop_list,json=dropList,proto3" json:"drop_list,omitempty"`
	BattleWaveId  uint32                 `protobuf:"varint,9,opt,name=battle_wave_id,json=battleWaveId,proto3" json:"battle_wave_id,omitempty"`    // 8
	BattleStageId uint32                 `protobuf:"varint,8,opt,name=battle_stage_id,json=battleStageId,proto3" json:"battle_stage_id,omitempty"` // 9
	MonsterList   []*SceneMonster        `protobuf:"bytes,12,rep,name=monster_list,json=monsterList,proto3" json:"monster_list,omitempty"`
	MonsterParam  *SceneMonsterWaveParam `protobuf:"bytes,6,opt,name=monster_param,json=monsterParam,proto3" json:"monster_param,omitempty"`
}

func (x *SceneMonsterWave) Reset() {
	*x = SceneMonsterWave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneMonsterWave_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneMonsterWave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneMonsterWave) ProtoMessage() {}

func (x *SceneMonsterWave) ProtoReflect() protoreflect.Message {
	mi := &file_SceneMonsterWave_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneMonsterWave.ProtoReflect.Descriptor instead.
func (*SceneMonsterWave) Descriptor() ([]byte, []int) {
	return file_SceneMonsterWave_proto_rawDescGZIP(), []int{0}
}

func (x *SceneMonsterWave) GetDropList() []*ItemList {
	if x != nil {
		return x.DropList
	}
	return nil
}

func (x *SceneMonsterWave) GetBattleWaveId() uint32 {
	if x != nil {
		return x.BattleWaveId
	}
	return 0
}

func (x *SceneMonsterWave) GetBattleStageId() uint32 {
	if x != nil {
		return x.BattleStageId
	}
	return 0
}

func (x *SceneMonsterWave) GetMonsterList() []*SceneMonster {
	if x != nil {
		return x.MonsterList
	}
	return nil
}

func (x *SceneMonsterWave) GetMonsterParam() *SceneMonsterWaveParam {
	if x != nil {
		return x.MonsterParam
	}
	return nil
}

var File_SceneMonsterWave_proto protoreflect.FileDescriptor

var file_SceneMonsterWave_proto_rawDesc = []byte{
	0x0a, 0x16, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61,
	0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x10, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61, 0x76, 0x65, 0x12, 0x26,
	0x0a, 0x09, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x64, 0x72,
	0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x5f, 0x77, 0x61, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x57, 0x61, 0x76, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61, 0x76, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneMonsterWave_proto_rawDescOnce sync.Once
	file_SceneMonsterWave_proto_rawDescData = file_SceneMonsterWave_proto_rawDesc
)

func file_SceneMonsterWave_proto_rawDescGZIP() []byte {
	file_SceneMonsterWave_proto_rawDescOnce.Do(func() {
		file_SceneMonsterWave_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneMonsterWave_proto_rawDescData)
	})
	return file_SceneMonsterWave_proto_rawDescData
}

var file_SceneMonsterWave_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneMonsterWave_proto_goTypes = []interface{}{
	(*SceneMonsterWave)(nil),      // 0: SceneMonsterWave
	(*ItemList)(nil),              // 1: ItemList
	(*SceneMonster)(nil),          // 2: SceneMonster
	(*SceneMonsterWaveParam)(nil), // 3: SceneMonsterWaveParam
}
var file_SceneMonsterWave_proto_depIdxs = []int32{
	1, // 0: SceneMonsterWave.drop_list:type_name -> ItemList
	2, // 1: SceneMonsterWave.monster_list:type_name -> SceneMonster
	3, // 2: SceneMonsterWave.monster_param:type_name -> SceneMonsterWaveParam
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SceneMonsterWave_proto_init() }
func file_SceneMonsterWave_proto_init() {
	if File_SceneMonsterWave_proto != nil {
		return
	}
	file_ItemList_proto_init()
	file_SceneMonsterWaveParam_proto_init()
	file_SceneMonster_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneMonsterWave_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneMonsterWave); i {
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
			RawDescriptor: file_SceneMonsterWave_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneMonsterWave_proto_goTypes,
		DependencyIndexes: file_SceneMonsterWave_proto_depIdxs,
		MessageInfos:      file_SceneMonsterWave_proto_msgTypes,
	}.Build()
	File_SceneMonsterWave_proto = out.File
	file_SceneMonsterWave_proto_rawDesc = nil
	file_SceneMonsterWave_proto_goTypes = nil
	file_SceneMonsterWave_proto_depIdxs = nil
}
