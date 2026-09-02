package embed

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/openai"
)

func EmbedText(ctx context.Context) {
	embedder, err := openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
		APIKey:  os.Getenv("COHERE_API_KEY"),
		BaseURL: os.Getenv("COHERE_BASE_URL"),
		Model:   os.Getenv("COHERE_EMBED_MODEL"),
	})
	if err != nil {
		panic(err)
	}
	texts := []string{
		"my name is John",
		"im 13 years old",
	}
	embeddings, err := embedder.EmbedStrings(ctx, texts)
	if err != nil {
		panic(err)
	}
	for i, embedding := range embeddings {
		fmt.Printf("Text %d: %s\nEmbedding: %d dims, first 5: %v\n", i+1, texts[i], len(embedding), embedding[:5])
	}
}
