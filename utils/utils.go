package utils

import (
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func GenerateOTP(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now())
}

func GenerateQRCode(otpCode string) ([]byte, error) {
	png, err := qrcode.Encode(otpCode, qrcode.Highest, 256)
	if err != nil {
		return nil, err
	}

	return png, nil
}
