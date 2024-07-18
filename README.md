**To run the server, execute:**

go run main.go

**Generate OTP QR Code:**

curl -X POST -H "Content-Type: application/json" -d '{"secret":"your-secret-key"}' http://localhost:8080/otp

**Verify OTP Code:**

curl -X POST -H "Content-Type: application/json" -d '{"secret":"your-secret-key", "code":"123456"}' http://localhost:8080/verify

Disclaimer: Please note that this project is for demonstrational and educational purposes only.
