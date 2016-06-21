package main

import (
	//"bufio"
	//"fmt"
	"iespec"
	"net"
)

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:4739")

	conn, _ := net.DialUDP("udp", nil, serverAddr)
	defer conn.Close()

	packetBytes := iespec.ExampleBytes()
	for i := 0; i < 10; i++ {
		conn.Write(packetBytes)
	}

}
