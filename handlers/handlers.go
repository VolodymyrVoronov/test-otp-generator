package handlers

import (
	"encoding/base32"
	"net/http"
	"test-otp-generator/models"
	"test-otp-generator/utils"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

func OtpHandler(c *gin.Context) {
	var requestBody models.RequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if requestBody.Secret == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing secret key"})
		return
	}

	secret := base32.StdEncoding.EncodeToString([]byte(requestBody.Secret))

	otpCode, err := utils.GenerateOTP(secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating OTP"})
		return
	}

	qrCode, err := utils.GenerateQRCode(otpCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating QR code"})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Writer.Write(qrCode)
}

func VerifyOtpHandler(c *gin.Context) {
	var verifyRequestBody models.VerifyRequestBody
	if err := c.ShouldBindJSON(&verifyRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if verifyRequestBody.Secret == "" || verifyRequestBody.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing secret key or OTP code"})
		return
	}

	secret := base32.StdEncoding.EncodeToString([]byte(verifyRequestBody.Secret))

	valid := totp.Validate(verifyRequestBody.Code, secret)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP code", "valid": valid})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP code is valid", "valid": valid})
}
