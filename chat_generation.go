package vertexai

import (
	"context"
	"fmt"
	"net/http"
)

// ChatGenerationRequest is the request for the ChatGeneration method
type ChatGenerationRequest struct {
	ProjectID      string
	EndpointID     string
	Context        string
	ChatMessages   []ChatMessage
	Examples       []Example
	Temperature    *float64
	MaxDecodeSteps *int
	TopK           *int
	TopP           *float64
}

// ChatGenerationResponse is the response for the ChatGeneration method
type ChatGenerationResponse struct {
	ChatPredictions  []ChatPrediction `json:"predictions"`
	DeployedModelID  string           `json:"deployedModelId"`
	Model            string           `json:"model"`
	ModelDisplayName string           `json:"modelDisplayName"`
	ModelVersionID   string           `json:"modelVersionId"`
}

// ChatPrediction is the generated response
type ChatPrediction struct {
	ChatMessages     []ChatMessage    `json:"candidates"`
	SafetyAttributes SafetyAttributes `json:"safetyAttributes"`
}

const (
	chatGenerationEndpoint = ":predict"
)

func (c *client) ChatGeneration(
	ctx context.Context,
	req ChatGenerationRequest,
) (*ChatGenerationResponse, error) {
	payload := c.getChatGenerationPayload(req)
	httpReq, err := c.requestBuilder.build(
		ctx,
		http.MethodPost,
		req.ProjectID,
		req.EndpointID,
		chatGenerationEndpoint,
		payload,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request for predict endpoint: %v", err)
	}

	resp := &ChatGenerationResponse{}
	err = c.sendRequest(httpReq, resp)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to predict endpoint: %v", err)
	}

	return resp, nil
}

func (c *client) getChatGenerationPayload(
	req ChatGenerationRequest,
) payload {
	return payload{
		Instances: []inputInstances{{
			Context:  req.Context,
			Examples: req.Examples,
			Messages: req.ChatMessages,
		}},
		Parameters: parameters{
			Temperature:    req.Temperature,
			MaxDecodeSteps: req.MaxDecodeSteps,
			TopP:           req.TopP,
			TopK:           req.TopK,
		},
	}
}
