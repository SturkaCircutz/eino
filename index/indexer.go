package index

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/openai"
	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/cloudwego/eino/schema"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var collection = "test"

var fields = []*entity.Field{
	{
		Name:     "id",
		DataType: entity.FieldTypeVarChar,
		TypeParams: map[string]string{
			"max_length": "255",
		},
		PrimaryKey: true,
	},
	{
		Name:     "vector", // 确保字段名匹配
		DataType: entity.FieldTypeBinaryVector,
		TypeParams: map[string]string{
			"dim": "32768",
		},
	},
	{
		Name:     "content",
		DataType: entity.FieldTypeVarChar,
		TypeParams: map[string]string{
			"max_length": "8192",
		},
	},
	{
		Name:     "metadata",
		DataType: entity.FieldTypeJSON,
	},
}

/*
schema.Document is a container for a piece of text plus info about that text

	doc := &schema.Document{
	    ID:      "1",
	    Content: "My name is John and I am 13 years old.",
	}
*/
func IndexRAG(ctx context.Context, docs []*schema.Document) {
	embedder, err := openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
		APIKey:  os.Getenv("COHERE_API_KEY"),
		Model:   os.Getenv("COHERE_EMBED_MODEL"),
		BaseURL: os.Getenv("COHERE_BASE_URL"),
	})
	if err != nil {
		panic(err)
	}

	indexer, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:     MilvusCli,
		Collection: collection,
		Fields:     fields,
		Embedding:  embedder,
	})
	if err != nil {
		panic(err)
	}

	ids, err := indexer.Store(ctx, docs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("stored documents with IDs: %v\n", ids)
}
