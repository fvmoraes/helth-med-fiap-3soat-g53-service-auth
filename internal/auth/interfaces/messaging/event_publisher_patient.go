package messaging

import (
	"encoding/json"
	"helthmed-auth/internal/auth/domain"
	"helthmed-auth/internal/auth/infrastructure/rabbitmq"
	"log"

	"github.com/streadway/amqp"
)

func PublishPatientRegistered(patient domain.Patient) {
	body, err := json.Marshal(patient)
	if err != nil {
		log.Fatalf("failed to marshal patient: %v", err)
	}

	err = rabbitmq.Channel.Publish(
		"",
		"patient_registered",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("failed to publish message: %v", err)
	}
}
