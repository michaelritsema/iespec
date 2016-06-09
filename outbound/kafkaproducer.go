package outbound

import (
	"encoding/base64"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"iespec/protomsg"
)

func encode(msg *protomsg.ZFlow) string {
	bytes, _ := proto.Marshal(msg)
	return "<pb type=\"ZFlow\">" + base64.StdEncoding.EncodeToString(bytes) + "</pb>"

}
func Kafka(c chan *protomsg.ZFlow) {
	brokers := []string{"54.84.148.16:9092"}
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("kafka producer init")
	go func() {
		for {
			msg := <-c
			xmlmsg := encode(msg)
			fmt.Println(xmlmsg)
			kafkaMsg := &sarama.ProducerMessage{
				Topic: "ZIFTEN.DATACOLLECTION_",
				Value: sarama.StringEncoder(xmlmsg),
			}
			fmt.Println(kafkaMsg)
			producer.SendMessage(kafkaMsg)

		}
	}()
}
