package redash_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go/v2"
)

func Test_GetStatus_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/status.json", func(req *http.Request) (*http.Response, error) {
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
				"dashboards_count": 2,
				"database_metrics": {
					"metrics": [
						[
							"Query Results Size",
							49152
						],
						[
							"Redash DB Size",
							9371820
						]
					]
				},
				"manager": {
					"last_refresh_at": "1676217265.221729",
					"outdated_queries_count": "0",
					"query_ids": "[]",
					"queues": {
						"celery": {
							"size": 0
						},
						"queries": {
							"size": 0
						},
						"scheduled_queries": {
							"size": 0
						}
					}
				},
				"queries_count": 5,
				"query_results_count": 4,
				"redis_used_memory": 2383280,
				"redis_used_memory_human": "2.27M",
				"unused_query_results_count": 0,
				"version": "8.0.0+b32245",
				"widgets_count": 0,
				"workers": []
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetStatus(context.Background())
	assert.NoError(err)
	assert.Equal(&redash.Status{
		DashboardsCount: 2,
		DatabaseMetrics: redash.StatusDatabaseMetrics{
			Metrics: [][]any{
				{"Query Results Size", float64(49152)},
				{"Redash DB Size", 9.37182e+06},
			},
		},
		Manager: redash.StatusManager{
			LastRefreshAt:        "1676217265.221729",
			OutdatedQueriesCount: "0",
			QueryIds:             "[]",
			Queues: redash.StatusManagerQueues{
				Celery: redash.StatusManagerQueuesCelery{
					Size: 0,
				},
				Queries: redash.StatusManagerQueuesQueries{
					Size: 0,
				},
				ScheduledQueries: redash.StatusManagerQueuesScheduledQueries{
					Size: 0,
				},
			},
		},
		QueriesCount:            5,
		QueryResultsCount:       4,
		RedisUsedMemory:         2383280,
		RedisUsedMemoryHuman:    "2.27M",
		UnusedQueryResultsCount: 0,
		Version:                 "8.0.0+b32245",
		WidgetsCount:            0,
		Workers:                 []any{},
	}, res)
}

func Test_GetStatus_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	status, err := client.GetStatus(context.Background())
	require.NoError(err)
	assert.NotEmpty(status.Version)
}
