package main

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/config"
	"shopping-cart/routes"
)

func main() {
	config.Connect() // connect to DB
	r := gin.Default()
	routes.SetupRoutes(r) // or routes.SetupRoutes(r) based on your function name
	r.Run(":8080") // run server on port 8080
}

