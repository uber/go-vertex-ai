package vertexai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// TextEmbeddingRequest is the request for the TextEmbedding method
type TextEmbeddingRequest struct {
	ProjectID  string
	EndpointID string
	/*
		Content is Deprecated. Use Inputs instead
	*/
	Content string
	Inputs  []string
}

// TextEmbeddingResponse is the response for the TextEmbedding method
type TextEmbeddingResponse struct {
	Predictions []Predictions `json:"predictions"`
}

// Predictions is the generated response
type Predictions struct {
	Embedding Embedding `json:"embeddings"`
}

// Embedding is the embedding for text
type Embedding struct {
	Values     []float64      `json:"values"`
	Statistics EmbeddingStats `json:"statistics"`
}

// EmbeddingStats define the statistics for a text embedding
type EmbeddingStats struct {
	TokenCount int  `json:"token_count"`
	Truncated  bool `json:"truncated"`
}

const (
	textEmbeddingEndpoint = ":predict"
)

func (c *client) TextEmbedding(
	ctx context.Context,
	req TextEmbeddingRequest,
) (*TextEmbeddingResponse, error) {
	payload, err := c.getTextEmbeddingPayload(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := c.requestBuilder.build(
		ctx,
		http.MethodPost,
		req.ProjectID,
		req.EndpointID,
		textGenerationEndpoint,
		payload,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request for predict endpoint: %v", err)
	}

	resp := &TextEmbeddingResponse{}
	err = c.sendRequest(httpReq, resp)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to predict endpoint: %v", err)
	}

	return resp, nil
}

func (c *client) getTextEmbeddingPayload(
	req TextEmbeddingRequest,
) (payload, error) {
	payload := payload{}
	if req.Content != "" && len(req.Inputs) > 0 {
		return payload, errors.New("received Content and Inputs both, only expected one")
	}

	if req.Content != "" {
		payload.Instances = []inputInstances{{Content: req.Content}}
	} else { // Use Inputs instead of Content
		payload.Instances = make([]inputInstances, 0, len(req.Inputs))
		for _, input := range req.Inputs {
			payload.Instances = append(payload.Instances, inputInstances{
				Content: input,
			})
		}
	}
	return payload, nil
}
