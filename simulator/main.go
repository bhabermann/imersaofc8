package main

import (
	"fmt"
	"log"

	kafka2 "github.com/bhabermann/imersaofc8-simulator/app/kafka"
	"github.com/bhabermann/imersaofc8-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// {"clientId":"1","routeId":"1"}
func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("Hello World", "route.new-direction", producer)

	// for {
	// 	_ = 1
	// }

	// route := myroute.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// route.LoadPosition()
	// json, _ := route.ExportJsonPositions()
	// fmt.Println(json[1])
}
