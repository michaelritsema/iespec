// Code generated by protoc-gen-go.
// source: PostgresType.proto
// DO NOT EDIT!

package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "iespec/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var E_PostgresType = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "postgresType",
	Tag:           "bytes,50000,opt,name=postgresType",
}

func init() {
	proto.RegisterExtension(E_PostgresType)
}

var fileDescriptor1 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x0a, 0xc8, 0x2f, 0x2e,
	0x49, 0x2f, 0x4a, 0x2d, 0x0e, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x52,
	0x48, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x52, 0x8b,
	0x93, 0x8b, 0x32, 0x0b, 0x4a, 0xf2, 0x8b, 0x20, 0x2a, 0xac, 0x4c, 0xb9, 0x78, 0x0a, 0x90, 0xf4,
	0x09, 0xc9, 0xea, 0x41, 0xb4, 0xe8, 0xc1, 0xb4, 0xe8, 0xb9, 0x65, 0xa6, 0xe6, 0xa4, 0xf8, 0x17,
	0x94, 0x64, 0xe6, 0xe7, 0x15, 0x4b, 0x5c, 0x68, 0x63, 0x56, 0x60, 0xd4, 0xe0, 0x74, 0x32, 0xe2,
	0x52, 0x4a, 0xce, 0xcf, 0xd5, 0xab, 0xca, 0x4c, 0x2b, 0x49, 0xcd, 0xd3, 0x2b, 0x4e, 0x2d, 0x2a,
	0x4b, 0x85, 0x9a, 0x98, 0x9c, 0x9f, 0xa3, 0x97, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0xea, 0xc4,
	0x83, 0xec, 0x24, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x3a, 0xca, 0x57, 0xa1, 0x00, 0x00,
	0x00,
}
