package main

import (
	"log"

	"github.com/gin-gonic/gin"

	config "github.com/Witaly3/golang-rest/config"
	routes "github.com/Witaly3/golang-rest/routes"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run("0.0.0.0:8080"))
}
