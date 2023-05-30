// SpaceTraders v2.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: register_faction.proto

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

type RegisterFaction int32

const (
	RegisterFaction_COSMIC   RegisterFaction = 0
	RegisterFaction_VOID     RegisterFaction = 1
	RegisterFaction_GALACTIC RegisterFaction = 2
	RegisterFaction_QUANTUM  RegisterFaction = 3
	RegisterFaction_DOMINION RegisterFaction = 4
)

// Enum value maps for RegisterFaction.
var (
	RegisterFaction_name = map[int32]string{
		0: "COSMIC",
		1: "VOID",
		2: "GALACTIC",
		3: "QUANTUM",
		4: "DOMINION",
	}
	RegisterFaction_value = map[string]int32{
		"COSMIC":   0,
		"VOID":     1,
		"GALACTIC": 2,
		"QUANTUM":  3,
		"DOMINION": 4,
	}
)

func (x RegisterFaction) Enum() *RegisterFaction {
	p := new(RegisterFaction)
	*p = x
	return p
}

func (x RegisterFaction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterFaction) Descriptor() protoreflect.EnumDescriptor {
	return file_register_faction_proto_enumTypes[0].Descriptor()
}

func (RegisterFaction) Type() protoreflect.EnumType {
	return &file_register_faction_proto_enumTypes[0]
}

func (x RegisterFaction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterFaction.Descriptor instead.
func (RegisterFaction) EnumDescriptor() ([]byte, []int) {
	return file_register_faction_proto_rawDescGZIP(), []int{0}
}

var File_register_faction_proto protoreflect.FileDescriptor

var file_register_faction_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e,
	0x69, 0x75, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2a, 0x50, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x53, 0x4d, 0x49, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x56, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x41, 0x4c, 0x41, 0x43, 0x54,
	0x49, 0x43, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x55, 0x4d, 0x10,
	0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x4f, 0x4d, 0x49, 0x4e, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_register_faction_proto_rawDescOnce sync.Once
	file_register_faction_proto_rawDescData = file_register_faction_proto_rawDesc
)

func file_register_faction_proto_rawDescGZIP() []byte {
	file_register_faction_proto_rawDescOnce.Do(func() {
		file_register_faction_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_faction_proto_rawDescData)
	})
	return file_register_faction_proto_rawDescData
}

var file_register_faction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_register_faction_proto_goTypes = []interface{}{
	(RegisterFaction)(0), // 0: macgenius.spacetraders.api.register.RegisterFaction
}
var file_register_faction_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_register_faction_proto_init() }
func file_register_faction_proto_init() {
	if File_register_faction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_register_faction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_register_faction_proto_goTypes,
		DependencyIndexes: file_register_faction_proto_depIdxs,
		EnumInfos:         file_register_faction_proto_enumTypes,
	}.Build()
	File_register_faction_proto = out.File
	file_register_faction_proto_rawDesc = nil
	file_register_faction_proto_goTypes = nil
	file_register_faction_proto_depIdxs = nil
}