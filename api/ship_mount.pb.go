// SpaceTraders v2.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: ship_mount.proto

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

type ShipMountSymbol int32

const (
	ShipMountSymbol_MOUNT_GAS_SIPHON_I       ShipMountSymbol = 0
	ShipMountSymbol_MOUNT_GAS_SIPHON_II      ShipMountSymbol = 1
	ShipMountSymbol_MOUNT_GAS_SIPHON_III     ShipMountSymbol = 2
	ShipMountSymbol_MOUNT_SURVEYOR_I         ShipMountSymbol = 3
	ShipMountSymbol_MOUNT_SURVEYOR_II        ShipMountSymbol = 4
	ShipMountSymbol_MOUNT_SURVEYOR_III       ShipMountSymbol = 5
	ShipMountSymbol_MOUNT_SENSOR_ARRAY_I     ShipMountSymbol = 6
	ShipMountSymbol_MOUNT_SENSOR_ARRAY_II    ShipMountSymbol = 7
	ShipMountSymbol_MOUNT_SENSOR_ARRAY_III   ShipMountSymbol = 8
	ShipMountSymbol_MOUNT_MINING_LASER_I     ShipMountSymbol = 9
	ShipMountSymbol_MOUNT_MINING_LASER_II    ShipMountSymbol = 10
	ShipMountSymbol_MOUNT_MINING_LASER_III   ShipMountSymbol = 11
	ShipMountSymbol_MOUNT_LASER_CANNON_I     ShipMountSymbol = 12
	ShipMountSymbol_MOUNT_MISSILE_LAUNCHER_I ShipMountSymbol = 13
	ShipMountSymbol_MOUNT_TURRET_I           ShipMountSymbol = 14
)

// Enum value maps for ShipMountSymbol.
var (
	ShipMountSymbol_name = map[int32]string{
		0:  "MOUNT_GAS_SIPHON_I",
		1:  "MOUNT_GAS_SIPHON_II",
		2:  "MOUNT_GAS_SIPHON_III",
		3:  "MOUNT_SURVEYOR_I",
		4:  "MOUNT_SURVEYOR_II",
		5:  "MOUNT_SURVEYOR_III",
		6:  "MOUNT_SENSOR_ARRAY_I",
		7:  "MOUNT_SENSOR_ARRAY_II",
		8:  "MOUNT_SENSOR_ARRAY_III",
		9:  "MOUNT_MINING_LASER_I",
		10: "MOUNT_MINING_LASER_II",
		11: "MOUNT_MINING_LASER_III",
		12: "MOUNT_LASER_CANNON_I",
		13: "MOUNT_MISSILE_LAUNCHER_I",
		14: "MOUNT_TURRET_I",
	}
	ShipMountSymbol_value = map[string]int32{
		"MOUNT_GAS_SIPHON_I":       0,
		"MOUNT_GAS_SIPHON_II":      1,
		"MOUNT_GAS_SIPHON_III":     2,
		"MOUNT_SURVEYOR_I":         3,
		"MOUNT_SURVEYOR_II":        4,
		"MOUNT_SURVEYOR_III":       5,
		"MOUNT_SENSOR_ARRAY_I":     6,
		"MOUNT_SENSOR_ARRAY_II":    7,
		"MOUNT_SENSOR_ARRAY_III":   8,
		"MOUNT_MINING_LASER_I":     9,
		"MOUNT_MINING_LASER_II":    10,
		"MOUNT_MINING_LASER_III":   11,
		"MOUNT_LASER_CANNON_I":     12,
		"MOUNT_MISSILE_LAUNCHER_I": 13,
		"MOUNT_TURRET_I":           14,
	}
)

func (x ShipMountSymbol) Enum() *ShipMountSymbol {
	p := new(ShipMountSymbol)
	*p = x
	return p
}

func (x ShipMountSymbol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShipMountSymbol) Descriptor() protoreflect.EnumDescriptor {
	return file_ship_mount_proto_enumTypes[0].Descriptor()
}

func (ShipMountSymbol) Type() protoreflect.EnumType {
	return &file_ship_mount_proto_enumTypes[0]
}

func (x ShipMountSymbol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShipMountSymbol.Descriptor instead.
func (ShipMountSymbol) EnumDescriptor() ([]byte, []int) {
	return file_ship_mount_proto_rawDescGZIP(), []int{0}
}

type ShipMountDeposit int32

const (
	ShipMountDeposit_QUARTZ_SAND      ShipMountDeposit = 0
	ShipMountDeposit_SILICON_CRYSTALS ShipMountDeposit = 1
	ShipMountDeposit_PRECIOUS_STONES  ShipMountDeposit = 2
	ShipMountDeposit_ICE_WATER        ShipMountDeposit = 3
	ShipMountDeposit_AMMONIA_ICE      ShipMountDeposit = 4
	ShipMountDeposit_IRON_ORE         ShipMountDeposit = 5
	ShipMountDeposit_COPPER_ORE       ShipMountDeposit = 6
	ShipMountDeposit_SILVER_ORE       ShipMountDeposit = 7
	ShipMountDeposit_ALUMINUM_ORE     ShipMountDeposit = 8
	ShipMountDeposit_GOLD_ORE         ShipMountDeposit = 9
	ShipMountDeposit_PLATINUM_ORE     ShipMountDeposit = 10
	ShipMountDeposit_DIAMOND          ShipMountDeposit = 11
	ShipMountDeposit_SURANITE_ORE     ShipMountDeposit = 12
	ShipMountDeposit_MERITIUM_ORE     ShipMountDeposit = 13
)

// Enum value maps for ShipMountDeposit.
var (
	ShipMountDeposit_name = map[int32]string{
		0:  "QUARTZ_SAND",
		1:  "SILICON_CRYSTALS",
		2:  "PRECIOUS_STONES",
		3:  "ICE_WATER",
		4:  "AMMONIA_ICE",
		5:  "IRON_ORE",
		6:  "COPPER_ORE",
		7:  "SILVER_ORE",
		8:  "ALUMINUM_ORE",
		9:  "GOLD_ORE",
		10: "PLATINUM_ORE",
		11: "DIAMOND",
		12: "SURANITE_ORE",
		13: "MERITIUM_ORE",
	}
	ShipMountDeposit_value = map[string]int32{
		"QUARTZ_SAND":      0,
		"SILICON_CRYSTALS": 1,
		"PRECIOUS_STONES":  2,
		"ICE_WATER":        3,
		"AMMONIA_ICE":      4,
		"IRON_ORE":         5,
		"COPPER_ORE":       6,
		"SILVER_ORE":       7,
		"ALUMINUM_ORE":     8,
		"GOLD_ORE":         9,
		"PLATINUM_ORE":     10,
		"DIAMOND":          11,
		"SURANITE_ORE":     12,
		"MERITIUM_ORE":     13,
	}
)

func (x ShipMountDeposit) Enum() *ShipMountDeposit {
	p := new(ShipMountDeposit)
	*p = x
	return p
}

func (x ShipMountDeposit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShipMountDeposit) Descriptor() protoreflect.EnumDescriptor {
	return file_ship_mount_proto_enumTypes[1].Descriptor()
}

func (ShipMountDeposit) Type() protoreflect.EnumType {
	return &file_ship_mount_proto_enumTypes[1]
}

func (x ShipMountDeposit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShipMountDeposit.Descriptor instead.
func (ShipMountDeposit) EnumDescriptor() ([]byte, []int) {
	return file_ship_mount_proto_rawDescGZIP(), []int{1}
}

// A mount is installed on the exterier of a ship.
type ShipMount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol       ShipMountSymbol    `protobuf:"varint,1,opt,name=symbol,proto3,enum=macgenius.spacetraders.api.ShipMountSymbol" json:"symbol,omitempty"`
	Name         string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string             `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Strength     int32              `protobuf:"varint,4,opt,name=strength,proto3" json:"strength,omitempty"`
	Deposits     []ShipMountDeposit `protobuf:"varint,5,rep,packed,name=deposits,proto3,enum=macgenius.spacetraders.api.ShipMountDeposit" json:"deposits,omitempty"`
	Requirements *ShipRequirements  `protobuf:"bytes,6,opt,name=requirements,proto3" json:"requirements,omitempty"`
}

func (x *ShipMount) Reset() {
	*x = ShipMount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ship_mount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipMount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipMount) ProtoMessage() {}

func (x *ShipMount) ProtoReflect() protoreflect.Message {
	mi := &file_ship_mount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipMount.ProtoReflect.Descriptor instead.
func (*ShipMount) Descriptor() ([]byte, []int) {
	return file_ship_mount_proto_rawDescGZIP(), []int{0}
}

func (x *ShipMount) GetSymbol() ShipMountSymbol {
	if x != nil {
		return x.Symbol
	}
	return ShipMountSymbol_MOUNT_GAS_SIPHON_I
}

func (x *ShipMount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShipMount) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ShipMount) GetStrength() int32 {
	if x != nil {
		return x.Strength
	}
	return 0
}

func (x *ShipMount) GetDeposits() []ShipMountDeposit {
	if x != nil {
		return x.Deposits
	}
	return nil
}

func (x *ShipMount) GetRequirements() *ShipRequirements {
	if x != nil {
		return x.Requirements
	}
	return nil
}

var File_ship_mount_proto protoreflect.FileDescriptor

var file_ship_mount_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x17,
	0x73, 0x68, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x09, 0x53, 0x68, 0x69, 0x70,
	0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e, 0x69, 0x75,
	0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x48, 0x0a, 0x08,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2c,
	0x2e, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70,
	0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x08, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x12, 0x50, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d,
	0x61, 0x63, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2a, 0x8f, 0x03, 0x0a, 0x0f, 0x53, 0x68, 0x69,
	0x70, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x12,
	0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x47, 0x41, 0x53, 0x5f, 0x53, 0x49, 0x50, 0x48, 0x4f, 0x4e,
	0x5f, 0x49, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x47, 0x41,
	0x53, 0x5f, 0x53, 0x49, 0x50, 0x48, 0x4f, 0x4e, 0x5f, 0x49, 0x49, 0x10, 0x01, 0x12, 0x18, 0x0a,
	0x14, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x47, 0x41, 0x53, 0x5f, 0x53, 0x49, 0x50, 0x48, 0x4f,
	0x4e, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x55, 0x4e, 0x54,
	0x5f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59, 0x4f, 0x52, 0x5f, 0x49, 0x10, 0x03, 0x12, 0x15, 0x0a,
	0x11, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59, 0x4f, 0x52, 0x5f,
	0x49, 0x49, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x55,
	0x52, 0x56, 0x45, 0x59, 0x4f, 0x52, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14,
	0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x41, 0x52, 0x52,
	0x41, 0x59, 0x5f, 0x49, 0x10, 0x06, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x5f, 0x49, 0x49, 0x10,
	0x07, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x4f,
	0x52, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x08, 0x12, 0x18, 0x0a,
	0x14, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x41,
	0x53, 0x45, 0x52, 0x5f, 0x49, 0x10, 0x09, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x55, 0x4e, 0x54,
	0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x49,
	0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x0b, 0x12, 0x18,
	0x0a, 0x14, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x41,
	0x4e, 0x4e, 0x4f, 0x4e, 0x5f, 0x49, 0x10, 0x0c, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4c, 0x45, 0x5f, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48,
	0x45, 0x52, 0x5f, 0x49, 0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x54, 0x55, 0x52, 0x52, 0x45, 0x54, 0x5f, 0x49, 0x10, 0x0e, 0x2a, 0xff, 0x01, 0x0a, 0x10, 0x53,
	0x68, 0x69, 0x70, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12,
	0x0f, 0x0a, 0x0b, 0x51, 0x55, 0x41, 0x52, 0x54, 0x5a, 0x5f, 0x53, 0x41, 0x4e, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x53, 0x49, 0x4c, 0x49, 0x43, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x59, 0x53,
	0x54, 0x41, 0x4c, 0x53, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x45, 0x43, 0x49, 0x4f,
	0x55, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x4e, 0x45, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x49,
	0x43, 0x45, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x4d,
	0x4d, 0x4f, 0x4e, 0x49, 0x41, 0x5f, 0x49, 0x43, 0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x49,
	0x52, 0x4f, 0x4e, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x50,
	0x50, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x49, 0x4c,
	0x56, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4c, 0x55,
	0x4d, 0x49, 0x4e, 0x55, 0x4d, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x47,
	0x4f, 0x4c, 0x44, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41,
	0x54, 0x49, 0x4e, 0x55, 0x4d, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x49, 0x41, 0x4d, 0x4f, 0x4e, 0x44, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x55, 0x52, 0x41,
	0x4e, 0x49, 0x54, 0x45, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45,
	0x52, 0x49, 0x54, 0x49, 0x55, 0x4d, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x0d, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ship_mount_proto_rawDescOnce sync.Once
	file_ship_mount_proto_rawDescData = file_ship_mount_proto_rawDesc
)

func file_ship_mount_proto_rawDescGZIP() []byte {
	file_ship_mount_proto_rawDescOnce.Do(func() {
		file_ship_mount_proto_rawDescData = protoimpl.X.CompressGZIP(file_ship_mount_proto_rawDescData)
	})
	return file_ship_mount_proto_rawDescData
}

var file_ship_mount_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ship_mount_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ship_mount_proto_goTypes = []interface{}{
	(ShipMountSymbol)(0),     // 0: macgenius.spacetraders.api.ShipMountSymbol
	(ShipMountDeposit)(0),    // 1: macgenius.spacetraders.api.ShipMountDeposit
	(*ShipMount)(nil),        // 2: macgenius.spacetraders.api.ShipMount
	(*ShipRequirements)(nil), // 3: macgenius.spacetraders.api.ShipRequirements
}
var file_ship_mount_proto_depIdxs = []int32{
	0, // 0: macgenius.spacetraders.api.ShipMount.symbol:type_name -> macgenius.spacetraders.api.ShipMountSymbol
	1, // 1: macgenius.spacetraders.api.ShipMount.deposits:type_name -> macgenius.spacetraders.api.ShipMountDeposit
	3, // 2: macgenius.spacetraders.api.ShipMount.requirements:type_name -> macgenius.spacetraders.api.ShipRequirements
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ship_mount_proto_init() }
func file_ship_mount_proto_init() {
	if File_ship_mount_proto != nil {
		return
	}
	file_ship_requirements_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ship_mount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipMount); i {
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
			RawDescriptor: file_ship_mount_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ship_mount_proto_goTypes,
		DependencyIndexes: file_ship_mount_proto_depIdxs,
		EnumInfos:         file_ship_mount_proto_enumTypes,
		MessageInfos:      file_ship_mount_proto_msgTypes,
	}.Build()
	File_ship_mount_proto = out.File
	file_ship_mount_proto_rawDesc = nil
	file_ship_mount_proto_goTypes = nil
	file_ship_mount_proto_depIdxs = nil
}
