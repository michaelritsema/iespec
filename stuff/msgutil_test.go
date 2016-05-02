package stuff

import (
	"testing"
	"github.com/golang/protobuf/proto"
	"iespec"
	"log"
)

func TestWrap(t *testing.T) {
		msg := &iespec.ZFlow {
			AgentGUID: proto.String("xxxx"),
		}

		data, _ := proto.Marshal(msg)
		msgstr := Wrap("bad_type", data)
		log.Println(msgstr)
 }