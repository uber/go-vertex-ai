package vertexai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	predictURLTemplate = "https://us-central1-aiplatform.googleapis.com/v1/projects/%v/locations/us-central1/publishers/google/models/%v" // projectID, endpointID
)

type requestBuilder interface {
	build(
		ctx context.Context,
		method string,
		projectID string,
		endpointID string,
		urlSuffix string,
		request any,
	) (*http.Request, error)
}

type httpRequestBuilder struct{}

func newRequestBuilder() requestBuilder {
	return &httpRequestBuilder{}
}

func (b *httpRequestBuilder) build(
	ctx context.Context,
	method string,
	projectID string,
	endpointID string,
	urlSuffix string,
	request any,
) (*http.Request, error) {
	url := fmt.Sprintf(predictURLTemplate, projectID, endpointID) + urlSuffix

	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBytes))
}
