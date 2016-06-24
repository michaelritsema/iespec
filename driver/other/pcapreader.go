package main

import (
	"flag"
	"fmt"
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

	if *file != "" {
		h, _ = pcap.OpenOffline(*file)
		if h == nil {
			//fmt.Printf("Openoffline(%s) failed: %s\n", *file, err)
			return
		}
	}

	for pkt := h.Next(); pkt != nil; pkt = h.Next() {
		fmt.Printf("%s", pkt.Data)
		buf := make([]byte, 64000)

		pmsgList := converter.Convert(buf[:pkt.Caplen])
		for _, msg := range pmsgList {
			fmt.Println(msg)
		}

	}
}
func main() {
	readpcapfile()
	fmt.Println("PCAP Reader complete")

}
