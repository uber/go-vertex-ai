package vertexai

import (
	"context"
	"fmt"

	"golang.org/x/oauth2/google"
)

type tokenizer interface {
	getToken(key string) (string, error)
}

type defaultTokenizer struct{}

// getToken returns the access token.
func (*defaultTokenizer) getToken(key string) (string, error) {
	scopes := []string{"https://www.googleapis.com/auth/cloud-platform"}

	credentials, err := google.CredentialsFromJSON(context.Background(), []byte(key), scopes...)
	if err != nil {
		return "", fmt.Errorf("GetToken: Unable to get credentials from the JSON: %v", err)
	}

	// Create a new access token from the retrieved credentials.
	token, err := credentials.TokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("GetToken: Unable to retrieve token from token source: %v", err)
	}

	return token.AccessToken, nil
}
