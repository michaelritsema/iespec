package inbound

import (
	//"encoding/base64"
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	//"github.com/golang/protobuf/proto"
	//"iespec"
	//"iespec/protomsg"
	//"io"
	//	"time"
)

/*
	takes in <pb> </pb>
	returns array of ipfix messages as [][]byte
*/
/*
func stripXML(msg []byte) [][]byte {
	payload := msg[66 : len(msg)-5]
	bytes, _ := base64.StdEncoding.DecodeString(string(payload))
	pmsg := &protomsg.IpfixEncapsulation{}
	proto.Unmarshal(bytes, pmsg)

	return pmsg.GetIpfixMessage()
}
*/
func Kafka() {
	//converter := iespec.InitIpfix()

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

	msgchan := make(chan *sarama.ConsumerMessage, 256)

	for x, c := range consumers {
		fmt.Printf("Init consumer partition %d\n", x)

		go func(pc sarama.PartitionConsumer) {
			msgCounter := 0

			for {
				message := <-pc.Messages()
				msgchan <- message

				msgCounter++

				fmt.Println(msgCounter)
			}
		}(c)

		go func() {
			for msg := range msgchan {
				//for _, ipfixList := range stripXML(msg.Value) {
				//for _, zflowMsg := range converter.Convert(ipfixList) {
				//fmt.Println(zflowMsg)
				//protoMsgChan <- zflowMsg
				fmt.Printf("Partition:\t%d\n", msg.Partition)
				fmt.Printf("Offset:\t%d\n", msg.Offset)
				fmt.Printf("Key:\t%s\n", string(msg.Key))
				fmt.Printf("Value:\t%s\n", string(msg.Value))
				fmt.Println()

				//}
				//}
			}
		}()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
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
