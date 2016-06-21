package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"iespec/inbound"
	"iespec/outbound"
	"iespec/protomsg"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"
)

/*
	--inbound--

	--inbound-http
		Enables HTTP input
		Binds to first parameter in form <hostname:port> or "localhost:9080"

	--inbound-https
		Enables HTTPS input.
		Takes in a parameter in this form <hostname:port> or "localhost:9443"
		Requires file in folder cert.pem and key.pem

	--inbound-kafka
		Enables input from kafka queue
		Reads off topic ZIFTEN.IPFIX
		Messages are serialized with IpfixEncapsulation.proto

	--outbound-splunk
		Enables Splunk HTTP Endpoint modular
		POSTS message from first parameter in this format: <hostname:port:token>
			where token is the api key token provided by splunk

	--outbound-kafka
		Enables output to Kafka Topic ZIFTEN.DATACOLLOCTION
		Message is serialized into ZFlow.proto wrapped in the <pb type hhmac/> xml annotation

	--output-stdout-json
		Enabled STDOUT printing in JSON


*/

func isSet(flagName string) bool {
	// This function doesn't properly support the = assignment for command lines
	found := false
	for _, x := range os.Args[1:] {
		if !strings.HasPrefix(flagName, "-") {
			flagName = "-" + flagName
		}
		sepIndex := strings.IndexAny(x, "=")
		if sepIndex > 1 {
			x = x[:sepIndex]
		}

		if x == flagName {
			found = true
			break
		}
	}
	return found
}

const (
	INBOUND_HTTP    = "inbound-http"
	INBOUND_HTTPS   = "inbound-https"
	INBOUND_KAFKA   = "inbound-kafka"
	INBOUND_UDP     = "inbound-udp"
	INBOUND_PCAP    = "inbound-pcap"
	OUTBOUND_SPLUNK = "outbound-splunk"
	OUTBOUND_KAFKA  = "outbound-kafka"
	OUTBOUND_JSON   = "outbound-json"
	OUTBOUND_CSV    = "outbound-csv"
)

func servehttp(address string) {
	fmt.Printf("Listening for HTTP on %s\n", address)
}

func servehttps(address string) {
	// found cert and key?
	fmt.Printf("Listening for HTTPS on %s\n", address)
}

func readfromkafka(address string) {
	fmt.Printf("Reading messages for Kafka: %s", address)
	go func() { inbound.Kafka(broadcast) }()
}

func readfrompcap(file string) {
	fmt.Printf("Reading from pcap")
	inbound.Readpcapfile(broadcast, file)
}

func serveudp(param string) {
	fmt.Printf("Listening on UDP %s\n", param)
	parts := strings.Split(param, ":")
	host := parts[0]
	port := parts[1]

	inbound.UDPServer(broadcast, host, port)
}

var broadcast chan *protomsg.ZFlow = make(chan *protomsg.ZFlow)
var outChannels map[string]chan *protomsg.ZFlow = make(map[string]chan *protomsg.ZFlow, 0)

func route() {
	go func() {
		for {
			msg := <-broadcast
			//fmt.Println(msg)
			for _, v := range outChannels {
				//fmt.Printf("routing to outbound channel [%v,%v]\n", k, v)
				v <- msg
			}
		}
	}()
}

func printJSON(c chan *protomsg.ZFlow) {

	go func() {
		for {
			msg := <-outChannels[OUTBOUND_JSON]

			marshaler := jsonpb.Marshaler{}
			jsonMsg, _ := marshaler.MarshalToString(msg)
			fmt.Println(jsonMsg)
		}
	}()
}

func printCSV(c chan *protomsg.ZFlow) {

	/*go func() {
		for {
			msg := <-outChannels[OUTBOUND_CSV]
			//bytes := make([]byte, 0)
			//marshaler := jsonpb.Marshaler{}
			//jsonMsg, _ := marshaler.Marshal(bytes, msg)
			//fmt.Println(bytes)
		}
	}()*/
}

func kafka(c chan *protomsg.ZFlow) {
	outbound.Kafka(c)
}

func splunk(param string) {
	c := make(chan *protomsg.ZFlow)
	outChannels[OUTBOUND_SPLUNK] = c

	go func() {
		for {
			msg := <-outChannels[OUTBOUND_SPLUNK]
			marshaler := jsonpb.Marshaler{}
			jsonMsg, _ := marshaler.MarshalToString(msg)
			outbound.SplunkPOST(param, jsonMsg)
		}
	}()
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	inboundHTTP := flag.String(INBOUND_HTTP, "", "Binds HTTP to <host:port>")
	inboundHTTPS := flag.String(INBOUND_HTTPS, "0.0.0.0:9443", "Binds HTTPS to <host:port> (cert.pem and key.pem required in same folder)")
	inboundKafka := flag.String(INBOUND_KAFKA, "0.0.0.0:9092:ZIFTEN.IPFIX", "Enables Inbound Kafka Topic <host:port:topic>")
	inboundUDP := flag.String(INBOUND_UDP, "0.0.0.0:4739", "Enables UDP <host:port>")
	inboundPCAP := flag.String(INBOUND_PCAP, "file.pcap", "PCAP file")

	outboundSplunk := flag.String(OUTBOUND_SPLUNK, "http://localhost:8088/services/collector", "Forwards data to Splunk (http://localhost:8088/services/collector) is default")
	outboundKafka := flag.String(OUTBOUND_KAFKA, "localhost:9092:ZIFTEN.DATACOLLECTION", "Forwards data to Kafka <host:port:topic>")
	outboundJSON := flag.Bool(OUTBOUND_JSON, false, "Prints to STDOUT in JSON")
	outboundCSV := flag.Bool(OUTBOUND_CSV, false, "Prints to STDOUT in CSV")

	flag.Parse()
	daemonize := false

	if isSet(INBOUND_HTTP) {
		fmt.Println("http")
		servehttp(*inboundHTTP)
		daemonize = true
	}
	if isSet(INBOUND_HTTPS) {
		servehttps(*inboundHTTPS)
		daemonize = true
	}
	if isSet(INBOUND_KAFKA) {
		readfromkafka(*inboundKafka)
		daemonize = true
	}
	if isSet(INBOUND_UDP) {
		go serveudp(*inboundUDP)
		daemonize = true
	}
	if isSet(INBOUND_PCAP) {
		go readfrompcap(*inboundPCAP)
		daemonize = true
	}
	if isSet(OUTBOUND_SPLUNK) {
		splunk(*outboundSplunk)
	}
	if isSet(OUTBOUND_KAFKA) {
		fmt.Println(*outboundKafka)
		c := make(chan *protomsg.ZFlow)
		outChannels[OUTBOUND_KAFKA] = c
		kafka(c)
	}

	if *outboundJSON {
		fmt.Println("...")
		c := make(chan *protomsg.ZFlow)
		outChannels[OUTBOUND_JSON] = c

		printJSON(c)
	}

	if *outboundCSV {
		c := make(chan *protomsg.ZFlow)
		outChannels[OUTBOUND_CSV] = c

		printCSV(c)
	}

	fmt.Printf("out channels %v\n", outChannels)
	route()

	if daemonize {
		for {
			time.Sleep(100 * time.Millisecond)
		}
	}

}
