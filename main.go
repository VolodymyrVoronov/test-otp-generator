package main

import (
	"test-otp-generator/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
