package protomsg_test

import (
	"iespec/protomsg"
	"testing"
	"github.com/golang/protobuf/proto"
	"reflect"
)

func TestProto(t *testing.T) {
	test := &protomsg.ZFlow {
		ZflowImagePath: proto.String("/image/path"),

	}

	if(*test.ZflowImagePath != "/image/path") {
		t.Error("ZFlowImagePath not set")
	}

	test.AgentGUID = proto.String("xxx")
	
	if(*test.AgentGUID != "xxx") {
		t.Error("agentGUID not set")
	}	
 }