package main

import (
	"context"
	"fmt"

	"vertexai"
	"vertexai/examples"
)

var (
	maxOutputTokens int     = 256
	temperature     float64 = 0.2
	topK            int     = 40
	topP            float64 = 0.8
)

func main() {
	key, err := examples.ReadKey()
	if err != nil {
		fmt.Printf("error getting key for text generation: %v\n", err)
		return
	}
	client, err := vertexai.New(key)
	if err != nil {
		fmt.Printf("failed to create vertex ai client: %v\n", err)
		return
	}

	resp, err := client.TextGeneration(context.Background(), vertexai.TextGenerationRequest{
		ProjectID:      examples.ProjectID,
		EndpointID:     "text-bison@001",
		Content:        "Hello Google! What can you do?",
		Temperature:    &temperature,
		MaxDecodeSteps: &maxOutputTokens,
		TopK:           &topK,
		TopP:           &topP,
	})
	if err != nil {
		fmt.Printf("error in text generation: %v\n", err)
		return
	}

	fmt.Printf("Generated Text: %v\n", *resp)
	return
}
