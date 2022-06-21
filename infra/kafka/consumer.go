package kafka

import(
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
) 

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

func(k *KafkaConsumer) Consume() {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupID"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("Error consuming kafka message: " + err.Error())
	}
	topics := []string{os.Getenv("KafkaReadTopic")}
	fmt.Println("Kafka consumer has been starter")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}