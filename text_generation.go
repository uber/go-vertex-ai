package vertexai

import (
	"context"
	"fmt"
	"net/http"
)

// TextGenerationRequest is the request for the TextGeneration method
type TextGenerationRequest struct {
	ProjectID      string
	EndpointID     string
	Content        string
	Temperature    *float64
	MaxDecodeSteps *int
	TopK           *int
	TopP           *float64
}

// TextGenerationResponse is the response for the TextGeneration method
type TextGenerationResponse struct {
	Candidates       []TextCandidate `json:"predictions"`
	DeployedModelID  string          `json:"deployedModelId"`
	Model            string          `json:"model"`
	ModelDisplayName string          `json:"modelDisplayName"`
	ModelVersionID   string          `json:"modelVersionId"`
}

// TextCandidate is the generated response
type TextCandidate struct {
	Content          string           `json:"content"`
	SafetyAttributes SafetyAttributes `json:"safetyAttributes"`
}

const (
	textGenerationEndpoint = ":predict"
)

func (c *client) TextGeneration(
	ctx context.Context,
	req TextGenerationRequest,
) (*TextGenerationResponse, error) {
	payload := c.getTextGenerationPayload(req)
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

	resp := &TextGenerationResponse{}
	err = c.sendRequest(httpReq, resp)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to predict endpoint: %v", err)
	}

	return resp, nil
}

func (c *client) getTextGenerationPayload(
	req TextGenerationRequest,
) payload {
	return payload{
		Instances: []inputInstances{{Content: req.Content}},
		Parameters: parameters{
			Temperature:    req.Temperature,
			MaxDecodeSteps: req.MaxDecodeSteps,
			TopP:           req.TopP,
			TopK:           req.TopK,
		},
	}
}
