package vertexai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *client) sendRequest(req *http.Request, val any) error {
	token, err := c.tokenizer.getToken(c.clientConfig.authToken)
	if err != nil {
		return fmt.Errorf("failed to get auth token: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		var errResp []byte
		defer res.Body.Close()
		_, err := res.Body.Read(errResp)
		if err != nil {
			return fmt.Errorf("status code: %s, message: %w", res.Status, err)
		}
		return fmt.Errorf("status code: %s, message: %s", res.Status, string(errResp))
	}

	if val != nil {
		if err = json.NewDecoder(res.Body).Decode(val); err != nil {
			return err
		}
	}

	return nil
}
