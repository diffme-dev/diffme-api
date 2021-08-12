package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

func NewKafkaClient() *kafka.Producer {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "host1:9092,host2:9092",
		"client.id":         "some_id",
		"acks":              "all",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	return p
}

func Producer(p *kafka.Producer) {
	deliveryChan := make(chan kafka.Event, 10000)
	topic := "topic"
	value := "hi"

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value)},
		deliveryChan,
	)

	if err != nil {
		// TODO:
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}
