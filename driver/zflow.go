package main

import (
	"flag"
	"fmt"
	"iespec"
	"os"
	"strings"
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
	OUTBOUND_SPLUNK = "outbound-splunk"
	OUTBOUND_KAFKA  = "outbound-kafka"
	OUTBOUND_JSON   = "outbound-json"
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
}

func serveudp(address string) {
	fmt.Printf("Listening on UDP %s", address)
	iespec.UDPServer()
}

func main() {
	inboundHTTP := flag.String(INBOUND_HTTP, "", "Binds HTTP to <host:port>")
	inboundHTTPS := flag.String(INBOUND_HTTPS, "localhost:9443", "Binds HTTPS to <host:port> (cert.pem and key.pem required in same folder)")
	inboundKafka := flag.String(INBOUND_KAFKA, "localhost:9092:ZIFTEN.IPFIX", "Enables Inbound Kafka Topic <host:port:topic>")
	inboundUDP := flag.String(INBOUND_UDP, "4739", "Enables UDP <host:port>")

	outboundSplunk := flag.String(OUTBOUND_SPLUNK, "http://localhost:8088/services/collector", "Forwards data to Splunk (http://localhost:8088/services/collector) is default")
	outboundKafka := flag.String(OUTBOUND_KAFKA, "localhost:9092:ZIFTEN.DATACOLLECTION", "Forwards data to Kafka <host:port:topic>")
	outboundJSON := flag.Bool(OUTBOUND_JSON, false, "Prints to STDOUT in JSON")

	flag.Parse()

	if isSet(INBOUND_HTTP) {
		servehttp(*inboundHTTP)
	}
	if isSet(INBOUND_HTTPS) {
		servehttps(*inboundHTTPS)
	}
	if isSet(INBOUND_KAFKA) {
		readfromkafka(*inboundKafka)
	}
	if isSet(INBOUND_UDP) {
		serveudp(*inboundUDP)
	}
	if isSet(OUTBOUND_SPLUNK) {
		fmt.Println(*outboundSplunk)
	}
	if isSet(OUTBOUND_KAFKA) {
		fmt.Println(*outboundKafka)
	}
	if isSet(OUTBOUND_JSON) {
		fmt.Println(*outboundJSON)
	}
}
