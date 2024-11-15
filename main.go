package main

import (
	"github.com/bsorawit1234/expense-tracker-backend/config"
	"github.com/bsorawit1234/expense-tracker-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	routes.InitializeRoutes(router)

	router.Run(":8080")
}
