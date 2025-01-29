package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}
	apiKey := os.Getenv("API_KEY")
	fmt.Printf("API Key: %s\n", apiKey)
}
