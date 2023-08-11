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
		fmt.Printf("error getting key for chat generation: %v\n", err)
		return
	}
	client, err := vertexai.New(key)
	if err != nil {
		fmt.Printf("failed to create vertex ai client: %v\n", err)
		return
	}

	resp, err := client.ChatGeneration(context.Background(), vertexai.ChatGenerationRequest{
		ProjectID:      examples.ProjectID,
		EndpointID:     "chat-bison@001",
		Temperature:    &temperature,
		MaxDecodeSteps: &maxOutputTokens,
		TopK:           &topK,
		TopP:           &topP,
		Context:        "Translate to French",
		ChatMessages: []vertexai.ChatMessage{
			{
				Author:  "user",
				Content: "hello",
			},
		},
	})
	if err != nil {
		fmt.Printf("error in chat generation: %v\n", err)
		return
	}

	fmt.Printf("Generated Chat: %v\n", *resp)
	return
}
