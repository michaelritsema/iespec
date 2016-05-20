package main

import (
	"iespec/stuff"
	//"encoding/json"
	"flag"
	"github.com/calmh/ipfix"
	"github.com/golang/protobuf/proto"
	"iespec/protomsg"
	"io"
	"log"
	//"os"
	"fmt"
	"net"
	"time"
	//"encoding/binary"
)

var ipfixcatVersion string

type InterpretedRecord struct {
	ExportTime uint32               `json:"exportTime"`
	TemplateId uint16               `json:"templateId"`
	Fields     []myInterpretedField `json:"fields"`
}

// Because we want to control JSON serialization
type myInterpretedField struct {
	Name         string      `json:"name"`
	EnterpriseId uint32      `json:"enterprise,omitempty"`
	FieldId      uint16      `json:"field"`
	Value        interface{} `json:"value,omitempty"`
	RawValue     []int       `json:"raw,omitempty"`
}

func messagesGenerator(r io.Reader, s *ipfix.Session, i *ipfix.Interpreter) <-chan []InterpretedRecord {
	c := make(chan []InterpretedRecord)

	errors := 0
	go func() {
		for {
			msg, err := s.ParseReader(r)
			if err == io.EOF {
				close(c)
				return
			}
			if err != nil {
				errors++
				if errors > 3 {
					panic(err)
				} else {
					log.Println(err)
				}
				continue
			} else {
				errors = 0
			}

			irecs := make([]InterpretedRecord, len(msg.DataRecords))
			for j, record := range msg.DataRecords {
				fmt.Println("...")
				// ifs = list of IntereptedFields
				ifs := i.Interpret(record)
				mfs := make([]myInterpretedField, len(ifs))
				pmsg := &protomsg.ZFlow{}
				//log.Println(ifs["zflowVerProductName"])
				log.Println("***start\n")
				for k, iif := range ifs {
					//log.Printf("case \"%s\":\n pmsg.%s = proto.String(iif.Value.(string))\n", iif.Name, iif.Name)
					//log.Println(iif.Name)
					//log.Printf("%s %s", iif.Name, iif.Value)
					switch iif.Name {
					// numeric fields
					case "protocolIdentifier":
						pmsg.ProtocolIdentifier = proto.Int32(int32(iif.Value.(uint8)))
					case "zflowPid":
						pmsg.ZflowPid = proto.Int32(int32(iif.Value.(uint32)))
					case "zflowParentPid":
						pmsg.ZflowParentPid = proto.Int32(int32(iif.Value.(uint32)))
					case "zflowCantHash":
						pmsg.ZflowCantHash = proto.Bool(iif.Value.(bool))
					case "zflowParentCantHash":
						pmsg.ZflowParentCantHash = proto.Bool(iif.Value.(bool))
					case "destinationTransportPort":
						pmsg.DestinationTransportPort = proto.Int32(int32(iif.Value.(uint16)))
					case "octetDeltaCount":
						pmsg.OctetDeltaCount = proto.Int64(int64(iif.Value.(uint64)))
					case "packetDeltaCount":
						pmsg.PacketDeltaCount = proto.Int64(int64(iif.Value.(uint64)))
					case "flowStartMilliseconds":
						pmsg.FlowStartMilliseconds = proto.Int64(int64(iif.Value.(uint64)))
					case "flowEndMilliseconds":
						pmsg.FlowEndMilliseconds = proto.Int64(int64(iif.Value.(uint64)))
					case "flowDirection":
						pmsg.FlowDirection = proto.Int64(int64(iif.Value.(uint8)))
					// string fields
					case "userName":
						pmsg.UserName = proto.String(iif.Value.(string))
					case "zflowAccount":
						pmsg.ZflowAccount = proto.String(iif.Value.(string))
					case "zflowImagePath":
						pmsg.ZflowImagePath = proto.String(iif.Value.(string))
					case "zflowCommandLine":
						pmsg.ZflowCommandLine = proto.String(iif.Value.(string))
					case "zflowImageBaseFileName":
						pmsg.ZflowImageBaseFileName = proto.String(iif.Value.(string))
					case "zflowMD5":
						pmsg.ZflowMD5 = proto.String(iif.Value.(string))
					case "zflowVerVersionString":
						pmsg.ZflowVerVersionString = proto.String(iif.Value.(string))
					case "zflowVerCompanyName":
						pmsg.ZflowVerCompanyName = proto.String(iif.Value.(string))
					case "zflowVerFileDescription":
						pmsg.ZflowVerFileDescription = proto.String(iif.Value.(string))
					case "zflowVerProductName":
						pmsg.ZflowVerProductName = proto.String(iif.Value.(string))
					case "zflowVerInternalName":
						pmsg.ZflowVerInternalName = proto.String(iif.Value.(string))
					case "zflowVerLegalCopyright":
						pmsg.ZflowVerLegalCopyright = proto.String(iif.Value.(string))
					case "zflowVerLegalTrademark":
						pmsg.ZflowVerLegalTrademark = proto.String(iif.Value.(string))
					case "zflowVerOriginalFilename":
						pmsg.ZflowVerOriginalFilename = proto.String(iif.Value.(string))
					case "zflowVerProductVersion":
						pmsg.ZflowVerProductVersion = proto.String(iif.Value.(string))
					case "flowVerFileTextVersion":
						pmsg.FlowVerFileTextVersion = proto.String(iif.Value.(string))
					case "zflowVerProductTextVersion":
						pmsg.ZflowVerProductTextVersion = proto.String(iif.Value.(string))
					case "zflowParentImagePath":
						pmsg.ZflowParentImagePath = proto.String(iif.Value.(string))
					case "zflowParentImageBaseFileName":
						pmsg.ZflowParentImageBaseFileName = proto.String(iif.Value.(string))
					case "zflowParentMD5":
						pmsg.ZflowParentMD5 = proto.String(iif.Value.(string))
					case "zflowParentVerVersionString":
						pmsg.ZflowParentVerVersionString = proto.String(iif.Value.(string))
					case "zflowParentVerCompanyName":
						pmsg.ZflowParentVerCompanyName = proto.String(iif.Value.(string))
					case "zflowParentVerFileDescription":
						pmsg.ZflowParentVerFileDescription = proto.String(iif.Value.(string))
					case "zflowParentVerProductName":
						pmsg.ZflowParentVerProductName = proto.String(iif.Value.(string))
					case "zflowParentVerInternalName":
						pmsg.ZflowParentVerInternalName = proto.String(iif.Value.(string))
					case "zflowParentVerLegalCopyright":
						pmsg.ZflowParentVerLegalCopyright = proto.String(iif.Value.(string))
					case "zflowParentVerLegalTrademark":
						pmsg.ZflowParentVerLegalTrademark = proto.String(iif.Value.(string))
					case "zflowParentVerOriginalFilename":
						pmsg.ZflowParentVerOriginalFilename = proto.String(iif.Value.(string))
					case "zflowParentVerProductVersion":
						pmsg.ZflowParentVerProductVersion = proto.String(iif.Value.(string))
					case "zflowParentVerFileTextVersion":
						pmsg.ZflowParentVerFileTextVersion = proto.String(iif.Value.(string))
					case "zflowParentVerProductTextVersion":
						pmsg.ZflowParentVerProductTextVersion = proto.String(iif.Value.(string))
					case "zflowAgentGuid":
						pmsg.ZflowAgentGuid = proto.String(iif.Value.(string))
					case "zflowMachineName":
						pmsg.ZflowMachineName = proto.String(iif.Value.(string))
					case "zflowOSName":
						pmsg.ZflowOSName = proto.String(iif.Value.(string))
					case "zflowOSVersion":
						pmsg.ZflowOSVersion = proto.String(iif.Value.(string))
					case "sourceIPv4Address":
						pmsg.SourceIPv4Address = iif.RawValue
					case "sourceTransportPort":
						pmsg.SourceTransportPort = proto.Int32(int32(iif.Value.(uint16)))
					case "destinationIPv4Address":
						pmsg.DestinationIPv4Address = iif.RawValue

					}

					mfs[k] = myInterpretedField{iif.Name, iif.EnterpriseID, iif.FieldID, iif.Value, integers(iif.RawValue)}
				}
				//log.Println("***end")

				mdata, _ := proto.Marshal(pmsg)
				data := stuff.Wrap("ZFlow", mdata)
				log.Println(data)
				fmt.Println(data)
				//stuff.Post(data)

				ir := InterpretedRecord{msg.Header.ExportTime, record.TemplateID, mfs}
				irecs[j] = ir
			}

			c <- irecs
		}
	}()

	return c
}

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", ":4739")
	ln, _ := net.ListenUDP("udp", serverAddr)
	defer ln.Close()

	r, w := io.Pipe()

	go func(writer io.Writer) {
		buf := make([]byte, 64000)
		for {
			n, _, _ := ln.ReadFromUDP(buf)
			fmt.Println(n)
			//fmt.Println("Received ",string(buf[0:n]))
			w.Write(buf[0:n])
		}
	}(w)

	//log.Println("ipfixcat", ipfixcatVersion)
	dictFile := flag.String("dict", "", "User dictionary file")
	messageStats := flag.Bool("mstats", false, "Log IPFIX message statistics")
	trafficStats := flag.Bool("acc", false, "Log traffic rates (Procera)")
	output := flag.Bool("output", true, "Display received flow records in JSON format")
	statsIntv := flag.Int("statsintv", 60, "Statistics log interval (s)")
	flag.Parse()

	s := ipfix.NewSession()
	i := ipfix.NewInterpreter(s)

	if *dictFile != "" {
		if err := stuff.LoadUserDictionary(*dictFile, i); err != nil {
			log.Fatal(err)
		}
	}

	msgs := messagesGenerator(r, s, i)
	tick := time.Tick(time.Duration(*statsIntv) * time.Second)
	//enc := json.NewEncoder(os.Stdout)
	for {
		select {
		case irecs, ok := <-msgs:
			if !ok {
				return
			}
			if *messageStats {
				//accountMsgStats(irecs)
			}

			for _, rec := range irecs {
				if *trafficStats {
					//accountTraffic(rec)
				}

				if *output {
					for i := range rec.Fields {
						f := &rec.Fields[i]
						switch v := f.Value.(type) {
						case []byte:
							f.RawValue = integers(v)
							f.Value = nil
						}
					}
					//enc.Encode(rec)
				}
			}
		case <-tick:
			if *messageStats {
				//logMsgStats()
			}

			if *trafficStats {
				//logAccountedTraffic()
			}
		}
	}
}

func integers(s []byte) []int {
	if s == nil {
		return nil
	}

	r := make([]int, len(s))
	for i := range s {
		r[i] = int(s[i])
	}
	return r
}
