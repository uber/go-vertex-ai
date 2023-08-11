package main

import (
	"context"
	"fmt"

	"vertexai"
	"vertexai/examples"
)

func main() {
	key, err := examples.ReadKey()
	if err != nil {
		fmt.Printf("error getting key for text embedding: %v\n", err)
		return
	}
	client, err := vertexai.New(key)
	if err != nil {
		fmt.Printf("failed to create vertex ai client: %v\n", err)
		return
	}

	resp, err := client.TextEmbedding(context.Background(), vertexai.TextEmbeddingRequest{
		ProjectID:  examples.ProjectID,
		EndpointID: "textembedding-gecko",
		Inputs:     []string{"Hello Google! What can you do?"},
	})
	if err != nil {
		fmt.Printf("error in text embedding: %v\n", err)
		return
	}

	fmt.Printf("Text Embedding: %v\n", *resp)
	return
}
