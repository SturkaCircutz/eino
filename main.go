package main

import (
	"context"
	"log"

	"my-agent/index"
	"my-agent/retriever"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()

	index.InitClient(ctx)
	index.IndexRAG(ctx, index.Docs)
	retriever.Retrieve(ctx, "yuanshen")
	// flow.Chat(ctx)          // add "my-agent/flow" to the imports above
	// embed.EmbedText(ctx)    // add "my-agent/embed" to the imports above
}
