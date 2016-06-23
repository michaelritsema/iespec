package outbound

import (
	"encoding/base64"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"iespec/protomsg"
)

func encode(msg *protomsg.ZFlow) string {
	bytes, _ := proto.Marshal(msg)
	return "<pb type=\"ZFlow\">" + base64.StdEncoding.EncodeToString(bytes) + "</pb>"

}
func Kafka(c chan *protomsg.ZFlow, connection string) {
	brokers := []string{connection}
	producer, _ := sarama.NewSyncProducer(brokers, nil)

	go func() {
		for {
			msg := <-c
			xmlmsg := encode(msg)
			kafkaMsg := &sarama.ProducerMessage{
				Topic: "ZIFTEN.DATACOLLECTION_",
				Value: sarama.StringEncoder(xmlmsg),
			}
			producer.SendMessage(kafkaMsg)
		}
	}()
}
