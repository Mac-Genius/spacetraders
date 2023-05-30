// SpaceTraders v2.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: shipyard_ship.proto

package api

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

type ShipyardShip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          ShipType      `protobuf:"varint,1,opt,name=type,proto3,enum=macgenius.spacetraders.api.ShipType" json:"type,omitempty"`
	Name          string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PurchasePrice int32         `protobuf:"varint,4,opt,name=purchase_price,json=purchasePrice,proto3" json:"purchase_price,omitempty"`
	Frame         *ShipFrame    `protobuf:"bytes,5,opt,name=frame,proto3" json:"frame,omitempty"`
	Reactor       *ShipReactor  `protobuf:"bytes,6,opt,name=reactor,proto3" json:"reactor,omitempty"`
	Engine        *ShipEngine   `protobuf:"bytes,7,opt,name=engine,proto3" json:"engine,omitempty"`
	Modules       []*ShipModule `protobuf:"bytes,8,rep,name=modules,proto3" json:"modules,omitempty"`
	Mounts        []*ShipMount  `protobuf:"bytes,9,rep,name=mounts,proto3" json:"mounts,omitempty"`
}

func (x *ShipyardShip) Reset() {
	*x = ShipyardShip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipyard_ship_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipyardShip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipyardShip) ProtoMessage() {}

func (x *ShipyardShip) ProtoReflect() protoreflect.Message {
	mi := &file_shipyard_ship_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipyardShip.ProtoReflect.Descriptor instead.
func (*ShipyardShip) Descriptor() ([]byte, []int) {
	return file_shipyard_ship_proto_rawDescGZIP(), []int{0}
}

func (x *ShipyardShip) GetType() ShipType {
	if x != nil {
		return x.Type
	}
	return ShipType_SHIP_PROBE
}

func (x *ShipyardShip) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShipyardShip) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ShipyardShip) GetPurchasePrice() int32 {
	if x != nil {
		return x.PurchasePrice
	}
	return 0
}

func (x *ShipyardShip) GetFrame() *ShipFrame {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *ShipyardShip) GetReactor() *ShipReactor {
	if x != nil {
		return x.Reactor
	}
	return nil
}

func (x *ShipyardShip) GetEngine() *ShipEngine {
	if x != nil {
		return x.Engine
	}
	return nil
}

func (x *ShipyardShip) GetModules() []*ShipModule {
	if x != nil {
		return x.Modules
	}
	return nil
}

func (x *ShipyardShip) GetMounts() []*ShipMount {
	if x != nil {
		return x.Mounts
	}
	return nil
}

var File_shipyard_ship_proto protoreflect.FileDescriptor

var file_shipyard_ship_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x68, 0x69, 0x70, 0x79, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x1a, 0x11, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x68, 0x69, 0x70, 0x5f,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe6, 0x03, 0x0a, 0x0c, 0x53, 0x68, 0x69, 0x70, 0x79, 0x61, 0x72, 0x64, 0x53, 0x68, 0x69,
	0x70, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e,
	0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x05,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e, 0x69,
	0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x07, 0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3e, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x61, 0x63, 0x67, 0x65,
	0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x61, 0x63, 0x67,
	0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x61, 0x63,
	0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x06, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x69,
	0x75, 0x73, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipyard_ship_proto_rawDescOnce sync.Once
	file_shipyard_ship_proto_rawDescData = file_shipyard_ship_proto_rawDesc
)

func file_shipyard_ship_proto_rawDescGZIP() []byte {
	file_shipyard_ship_proto_rawDescOnce.Do(func() {
		file_shipyard_ship_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipyard_ship_proto_rawDescData)
	})
	return file_shipyard_ship_proto_rawDescData
}

var file_shipyard_ship_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shipyard_ship_proto_goTypes = []interface{}{
	(*ShipyardShip)(nil), // 0: macgenius.spacetraders.api.ShipyardShip
	(ShipType)(0),        // 1: macgenius.spacetraders.api.ShipType
	(*ShipFrame)(nil),    // 2: macgenius.spacetraders.api.ShipFrame
	(*ShipReactor)(nil),  // 3: macgenius.spacetraders.api.ShipReactor
	(*ShipEngine)(nil),   // 4: macgenius.spacetraders.api.ShipEngine
	(*ShipModule)(nil),   // 5: macgenius.spacetraders.api.ShipModule
	(*ShipMount)(nil),    // 6: macgenius.spacetraders.api.ShipMount
}
var file_shipyard_ship_proto_depIdxs = []int32{
	1, // 0: macgenius.spacetraders.api.ShipyardShip.type:type_name -> macgenius.spacetraders.api.ShipType
	2, // 1: macgenius.spacetraders.api.ShipyardShip.frame:type_name -> macgenius.spacetraders.api.ShipFrame
	3, // 2: macgenius.spacetraders.api.ShipyardShip.reactor:type_name -> macgenius.spacetraders.api.ShipReactor
	4, // 3: macgenius.spacetraders.api.ShipyardShip.engine:type_name -> macgenius.spacetraders.api.ShipEngine
	5, // 4: macgenius.spacetraders.api.ShipyardShip.modules:type_name -> macgenius.spacetraders.api.ShipModule
	6, // 5: macgenius.spacetraders.api.ShipyardShip.mounts:type_name -> macgenius.spacetraders.api.ShipMount
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_shipyard_ship_proto_init() }
func file_shipyard_ship_proto_init() {
	if File_shipyard_ship_proto != nil {
		return
	}
	file_ship_engine_proto_init()
	file_ship_frame_proto_init()
	file_ship_module_proto_init()
	file_ship_mount_proto_init()
	file_ship_reactor_proto_init()
	file_ship_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shipyard_ship_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipyardShip); i {
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
			RawDescriptor: file_shipyard_ship_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shipyard_ship_proto_goTypes,
		DependencyIndexes: file_shipyard_ship_proto_depIdxs,
		MessageInfos:      file_shipyard_ship_proto_msgTypes,
	}.Build()
	File_shipyard_ship_proto = out.File
	file_shipyard_ship_proto_rawDesc = nil
	file_shipyard_ship_proto_goTypes = nil
	file_shipyard_ship_proto_depIdxs = nil
}