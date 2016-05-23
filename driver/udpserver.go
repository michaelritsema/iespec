package main

import (
	"encoding/base64"
	"fmt"
	"github.com/calmh/ipfix"
	"github.com/golang/protobuf/jsonpb"
	"iespec"
	"iespec/protomsg"
	"net"
	//"reflect"
)

func handleConnection(conn net.Conn) {

}

func handleZflow(b []byte, s *ipfix.Session, i *ipfix.Interpreter) []*protomsg.ZFlow {

	msg, _ := s.ParseBuffer(b)
	//if len(msg.DataRecords) > 0 {
	fmt.Printf("Template Records: %d DataRecords: %d\n", len(msg.TemplateRecords), len(msg.DataRecords))
	//}

	pmsgList := make([]*protomsg.ZFlow, 0)

	for _, rec := range msg.DataRecords {
		fmt.Printf("\n\n")
		//fmt.Printf("\nRecord %d : %s\n\n", j, rec)
		ifs := i.Interpret(rec)
		/*
			for _, ifield := range ifs {
				if ifield.Value != nil {
					fmt.Printf("Field: %s, Value: %s Type: %s \n", ifield.Name, ifield.Value, reflect.TypeOf(ifield.Value))
				}
			}
		*/
		pmsg := iespec.ConvertFieldListToProtobuf(ifs)
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
		iespec.SplunkPOST(jsonString)
	}
}

func main() {
	doPrintEncoded := false
	doZflow := true
	s := ipfix.NewSession()
	i := ipfix.NewInterpreter(s)
	for _, entry := range iespec.MyFields {
		//fmt.Println(entry)
		i.AddDictionaryEntry(entry)
	}

	serverAddr, _ := net.ResolveUDPAddr("udp", ":4739")
	ln, _ := net.ListenUDP("udp", serverAddr)
	defer ln.Close()

	for {
		fmt.Print(".")
		buf := make([]byte, 64000)
		n, _, _ := ln.ReadFromUDP(buf)

		if doZflow {
			pmsgList := handleZflow(buf[:n], s, i)
			for _, msg := range pmsgList {
				processMsg(msg, outputOptions{STDOUT_JSON: false, SPLUNK_HTTP: true})
			}
		}
		if doPrintEncoded {
			printEncoded(buf[:n])
		}
	}
}
