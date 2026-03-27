package acex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (a *AcexPlugin) sendRequest(req *http.Request, v any) error {
	// Set default headers
	req.Header.Set("Accept", "application/json; charset=utf-8")

	// Execute request
	res, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Validate status codes
	ok := false
	for _, code := range a.SuccessStatusCodes {
		if res.StatusCode == code {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// Parse JSON
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bodyBytes, &v); err != nil {
		return err
	}

	return nil
}
