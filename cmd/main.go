package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type RequestBody struct {
	Prompt string `json:"prompt"`
}

type ResponseBody struct {
	Output string `json:"output"`
}

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}
	apiKey := os.Getenv("API_KEY")
	fmt.Printf("API Key: %s\n", apiKey)

	client := resty.New()

	requestBody := RequestBody{
		Prompt: "Привет, как дела?",
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+apiKey).
		SetBody(requestBody).
		Post("https://dashscope-intl.aliyuncs.com/compatible-mode/v1")

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	var responseBody ResponseBody
	body := resp.Body()
	err = json.Unmarshal(body, &responseBody) //
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}

	fmt.Println("Response:", responseBody.Output)

}
