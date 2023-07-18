package vertexai

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type httpMockPredictTextEmbedding struct{}

func (h *httpMockPredictTextEmbedding) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
			"predictions": [
			  {
				"embeddings": {
				  "statistics": {
					"truncated": true,
					"token_count": 6
				  },
				  "values": [
					0.014090980403125286,
					-0.016960423439741135,
					-0.033782251179218292,
					0.0079174842685461044
				  ]
				}
			  }
			]
		  }`))),
	}, nil
}

func Test_client_TextEmbedding(t *testing.T) {
	type fields struct {
		clientConfig   clientConfig
		requestBuilder requestBuilder
		httpClient     httpClient
	}
	type args struct {
		ctx context.Context
		req TextEmbeddingRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TextEmbeddingResponse
		wantErr bool
	}{
		{
			name: "error both content and inputs passed",
			fields: fields{
				clientConfig: clientConfig{
					authToken: "test_token",
				},
			},
			args: args{
				ctx: context.Background(),
				req: TextEmbeddingRequest{
					Content: "sample content",
					Inputs: []string{
						"sample content",
						"sample content",
					},
				},
			},
			wantErr: true,
		},
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
				req: TextEmbeddingRequest{},
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
				req: TextEmbeddingRequest{},
			},
			wantErr: true,
		},
		{
			name: "success - Content",
			fields: fields{
				clientConfig: clientConfig{
					authToken: "test_token",
				},
				requestBuilder: &requestBuilderImpl2{},
				httpClient:     &httpMockPredictTextEmbedding{},
			},
			args: args{
				ctx: context.Background(),
				req: TextEmbeddingRequest{
					Content: "sample content",
				},
			},
			want: &TextEmbeddingResponse{
				Predictions: []Predictions{
					Predictions{
						Embedding: Embedding{
							Values: []float64{
								0.014090980403125286,
								-0.016960423439741135,
								-0.033782251179218292,
								0.0079174842685461044,
							},
							Statistics: EmbeddingStats{
								TokenCount: 6,
								Truncated:  true,
							},
						},
					},
				},
			},
		},
		{
			name: "success - Inputs",
			fields: fields{
				clientConfig: clientConfig{
					authToken: "test_token",
				},
				requestBuilder: &requestBuilderImpl2{},
				httpClient:     &httpMockPredictTextEmbedding{},
			},
			args: args{
				ctx: context.Background(),
				req: TextEmbeddingRequest{
					Inputs: []string{
						"sample content",
						"sample content",
					},
				},
			},
			want: &TextEmbeddingResponse{
				Predictions: []Predictions{
					Predictions{
						Embedding: Embedding{
							Values: []float64{
								0.014090980403125286,
								-0.016960423439741135,
								-0.033782251179218292,
								0.0079174842685461044,
							},
							Statistics: EmbeddingStats{
								TokenCount: 6,
								Truncated:  true,
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
			got, err := c.TextEmbedding(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
