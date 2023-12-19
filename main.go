package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	//"context"
	//"math/rand"
	//"time"

	// External

	"github.com/joho/godotenv"
	"github.com/rodaxx/ecommerce-catalogue/config"
	"github.com/rodaxx/ecommerce-catalogue/database"
	"github.com/rodaxx/ecommerce-catalogue/server"
	"github.com/rodaxx/ecommerce-catalogue/utils"
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
		 log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	DATABASE_URL := os.Getenv("DATABASE_URL")
	
	repo,err:= database.NewMySQLRepository(DATABASE_URL)
	
	if err!= nil {
		log.Fatal(err)
	}

	server:= server.NewCatalogueServer(repo)

	if err!= nil {
		log.Fatal(err)
	}
	
	log.Println("Starting server on port", PORT)
	config.SetupRabbitMQ()

	// Procesar mensajes
	forever := make(chan bool)
	//consumir los mensajes de la cola
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
			case "FindAllProducts":
				products,err:=server.GetProducts(context.Background())
				if err != nil {
					fmt.Println("Error al obtener productos:", err)
					return
				}
				resultJSON, err := json.MarshalIndent(products, "", "  ")
				if err != nil {
					fmt.Println("Error al convertir a JSON:", err)
					return
				}
				utils.SendResponse(config.RMQChannel, msg, string(resultJSON));
			case "GetProductById":
				productID, ok := message.Data.(map[string]interface{})["productID"].(float64)
				if !ok {
					fmt.Println("Error: productID no encontrado o no es un número")
					return
				}
				// Obtener el producto por ID
				product, err := server.GetProductById(context.Background(),uint(productID))
				if err != nil {
					fmt.Println("Error al obtener el producto por ID:", err)
					return
				}
				resultJSON, err := json.MarshalIndent(product, "", "  ")
				if err != nil {
					fmt.Println("Error al convertir a JSON:", err)
					return
				}
				utils.SendResponse(config.RMQChannel, msg, string(resultJSON));
			case "Purchase":
				productID, ok := message.Data.(map[string]interface{})["productID"].(float64)
				if !ok {
					fmt.Println("Error: productID no encontrado o no es un número")
					return
				}
				newQuantity, ok := message.Data.(map[string]interface{})["newQuantity"].(float64)
				if !ok {
					fmt.Println("Error: newQuantity no encontrado o no es un número")
					return
				}
				err=server.UpdateProductById(context.Background(),uint(productID),uint(newQuantity))
				if err != nil{
					log.Fatal(err)
				}
				product, err := server.GetProductById(context.Background(),uint(productID))
				if err != nil {
					fmt.Println("Error al obtener el producto por ID:", err)
					return
				}
				resultJSON, err := json.MarshalIndent(product, "", "  ")
				if err != nil {
					fmt.Println("Error al convertir a JSON:", err)
					return
				}
				utils.SendResponse(config.RMQChannel, msg, string(resultJSON));
			case "FindProductsForBranch":
			}
		}
	}()
	fmt.Println("Microservicio escuchando. Presiona CTRL+C para salir.")
	<-forever
}
