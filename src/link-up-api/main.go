package main

import (
	"linkupapi/config"
	"linkupapi/controllers"
	"linkupapi/routes"
	"log"
	"net/http"
)

func main() {
	// Conectando ao MongoDB
	config.ConnectDatabase("mongodb://localhost:27017")

	availabilityController := controllers.InitAvailabilityController()

	// Registrando rotas
	routes.RegisterRoutes(availabilityController)

	// Iniciando o servidor
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
