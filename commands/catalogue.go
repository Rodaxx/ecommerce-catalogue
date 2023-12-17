package commands

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

func FindAll() string {
	// Crear una lista de productos de ejemplo
	products := []Product{
		{ID: 1, Name: "Producto1", Price: 19990.0},
		{ID: 2, Name: "Producto2", Price: 29990.0},
		{ID: 3, Name: "Producto3", Price: 39990.0},
	}

	// Convertir la lista de productos a formato JSON
	jsonData, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error al convertir productos a JSON:", err)
		return ""
	}

	return string(jsonData)
}