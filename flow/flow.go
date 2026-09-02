package flow

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/schema"
)

func Chat(ctx context.Context) {
	model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: os.Getenv("COHERE_BASE_URL"),
		APIKey:  os.Getenv("COHERE_API_KEY"),
		Model:   os.Getenv("COHERE_MODEL"),
	})
	if err != nil {
		panic(err)
	}
	messages := []*schema.Message{
		schema.SystemMessage("you are a good human helper"),
		schema.UserMessage("which year is it"),
	}

	reader, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)

	}
	println(reader.Content)
}
