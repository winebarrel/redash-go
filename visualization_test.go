package redash_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/araddon/dateparse"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go/v2"
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

func Test_UpdateVisualization_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/visualizations/1", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.UpdateVisualization(context.Background(), 1, &redash.UpdateVisualizationInput{
		Description: "description",
		Name:        "name",
		Type:        "TABLE",
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_UpdateVisualization_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/visualizations/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.UpdateVisualization(context.Background(), 1, &redash.UpdateVisualizationInput{
		Description: "description",
		Name:        "name",
		Type:        "TABLE",
	})
	assert.ErrorContains(err, "POST api/visualizations/1 failed: HTTP status code not OK: 503\nerror")
}

func Test_Visualization_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	ds, err := client.CreateDataSource(context.Background(), &redash.CreateDataSourceInput{
		Name: "test-postgres-1",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
	})
	require.NoError(err)

	defer func() {
		client.DeleteDataSource(context.Background(), ds.ID) //nolint:errcheck
	}()

	query, _ := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "test-query-1",
		Query:        "select 1",
	})

	defer func() {
		client.ArchiveQuery(context.Background(), query.ID) //nolint:errcheck
	}()

	if len(query.Visualizations) < 1 {
		assert.FailNow("len(query.Visualizations) < 1")
	}

	vizId := query.Visualizations[0].ID
	viz, err := client.UpdateVisualization(context.Background(), vizId, &redash.UpdateVisualizationInput{
		Name:        "test-viz-1",
		Description: "test-viz-1-desc",
	})
	require.NoError(err)
	assert.Equal("test-viz-1", viz.Name)
	assert.Equal("test-viz-1-desc", viz.Description)
}
