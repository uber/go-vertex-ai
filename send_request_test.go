package vertexai

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type httpMock1 struct{}

func (hm *httpMock1) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("{\"success\": true}")),
	}, nil
}

type httpMock2 struct{}

func (hm *httpMock2) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(strings.NewReader("error calling http client")),
	}, nil
}

type httpMock3 struct{}

func (hm *httpMock3) Do(req *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("error calling http client")
}

type httpMock4 struct{}

func (hm *httpMock4) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("{\"success\": \"true\"}")),
	}, nil
}

type tokenMockSuccess struct{}

func (*tokenMockSuccess) getToken(string) (string, error) {
	return "", nil
}

func Test_client_sendRequest(t *testing.T) {
	httpReq, err := http.NewRequest("method", "url", nil)
	if err != nil {
		t.Errorf("failed to create http request: %v", err)
	}
	type respStruct struct {
		Success bool `json:"success"`
	}
	tests := []struct {
		name           string
		req            *http.Request
		want           *respStruct
		wantErr        bool
		mockHTTPClient httpClient
		mockTokenizer  tokenizer
	}{
		{
			name: "success",
			req:  httpReq,
			want: &respStruct{
				Success: true,
			},
			wantErr:        false,
			mockHTTPClient: &httpMock1{},
			mockTokenizer:  &tokenMockSuccess{},
		},
		{
			name:           "error response from http client",
			req:            httpReq,
			wantErr:        true,
			mockHTTPClient: &httpMock2{},
			mockTokenizer:  &tokenMockSuccess{},
		},
		{
			name:           "error in calling http client",
			req:            httpReq,
			wantErr:        true,
			mockHTTPClient: &httpMock3{},
			mockTokenizer:  &tokenMockSuccess{},
		},
		{
			name: "error decoding response",
			req:  httpReq,
			want: &respStruct{
				Success: false,
			},
			wantErr:        true,
			mockHTTPClient: &httpMock4{},
			mockTokenizer:  &tokenMockSuccess{},
		},
		{
			name:          "error get token",
			req:           httpReq,
			want:          nil,
			wantErr:       true,
			mockTokenizer: &defaultTokenizer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				clientConfig: clientConfig{
					authToken: "test_token",
				},
				httpClient: tt.mockHTTPClient,
				tokenizer:  tt.mockTokenizer,
			}
			got := &respStruct{}
			if tt.want == nil {
				got = nil
			}
			err := c.sendRequest(tt.req, got)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
