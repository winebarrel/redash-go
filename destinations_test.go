package redash_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go/v2"
)

func Test_ListDestinations_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/destinations", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, `
			[
				{
					"icon": "fa-envelope",
					"id": 1,
					"name": "alert@example.com",
					"type": "email"
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListDestinations(context.Background())
	assert.NoError(err)
	assert.Equal([]redash.Destination{
		{
			Icon:    "fa-envelope",
			ID:      1,
			Name:    "alert@example.com",
			Options: nil,
			Type:    "email",
		},
	}, res)
}

func Test_ListDestinations_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/destinations", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListDestinations(context.Background())
	assert.ErrorContains(err, "GET api/destinations failed: HTTP status code not OK: 503 Service Unavailable\nerror")
}

func Test_ListDestinations_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/destinations", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListDestinations(context.Background())
	assert.ErrorContains(err, "read response body failed: IO error")
}

func Test_GetDestination_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/destinations/1", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"icon": "fa-envelope",
				"id": 1,
				"name": "alert@example.com",
				"options": {
					"addresses": "alert@example.com"
				},
				"type": "email"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetDestination(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.Destination{
		Icon: "fa-envelope",
		ID:   1,
		Name: "alert@example.com",
		Options: map[string]any{
			"addresses": "alert@example.com",
		},
		Type: "email",
	}, res)
}

func Test_GetDestination_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/destinations/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetDestination(context.Background(), 1)
	assert.ErrorContains(err, "GET api/destinations/1 failed: HTTP status code not OK: 503 Service Unavailable\nerror")
}

func Test_GetDestination_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/destinations/1", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetDestination(context.Background(), 1)
	assert.ErrorContains(err, "read response body failed: IO error")
}

func Test_CreateDestination_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/destinations", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		if req.Body == nil {
			assert.FailNow("req.Body is nil")
		}
		body, _ := io.ReadAll(req.Body)
		assert.Equal(`{"name":"alert@example.com","options":{"addresses":"alert@example.com"},"type":"email"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"icon": "fa-envelope",
				"id": 1,
				"name": "alert@example.com",
				"options": {
					"addresses": "alert@example.com"
				},
				"type": "email"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.CreateDestination(context.Background(), &redash.CreateDestinationInput{
		Name: "alert@example.com",
		Options: map[string]any{
			"addresses": "alert@example.com",
		},
		Type: "email",
	})
	assert.NoError(err)
	assert.Equal(&redash.Destination{
		Icon: "fa-envelope",
		ID:   1,
		Name: "alert@example.com",
		Options: map[string]any{
			"addresses": "alert@example.com",
		},
		Type: "email",
	}, res)
}

func Test_CreateDestination_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/destinations", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.CreateDestination(context.Background(), &redash.CreateDestinationInput{
		Name: "alert@example.com",
		Options: map[string]any{
			"addresses": "alert@example.com",
		},
		Type: "email",
	})
	assert.ErrorContains(err, "POST api/destinations failed: HTTP status code not OK: 503 Service Unavailable\nerror")
}

func Test_CreateDestination_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/destinations", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.CreateDestination(context.Background(), &redash.CreateDestinationInput{
		Name: "alert@example.com",
		Options: map[string]any{
			"addresses": "alert@example.com",
		},
		Type: "email",
	})
	assert.ErrorContains(err, "read response body failed: IO error")
}

func Test_DeleteDestination_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/destinations/1", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.DeleteDestination(context.Background(), 1)
	assert.NoError(err)
}

func Test_DeleteDestination_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/destinations/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.DeleteDestination(context.Background(), 1)
	assert.ErrorContains(err, "DELETE api/destinations/1 failed: HTTP status code not OK: 503 Service Unavailable\nerror")
}

func Test_GetDestinationTypes_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/destinations/types", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, `
			[
				{
					"configuration_schema": {
						"extra_options": [
							"subject_template"
						],
						"properties": {
							"icon_url": {
								"title": "Icon URL (32x32 or multiple, png format)",
								"type": "string"
							},
							"url": {
								"title": "Webhook URL (get it from the room settings)",
								"type": "string"
							},
							"subject_template": {
								"default": "({state}) {alert_name}",
								"title": "Subject Template",
								"type": "string"
							}
						},
						"required": [
							"url"
						],
						"secret": [
							"url"
						],
						"type": "object"
					},
					"icon": "fa-bolt",
					"name": "Google Hangouts Chat",
					"type": "hangouts_chat"
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetDestinationTypes(context.Background())
	assert.NoError(err)
	assert.Equal([]redash.DestinationType{
		{
			ConfigurationSchema: redash.DestinationTypeConfigurationSchema{
				ExtraOptions: []string{
					"subject_template",
				},
				Properties: map[string]redash.DestinationTypeConfigurationSchemaProperty{
					"icon_url": {
						Title: "Icon URL (32x32 or multiple, png format)",
						Type:  "string",
					},
					"url": {
						Title: "Webhook URL (get it from the room settings)",
						Type:  "string",
					},
					"subject_template": {
						Default: "({state}) {alert_name}",
						Title:   "Subject Template",
						Type:    "string",
					},
				},
				Required: []string{
					"url",
				},
				Secret: []any{
					"url",
				},
				Type: "object",
			},
			Icon: "fa-bolt",
			Name: "Google Hangouts Chat",
			Type: "hangouts_chat",
		},
	}, res)
}

func Test_GetDestinationTypes_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/destinations/types", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetDestinationTypes(context.Background())
	assert.ErrorContains(err, "GET api/destinations/types failed: HTTP status code not OK: 503 Service Unavailable\nerror")
}

func Test_GetDestinationTypes_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/destinations/types", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetDestinationTypes(context.Background())
	assert.ErrorContains(err, "read response body failed: IO error")
}

func Test_Destination_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)

	_, err := client.ListDestinations(context.Background())
	assert.NoError(err)

	dest, err := client.CreateDestination(context.Background(), &redash.CreateDestinationInput{
		Name: "test-dest-1",
		Options: map[string]any{
			"addresses": "alert@example.com",
		},
		Type: "email",
	})
	require.NoError(err)
	assert.Equal("test-dest-1", dest.Name)

	dest, err = client.GetDestination(context.Background(), dest.ID)
	require.NoError(err)
	assert.Equal("test-dest-1", dest.Name)

	err = client.DeleteDestination(context.Background(), dest.ID)
	require.NoError(err)

	_, err = client.GetDestination(context.Background(), dest.ID)
	assert.Error(err)

	types, err := client.GetDestinationTypes(context.Background())
	require.NoError(err)
	assert.GreaterOrEqual(len(types), 1)
}
