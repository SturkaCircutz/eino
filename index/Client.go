package main

import (
	"context"
	"log"

	cli "github.com/milvus-io/milvus-sdk-go/v2/client"
)

var MilvusCli cli.Client

func InitClient() {

	ctx := context.Background()
	client, err := cli.NewClient(ctx, cli.Config{
		Address: "localhost:19530",
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	MilvusCli = client
}
