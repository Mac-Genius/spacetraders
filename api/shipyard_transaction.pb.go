// SpaceTraders v2.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: shipyard_transaction.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShipyardTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WaypointSymbol string                 `protobuf:"bytes,1,opt,name=waypoint_symbol,json=waypointSymbol,proto3" json:"waypoint_symbol,omitempty"` // The symbol of the waypoint where the transaction took place.
	ShipSymbol     string                 `protobuf:"bytes,2,opt,name=ship_symbol,json=shipSymbol,proto3" json:"ship_symbol,omitempty"`             // The symbol of the ship that was purchased.
	Price          int32                  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`                                        // The price of the transaction.
	AgentSymbol    string                 `protobuf:"bytes,4,opt,name=agent_symbol,json=agentSymbol,proto3" json:"agent_symbol,omitempty"`          // The symbol of the agent that made the transaction.
	Timestamp      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                 // The timestamp of the transaction.
}

func (x *ShipyardTransaction) Reset() {
	*x = ShipyardTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipyard_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipyardTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipyardTransaction) ProtoMessage() {}

func (x *ShipyardTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_shipyard_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipyardTransaction.ProtoReflect.Descriptor instead.
func (*ShipyardTransaction) Descriptor() ([]byte, []int) {
	return file_shipyard_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *ShipyardTransaction) GetWaypointSymbol() string {
	if x != nil {
		return x.WaypointSymbol
	}
	return ""
}

func (x *ShipyardTransaction) GetShipSymbol() string {
	if x != nil {
		return x.ShipSymbol
	}
	return ""
}

func (x *ShipyardTransaction) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ShipyardTransaction) GetAgentSymbol() string {
	if x != nil {
		return x.AgentSymbol
	}
	return ""
}

func (x *ShipyardTransaction) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_shipyard_transaction_proto protoreflect.FileDescriptor

var file_shipyard_transaction_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x68, 0x69, 0x70, 0x79, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x61,
	0x63, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x13, 0x53, 0x68,
	0x69, 0x70, 0x79, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x61, 0x79, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68,
	0x69, 0x70, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x68, 0x69, 0x70, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x28,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipyard_transaction_proto_rawDescOnce sync.Once
	file_shipyard_transaction_proto_rawDescData = file_shipyard_transaction_proto_rawDesc
)

func file_shipyard_transaction_proto_rawDescGZIP() []byte {
	file_shipyard_transaction_proto_rawDescOnce.Do(func() {
		file_shipyard_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipyard_transaction_proto_rawDescData)
	})
	return file_shipyard_transaction_proto_rawDescData
}

var file_shipyard_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shipyard_transaction_proto_goTypes = []interface{}{
	(*ShipyardTransaction)(nil),   // 0: macgenius.spacetraders.api.ShipyardTransaction
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_shipyard_transaction_proto_depIdxs = []int32{
	1, // 0: macgenius.spacetraders.api.ShipyardTransaction.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shipyard_transaction_proto_init() }
func file_shipyard_transaction_proto_init() {
	if File_shipyard_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shipyard_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipyardTransaction); i {
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
			RawDescriptor: file_shipyard_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shipyard_transaction_proto_goTypes,
		DependencyIndexes: file_shipyard_transaction_proto_depIdxs,
		MessageInfos:      file_shipyard_transaction_proto_msgTypes,
	}.Build()
	File_shipyard_transaction_proto = out.File
	file_shipyard_transaction_proto_rawDesc = nil
	file_shipyard_transaction_proto_goTypes = nil
	file_shipyard_transaction_proto_depIdxs = nil
}