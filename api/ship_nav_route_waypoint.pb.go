// SpaceTraders v2.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: ship_nav_route_waypoint.proto

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

// The destination or departure of a ships nav route.
type ShipNavRouteWaypoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol       string       `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Type         WaypointType `protobuf:"varint,2,opt,name=type,proto3,enum=macgenius.spacetraders.api.wp.WaypointType" json:"type,omitempty"` // The type of waypoint.
	SystemSymbol string       `protobuf:"bytes,3,opt,name=system_symbol,json=systemSymbol,proto3" json:"system_symbol,omitempty"`
	X            int32        `protobuf:"varint,4,opt,name=x,proto3" json:"x,omitempty"`
	Y            int32        `protobuf:"varint,5,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *ShipNavRouteWaypoint) Reset() {
	*x = ShipNavRouteWaypoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ship_nav_route_waypoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipNavRouteWaypoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipNavRouteWaypoint) ProtoMessage() {}

func (x *ShipNavRouteWaypoint) ProtoReflect() protoreflect.Message {
	mi := &file_ship_nav_route_waypoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipNavRouteWaypoint.ProtoReflect.Descriptor instead.
func (*ShipNavRouteWaypoint) Descriptor() ([]byte, []int) {
	return file_ship_nav_route_waypoint_proto_rawDescGZIP(), []int{0}
}

func (x *ShipNavRouteWaypoint) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ShipNavRouteWaypoint) GetType() WaypointType {
	if x != nil {
		return x.Type
	}
	return WaypointType_PLANET
}

func (x *ShipNavRouteWaypoint) GetSystemSymbol() string {
	if x != nil {
		return x.SystemSymbol
	}
	return ""
}

func (x *ShipNavRouteWaypoint) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ShipNavRouteWaypoint) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_ship_nav_route_waypoint_proto protoreflect.FileDescriptor

var file_ship_nav_route_waypoint_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6e, 0x61, 0x76, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x13, 0x77, 0x61, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb0, 0x01, 0x0a, 0x14, 0x53, 0x68, 0x69, 0x70, 0x4e, 0x61, 0x76, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x3f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2b, 0x2e, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x70, 0x2e,
	0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x79, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x61, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ship_nav_route_waypoint_proto_rawDescOnce sync.Once
	file_ship_nav_route_waypoint_proto_rawDescData = file_ship_nav_route_waypoint_proto_rawDesc
)

func file_ship_nav_route_waypoint_proto_rawDescGZIP() []byte {
	file_ship_nav_route_waypoint_proto_rawDescOnce.Do(func() {
		file_ship_nav_route_waypoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_ship_nav_route_waypoint_proto_rawDescData)
	})
	return file_ship_nav_route_waypoint_proto_rawDescData
}

var file_ship_nav_route_waypoint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ship_nav_route_waypoint_proto_goTypes = []interface{}{
	(*ShipNavRouteWaypoint)(nil), // 0: macgenius.spacetraders.api.ShipNavRouteWaypoint
	(WaypointType)(0),            // 1: macgenius.spacetraders.api.wp.WaypointType
}
var file_ship_nav_route_waypoint_proto_depIdxs = []int32{
	1, // 0: macgenius.spacetraders.api.ShipNavRouteWaypoint.type:type_name -> macgenius.spacetraders.api.wp.WaypointType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ship_nav_route_waypoint_proto_init() }
func file_ship_nav_route_waypoint_proto_init() {
	if File_ship_nav_route_waypoint_proto != nil {
		return
	}
	file_waypoint_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ship_nav_route_waypoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipNavRouteWaypoint); i {
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
			RawDescriptor: file_ship_nav_route_waypoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ship_nav_route_waypoint_proto_goTypes,
		DependencyIndexes: file_ship_nav_route_waypoint_proto_depIdxs,
		MessageInfos:      file_ship_nav_route_waypoint_proto_msgTypes,
	}.Build()
	File_ship_nav_route_waypoint_proto = out.File
	file_ship_nav_route_waypoint_proto_rawDesc = nil
	file_ship_nav_route_waypoint_proto_goTypes = nil
	file_ship_nav_route_waypoint_proto_depIdxs = nil
}
