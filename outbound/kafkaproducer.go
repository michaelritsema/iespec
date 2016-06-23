package outbound

import (
	"encoding/base64"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"iespec/protomsg"
	"log"
	"time"
)

func encode(msg *protomsg.ZFlow) string {
	bytes, _ := proto.Marshal(msg)
	return "<pb type=\"ZFlow\">" + base64.StdEncoding.EncodeToString(bytes) + "</pb>"

}
func Kafka(c chan *protomsg.ZFlow, connection string) {
	brokers := []string{connection}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal        // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy    // Compress messages
	config.Producer.Flush.Frequency = 5000 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Printf("Error when trying to connect to Kafka: %s", err)
	}
	go func() {
		for {
			msg := <-c
			log.Println("Kafka Producer received message.")
			xmlmsg := encode(msg)
			log.Println("Kafka Producer converted message to <pb> format")
			kafkaMsg := &sarama.ProducerMessage{
				Topic: "ZIFTEN.DATACOLLECTION_",
				Value: sarama.StringEncoder(xmlmsg),
			}

			producer.Input() <- kafkaMsg

			log.Println("Kafka Producer sent message.")
		}
	}()
}
