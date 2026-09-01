package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  os.Getenv("COHERE_API_KEY"),
		BaseURL: os.Getenv("COHERE_BASE_URL"),
		Model:   os.Getenv("COHERE_MODEL"),
	})
	if err != nil {
		panic(err)
	}
	// input messages for the chat model
	input := []*schema.Message{
		schema.SystemMessage("You are a helpful assistant."),
		schema.UserMessage("How are you?"),
	}
	response, err := model.Generate(ctx, input)
	if err != nil {
		panic(err)
	}
	println(response.Content)
}
