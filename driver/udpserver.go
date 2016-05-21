package main

import (
	"encoding/base64"
	"fmt"
	"github.com/calmh/ipfix"
	"iespec"
	"net"
	"reflect"
)

func handleConnection(conn net.Conn) {

}

func handleZflow(b []byte, s *ipfix.Session, i *ipfix.Interpreter) {

	msg, _ := s.ParseBuffer(b)
	//if len(msg.DataRecords) > 0 {
	fmt.Printf("Template Records: %d DataRecords: %d\n", len(msg.TemplateRecords), len(msg.DataRecords))
	//}
	for _, rec := range msg.DataRecords {
		fmt.Printf("\n\n")
		//fmt.Printf("\nRecord %d : %s\n\n", j, rec)
		ifs := i.Interpret(rec)
		for _, ifield := range ifs {
			if ifield.Value != nil {
				fmt.Printf("Field: %s, Value: %s Type: %s \n", ifield.Name, ifield.Value, reflect.TypeOf(ifield.Value))
			}
		}
	}
}

func printEncoded(buf []byte) {
	encoded := base64.StdEncoding.EncodeToString(buf)
	fmt.Println(encoded)
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
			handleZflow(buf[:n], s, i)
		}
		if doPrintEncoded {
			printEncoded(buf[:n])
		}
	}
}
