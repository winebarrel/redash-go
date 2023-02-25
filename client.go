package redash

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"

	"github.com/winebarrel/redash-go/internal/util"
)

const (
	UserAgent = "redash-go"
)

type Client struct {
	httpCli  *http.Client
	endpoint string
	apiKey   string
	Debug    bool
}

type ClientWithoutContext struct {
	withCtx *Client
}

func NewClient(endpoint string, apiKey string) (*Client, error) {
	return NewClientWithHTTPClient(endpoint, apiKey, nil)
}

func MustNewClient(endpoint string, apiKey string) *Client {
	client, err := NewClient(endpoint, apiKey)

	if err != nil {
		panic("MustNewClient(" + strconv.Quote(endpoint) + `, ...): ` + err.Error())
	}

	return client
}

func MustNewClientWithHTTPClient(endpoint string, apiKey string, httpClient *http.Client) *Client {
	client, err := NewClientWithHTTPClient(endpoint, apiKey, httpClient)

	if err != nil {
		panic("MustNewClientWithHTTPClient(" + strconv.Quote(endpoint) + `, ...): ` + err.Error())
	}

	return client
}

func NewClientWithHTTPClient(endpoint string, apiKey string, httpClient *http.Client) (*Client, error) {
	_, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		httpCli:  httpClient,
		endpoint: endpoint,
		apiKey:   apiKey,
	}

	return client, nil
}

func (client *Client) WithoutContext() *ClientWithoutContext {
	return &ClientWithoutContext{
		withCtx: client,
	}
}

type responseCloser func()

func (client *Client) Get(ctx context.Context, path string, params any) (*http.Response, responseCloser, error) {
	res, err := client.sendRequest(ctx, http.MethodGet, path, params, nil)

	if err != nil {
		return nil, func() {}, fmt.Errorf("GET %s failed: %w", path, err)
	}

	return res, func() { util.CloseResponse(res) }, nil
}

func (client *Client) Post(ctx context.Context, path string, body any) (*http.Response, responseCloser, error) {
	res, err := client.sendRequest(ctx, http.MethodPost, path, nil, body)

	if err != nil {
		return nil, func() {}, fmt.Errorf("POST %s failed: %w", path, err)
	}

	return res, func() { util.CloseResponse(res) }, nil
}

func (client *Client) Delete(ctx context.Context, path string) (*http.Response, responseCloser, error) {
	res, err := client.sendRequest(ctx, http.MethodDelete, path, nil, nil)

	if err != nil {
		return nil, func() {}, fmt.Errorf("DELETE %s failed: %w", path, err)
	}

	return res, func() { util.CloseResponse(res) }, nil
}

func (client *Client) sendRequest(ctx context.Context, method string, path string, params any, body any) (*http.Response, error) {
	url, err := url.JoinPath(client.endpoint, path)

	if err != nil {
		return nil, err
	}

	var reader io.Reader

	if body != nil {
		rawBody, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}

		reader = bytes.NewReader(rawBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reader)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Set("Authorization", "Key "+client.apiKey)

	if params != nil {
		values, err := util.URLValuesFrom(params)

		if err != nil {
			return nil, err
		}

		req.URL.RawQuery = values.Encode()
	}

	if client.Debug {
		b, _ := httputil.DumpRequest(req, true)
		fmt.Fprintf(os.Stderr, "---request begin---\n%s\n---request end---\n", b)
	}

	res, err := client.httpCli.Do(req)

	if err != nil {
		return nil, err
	}

	if client.Debug {
		b, _ := httputil.DumpResponse(res, true)
		fmt.Fprintf(os.Stderr, "---response begin---\n%s\n---response end---\n", b)
	}

	if err := util.CheckStatus(res); err != nil {
		return nil, err
	}

	return res, err
}
