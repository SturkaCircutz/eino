package retriever

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/openai"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
	"github.com/cloudwego/eino/schema"

	"my-agent/index"
)

func Retrieve(ctx context.Context, query string) []*schema.Document {
	embedder, err := openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
		APIKey:  os.Getenv("COHERE_API_KEY"),
		Model:   os.Getenv("COHERE_EMBED_MODEL"),
		BaseURL: os.Getenv("COHERE_BASE_URL"),
	})
	if err != nil {
		panic(err)
	}

	r, err := milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:      index.MilvusCli,
		Collection:  "test",
		VectorField: "vector",
		OutputFields: []string{
			"id",
			"content",
			"metadata",
		},
		TopK:      2,
		Embedding: embedder,
	})
	if err != nil {
		panic(err)
	}

	docs, err := r.Retrieve(ctx, query)
	if err != nil {
		panic(err)
	}
	for _, doc := range docs {
		fmt.Printf("%s: %s\n", doc.ID, doc.Content)
	}
	return docs
}
