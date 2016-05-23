// Code generated by protoc-gen-go.
// source: IpfixEncapsulation.proto
// DO NOT EDIT!

/*
Package IpfixEncapsulation is a generated protocol buffer package.

It is generated from these files:
	IpfixEncapsulation.proto

It has these top-level messages:
	IpfixEncapsulation
*/
package ipfix

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
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xf0, 0x2c, 0x48, 0xcb,
	0xac, 0x70, 0xcd, 0x4b, 0x4e, 0x2c, 0x28, 0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0xca, 0xe3, 0x12, 0xc2, 0x94, 0x13, 0x12, 0xe4, 0xe2, 0x2c, 0xc9,
	0xcc, 0x4d, 0x0d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x54, 0x60, 0xd2, 0x60, 0x06, 0x09, 0x25,
	0xa6, 0xa7, 0xe6, 0x95, 0xb8, 0x87, 0x7a, 0xba, 0x48, 0x30, 0x01, 0x85, 0x38, 0x85, 0x44, 0xb8,
	0x78, 0x32, 0x41, 0x7a, 0x7d, 0x53, 0x8b, 0x8b, 0x81, 0x72, 0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0x3c,
	0x42, 0x7c, 0x5c, 0x6c, 0xc5, 0x99, 0x25, 0xa9, 0x9e, 0x29, 0x12, 0x2c, 0x0a, 0x8c, 0x40, 0x55,
	0x3c, 0x5c, 0x2c, 0xa5, 0xa5, 0x99, 0x29, 0x12, 0xac, 0x20, 0x9e, 0x93, 0x3d, 0x97, 0x52, 0x72,
	0x7e, 0xae, 0x5e, 0x55, 0x66, 0x5a, 0x49, 0x6a, 0x9e, 0x5e, 0x71, 0x6a, 0x51, 0x59, 0x6a, 0x11,
	0xc4, 0x29, 0xc9, 0xf9, 0x39, 0x7a, 0xb9, 0x10, 0x93, 0x9c, 0x24, 0x31, 0xdd, 0x04, 0xb5, 0x04,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xd0, 0x29, 0x03, 0xcb, 0x00, 0x00, 0x00,
}