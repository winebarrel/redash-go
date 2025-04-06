package redash

import (
	"bytes"
	"context"
	"io"
	"math"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_sendRequest_OK(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key <secret>"},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		assert.Equal("foo=bar", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `{"zoo":"baz"}`), nil
	})

	client, _ := NewClient("https://redash.example.com", "<secret>")
	res, err := client.sendRequest(context.Background(), http.MethodGet, "api/queries/1", map[string]string{"foo": "bar"}, nil)
	assert.NoError(err)
	assert.Equal("200 OK", res.Status)
	require.NotNil(res.Body)
	body, _ := io.ReadAll(res.Body)
	assert.Equal(`{"zoo":"baz"}`, string(body))
}

func Test_sendRequest_Err_JoinPath(t *testing.T) {
	assert := assert.New(t)
	client, _ := NewClient("https://redash.example.com", "<secret>")
	client.endpoint = "\b"
	_, err := client.sendRequest(context.Background(), http.MethodGet, "api/queries/1", map[string]string{"foo": "bar"}, nil)
	assert.ErrorContains(err, "parse \"\\b\": net/url: invalid control character in URL")
}

func Test_sendRequest_Err_NewRequestWithContext(t *testing.T) {
	assert := assert.New(t)
	client, _ := NewClient("https://redash.example.com", "<secret>")
	_, err := client.sendRequest(context.Background(), "あいうえお", "api/queries/1", map[string]string{"foo": "bar"}, nil)
	assert.ErrorContains(err, "net/http: invalid method \"あいうえお\"")
}

func Test_sendRequest_Err_Marshal(t *testing.T) {
	assert := assert.New(t)
	client, _ := NewClient("https://redash.example.com", "<secret>")
	_, err := client.sendRequest(context.Background(), http.MethodGet, "api/queries/1", map[string]string{"foo": "bar"}, math.NaN())
	assert.ErrorContains(err, "json: unsupported value: NaN")
}

func Test_sendRequest_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := NewClient("https://redash.example.com", "<secret>")
	_, err := client.sendRequest(context.Background(), http.MethodGet, "api/queries/1", map[string]string{"foo": "bar"}, nil)
	assert.ErrorContains(err, "HTTP status code not OK: 503 Service Unavailable\nerror")
}

func Test_sendRequest_Err_params(t *testing.T) {
	assert := assert.New(t)
	client, _ := NewClient("https://redash.example.com", "<secret>")
	_, err := client.sendRequest(context.Background(), http.MethodGet, "api/queries/1", "bad params", nil)
	assert.ErrorContains(err, "query: Values() expects struct input. Got string")
}

func Test_sendRequest_Err_HTTPRequest(t *testing.T) {
	assert := assert.New(t)
	client, _ := NewClient("https://redash.example.com", "<secret>")
	client.endpoint = "x"
	_, err := client.sendRequest(context.Background(), http.MethodGet, "api/queries/1", map[string]string{"foo": "bar"}, nil)
	assert.ErrorContains(err, `Get "x/api/queries/1?foo=bar": unsupported protocol scheme ""`)
}

func Test_sendRequest_Debug(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	origDebugOut := _debugOut
	var buf bytes.Buffer
	_debugOut = &buf
	defer func() { _debugOut = origDebugOut }()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusOK, `{"zoo":"baz"}`), nil
	})

	client, _ := NewClient("https://redash.example.com", "<secret>")
	client.SetDebug(true)
	res, err := client.sendRequest(context.Background(), http.MethodGet, "api/queries/1", map[string]string{"foo": "bar"}, nil)
	require.NoError(err)
	assert.Equal("200 OK", res.Status)
	assert.Equal("---request begin---\nGET /api/queries/1?foo=bar HTTP/1.1\r\nHost: redash.example.com\r\nAuthorization: Key <secret>\r\nContent-Type: application/json\r\nUser-Agent: redash-go\r\n\r\n\n---request end---\n---response begin---\nHTTP/0.0 200 OK\r\n\r\n{\"zoo\":\"baz\"}\n---response end---\n", buf.String())
}
