package vertexai

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newRequestBuilder(t *testing.T) {
	got := newRequestBuilder()
	assert.NotNil(t, got)
}

func Test_buildRequest(t *testing.T) {
	type testBody struct {
		Content string `json:"content"`
	}
	content := "content"
	type args struct {
		method     string
		projectID  string
		endpointID string
		urlSuffix  string
		req        any
	}
	rawReq := testBody{Content: content}
	jsonReq, err := json.Marshal(rawReq)
	if err != nil {
		t.Errorf("failed to setup test request: %v", err)
	}
	httpReq, err := http.NewRequestWithContext(context.Background(),
		"POST",
		"https://us-central1-aiplatform.googleapis.com/v1/projects/cloud-large-language-models/locations/us-central1/publishers/google/models/4511608470067216384:predict",
		bytes.NewBuffer(jsonReq),
	)
	if err != nil {
		t.Errorf("failed to setup test httpRequest: %v", err)
	}

	httpReqNilBody, err := http.NewRequestWithContext(context.Background(),
		"POST",
		"https://us-central1-aiplatform.googleapis.com/v1/projects/cloud-large-language-models/locations/us-central1/publishers/google/models/4511608470067216384:predict",
		nil,
	)

	if err != nil {
		t.Errorf("failed to setup test httpRequest: %v", err)
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				req:        testBody{Content: content},
				method:     "POST",
				projectID:  "cloud-large-language-models",
				endpointID: "4511608470067216384",
				urlSuffix:  ":predict",
			},
			want: httpReq,
		},
		{
			name: "nil request",
			args: args{
				method:     "POST",
				projectID:  "cloud-large-language-models",
				endpointID: "4511608470067216384",
				urlSuffix:  ":predict",
			},
			want: httpReqNilBody,
		},
		{
			name: "invalid method",
			args: args{
				req:        testBody{Content: content},
				method:     "!@#$",
				projectID:  "cloud-large-language-models",
				endpointID: "4511608470067216384",
				urlSuffix:  ":predict",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&httpRequestBuilder{}).build(
				context.Background(),
				tt.args.method,
				tt.args.projectID,
				tt.args.endpointID,
				tt.args.urlSuffix, tt.args.req)
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				assert.Equal(t, tt.want.Method, got.Method)
				assert.Equal(t, tt.want.URL, got.URL)
				assert.Equal(t, tt.want.Header, got.Header)
				assert.Equal(t, tt.want.Body, got.Body)
			}
		})
	}
}
