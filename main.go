package main

import (
	"context"
	"log"

	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"my-agent/index"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()

	index.InitClient(ctx)
	index.IndexRAG(ctx, []*schema.Document{
		{ID: "1", Content: "My name is John and I am 13 years old."},
		{ID: "2", Content: "The capital of France is Paris."},
	})

	// flow.Chat(ctx)          // add "my-agent/flow" to the imports above
	// embed.EmbedText(ctx)    // add "my-agent/embed" to the imports above
}
