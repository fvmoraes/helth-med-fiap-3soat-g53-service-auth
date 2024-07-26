package messaging

import (
	"encoding/json"
	"helthmed-auth/internal/auth/domain"
	"helthmed-auth/internal/auth/infrastructure/rabbitmq"
	"log"

	"github.com/streadway/amqp"
)

func PublishDoctorRegistered(doctor *domain.Doctor) {
	body, err := json.Marshal(doctor)
	if err != nil {
		log.Fatalf("failed to marshal doctor: %v", err)
	}

	err = rabbitmq.Channel.Publish(
		"",
		"doctor_registered",
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
