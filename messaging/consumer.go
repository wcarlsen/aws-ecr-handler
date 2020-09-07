package messaging

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Initialise Kafka consumer
func KafkaConsumer() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("KAFKA_SERVER"),
		"group.id":           os.Getenv("KAFKA_CONSUMERGROUP"),
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	})

	if err != nil {
		panic(err)
	}

	topic := os.Getenv("KAFKA_TOPIC")

	err = c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	if err != nil {
		panic(err)
	}

	return c
}
