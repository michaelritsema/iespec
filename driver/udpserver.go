package main

import (
	"encoding/base64"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {

}
func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", ":4739")
	ln, _ := net.ListenUDP("udp", serverAddr)
	defer ln.Close()

	for {
		buf := make([]byte, 64000)

		n, _, _ := ln.ReadFromUDP(buf)
		//fmt.Println(buf[:n])
		encoded := base64.StdEncoding.EncodeToString(buf[:n])
		fmt.Println(encoded)
		//fmt.Println(buf)
		//fmt.Println("Received ",string(buf[0:n]))

	}

}
