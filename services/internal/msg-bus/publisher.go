package message_bus

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaPublisher struct {
	producer *kafka.Producer
}

type Publisher interface {
	PublishMessage(topic string, message any) error
	Close()
}

func InitPublisher() *KafkaPublisher {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}
	return &KafkaPublisher{producer: p}
}

func (kp *KafkaPublisher) PublishMessage(topic string, message any) error {
	go func() {
		for e := range kp.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	js, err := json.Marshal(message)
	if err != nil {
		return err
	}

	kp.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          js,
		},
		nil,
	)

	return nil
}

func (kp *KafkaPublisher) Close()  {
	kp.producer.Close()	
}
