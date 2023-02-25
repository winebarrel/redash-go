package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
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

func URLValuesFrom(params any) (url.Values, error) {
	if m, ok := params.(map[string]string); ok {
		values := url.Values{}

		for k, v := range m {
			values.Add(k, v)
		}

		return values, nil
	} else {
		values, err := query.Values(params)

		if err != nil {
			return nil, err
		}

		return values, nil
	}
}
