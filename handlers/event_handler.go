package handlers

import (
	"aws-ecr-handler/aws"
	"aws-ecr-handler/events"
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func EventHandler(msg *kafka.Message) error {
	var ge events.GeneralEvent
	err := json.Unmarshal(msg.Value, &ge)

	if err != nil {
		//
	}

	switch ge.Type {

	case "aws_ecr_added":
		// se := ge.Payload.(events.ECRAddedPayload)
		var se events.ECRAddedEvent
		err := json.Unmarshal(msg.Value, &se)

		_, err = aws.CreateECR(se.Payload.Properties.RepoName)

		if err != nil {
			if err.Error() == ecr.ErrCodeImageAlreadyExistsException {
				err = nil
			}
			return err
		}

		_, err = aws.SetECRPolicy(se.Payload.Properties.RepoName, se.Payload.Properties.AccountId)

		if err != nil {
			return err
		}

	case "aws_ecr_deleted":
		se := ge.Payload.(events.ECRDeletedEvent)
		_, err := aws.DeleteECR(se.Payload.Properties.RepoName, se.Payload.Properties.AccountId)

		if err != nil {
			return err
		}

	default:
		// Log warning about unscoped event
	}

	return err
}
