package main

import (
	"encoding/json"
	"fmt"
	"log"

	//"context"
	//"math/rand"
	//"time"

	// External

	"github.com/rodaxx/ecommerce-catalogue/commands"
	"github.com/rodaxx/ecommerce-catalogue/utils"
	"github.com/rodaxx/ecommerce-catalogue/config"
)

type MessagePattern struct {
	Pattern struct {
		Cmd string `json:"cmd"`
	} `json:"pattern"`
	Data interface{} `json:"data"`
	ID   string      `json:"id"`
}

func main() {
	config.SetupRabbitMQ()

	// Procesar mensajes
	forever := make(chan bool)
	
	messages, err := config.RMQChannel.Consume(
		config.RMQQueue.Name, // Nombre de la cola
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

	go func() {
		for msg := range messages {
			// Convertir el cuerpo del mensaje a la estructura Mensaje
			var message MessagePattern
			if err := json.Unmarshal(msg.Body, &message); err != nil {
				log.Printf("Error al decodificar el mensaje: %v", err)
				continue
			}

			switch message.Pattern.Cmd {
			case "createTransaction":
				fmt.Println("Nueva solicitud RPC recibida: findAllProducts")
				utils.SendResponse(config.RMQChannel, msg, commands.FindAllProducts());
			}
		}
	}()

	fmt.Println("Microservicio escuchando. Presiona CTRL+C para salir.")
	<-forever
}
