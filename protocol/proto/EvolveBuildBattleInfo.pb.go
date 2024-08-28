// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EvolveBuildBattleInfo.proto

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

type EvolveBuildBattleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurLevelId            uint32                       `protobuf:"varint,1,opt,name=cur_level_id,json=curLevelId,proto3" json:"cur_level_id,omitempty"`
	CurPeriod             uint32                       `protobuf:"varint,2,opt,name=cur_period,json=curPeriod,proto3" json:"cur_period,omitempty"`
	CurCoin               uint32                       `protobuf:"varint,3,opt,name=cur_coin,json=curCoin,proto3" json:"cur_coin,omitempty"`
	WeaponSlotList        []*MOGAOBCPBHK               `protobuf:"bytes,4,rep,name=weapon_slot_list,json=weaponSlotList,proto3" json:"weapon_slot_list,omitempty"`
	AccessorySlotList     []*MOGAOBCPBHK               `protobuf:"bytes,5,rep,name=accessory_slot_list,json=accessorySlotList,proto3" json:"accessory_slot_list,omitempty"`
	BanGearList           []uint32                     `protobuf:"varint,6,rep,packed,name=ban_gear_list,json=banGearList,proto3" json:"ban_gear_list,omitempty"`
	Collection            *LGBDMBJDNLK                 `protobuf:"bytes,7,opt,name=collection,proto3" json:"collection,omitempty"`
	AllowedGearList       []uint32                     `protobuf:"varint,8,rep,packed,name=allowed_gear_list,json=allowedGearList,proto3" json:"allowed_gear_list,omitempty"`
	CurExp                uint32                       `protobuf:"varint,9,opt,name=cur_exp,json=curExp,proto3" json:"cur_exp,omitempty"`
	CurReroll             uint32                       `protobuf:"varint,10,opt,name=cur_reroll,json=curReroll,proto3" json:"cur_reroll,omitempty"`
	CurTreasureMissCnt    uint32                       `protobuf:"varint,11,opt,name=cur_treasure_miss_cnt,json=curTreasureMissCnt,proto3" json:"cur_treasure_miss_cnt,omitempty"`
	PeriodIdList          []uint32                     `protobuf:"varint,12,rep,packed,name=period_id_list,json=periodIdList,proto3" json:"period_id_list,omitempty"`
	CurGearLostCnt        uint32                       `protobuf:"varint,13,opt,name=cur_gear_lost_cnt,json=curGearLostCnt,proto3" json:"cur_gear_lost_cnt,omitempty"`
	CurWave               uint32                       `protobuf:"varint,14,opt,name=cur_wave,json=curWave,proto3" json:"cur_wave,omitempty"`
	IsUnlockGearReroll    bool                         `protobuf:"varint,15,opt,name=is_unlock_gear_reroll,json=isUnlockGearReroll,proto3" json:"is_unlock_gear_reroll,omitempty"`
	IsUnlockGearBan       bool                         `protobuf:"varint,16,opt,name=is_unlock_gear_ban,json=isUnlockGearBan,proto3" json:"is_unlock_gear_ban,omitempty"`
	CardList              []*BHJEIEBCMOL               `protobuf:"bytes,17,rep,name=card_list,json=cardList,proto3" json:"card_list,omitempty"`
	GearDamageList        []*EvolveBuildGearDamageInfo `protobuf:"bytes,18,rep,name=gear_damage_list,json=gearDamageList,proto3" json:"gear_damage_list,omitempty"`
	StatParams            []uint32                     `protobuf:"varint,19,rep,packed,name=stat_params,json=statParams,proto3" json:"stat_params,omitempty"`
	IsGiveup              bool                         `protobuf:"varint,20,opt,name=is_giveup,json=isGiveup,proto3" json:"is_giveup,omitempty"`
	CurUnusedRoundCnt     uint32                       `protobuf:"varint,21,opt,name=cur_unused_round_cnt,json=curUnusedRoundCnt,proto3" json:"cur_unused_round_cnt,omitempty"`
	StatLogInfo           *HEDDDDANPKB                 `protobuf:"bytes,22,opt,name=stat_log_info,json=statLogInfo,proto3" json:"stat_log_info,omitempty"`
	PeriodFirstRandomSeed uint32                       `protobuf:"varint,23,opt,name=period_first_random_seed,json=periodFirstRandomSeed,proto3" json:"period_first_random_seed,omitempty"`
}

func (x *EvolveBuildBattleInfo) Reset() {
	*x = EvolveBuildBattleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvolveBuildBattleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvolveBuildBattleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvolveBuildBattleInfo) ProtoMessage() {}

func (x *EvolveBuildBattleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_EvolveBuildBattleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvolveBuildBattleInfo.ProtoReflect.Descriptor instead.
func (*EvolveBuildBattleInfo) Descriptor() ([]byte, []int) {
	return file_EvolveBuildBattleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EvolveBuildBattleInfo) GetCurLevelId() uint32 {
	if x != nil {
		return x.CurLevelId
	}
	return 0
}

func (x *EvolveBuildBattleInfo) GetCurPeriod() uint32 {
	if x != nil {
		return x.CurPeriod
	}
	return 0
}

func (x *EvolveBuildBattleInfo) GetCurCoin() uint32 {
	if x != nil {
		return x.CurCoin
	}
	return 0
}

func (x *EvolveBuildBattleInfo) GetWeaponSlotList() []*MOGAOBCPBHK {
	if x != nil {
		return x.WeaponSlotList
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetAccessorySlotList() []*MOGAOBCPBHK {
	if x != nil {
		return x.AccessorySlotList
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetBanGearList() []uint32 {
	if x != nil {
		return x.BanGearList
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetCollection() *LGBDMBJDNLK {
	if x != nil {
		return x.Collection
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetAllowedGearList() []uint32 {
	if x != nil {
		return x.AllowedGearList
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetCurExp() uint32 {
	if x != nil {
		return x.CurExp
	}
	return 0
}

func (x *EvolveBuildBattleInfo) GetCurReroll() uint32 {
	if x != nil {
		return x.CurReroll
	}
	return 0
}

func (x *EvolveBuildBattleInfo) GetCurTreasureMissCnt() uint32 {
	if x != nil {
		return x.CurTreasureMissCnt
	}
	return 0
}

func (x *EvolveBuildBattleInfo) GetPeriodIdList() []uint32 {
	if x != nil {
		return x.PeriodIdList
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetCurGearLostCnt() uint32 {
	if x != nil {
		return x.CurGearLostCnt
	}
	return 0
}

func (x *EvolveBuildBattleInfo) GetCurWave() uint32 {
	if x != nil {
		return x.CurWave
	}
	return 0
}

func (x *EvolveBuildBattleInfo) GetIsUnlockGearReroll() bool {
	if x != nil {
		return x.IsUnlockGearReroll
	}
	return false
}

func (x *EvolveBuildBattleInfo) GetIsUnlockGearBan() bool {
	if x != nil {
		return x.IsUnlockGearBan
	}
	return false
}

func (x *EvolveBuildBattleInfo) GetCardList() []*BHJEIEBCMOL {
	if x != nil {
		return x.CardList
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetGearDamageList() []*EvolveBuildGearDamageInfo {
	if x != nil {
		return x.GearDamageList
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetStatParams() []uint32 {
	if x != nil {
		return x.StatParams
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetIsGiveup() bool {
	if x != nil {
		return x.IsGiveup
	}
	return false
}

func (x *EvolveBuildBattleInfo) GetCurUnusedRoundCnt() uint32 {
	if x != nil {
		return x.CurUnusedRoundCnt
	}
	return 0
}

func (x *EvolveBuildBattleInfo) GetStatLogInfo() *HEDDDDANPKB {
	if x != nil {
		return x.StatLogInfo
	}
	return nil
}

func (x *EvolveBuildBattleInfo) GetPeriodFirstRandomSeed() uint32 {
	if x != nil {
		return x.PeriodFirstRandomSeed
	}
	return 0
}

var File_EvolveBuildBattleInfo_proto protoreflect.FileDescriptor

var file_EvolveBuildBattleInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c,
	0x47, 0x42, 0x44, 0x4d, 0x42, 0x4a, 0x44, 0x4e, 0x4c, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x4d, 0x4f, 0x47, 0x41, 0x4f, 0x42, 0x43, 0x50, 0x42, 0x48, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x47, 0x65, 0x61, 0x72, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x48, 0x4a, 0x45, 0x49, 0x45, 0x42, 0x43, 0x4d, 0x4f,
	0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x45, 0x44, 0x44, 0x44, 0x44, 0x41,
	0x4e, 0x50, 0x4b, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x07, 0x0a, 0x15, 0x45,
	0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x5f, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x75, 0x72, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x6f, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x75, 0x72, 0x43, 0x6f, 0x69, 0x6e,
	0x12, 0x36, 0x0a, 0x10, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x4f, 0x47,
	0x41, 0x4f, 0x42, 0x43, 0x50, 0x42, 0x48, 0x4b, 0x52, 0x0e, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x53, 0x6c, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x13, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x4f, 0x47, 0x41, 0x4f, 0x42, 0x43, 0x50,
	0x42, 0x48, 0x4b, 0x52, 0x11, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x79, 0x53, 0x6c,
	0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x61, 0x6e, 0x5f, 0x67, 0x65,
	0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x62,
	0x61, 0x6e, 0x47, 0x65, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0a, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4c, 0x47, 0x42, 0x44, 0x4d, 0x42, 0x4a, 0x44, 0x4e, 0x4c, 0x4b, 0x52, 0x0a, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x47, 0x65, 0x61, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x75, 0x72, 0x45, 0x78, 0x70, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x75, 0x72, 0x5f, 0x72, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x63, 0x75, 0x72, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x31, 0x0a, 0x15,
	0x63, 0x75, 0x72, 0x5f, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x69, 0x73,
	0x73, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x63, 0x75, 0x72,
	0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x43, 0x6e, 0x74, 0x12,
	0x24, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x75, 0x72, 0x5f, 0x67, 0x65, 0x61,
	0x72, 0x5f, 0x6c, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x63, 0x75, 0x72, 0x47, 0x65, 0x61, 0x72, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x6e, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x5f, 0x77, 0x61, 0x76, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x63, 0x75, 0x72, 0x57, 0x61, 0x76, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x69,
	0x73, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x72, 0x65,
	0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x47, 0x65, 0x61, 0x72, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x2b,
	0x0a, 0x12, 0x69, 0x73, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x67, 0x65, 0x61, 0x72,
	0x5f, 0x62, 0x61, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x47, 0x65, 0x61, 0x72, 0x42, 0x61, 0x6e, 0x12, 0x29, 0x0a, 0x09, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x42, 0x48, 0x4a, 0x45, 0x49, 0x45, 0x42, 0x43, 0x4d, 0x4f, 0x4c, 0x52, 0x08, 0x63, 0x61,
	0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x64,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x47, 0x65,
	0x61, 0x72, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x67, 0x65,
	0x61, 0x72, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x75, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x47, 0x69, 0x76, 0x65, 0x75, 0x70, 0x12, 0x2f, 0x0a, 0x14, 0x63, 0x75,
	0x72, 0x5f, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63,
	0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x75, 0x72, 0x55, 0x6e, 0x75,
	0x73, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x45, 0x44, 0x44, 0x44, 0x44, 0x41, 0x4e, 0x50, 0x4b, 0x42,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a,
	0x18, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x72, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x15, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x46, 0x69, 0x72, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x53, 0x65, 0x65, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvolveBuildBattleInfo_proto_rawDescOnce sync.Once
	file_EvolveBuildBattleInfo_proto_rawDescData = file_EvolveBuildBattleInfo_proto_rawDesc
)

func file_EvolveBuildBattleInfo_proto_rawDescGZIP() []byte {
	file_EvolveBuildBattleInfo_proto_rawDescOnce.Do(func() {
		file_EvolveBuildBattleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvolveBuildBattleInfo_proto_rawDescData)
	})
	return file_EvolveBuildBattleInfo_proto_rawDescData
}

var file_EvolveBuildBattleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvolveBuildBattleInfo_proto_goTypes = []interface{}{
	(*EvolveBuildBattleInfo)(nil),     // 0: EvolveBuildBattleInfo
	(*MOGAOBCPBHK)(nil),               // 1: MOGAOBCPBHK
	(*LGBDMBJDNLK)(nil),               // 2: LGBDMBJDNLK
	(*BHJEIEBCMOL)(nil),               // 3: BHJEIEBCMOL
	(*EvolveBuildGearDamageInfo)(nil), // 4: EvolveBuildGearDamageInfo
	(*HEDDDDANPKB)(nil),               // 5: HEDDDDANPKB
}
var file_EvolveBuildBattleInfo_proto_depIdxs = []int32{
	1, // 0: EvolveBuildBattleInfo.weapon_slot_list:type_name -> MOGAOBCPBHK
	1, // 1: EvolveBuildBattleInfo.accessory_slot_list:type_name -> MOGAOBCPBHK
	2, // 2: EvolveBuildBattleInfo.collection:type_name -> LGBDMBJDNLK
	3, // 3: EvolveBuildBattleInfo.card_list:type_name -> BHJEIEBCMOL
	4, // 4: EvolveBuildBattleInfo.gear_damage_list:type_name -> EvolveBuildGearDamageInfo
	5, // 5: EvolveBuildBattleInfo.stat_log_info:type_name -> HEDDDDANPKB
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_EvolveBuildBattleInfo_proto_init() }
func file_EvolveBuildBattleInfo_proto_init() {
	if File_EvolveBuildBattleInfo_proto != nil {
		return
	}
	file_LGBDMBJDNLK_proto_init()
	file_MOGAOBCPBHK_proto_init()
	file_EvolveBuildGearDamageInfo_proto_init()
	file_BHJEIEBCMOL_proto_init()
	file_HEDDDDANPKB_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvolveBuildBattleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvolveBuildBattleInfo); i {
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
			RawDescriptor: file_EvolveBuildBattleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvolveBuildBattleInfo_proto_goTypes,
		DependencyIndexes: file_EvolveBuildBattleInfo_proto_depIdxs,
		MessageInfos:      file_EvolveBuildBattleInfo_proto_msgTypes,
	}.Build()
	File_EvolveBuildBattleInfo_proto = out.File
	file_EvolveBuildBattleInfo_proto_rawDesc = nil
	file_EvolveBuildBattleInfo_proto_goTypes = nil
	file_EvolveBuildBattleInfo_proto_depIdxs = nil
}
