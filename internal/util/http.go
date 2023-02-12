package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func UnmarshalBody(res *http.Response, v any) error {
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Errorf("Read response body failed: %w", err)
	}

	err = json.Unmarshal(body, v)

	if err != nil {
		return fmt.Errorf("Unmarshal response body failed: %w", err)
	}

	return nil
}

func CheckStatus(res *http.Response) error {
	if 200 <= res.StatusCode && res.StatusCode <= 299 {
		return nil
	}

	msg := fmt.Sprintf("HTTP status code not OK: %s", res.Status)

	if res.Body != nil {
		body, err := io.ReadAll(res.Body)

		if err == nil && len(body) > 0 {
			msg += "\n" + string(body)
		}
	}

	return errors.New(msg)
}

func CloseResponse(res *http.Response) {
	if res == nil || res.Body == nil {
		return
	}

	io.Copy(io.Discard, res.Body) //nolint:errcheck
	res.Body.Close()
}
