package main

import (
	"github.com/dario-nihil/go_rest_api/db"
	"github.com/dario-nihil/go_rest_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
