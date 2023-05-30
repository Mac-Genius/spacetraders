// SpaceTraders v2.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: trade_symbol.proto

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

type TradeSymbol int32

const (
	TradeSymbol_PRECIOUS_STONES            TradeSymbol = 0
	TradeSymbol_QUARTZ_SAND                TradeSymbol = 1
	TradeSymbol_SILICON_CRYSTALS           TradeSymbol = 2
	TradeSymbol_AMMONIA_ICE                TradeSymbol = 3
	TradeSymbol_LIQUID_HYDROGEN            TradeSymbol = 4
	TradeSymbol_LIQUID_NITROGEN            TradeSymbol = 5
	TradeSymbol_ICE_WATER                  TradeSymbol = 6
	TradeSymbol_EXOTIC_MATTER              TradeSymbol = 7
	TradeSymbol_ADVANCED_CIRCUITRY         TradeSymbol = 8
	TradeSymbol_GRAVITON_EMITTERS          TradeSymbol = 9
	TradeSymbol_IRON                       TradeSymbol = 10
	TradeSymbol_IRON_ORE                   TradeSymbol = 11
	TradeSymbol_COPPER                     TradeSymbol = 12
	TradeSymbol_COPPER_ORE                 TradeSymbol = 13
	TradeSymbol_ALUMINUM                   TradeSymbol = 14
	TradeSymbol_ALUMINUM_ORE               TradeSymbol = 15
	TradeSymbol_SILVER                     TradeSymbol = 16
	TradeSymbol_SILVER_ORE                 TradeSymbol = 17
	TradeSymbol_GOLD                       TradeSymbol = 18
	TradeSymbol_GOLD_ORE                   TradeSymbol = 19
	TradeSymbol_PLATINUM                   TradeSymbol = 20
	TradeSymbol_PLATINUM_ORE               TradeSymbol = 21
	TradeSymbol_DIAMONDS                   TradeSymbol = 22
	TradeSymbol_URANITE                    TradeSymbol = 23
	TradeSymbol_URANITE_ORE                TradeSymbol = 24
	TradeSymbol_MERITIUM                   TradeSymbol = 25
	TradeSymbol_MERITIUM_ORE               TradeSymbol = 26
	TradeSymbol_HYDROCARBON                TradeSymbol = 27
	TradeSymbol_ANTIMATTER                 TradeSymbol = 28
	TradeSymbol_FERTILIZERS                TradeSymbol = 29
	TradeSymbol_FABRICS                    TradeSymbol = 30
	TradeSymbol_FOOD                       TradeSymbol = 31
	TradeSymbol_JEWELRY                    TradeSymbol = 32
	TradeSymbol_MACHINERY                  TradeSymbol = 33
	TradeSymbol_FIREARMS                   TradeSymbol = 34
	TradeSymbol_ASSAULT_RIFLES             TradeSymbol = 35
	TradeSymbol_MILITARY_EQUIPMENT         TradeSymbol = 36
	TradeSymbol_EXPLOSIVES                 TradeSymbol = 37
	TradeSymbol_LAB_INSTRUMENTS            TradeSymbol = 38
	TradeSymbol_AMMUNITION                 TradeSymbol = 39
	TradeSymbol_ELECTRONICS                TradeSymbol = 40
	TradeSymbol_SHIP_PLATING               TradeSymbol = 41
	TradeSymbol_EQUIPMENT                  TradeSymbol = 42
	TradeSymbol_FUEL                       TradeSymbol = 43
	TradeSymbol_MEDICINE                   TradeSymbol = 44
	TradeSymbol_DRUGS                      TradeSymbol = 45
	TradeSymbol_CLOTHING                   TradeSymbol = 46
	TradeSymbol_MICROPROCESSORS            TradeSymbol = 47
	TradeSymbol_PLASTICS                   TradeSymbol = 48
	TradeSymbol_POLYNUCLEOTIDES            TradeSymbol = 49
	TradeSymbol_BIOCOMPOSITES              TradeSymbol = 50
	TradeSymbol_NANOBOTS                   TradeSymbol = 51
	TradeSymbol_AI_MAINFRAMES              TradeSymbol = 52
	TradeSymbol_QUANTUM_DRIVES             TradeSymbol = 53
	TradeSymbol_ROBOTIC_DRONES             TradeSymbol = 54
	TradeSymbol_CYBER_IMPLANTS             TradeSymbol = 55
	TradeSymbol_GENE_THERAPEUTICS          TradeSymbol = 56
	TradeSymbol_NEURAL_CHIPS               TradeSymbol = 57
	TradeSymbol_MOOD_REGULATORS            TradeSymbol = 58
	TradeSymbol_VIRAL_AGENTS               TradeSymbol = 59
	TradeSymbol_MICRO_FUSION_GENERATORS    TradeSymbol = 60
	TradeSymbol_SUPERGRAINS                TradeSymbol = 61
	TradeSymbol_LASER_RIFLES               TradeSymbol = 62
	TradeSymbol_HOLOGRAPHICS               TradeSymbol = 63
	TradeSymbol_SHIP_SALVAGE               TradeSymbol = 64
	TradeSymbol_RELIC_TECH                 TradeSymbol = 65
	TradeSymbol_NOVEL_LIFEFORMS            TradeSymbol = 66
	TradeSymbol_BOTANICAL_SPECIMENS        TradeSymbol = 67
	TradeSymbol_CULTURAL_ARTIFACTS         TradeSymbol = 68
	TradeSymbol_REACTOR_SOLAR_I            TradeSymbol = 69
	TradeSymbol_REACTOR_FUSION_I           TradeSymbol = 70
	TradeSymbol_REACTOR_FISSION_I          TradeSymbol = 71
	TradeSymbol_REACTOR_CHEMICAL_I         TradeSymbol = 72
	TradeSymbol_REACTOR_ANTIMATTER_I       TradeSymbol = 73
	TradeSymbol_ENGINE_IMPULSE_DRIVE_I     TradeSymbol = 74
	TradeSymbol_ENGINE_ION_DRIVE_I         TradeSymbol = 75
	TradeSymbol_ENGINE_ION_DRIVE_II        TradeSymbol = 76
	TradeSymbol_ENGINE_HYPER_DRIVE_I       TradeSymbol = 77
	TradeSymbol_MODULE_MINERAL_PROCESSOR_I TradeSymbol = 78
	TradeSymbol_MODULE_CARGO_HOLD_I        TradeSymbol = 79
	TradeSymbol_MODULE_CREW_QUARTERS_I     TradeSymbol = 80
	TradeSymbol_MODULE_ENVOY_QUARTERS_I    TradeSymbol = 81
	TradeSymbol_MODULE_PASSENGER_CABIN_I   TradeSymbol = 82
	TradeSymbol_MODULE_MICRO_REFINERY_I    TradeSymbol = 83
	TradeSymbol_MODULE_ORE_REFINERY_I      TradeSymbol = 84
	TradeSymbol_MODULE_FUEL_REFINERY_I     TradeSymbol = 85
	TradeSymbol_MODULE_SCIENCE_LAB_I       TradeSymbol = 86
	TradeSymbol_MODULE_JUMP_DRIVE_I        TradeSymbol = 87
	TradeSymbol_MODULE_JUMP_DRIVE_II       TradeSymbol = 88
	TradeSymbol_MODULE_JUMP_DRIVE_III      TradeSymbol = 89
	TradeSymbol_MODULE_WARP_DRIVE_I        TradeSymbol = 90
	TradeSymbol_MODULE_WARP_DRIVE_II       TradeSymbol = 91
	TradeSymbol_MODULE_WARP_DRIVE_III      TradeSymbol = 92
	TradeSymbol_MODULE_SHIELD_GENERATOR_I  TradeSymbol = 93
	TradeSymbol_MODULE_SHIELD_GENERATOR_II TradeSymbol = 94
	TradeSymbol_MOUNT_GAS_SIPHON_I         TradeSymbol = 95
	TradeSymbol_MOUNT_GAS_SIPHON_II        TradeSymbol = 96
	TradeSymbol_MOUNT_GAS_SIPHON_III       TradeSymbol = 97
	TradeSymbol_MOUNT_SURVEYOR_I           TradeSymbol = 98
	TradeSymbol_MOUNT_SURVEYOR_II          TradeSymbol = 99
	TradeSymbol_MOUNT_SURVEYOR_III         TradeSymbol = 100
	TradeSymbol_MOUNT_SENSOR_ARRAY_I       TradeSymbol = 101
	TradeSymbol_MOUNT_SENSOR_ARRAY_II      TradeSymbol = 102
	TradeSymbol_MOUNT_SENSOR_ARRAY_III     TradeSymbol = 103
	TradeSymbol_MOUNT_MINING_LASER_I       TradeSymbol = 104
	TradeSymbol_MOUNT_MINING_LASER_II      TradeSymbol = 105
	TradeSymbol_MOUNT_MINING_LASER_III     TradeSymbol = 106
	TradeSymbol_MOUNT_LASER_CANNON_I       TradeSymbol = 107
	TradeSymbol_MOUNT_MISSILE_LAUNCHER_I   TradeSymbol = 108
	TradeSymbol_MOUNT_TURRET_I             TradeSymbol = 109
)

// Enum value maps for TradeSymbol.
var (
	TradeSymbol_name = map[int32]string{
		0:   "PRECIOUS_STONES",
		1:   "QUARTZ_SAND",
		2:   "SILICON_CRYSTALS",
		3:   "AMMONIA_ICE",
		4:   "LIQUID_HYDROGEN",
		5:   "LIQUID_NITROGEN",
		6:   "ICE_WATER",
		7:   "EXOTIC_MATTER",
		8:   "ADVANCED_CIRCUITRY",
		9:   "GRAVITON_EMITTERS",
		10:  "IRON",
		11:  "IRON_ORE",
		12:  "COPPER",
		13:  "COPPER_ORE",
		14:  "ALUMINUM",
		15:  "ALUMINUM_ORE",
		16:  "SILVER",
		17:  "SILVER_ORE",
		18:  "GOLD",
		19:  "GOLD_ORE",
		20:  "PLATINUM",
		21:  "PLATINUM_ORE",
		22:  "DIAMONDS",
		23:  "URANITE",
		24:  "URANITE_ORE",
		25:  "MERITIUM",
		26:  "MERITIUM_ORE",
		27:  "HYDROCARBON",
		28:  "ANTIMATTER",
		29:  "FERTILIZERS",
		30:  "FABRICS",
		31:  "FOOD",
		32:  "JEWELRY",
		33:  "MACHINERY",
		34:  "FIREARMS",
		35:  "ASSAULT_RIFLES",
		36:  "MILITARY_EQUIPMENT",
		37:  "EXPLOSIVES",
		38:  "LAB_INSTRUMENTS",
		39:  "AMMUNITION",
		40:  "ELECTRONICS",
		41:  "SHIP_PLATING",
		42:  "EQUIPMENT",
		43:  "FUEL",
		44:  "MEDICINE",
		45:  "DRUGS",
		46:  "CLOTHING",
		47:  "MICROPROCESSORS",
		48:  "PLASTICS",
		49:  "POLYNUCLEOTIDES",
		50:  "BIOCOMPOSITES",
		51:  "NANOBOTS",
		52:  "AI_MAINFRAMES",
		53:  "QUANTUM_DRIVES",
		54:  "ROBOTIC_DRONES",
		55:  "CYBER_IMPLANTS",
		56:  "GENE_THERAPEUTICS",
		57:  "NEURAL_CHIPS",
		58:  "MOOD_REGULATORS",
		59:  "VIRAL_AGENTS",
		60:  "MICRO_FUSION_GENERATORS",
		61:  "SUPERGRAINS",
		62:  "LASER_RIFLES",
		63:  "HOLOGRAPHICS",
		64:  "SHIP_SALVAGE",
		65:  "RELIC_TECH",
		66:  "NOVEL_LIFEFORMS",
		67:  "BOTANICAL_SPECIMENS",
		68:  "CULTURAL_ARTIFACTS",
		69:  "REACTOR_SOLAR_I",
		70:  "REACTOR_FUSION_I",
		71:  "REACTOR_FISSION_I",
		72:  "REACTOR_CHEMICAL_I",
		73:  "REACTOR_ANTIMATTER_I",
		74:  "ENGINE_IMPULSE_DRIVE_I",
		75:  "ENGINE_ION_DRIVE_I",
		76:  "ENGINE_ION_DRIVE_II",
		77:  "ENGINE_HYPER_DRIVE_I",
		78:  "MODULE_MINERAL_PROCESSOR_I",
		79:  "MODULE_CARGO_HOLD_I",
		80:  "MODULE_CREW_QUARTERS_I",
		81:  "MODULE_ENVOY_QUARTERS_I",
		82:  "MODULE_PASSENGER_CABIN_I",
		83:  "MODULE_MICRO_REFINERY_I",
		84:  "MODULE_ORE_REFINERY_I",
		85:  "MODULE_FUEL_REFINERY_I",
		86:  "MODULE_SCIENCE_LAB_I",
		87:  "MODULE_JUMP_DRIVE_I",
		88:  "MODULE_JUMP_DRIVE_II",
		89:  "MODULE_JUMP_DRIVE_III",
		90:  "MODULE_WARP_DRIVE_I",
		91:  "MODULE_WARP_DRIVE_II",
		92:  "MODULE_WARP_DRIVE_III",
		93:  "MODULE_SHIELD_GENERATOR_I",
		94:  "MODULE_SHIELD_GENERATOR_II",
		95:  "MOUNT_GAS_SIPHON_I",
		96:  "MOUNT_GAS_SIPHON_II",
		97:  "MOUNT_GAS_SIPHON_III",
		98:  "MOUNT_SURVEYOR_I",
		99:  "MOUNT_SURVEYOR_II",
		100: "MOUNT_SURVEYOR_III",
		101: "MOUNT_SENSOR_ARRAY_I",
		102: "MOUNT_SENSOR_ARRAY_II",
		103: "MOUNT_SENSOR_ARRAY_III",
		104: "MOUNT_MINING_LASER_I",
		105: "MOUNT_MINING_LASER_II",
		106: "MOUNT_MINING_LASER_III",
		107: "MOUNT_LASER_CANNON_I",
		108: "MOUNT_MISSILE_LAUNCHER_I",
		109: "MOUNT_TURRET_I",
	}
	TradeSymbol_value = map[string]int32{
		"PRECIOUS_STONES":            0,
		"QUARTZ_SAND":                1,
		"SILICON_CRYSTALS":           2,
		"AMMONIA_ICE":                3,
		"LIQUID_HYDROGEN":            4,
		"LIQUID_NITROGEN":            5,
		"ICE_WATER":                  6,
		"EXOTIC_MATTER":              7,
		"ADVANCED_CIRCUITRY":         8,
		"GRAVITON_EMITTERS":          9,
		"IRON":                       10,
		"IRON_ORE":                   11,
		"COPPER":                     12,
		"COPPER_ORE":                 13,
		"ALUMINUM":                   14,
		"ALUMINUM_ORE":               15,
		"SILVER":                     16,
		"SILVER_ORE":                 17,
		"GOLD":                       18,
		"GOLD_ORE":                   19,
		"PLATINUM":                   20,
		"PLATINUM_ORE":               21,
		"DIAMONDS":                   22,
		"URANITE":                    23,
		"URANITE_ORE":                24,
		"MERITIUM":                   25,
		"MERITIUM_ORE":               26,
		"HYDROCARBON":                27,
		"ANTIMATTER":                 28,
		"FERTILIZERS":                29,
		"FABRICS":                    30,
		"FOOD":                       31,
		"JEWELRY":                    32,
		"MACHINERY":                  33,
		"FIREARMS":                   34,
		"ASSAULT_RIFLES":             35,
		"MILITARY_EQUIPMENT":         36,
		"EXPLOSIVES":                 37,
		"LAB_INSTRUMENTS":            38,
		"AMMUNITION":                 39,
		"ELECTRONICS":                40,
		"SHIP_PLATING":               41,
		"EQUIPMENT":                  42,
		"FUEL":                       43,
		"MEDICINE":                   44,
		"DRUGS":                      45,
		"CLOTHING":                   46,
		"MICROPROCESSORS":            47,
		"PLASTICS":                   48,
		"POLYNUCLEOTIDES":            49,
		"BIOCOMPOSITES":              50,
		"NANOBOTS":                   51,
		"AI_MAINFRAMES":              52,
		"QUANTUM_DRIVES":             53,
		"ROBOTIC_DRONES":             54,
		"CYBER_IMPLANTS":             55,
		"GENE_THERAPEUTICS":          56,
		"NEURAL_CHIPS":               57,
		"MOOD_REGULATORS":            58,
		"VIRAL_AGENTS":               59,
		"MICRO_FUSION_GENERATORS":    60,
		"SUPERGRAINS":                61,
		"LASER_RIFLES":               62,
		"HOLOGRAPHICS":               63,
		"SHIP_SALVAGE":               64,
		"RELIC_TECH":                 65,
		"NOVEL_LIFEFORMS":            66,
		"BOTANICAL_SPECIMENS":        67,
		"CULTURAL_ARTIFACTS":         68,
		"REACTOR_SOLAR_I":            69,
		"REACTOR_FUSION_I":           70,
		"REACTOR_FISSION_I":          71,
		"REACTOR_CHEMICAL_I":         72,
		"REACTOR_ANTIMATTER_I":       73,
		"ENGINE_IMPULSE_DRIVE_I":     74,
		"ENGINE_ION_DRIVE_I":         75,
		"ENGINE_ION_DRIVE_II":        76,
		"ENGINE_HYPER_DRIVE_I":       77,
		"MODULE_MINERAL_PROCESSOR_I": 78,
		"MODULE_CARGO_HOLD_I":        79,
		"MODULE_CREW_QUARTERS_I":     80,
		"MODULE_ENVOY_QUARTERS_I":    81,
		"MODULE_PASSENGER_CABIN_I":   82,
		"MODULE_MICRO_REFINERY_I":    83,
		"MODULE_ORE_REFINERY_I":      84,
		"MODULE_FUEL_REFINERY_I":     85,
		"MODULE_SCIENCE_LAB_I":       86,
		"MODULE_JUMP_DRIVE_I":        87,
		"MODULE_JUMP_DRIVE_II":       88,
		"MODULE_JUMP_DRIVE_III":      89,
		"MODULE_WARP_DRIVE_I":        90,
		"MODULE_WARP_DRIVE_II":       91,
		"MODULE_WARP_DRIVE_III":      92,
		"MODULE_SHIELD_GENERATOR_I":  93,
		"MODULE_SHIELD_GENERATOR_II": 94,
		"MOUNT_GAS_SIPHON_I":         95,
		"MOUNT_GAS_SIPHON_II":        96,
		"MOUNT_GAS_SIPHON_III":       97,
		"MOUNT_SURVEYOR_I":           98,
		"MOUNT_SURVEYOR_II":          99,
		"MOUNT_SURVEYOR_III":         100,
		"MOUNT_SENSOR_ARRAY_I":       101,
		"MOUNT_SENSOR_ARRAY_II":      102,
		"MOUNT_SENSOR_ARRAY_III":     103,
		"MOUNT_MINING_LASER_I":       104,
		"MOUNT_MINING_LASER_II":      105,
		"MOUNT_MINING_LASER_III":     106,
		"MOUNT_LASER_CANNON_I":       107,
		"MOUNT_MISSILE_LAUNCHER_I":   108,
		"MOUNT_TURRET_I":             109,
	}
)

func (x TradeSymbol) Enum() *TradeSymbol {
	p := new(TradeSymbol)
	*p = x
	return p
}

func (x TradeSymbol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TradeSymbol) Descriptor() protoreflect.EnumDescriptor {
	return file_trade_symbol_proto_enumTypes[0].Descriptor()
}

func (TradeSymbol) Type() protoreflect.EnumType {
	return &file_trade_symbol_proto_enumTypes[0]
}

func (x TradeSymbol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TradeSymbol.Descriptor instead.
func (TradeSymbol) EnumDescriptor() ([]byte, []int) {
	return file_trade_symbol_proto_rawDescGZIP(), []int{0}
}

var File_trade_symbol_proto protoreflect.FileDescriptor

var file_trade_symbol_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6d, 0x61, 0x63, 0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2a, 0xe4, 0x11, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x45, 0x43, 0x49, 0x4f,
	0x55, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x4e, 0x45, 0x53, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x51,
	0x55, 0x41, 0x52, 0x54, 0x5a, 0x5f, 0x53, 0x41, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10,
	0x53, 0x49, 0x4c, 0x49, 0x43, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x59, 0x53, 0x54, 0x41, 0x4c, 0x53,
	0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x4d, 0x4d, 0x4f, 0x4e, 0x49, 0x41, 0x5f, 0x49, 0x43,
	0x45, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x51, 0x55, 0x49, 0x44, 0x5f, 0x48, 0x59,
	0x44, 0x52, 0x4f, 0x47, 0x45, 0x4e, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x51, 0x55,
	0x49, 0x44, 0x5f, 0x4e, 0x49, 0x54, 0x52, 0x4f, 0x47, 0x45, 0x4e, 0x10, 0x05, 0x12, 0x0d, 0x0a,
	0x09, 0x49, 0x43, 0x45, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d,
	0x45, 0x58, 0x4f, 0x54, 0x49, 0x43, 0x5f, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x52, 0x10, 0x07, 0x12,
	0x16, 0x0a, 0x12, 0x41, 0x44, 0x56, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x43, 0x49, 0x52, 0x43,
	0x55, 0x49, 0x54, 0x52, 0x59, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x52, 0x41, 0x56, 0x49,
	0x54, 0x4f, 0x4e, 0x5f, 0x45, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x52, 0x53, 0x10, 0x09, 0x12, 0x08,
	0x0a, 0x04, 0x49, 0x52, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x52, 0x4f, 0x4e,
	0x5f, 0x4f, 0x52, 0x45, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x50, 0x50, 0x45, 0x52,
	0x10, 0x0c, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x50, 0x50, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x45,
	0x10, 0x0d, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x4c, 0x55, 0x4d, 0x49, 0x4e, 0x55, 0x4d, 0x10, 0x0e,
	0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4c, 0x55, 0x4d, 0x49, 0x4e, 0x55, 0x4d, 0x5f, 0x4f, 0x52, 0x45,
	0x10, 0x0f, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x49, 0x4c, 0x56, 0x45, 0x52, 0x10, 0x10, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x49, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x11, 0x12, 0x08,
	0x0a, 0x04, 0x47, 0x4f, 0x4c, 0x44, 0x10, 0x12, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x4f, 0x4c, 0x44,
	0x5f, 0x4f, 0x52, 0x45, 0x10, 0x13, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4c, 0x41, 0x54, 0x49, 0x4e,
	0x55, 0x4d, 0x10, 0x14, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54, 0x49, 0x4e, 0x55, 0x4d,
	0x5f, 0x4f, 0x52, 0x45, 0x10, 0x15, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x41, 0x4d, 0x4f, 0x4e,
	0x44, 0x53, 0x10, 0x16, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x52, 0x41, 0x4e, 0x49, 0x54, 0x45, 0x10,
	0x17, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x52, 0x41, 0x4e, 0x49, 0x54, 0x45, 0x5f, 0x4f, 0x52, 0x45,
	0x10, 0x18, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45, 0x52, 0x49, 0x54, 0x49, 0x55, 0x4d, 0x10, 0x19,
	0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x52, 0x49, 0x54, 0x49, 0x55, 0x4d, 0x5f, 0x4f, 0x52, 0x45,
	0x10, 0x1a, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x59, 0x44, 0x52, 0x4f, 0x43, 0x41, 0x52, 0x42, 0x4f,
	0x4e, 0x10, 0x1b, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4e, 0x54, 0x49, 0x4d, 0x41, 0x54, 0x54, 0x45,
	0x52, 0x10, 0x1c, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x45, 0x52, 0x54, 0x49, 0x4c, 0x49, 0x5a, 0x45,
	0x52, 0x53, 0x10, 0x1d, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x42, 0x52, 0x49, 0x43, 0x53, 0x10,
	0x1e, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x4f, 0x4f, 0x44, 0x10, 0x1f, 0x12, 0x0b, 0x0a, 0x07, 0x4a,
	0x45, 0x57, 0x45, 0x4c, 0x52, 0x59, 0x10, 0x20, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41, 0x43, 0x48,
	0x49, 0x4e, 0x45, 0x52, 0x59, 0x10, 0x21, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x52, 0x45, 0x41,
	0x52, 0x4d, 0x53, 0x10, 0x22, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x53, 0x53, 0x41, 0x55, 0x4c, 0x54,
	0x5f, 0x52, 0x49, 0x46, 0x4c, 0x45, 0x53, 0x10, 0x23, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x49, 0x4c,
	0x49, 0x54, 0x41, 0x52, 0x59, 0x5f, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x24, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x58, 0x50, 0x4c, 0x4f, 0x53, 0x49, 0x56, 0x45, 0x53, 0x10,
	0x25, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x41, 0x42, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d,
	0x45, 0x4e, 0x54, 0x53, 0x10, 0x26, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4d, 0x4d, 0x55, 0x4e, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x27, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x52,
	0x4f, 0x4e, 0x49, 0x43, 0x53, 0x10, 0x28, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x48, 0x49, 0x50, 0x5f,
	0x50, 0x4c, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x29, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x51, 0x55,
	0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x2a, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x45, 0x4c,
	0x10, 0x2b, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45, 0x44, 0x49, 0x43, 0x49, 0x4e, 0x45, 0x10, 0x2c,
	0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x55, 0x47, 0x53, 0x10, 0x2d, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x4c, 0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x2e, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x49, 0x43,
	0x52, 0x4f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x4f, 0x52, 0x53, 0x10, 0x2f, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x4c, 0x41, 0x53, 0x54, 0x49, 0x43, 0x53, 0x10, 0x30, 0x12, 0x13, 0x0a, 0x0f,
	0x50, 0x4f, 0x4c, 0x59, 0x4e, 0x55, 0x43, 0x4c, 0x45, 0x4f, 0x54, 0x49, 0x44, 0x45, 0x53, 0x10,
	0x31, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x49, 0x4f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x45, 0x53, 0x10, 0x32, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x41, 0x4e, 0x4f, 0x42, 0x4f, 0x54, 0x53,
	0x10, 0x33, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x49, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x46, 0x52, 0x41,
	0x4d, 0x45, 0x53, 0x10, 0x34, 0x12, 0x12, 0x0a, 0x0e, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x55, 0x4d,
	0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x53, 0x10, 0x35, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x42,
	0x4f, 0x54, 0x49, 0x43, 0x5f, 0x44, 0x52, 0x4f, 0x4e, 0x45, 0x53, 0x10, 0x36, 0x12, 0x12, 0x0a,
	0x0e, 0x43, 0x59, 0x42, 0x45, 0x52, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x41, 0x4e, 0x54, 0x53, 0x10,
	0x37, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x45, 0x4e, 0x45, 0x5f, 0x54, 0x48, 0x45, 0x52, 0x41, 0x50,
	0x45, 0x55, 0x54, 0x49, 0x43, 0x53, 0x10, 0x38, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x55, 0x52,
	0x41, 0x4c, 0x5f, 0x43, 0x48, 0x49, 0x50, 0x53, 0x10, 0x39, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x4f,
	0x4f, 0x44, 0x5f, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x54, 0x4f, 0x52, 0x53, 0x10, 0x3a, 0x12,
	0x10, 0x0a, 0x0c, 0x56, 0x49, 0x52, 0x41, 0x4c, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x10,
	0x3b, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x5f, 0x46, 0x55, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x53, 0x10, 0x3c, 0x12, 0x0f,
	0x0a, 0x0b, 0x53, 0x55, 0x50, 0x45, 0x52, 0x47, 0x52, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x3d, 0x12,
	0x10, 0x0a, 0x0c, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x49, 0x46, 0x4c, 0x45, 0x53, 0x10,
	0x3e, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x4f, 0x4c, 0x4f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43,
	0x53, 0x10, 0x3f, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x53, 0x41, 0x4c, 0x56,
	0x41, 0x47, 0x45, 0x10, 0x40, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x4c, 0x49, 0x43, 0x5f, 0x54,
	0x45, 0x43, 0x48, 0x10, 0x41, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x4f, 0x56, 0x45, 0x4c, 0x5f, 0x4c,
	0x49, 0x46, 0x45, 0x46, 0x4f, 0x52, 0x4d, 0x53, 0x10, 0x42, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x4f,
	0x54, 0x41, 0x4e, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x4d, 0x45, 0x4e,
	0x53, 0x10, 0x43, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x55, 0x4c, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x5f,
	0x41, 0x52, 0x54, 0x49, 0x46, 0x41, 0x43, 0x54, 0x53, 0x10, 0x44, 0x12, 0x13, 0x0a, 0x0f, 0x52,
	0x45, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x53, 0x4f, 0x4c, 0x41, 0x52, 0x5f, 0x49, 0x10, 0x45,
	0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x46, 0x55, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x10, 0x46, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x41, 0x43, 0x54, 0x4f,
	0x52, 0x5f, 0x46, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x10, 0x47, 0x12, 0x16, 0x0a,
	0x12, 0x52, 0x45, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x43, 0x48, 0x45, 0x4d, 0x49, 0x43, 0x41,
	0x4c, 0x5f, 0x49, 0x10, 0x48, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x41, 0x43, 0x54, 0x4f, 0x52,
	0x5f, 0x41, 0x4e, 0x54, 0x49, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x52, 0x5f, 0x49, 0x10, 0x49, 0x12,
	0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x49, 0x4d, 0x50, 0x55, 0x4c, 0x53,
	0x45, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x10, 0x4a, 0x12, 0x16, 0x0a, 0x12, 0x45,
	0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f,
	0x49, 0x10, 0x4b, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x49, 0x10, 0x4c, 0x12, 0x18, 0x0a, 0x14,
	0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x48, 0x59, 0x50, 0x45, 0x52, 0x5f, 0x44, 0x52, 0x49,
	0x56, 0x45, 0x5f, 0x49, 0x10, 0x4d, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45,
	0x5f, 0x4d, 0x49, 0x4e, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x4f, 0x52, 0x5f, 0x49, 0x10, 0x4e, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45,
	0x5f, 0x43, 0x41, 0x52, 0x47, 0x4f, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x49, 0x10, 0x4f, 0x12,
	0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x57, 0x5f, 0x51,
	0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x53, 0x5f, 0x49, 0x10, 0x50, 0x12, 0x1b, 0x0a, 0x17, 0x4d,
	0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x56, 0x4f, 0x59, 0x5f, 0x51, 0x55, 0x41, 0x52,
	0x54, 0x45, 0x52, 0x53, 0x5f, 0x49, 0x10, 0x51, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x4f, 0x44, 0x55,
	0x4c, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x4e, 0x47, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x42,
	0x49, 0x4e, 0x5f, 0x49, 0x10, 0x52, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45,
	0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x5f, 0x52, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x52, 0x59, 0x5f,
	0x49, 0x10, 0x53, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x52,
	0x45, 0x5f, 0x52, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x52, 0x59, 0x5f, 0x49, 0x10, 0x54, 0x12, 0x1a,
	0x0a, 0x16, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x46, 0x55, 0x45, 0x4c, 0x5f, 0x52, 0x45,
	0x46, 0x49, 0x4e, 0x45, 0x52, 0x59, 0x5f, 0x49, 0x10, 0x55, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x4f,
	0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x43, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4c, 0x41, 0x42,
	0x5f, 0x49, 0x10, 0x56, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4a,
	0x55, 0x4d, 0x50, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x10, 0x57, 0x12, 0x18, 0x0a,
	0x14, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4a, 0x55, 0x4d, 0x50, 0x5f, 0x44, 0x52, 0x49,
	0x56, 0x45, 0x5f, 0x49, 0x49, 0x10, 0x58, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x44, 0x55, 0x4c,
	0x45, 0x5f, 0x4a, 0x55, 0x4d, 0x50, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x49, 0x49,
	0x10, 0x59, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x57, 0x41, 0x52,
	0x50, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x10, 0x5a, 0x12, 0x18, 0x0a, 0x14, 0x4d,
	0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x57, 0x41, 0x52, 0x50, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45,
	0x5f, 0x49, 0x49, 0x10, 0x5b, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f,
	0x57, 0x41, 0x52, 0x50, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x5c,
	0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x48, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x49, 0x10, 0x5d, 0x12,
	0x1e, 0x0a, 0x1a, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x48, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x49, 0x49, 0x10, 0x5e, 0x12,
	0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x47, 0x41, 0x53, 0x5f, 0x53, 0x49, 0x50,
	0x48, 0x4f, 0x4e, 0x5f, 0x49, 0x10, 0x5f, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f, 0x55, 0x4e, 0x54,
	0x5f, 0x47, 0x41, 0x53, 0x5f, 0x53, 0x49, 0x50, 0x48, 0x4f, 0x4e, 0x5f, 0x49, 0x49, 0x10, 0x60,
	0x12, 0x18, 0x0a, 0x14, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x47, 0x41, 0x53, 0x5f, 0x53, 0x49,
	0x50, 0x48, 0x4f, 0x4e, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x61, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59, 0x4f, 0x52, 0x5f, 0x49, 0x10, 0x62,
	0x12, 0x15, 0x0a, 0x11, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59,
	0x4f, 0x52, 0x5f, 0x49, 0x49, 0x10, 0x63, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x55, 0x4e, 0x54,
	0x5f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59, 0x4f, 0x52, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x64, 0x12,
	0x18, 0x0a, 0x14, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f,
	0x41, 0x52, 0x52, 0x41, 0x59, 0x5f, 0x49, 0x10, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x5f,
	0x49, 0x49, 0x10, 0x66, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x45,
	0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x67,
	0x12, 0x18, 0x0a, 0x14, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4e, 0x47,
	0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x10, 0x68, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52,
	0x5f, 0x49, 0x49, 0x10, 0x69, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d,
	0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x49, 0x49, 0x10,
	0x6a, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52,
	0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x4e, 0x5f, 0x49, 0x10, 0x6b, 0x12, 0x1c, 0x0a, 0x18, 0x4d,
	0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4c, 0x45, 0x5f, 0x4c, 0x41, 0x55,
	0x4e, 0x43, 0x48, 0x45, 0x52, 0x5f, 0x49, 0x10, 0x6c, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x54, 0x55, 0x52, 0x52, 0x45, 0x54, 0x5f, 0x49, 0x10, 0x6d, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trade_symbol_proto_rawDescOnce sync.Once
	file_trade_symbol_proto_rawDescData = file_trade_symbol_proto_rawDesc
)

func file_trade_symbol_proto_rawDescGZIP() []byte {
	file_trade_symbol_proto_rawDescOnce.Do(func() {
		file_trade_symbol_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_symbol_proto_rawDescData)
	})
	return file_trade_symbol_proto_rawDescData
}

var file_trade_symbol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trade_symbol_proto_goTypes = []interface{}{
	(TradeSymbol)(0), // 0: macgenius.spacetraders.api.trade.TradeSymbol
}
var file_trade_symbol_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trade_symbol_proto_init() }
func file_trade_symbol_proto_init() {
	if File_trade_symbol_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trade_symbol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trade_symbol_proto_goTypes,
		DependencyIndexes: file_trade_symbol_proto_depIdxs,
		EnumInfos:         file_trade_symbol_proto_enumTypes,
	}.Build()
	File_trade_symbol_proto = out.File
	file_trade_symbol_proto_rawDesc = nil
	file_trade_symbol_proto_goTypes = nil
	file_trade_symbol_proto_depIdxs = nil
}
