go run main.go

curl -X POST -H "Content-Type: application/json" -d '{"secret":"your-secret-key"}' http://localhost:8080/otp

curl -X POST -H "Content-Type: application/json" -d '{"secret":"your-secret-key", "code":"123456"}' http://localhost:8080/verify

Disclaimer: Please note that this project is for demonstrational and educational purposes only.
