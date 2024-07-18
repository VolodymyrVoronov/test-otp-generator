package routes

import (
	"test-otp-generator/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/otp", handlers.OtpHandler)
	server.POST("/verify", handlers.VerifyOtpHandler)
}
