package messaging

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_SERVER")})

	if err != nil {
		panic(err)
	}

	return p
}
