package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
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
	url := "https://dashscope-intl.aliyuncs.com/compatible-mode/v1"
	baseUrl := option.WithBaseURL(url)

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		baseUrl,
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Say this is a test"),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err.Error())
	}
	println(chatCompletion.Choices[0].Message.Content)
}
