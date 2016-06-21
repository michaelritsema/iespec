package iespec

import (
	//"fmt"
	"github.com/michaelritsema/ipfix"
	"iespec/protomsg"
)

type Converter struct {
	s *ipfix.Session
	i *ipfix.Interpreter
}

func InitIpfix() (c *Converter) {

	s := ipfix.NewSession()
	i := ipfix.NewInterpreter(s)

	for _, entry := range MyFields {
		i.AddDictionaryEntry(entry)
	}
	c = &Converter{
		s: s,
		i: i,
	}
	return c
}

func (c Converter) Convert(b []byte) []*protomsg.ZFlow {

	msg, _ := c.s.ParseBuffer(b)
	//fmt.Printf("templates %v", msg.TemplateRecords)
	pmsgList := make([]*protomsg.ZFlow, 0)

	for _, rec := range msg.DataRecords {
		//fmt.Printf("rec: %v", rec)
		ifs := c.i.Interpret(rec)
		pmsg := ConvertFieldListToProtobuf(ifs)
		pmsgList = append(pmsgList, pmsg)
	}
	//fmt.Println(pmsgList)
	return pmsgList
}
