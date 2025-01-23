package pkg

import (
	"log"
	"os"
)

var logger *log.Logger

// InitLogger inicializa o logger e cria o arquivo de log
func InitLogger() {
	file, err := os.OpenFile("logs/requests.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	logger = log.New(file, "LOG: ", log.LstdFlags)
}

// LogRequest loga detalhes de uma requisição
func LogRequest(method, url string) {
	logger.Printf("Request: %s %s\n", method, url)
}
