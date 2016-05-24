package inbound

import (
	"iespec"
	"iespec/protomsg"
	"net"
)

func UDPServer(c chan *protomsg.ZFlow, host string, port string) {

	converter := iespec.InitIpfix()

	serverAddr, _ := net.ResolveUDPAddr("udp", host+":"+port)
	ln, _ := net.ListenUDP("udp", serverAddr)
	defer ln.Close()

	for {
		buf := make([]byte, 64000)
		n, _, _ := ln.ReadFromUDP(buf)

		pmsgList := converter.Convert(buf[:n])
		for _, msg := range pmsgList {
			c <- msg
		}
	}
}
