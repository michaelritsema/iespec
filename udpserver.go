package iespec

import (
	"encoding/base64"
	"fmt"
	"github.com/calmh/ipfix"
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

func UDPServer(c chan *protomsg.ZFlow, host string, port string) {

	doPrintEncoded := false

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

		pmsgList := handleZflow(buf[:n], s, i)
		for _, msg := range pmsgList {
			c <- msg
		}

		if doPrintEncoded {
			printEncoded(buf[:n])
		}

	}
}
