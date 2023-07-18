package vertexai

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type requestBuilderImpl1 struct{}

func (r *requestBuilderImpl1) build(ctx context.Context, method string, projectID string, endpointID string, urlSuffix string, request any) (*http.Request, error) {
	return nil, errors.New("error building request")
}

type requestBuilderImpl2 struct{}

func (r *requestBuilderImpl2) build(ctx context.Context, method string, projectID string, endpointID string, urlSuffix string, request any) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, "POST", "http://google.com", io.NopCloser(bytes.NewReader([]byte{})))
}

type httpMockPredict struct{}

func (h *httpMockPredict) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"predictions": [{"content": "test", "safetyAttributes": {"blocked": false, "scores": [0.1, 0.2, 0.3], "categories": ["test1", "test2", "test3"]}}]}`))),
	}, nil
}

func Test_client_TextGeneration(t *testing.T) {
	type fields struct {
		clientConfig   clientConfig
		requestBuilder requestBuilder
		httpClient     httpClient
	}
	type args struct {
		ctx context.Context
		req TextGenerationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TextGenerationResponse
		wantErr bool
	}{
		{
			name: "error building request",
			fields: fields{
				clientConfig: clientConfig{
					authToken: "test_token",
				},
				requestBuilder: &requestBuilderImpl1{},
			},
			args: args{
				ctx: context.Background(),
				req: TextGenerationRequest{},
			},
			wantErr: true,
		},
		{
			name: "error making http call",
			fields: fields{
				clientConfig: clientConfig{
					authToken: "test_token",
				},
				requestBuilder: &requestBuilderImpl2{},
				httpClient:     &httpMock2{},
			},
			args: args{
				ctx: context.Background(),
				req: TextGenerationRequest{},
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				clientConfig: clientConfig{
					authToken: "test_token",
				},
				requestBuilder: &requestBuilderImpl2{},
				httpClient:     &httpMockPredict{},
			},
			args: args{
				ctx: context.Background(),
				req: TextGenerationRequest{},
			},
			want: &TextGenerationResponse{
				Candidates: []TextCandidate{
					{
						Content: "test",
						SafetyAttributes: SafetyAttributes{
							Blocked: false,
							Scores:  []float64{0.1, 0.2, 0.3},
							Categories: []string{
								"test1",
								"test2",
								"test3",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				clientConfig:   tt.fields.clientConfig,
				requestBuilder: tt.fields.requestBuilder,
				httpClient:     tt.fields.httpClient,
				tokenizer:      &tokenMockSuccess{},
			}
			got, err := c.TextGeneration(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
