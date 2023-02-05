package redash_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/araddon/dateparse"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/redash-go"
)

func Test_UpdateVisualization_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/visualizations/1", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"description":"description","name":"name","type":"TABLE"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"created_at": "2023-02-10T01:23:45.000Z",
				"description": "description",
				"id": 1,
				"name": "Table",
				"options": {},
				"type": "TABLE",
				"updated_at": "2023-02-10T01:23:45.000Z"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.UpdateVisualization(context.Background(), 1, &redash.UpdateVisualizationInput{
		Description: "description",
		Name:        "name",
		Type:        "TABLE",
	})
	assert.NoError(err)
	assert.Equal(&redash.Visualization{
		CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		Description: "description",
		ID:          1,
		Name:        "Table",
		Options:     map[string]any{},
		Query:       redash.Query{},
		Type:        "TABLE",
		UpdatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
	}, res)
}
