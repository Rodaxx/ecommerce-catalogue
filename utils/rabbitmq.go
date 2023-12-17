package utils

import (
	//"encoding/json"
	//"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func sendResponse(ch *amqp.Channel, request amqp.Delivery, response string) error{
	err := ch.Publish(
		"",          // Exchange
		request.ReplyTo, 	 // Routing key (cola de respuesta)
		false,       // Mandatory
		false,       // Immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: request.CorrelationId,
			Body:          []byte(response),
		},
	)
	if err != nil {
		log.Printf("Error al publicar la respuesta: %s", err)
		return err
	}

	return nil
}