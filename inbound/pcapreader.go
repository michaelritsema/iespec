package inbound

import (
	//"flag"
	//"fmt"
	"github.com/miekg/pcap"
	"iespec"
	"iespec/protomsg"
)

func min(x uint32, y uint32) uint32 {
	if x < y {
		return x
	}
	return y
}

func Readpcapfile(c chan *protomsg.ZFlow, file string) {

	converter := iespec.InitIpfix()
	//var file *string = flag.String("r", "", "file")

	//flag.Parse()

	var h *pcap.Pcap

	h, _ = pcap.OpenOffline(file)

	for pkt := h.Next(); pkt != nil; pkt = h.Next() {
		//fmt.Println(".")
		//fmt.Printf("%s\n\n", pkt.Type)
		//buf := make([]byte, 64000)
		pkt.Decode()
		pmsgList := converter.Convert(pkt.Payload)
		for _, msg := range pmsgList {
			//fmt.Println(msg)
			c <- msg
		}

	}
}
