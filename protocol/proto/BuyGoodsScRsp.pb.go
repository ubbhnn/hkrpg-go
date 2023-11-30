// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BuyGoodsScRsp.proto

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

type BuyGoodsScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnItemList *ItemList `protobuf:"bytes,5,opt,name=return_item_list,json=returnItemList,proto3" json:"return_item_list,omitempty"`
	Retcode        uint32    `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ShopId         uint32    `protobuf:"varint,15,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	GoodsId        uint32    `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	GoodsBuyTimes  uint32    `protobuf:"varint,11,opt,name=goods_buy_times,json=goodsBuyTimes,proto3" json:"goods_buy_times,omitempty"`
}

func (x *BuyGoodsScRsp) Reset() {
	*x = BuyGoodsScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BuyGoodsScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyGoodsScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyGoodsScRsp) ProtoMessage() {}

func (x *BuyGoodsScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_BuyGoodsScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyGoodsScRsp.ProtoReflect.Descriptor instead.
func (*BuyGoodsScRsp) Descriptor() ([]byte, []int) {
	return file_BuyGoodsScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *BuyGoodsScRsp) GetReturnItemList() *ItemList {
	if x != nil {
		return x.ReturnItemList
	}
	return nil
}

func (x *BuyGoodsScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *BuyGoodsScRsp) GetShopId() uint32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *BuyGoodsScRsp) GetGoodsId() uint32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *BuyGoodsScRsp) GetGoodsBuyTimes() uint32 {
	if x != nil {
		return x.GoodsBuyTimes
	}
	return 0
}

var File_BuyGoodsScRsp_proto protoreflect.FileDescriptor

var file_BuyGoodsScRsp_proto_rawDesc = []byte{
	0x0a, 0x13, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x0d, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x10, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0e, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x75, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BuyGoodsScRsp_proto_rawDescOnce sync.Once
	file_BuyGoodsScRsp_proto_rawDescData = file_BuyGoodsScRsp_proto_rawDesc
)

func file_BuyGoodsScRsp_proto_rawDescGZIP() []byte {
	file_BuyGoodsScRsp_proto_rawDescOnce.Do(func() {
		file_BuyGoodsScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_BuyGoodsScRsp_proto_rawDescData)
	})
	return file_BuyGoodsScRsp_proto_rawDescData
}

var file_BuyGoodsScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BuyGoodsScRsp_proto_goTypes = []interface{}{
	(*BuyGoodsScRsp)(nil), // 0: BuyGoodsScRsp
	(*ItemList)(nil),      // 1: ItemList
}
var file_BuyGoodsScRsp_proto_depIdxs = []int32{
	1, // 0: BuyGoodsScRsp.return_item_list:type_name -> ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BuyGoodsScRsp_proto_init() }
func file_BuyGoodsScRsp_proto_init() {
	if File_BuyGoodsScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BuyGoodsScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyGoodsScRsp); i {
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
			RawDescriptor: file_BuyGoodsScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BuyGoodsScRsp_proto_goTypes,
		DependencyIndexes: file_BuyGoodsScRsp_proto_depIdxs,
		MessageInfos:      file_BuyGoodsScRsp_proto_msgTypes,
	}.Build()
	File_BuyGoodsScRsp_proto = out.File
	file_BuyGoodsScRsp_proto_rawDesc = nil
	file_BuyGoodsScRsp_proto_goTypes = nil
	file_BuyGoodsScRsp_proto_depIdxs = nil
}
