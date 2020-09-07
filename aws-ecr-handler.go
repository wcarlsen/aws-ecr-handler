package main

import (
	"aws-ecr-handler/handlers"
	"aws-ecr-handler/messaging"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c := messaging.KafkaConsumer()
	pro := messaging.KafkaProducer()

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			err = handlers.EventHandler(msg)

			if err != nil {
				// Produce terminated event
				topic := "produce"

				value := "terminated"
				pro.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(value)}, nil)
				c.Commit()
			} else {
				// Produce success event
				topic := "produce"

				value := "success"
				pro.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(value)}, nil)
				c.Commit()
			}
		} else {
			// Error in consuming message
			c.Commit()
		}

	}
}
