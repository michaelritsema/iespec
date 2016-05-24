package iespec

import (
	"encoding/base64"
	"fmt"
	"github.com/calmh/ipfix"
	"github.com/golang/protobuf/jsonpb"
	"iespec/protomsg"
	"net"
)

func handleZflow(b []byte, s *ipfix.Session, i *ipfix.Interpreter) []*protomsg.ZFlow {

	msg, _ := s.ParseBuffer(b)
	pmsgList := make([]*protomsg.ZFlow, 0)

	for _, rec := range msg.DataRecords {
		ifs := i.Interpret(rec)
		pmsg := ConvertFieldListToProtobuf(ifs)
		pmsgList = append(pmsgList, pmsg)
	}
	return pmsgList
}

func printEncoded(buf []byte) {
	encoded := base64.StdEncoding.EncodeToString(buf)
	fmt.Println(encoded)
}

type outputOptions struct {
	STDOUT_JSON bool
	SPLUNK_HTTP bool
	BROADCAST   bool
}

func processMsg(msg *protomsg.ZFlow, output outputOptions) {
	if output.STDOUT_JSON {
		marshaler := jsonpb.Marshaler{}
		jsonMsg, _ := marshaler.MarshalToString(msg)
		fmt.Println(jsonMsg)
	}

	if output.SPLUNK_HTTP {
		marshaler := jsonpb.Marshaler{}
		jsonString, _ := marshaler.MarshalToString(msg)
		SplunkPOST(jsonString)
	}

}

func broadcast(c, msg *protomsg.ZFlow) {

}

func UDPServer(c chan *protomsg.ZFlow, host string, port string) {

	doPrintEncoded := false
	doZflow := true
	s := ipfix.NewSession()
	i := ipfix.NewInterpreter(s)
	for _, entry := range MyFields {
		i.AddDictionaryEntry(entry)
	}

	serverAddr, _ := net.ResolveUDPAddr("udp", host+":"+port)
	ln, _ := net.ListenUDP("udp", serverAddr)
	defer ln.Close()

	for {
		buf := make([]byte, 64000)
		n, _, _ := ln.ReadFromUDP(buf)

		if doZflow {
			pmsgList := handleZflow(buf[:n], s, i)
			for _, msg := range pmsgList {
				processMsg(msg, outputOptions{STDOUT_JSON: false, SPLUNK_HTTP: true})
				c <- msg
			}
		}
		if doPrintEncoded {
			printEncoded(buf[:n])
		}

	}
}
