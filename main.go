package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/openai"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// ctx := context.Background()
	// model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
	// 	APIKey:  os.Getenv("COHERE_API_KEY"),
	// 	BaseURL: os.Getenv("COHERE_BASE_URL"),
	// 	Model:   os.Getenv("COHERE_MODEL"),
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// //template for model
	// template := prompt.FromMessages(schema.FString,
	// 	schema.SystemMessage("You are a professor in {role}"),
	// 	&schema.Message{
	// 		Role:    schema.User,
	// 		Content: "What is the difference between {topic1} and {topic2}?",
	// 	},
	// )
	// params := map[string]any{
	// 	"role":   "computer science",
	// 	"topic1": "machine learning",
	// 	"topic2": "deep learning",
	// }
	// messages, err := template.Format(ctx, params)
	// if err != nil {
	// 	panic(err)
	// }

	// // input messages for the chat model
	// // input := []*schema.Message{
	// // 	schema.SystemMessage("You are a helpful assistant."),
	// // 	schema.UserMessage("How are you?"),
	// // }
	// response, err := model.Generate(ctx, messages)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(response.Content)
	ctx := context.Background()
	embedd, err := openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
		APIKey:  os.Getenv("COHERE_API_KEY"),
		BaseURL: os.Getenv("COHERE_BASE_URL"),
		Model:   os.Getenv("COHERE_EMBED_MODEL"),
	})
	if err != nil {
		panic(err)
	}
	text := []string{
		"Hello, world!",
		"How are you?",
		"This is a test.",
	}
	embedding, err := embedd.EmbedStrings(ctx, text)
	if err != nil {
		panic(err)
	}
	for i, e := range embedding {
		fmt.Printf("Text: %s\nEmbedding: %v\n", text[i], e)
	}
}
