package vertexai

import (
	"context"
	"net/http"
)

// Client is the interface for the Vertex AI client for Bard LLM
type Client interface {
	TextGeneration(ctx context.Context, req TextGenerationRequest) (*TextGenerationResponse, error)
	ChatGeneration(ctx context.Context, req ChatGenerationRequest) (*ChatGenerationResponse, error)
	TextEmbedding(ctx context.Context, req TextEmbeddingRequest) (*TextEmbeddingResponse, error)
}

type client struct {
	clientConfig   clientConfig
	requestBuilder requestBuilder
	httpClient     httpClient
	tokenizer      tokenizer
}

// New creates a new client for the Vertex AI client for Bard LLM
// key is the oauth key of the Google Service account for authentication
func New(key string) (Client, error) {
	return &client{
		clientConfig:   newClientConfig(setAuthToken(key)),
		requestBuilder: newRequestBuilder(),
		httpClient:     http.DefaultClient,
		tokenizer:      &defaultTokenizer{},
	}, nil
}
