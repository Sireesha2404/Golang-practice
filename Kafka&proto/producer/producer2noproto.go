package main

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	data := "implementing kafka"
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
	})
	if err != nil {
		panic(err)
	}

	defer p.Close()
	topic := "my-kafka-topic"

	// Message := "Take up one idea. Make that one idea your life--think of it, dream of it, live on that idea. Let the brain, muscles, nerves, every part of your body. Testing."
	// data := []byte(Message)
	deliveryChan := make(chan kafka.Event)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}, deliveryChan)

	select {
	case ev := <-deliveryChan:
		switch e := ev.(type) {
		case *kafka.Message:
			if e.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", e.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", e.TopicPartition)
				fmt.Printf("Delivered message to %v\n", string(e.Value))
			}
			close(deliveryChan)
		}
	}
}
