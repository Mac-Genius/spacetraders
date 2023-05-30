package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DataWrapper[T any] interface {
	GetData() T
}

type DataListWrapper[T any] interface {
	GetData() []T
	GetMeta() *Meta
	ProtoReflect() protoreflect.Message
}
