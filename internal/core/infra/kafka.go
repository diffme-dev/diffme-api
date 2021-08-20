package infra

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	"log"
)

var kafkaProducerConfig = &kafka.ConfigMap{
	"bootstrap.servers": "pkc-4nym6.us-east-1.aws.confluent.cloud:9092",
	"security.protocol": "SASL_SSL",
	"sasl.mechanisms":   "PLAIN",
	"sasl.username":     "HJI2ZYBNYYXF673A",
	"sasl.password":     "22+vy5+GeNGx0CU2IEiMy2s/tMWvQ1gBOPdBF1VJY90je0i9SfE8ACp+bupZ4EV7",
}

var kafkaConsumerConfig = &kafka.ConfigMap{
	"bootstrap.servers": "pkc-4nym6.us-east-1.aws.confluent.cloud:9092",
	"security.protocol": "SASL_SSL",
	"sasl.mechanisms":   "PLAIN",
	"sasl.username":     "HJI2ZYBNYYXF673A",
	"sasl.password":     "22+vy5+GeNGx0CU2IEiMy2s/tMWvQ1gBOPdBF1VJY90je0i9SfE8ACp+bupZ4EV7",
	"group.id":          "myGroup",
	"auto.offset.reset": "earliest",
}

type OnConsume = func(message proto.Message) error

type KafkaClient struct {
	TopicName string          `json:"topic"`
	Producer  *kafka.Producer `json:"producer"`
	Consumer  *kafka.Consumer `json:"consumer"`
	OnConsume OnConsume       `json:"on_consume"`
}

func NewKafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(kafkaProducerConfig)

	if err != nil {
		panic(err)
	}

	return p
}

func NewKafkaConsumer() *kafka.Consumer {
	c, err := kafka.NewConsumer(kafkaConsumerConfig)

	if err != nil {
		panic(err)
	}

	return c
}

func NewKafkaClient(
	topicName string,
	onConsume OnConsume,
	producer *kafka.Producer,
	consumer *kafka.Consumer,
) *KafkaClient {
	return &KafkaClient{
		TopicName: topicName,
		Producer:  producer,
		OnConsume: onConsume,
		Consumer:  consumer,
	}
}

func (c *KafkaClient) produce(message proto.Message) {

	data, err := proto.Marshal(message)

	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	err = c.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &c.TopicName, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}, nil)
}

func (c *KafkaClient) consume() {

	c.Consumer.SubscribeTopics([]string{c.TopicName}, nil)

	for {
		msg, err := c.Consumer.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)

			var payload *proto.Message

			err := proto.Unmarshal(msg.Value, *payload)

			if err != nil {
				println(err)
			}

			c.OnConsume(*payload)
		}
	}

	defer c.Consumer.Close()
}
