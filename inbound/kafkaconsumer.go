package inbound

import (
	"encoding/base64"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"iespec"
	"iespec/protomsg"
	//"io"
	"log"
	"os"
)

/*
	takes in <pb> </pb>
	returns array of ipfix messages as [][]byte
*/
func stripXML(msg []byte) [][]byte {
	payload := msg[66 : len(msg)-5]
	bytes, _ := base64.StdEncoding.DecodeString(string(payload))
	pmsg := &protomsg.IpfixEncapsulation{}
	proto.Unmarshal(bytes, pmsg)

	return pmsg.GetIpfixMessage()
}
func Kafka(protoMsgChan chan *protomsg.ZFlow) {
	converter := iespec.InitIpfix()

	log.SetOutput(os.Stdout)
	//config := sarama.NewConfig()

	//config.Consumer.Return.Errors = true

	// Specify brokers address. This is default one
	brokers := []string{"54.84.148.16:9092"}
	//ec2-54-210-70-189.compute-1.amazonaws.com:9092
	// Create new coeeensumer
	master, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	topic := "ZIFTEN.IPFIX"
	// How to decide partition, is it fixed value...?
	partitions, _ := master.Partitions(topic)
	fmt.Println(partitions)
	consumers := make([]sarama.PartitionConsumer, 0)
	for _, p := range partitions {
		consumer, err := master.ConsumePartition(topic, p, sarama.OffsetOldest)
		consumers = append(consumers, consumer)
		if err != nil {
			panic(err)
		}
	}

	go func() {
		for {
			for _, c := range consumers {
				select {
				case msg := <-c.Messages():
					for _, ipfixList := range stripXML(msg.Value) {
						for _, zflowMsg := range converter.Convert(ipfixList) {
							//fmt.Println(zflowMsg)
							protoMsgChan <- zflowMsg

						}

					}
				}
			}
		}
	}()

}

/*
func main() {
	c := make(chan *protomsg.ZFlow)
	go kafka(c)
	for {
		x := <-c
		fmt.Println(x)
	}
}

*/
