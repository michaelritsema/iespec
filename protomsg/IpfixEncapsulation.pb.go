// Code generated by protoc-gen-go.
// source: IpfixEncapsulation.proto
// DO NOT EDIT!

/*
Package protomsg is a generated protocol buffer package.

It is generated from these files:
	IpfixEncapsulation.proto
	PostgresType.proto
	ZFlow.proto

It has these top-level messages:
	IpfixEncapsulation
	ZFlow
*/
package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// This conveys a list of IPFIX-formatted messages.  This is used to implement a
// server-side 'concentrator' of messages generated by zFlow clients.
type IpfixEncapsulation struct {
	// The GetSystemTimeAsFileTime that the message was produced by the agent.
	// Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to
	// distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// a series of IPFIX-formatted messages; these may be sent verbatim to the
	// collecting process.
	IpfixMessage     [][]byte `protobuf:"bytes,3,rep,name=ipfixMessage" json:"ipfixMessage,omitempty"`
	SiteId           *string  `protobuf:"bytes,4,opt,name=siteId" json:"siteId,omitempty"`
	Uuid             *string  `protobuf:"bytes,5,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IpfixEncapsulation) Reset()                    { *m = IpfixEncapsulation{} }
func (m *IpfixEncapsulation) String() string            { return proto.CompactTextString(m) }
func (*IpfixEncapsulation) ProtoMessage()               {}
func (*IpfixEncapsulation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IpfixEncapsulation) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *IpfixEncapsulation) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *IpfixEncapsulation) GetIpfixMessage() [][]byte {
	if m != nil {
		return m.IpfixMessage
	}
	return nil
}

func (m *IpfixEncapsulation) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *IpfixEncapsulation) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*IpfixEncapsulation)(nil), "IpfixEncapsulation")
}

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xf0, 0x2c, 0x48, 0xcb,
	0xac, 0x70, 0xcd, 0x4b, 0x4e, 0x2c, 0x28, 0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x97, 0x12, 0x0a, 0xc8, 0x2f, 0x2e, 0x49, 0x2f, 0x4a, 0x2d, 0x0e, 0xa9,
	0x2c, 0x48, 0x85, 0x88, 0x29, 0x4d, 0x64, 0xe4, 0x12, 0xc2, 0xd4, 0x20, 0xa4, 0xc0, 0xc5, 0x59,
	0x92, 0x99, 0x9b, 0x1a, 0x5c, 0x92, 0x98, 0x5b, 0x20, 0xc1, 0xa8, 0xc0, 0xa4, 0xc1, 0xec, 0xc4,
	0xdb, 0xb4, 0x55, 0x02, 0x2c, 0x58, 0x0c, 0x12, 0x14, 0x92, 0xe6, 0xe2, 0x4c, 0x4c, 0x4f, 0xcd,
	0x2b, 0x71, 0x0f, 0xf5, 0x74, 0x91, 0x60, 0x02, 0xaa, 0xe0, 0x74, 0xe2, 0x00, 0xaa, 0x60, 0x29,
	0x2d, 0xcd, 0x4c, 0x11, 0x12, 0xe1, 0xe2, 0xc9, 0x04, 0x19, 0xea, 0x9b, 0x5a, 0x5c, 0x0c, 0x54,
	0x25, 0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x23, 0xc4, 0xc7, 0xc5, 0x56, 0x9c, 0x59, 0x92, 0xea, 0x99,
	0x22, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x29, 0x24, 0xc6, 0x05, 0x56, 0x2d, 0xc1, 0x0a, 0xe2, 0x21,
	0x74, 0x3b, 0xd9, 0x73, 0x29, 0x25, 0xe7, 0xe7, 0xea, 0x55, 0x65, 0xa6, 0x95, 0xa4, 0xe6, 0xe9,
	0x15, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x41, 0x9c, 0x9b, 0x9c, 0x9f, 0xa3, 0x97, 0x0b, 0x31, 0xd3,
	0x49, 0x12, 0xd3, 0xd9, 0x50, 0xeb, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0x44, 0xac, 0x53, 0xa4,
	0x03, 0x01, 0x00, 0x00,
}
