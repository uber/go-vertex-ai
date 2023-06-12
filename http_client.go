package vertexai

import "net/http"

// Client is the interface for making HTTP requests. This is specifically used for ease in mocking in unit tests
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
