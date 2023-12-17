package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	//"context"
	//"math/rand"
	//"time"

	// External
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"

	// Internal
	"github.com/Rodaxx/ecommerce-catalogue/mod/catalogue"
	"github.com/Rodaxx/ecommerce-catalogue/mod/utils"
)

type MessagePattern struct {
	Pattern struct {
		Cmd string `json:"cmd"`
	} `json:"pattern"`
	Data interface{} `json:"data"`
	ID   string      `json:"id"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	rabbitMQUser := os.Getenv("RABBITMQ_USERNAME")
	rabbitMQPass := os.Getenv("RABBITMQ_PASSWORD")
	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")

	fmt.Println(fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitMQUser, rabbitMQPass, rabbitMQHost, rabbitMQPort))
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitMQUser, rabbitMQPass, rabbitMQHost, rabbitMQPort))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"catalogue", // Nombre de la cola
		false,       // Durable
		false,       // Eliminar cuando no haya consumidores
		false,       // Exclusiva
		false,       // No esperar confirmación de servidor
		nil,         // Argumentos adicionales
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name, // Nombre de la cola
		"",     // Nombre del consumidor
		true,   // Auto-ack (auto acknowledge)
		false,  // Exclusiva
		false,  // No-local
		false,  // No-wait
		nil,    // Argumentos adicionales
	)
	if err != nil {
		log.Fatal(err)
	}

	// Procesar mensajes
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			// Convertir el cuerpo del mensaje a la estructura Mensaje
			var message MessagePattern
			if err := json.Unmarshal(d.Body, &message); err != nil {
				log.Printf("Error al decodificar el mensaje: %v", err)
				continue
			}

			switch message.Pattern.Cmd {
			case "createTransaction":
				//resultado := funcion1(mensaje.Datos)

				fmt.Printf("Mensaje Recibido:\n")
				fmt.Printf(" - Body: %s\n", d.Body)
				fmt.Printf(" - Routing Key: %s\n", d.RoutingKey)
				fmt.Printf(" - Delivery Tag: %d\n", d.DeliveryTag)
				fmt.Printf(" - Correlation ID: %s\n", d.CorrelationId)
				fmt.Printf(" - Reply To: %s\n", d.ReplyTo)

				responseMessage := map[string]interface{}{
					"status":  "success",
					"message": "Saludos desde go",
					// Puedes agregar más campos según tus necesidades
				}

				rabbitmq.sendResponse(ch, d, catalogue.findAll())
				catalogue.FindAll()
				/**
				// Convierte la respuesta a formato JSON
				jsonResponse, err := json.Marshal(responseMessage)
				if err != nil {
					// Maneja el error según tus necesidades
					log.Printf("Error al convertir la respuesta a JSON: %s", err)
					return
				}

				err = ch.Publish(
					"",          // Exchange
					d.ReplyTo, 	 // Routing key (cola de respuesta)
					false,       // Mandatory
					false,       // Immediate
					amqp.Publishing{
						ContentType:   "application/json",
						CorrelationId: d.CorrelationId,
						Body:          jsonResponse,
					},
				)

				fmt.Println(message)
				**/

				//fmt.Println(catalogue.FindAll())
				/*
					case "funcion2":
						resultado := funcion2(mensaje.Datos)
						fmt.Printf("Resultado de funcion2: %s\n", resultado)
				*/
			}
		}
	}()

	fmt.Println("Microservicio escuchando. Presiona CTRL+C para salir.")
	<-forever
}
