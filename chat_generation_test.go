package vertexai

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type httpMockPredictChat struct{}

func (h *httpMockPredictChat) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(bytes.NewReader([]byte(`{"predictions": [{
			"candidates": [
			  {
				"author": "1",
				"content": "Howdy?"
			  }
			],
			"safetyAttributes": [
				{
				"categories": [
					"test1",
					"test2",
					"test3"
				],
				"blocked": false,
				"scores": [0.1,0.2,0.3]
				}
			]
		  }]}`))),
	}, nil
}

func Test_client_ChatGeneration(t *testing.T) {
	type fields struct {
		clientConfig   clientConfig
		requestBuilder requestBuilder
		httpClient     httpClient
	}
	type args struct {
		ctx context.Context
		req ChatGenerationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ChatGenerationResponse
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
				req: ChatGenerationRequest{},
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
				req: ChatGenerationRequest{},
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
				httpClient:     &httpMockPredictChat{},
			},
			args: args{
				ctx: context.Background(),
				req: ChatGenerationRequest{},
			},
			want: &ChatGenerationResponse{
				ChatPredictions: []ChatPrediction{
					{
						ChatMessages: []ChatMessage{
							{
								Author:  "1",
								Content: "Howdy?",
							},
						},
						SafetyAttributes: []SafetyAttributes{
							{
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
			got, err := c.ChatGeneration(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
