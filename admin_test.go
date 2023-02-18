package redash_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/redash-go"
)

func Test_GetAdminQueriesOutdated_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/admin/queries/outdated", func(req *http.Request) (*http.Response, error) {
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
				"queries": [],
				"updated_at": "1676390846.0004709"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetAdminQueriesOutdated(context.Background())
	assert.NoError(err)
	assert.Equal(&redash.AdminQueriesOutdated{
		Queries:   []redash.Query{},
		UpdatedAt: "1676390846.0004709",
	}, res)
}

func Test_GetAdminQueriesRqStatus_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/admin/queries/rq_status", func(req *http.Request) (*http.Response, error) {
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
				"queues": {
					"default": {
						"name": "default",
						"queued": 0,
						"started": []
					},
					"emails": {
						"name": "emails",
						"queued": 0,
						"started": []
					},
					"periodic": {
						"name": "periodic",
						"queued": 0,
						"started": []
					},
					"queries": {
						"name": "queries",
						"queued": 0,
						"started": [
							{
								"enqueued_at": "2023-02-14T16:19:51.947",
								"id": "8cd29f9b-4227-4fd1-9197-3fe3ccdd5562",
								"meta": {
									"data_source_id": 4,
									"org_id": 1,
									"query_id": "7",
									"scheduled": false,
									"user_id": 1
								},
								"name": "redash.tasks.queries.execution.execute_query",
								"origin": "queries",
								"started_at": "2023-02-14T16:19:51.963"
							}
						]
					},
					"schemas": {
						"name": "schemas",
						"queued": 0,
						"started": []
					}
				},
				"workers": [
					{
						"birth_date": "2023-02-14T16:04:39.454",
						"current_job": null,
						"failed_jobs": 0,
						"hostname": "2e21d2a7bae6",
						"last_heartbeat": "2023-02-14T16:19:53.462",
						"name": "8fb4d57148254783b6e63a0e420738ac",
						"pid": 16,
						"queues": "periodic, emails, default, scheduled_queries, queries, schemas",
						"state": "busy",
						"successful_jobs": 49,
						"total_working_time": 2.224891
					},
					{
						"birth_date": "2023-02-14T16:04:38.736",
						"current_job": "8cd29f9b-4227-4fd1-9197-3fe3ccdd5562 (execute_query)",
						"failed_jobs": 0,
						"hostname": "2e21d2a7bae6",
						"last_heartbeat": "2023-02-14T16:19:51.963",
						"name": "980ad4d9b3994134abf281d93aaec1b7",
						"pid": 15,
						"queues": "periodic, emails, default, scheduled_queries, queries, schemas",
						"state": "busy",
						"successful_jobs": 44,
						"total_working_time": 62.508043
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetAdminQueriesRqStatus(context.Background())
	assert.NoError(err)
	assert.Equal(&redash.AdminQuerisRqStatus{
		Queues: redash.AdminQuerisRqStatusQueues{
			Default: &redash.AdminQuerisRqStatusDefault{
				Name:    "default",
				Queued:  0,
				Started: []any{},
			},
			Emails: &redash.AdminQuerisRqStatusEmails{
				Name:    "emails",
				Queued:  0,
				Started: []any{},
			},
			Periodic: &redash.AdminQuerisRqStatusPeriodic{
				Name:    "periodic",
				Queued:  0,
				Started: []any{},
			},
			Queries: &redash.AdminQuerisRqStatusQueries{
				Name:   "queries",
				Queued: 0,
				Started: []any{
					map[string]any{
						"enqueued_at": "2023-02-14T16:19:51.947",
						"id":          "8cd29f9b-4227-4fd1-9197-3fe3ccdd5562",
						"meta": map[string]any{
							"data_source_id": float64(4),
							"org_id":         float64(1),
							"query_id":       "7",
							"scheduled":      false,
							"user_id":        float64(1),
						},
						"name":       "redash.tasks.queries.execution.execute_query",
						"origin":     "queries",
						"started_at": "2023-02-14T16:19:51.963",
					},
				},
			},
			Schemas: &redash.AdminQuerisRqStatusSchemas{
				Name:    "schemas",
				Queued:  0,
				Started: []any{},
			},
		},
		Workers: []redash.AdminQuerisRqStatusWorker{
			{
				BirthDate:        "2023-02-14T16:04:39.454",
				CurrentJob:       "",
				FailedJobs:       0,
				Hostname:         "2e21d2a7bae6",
				LastHeartbeat:    "2023-02-14T16:19:53.462",
				Name:             "8fb4d57148254783b6e63a0e420738ac",
				Pid:              16,
				Queues:           "periodic, emails, default, scheduled_queries, queries, schemas",
				State:            "busy",
				SuccessfulJobs:   49,
				TotalWorkingTime: 2.224891,
			},
			{
				BirthDate:        "2023-02-14T16:04:38.736",
				CurrentJob:       "8cd29f9b-4227-4fd1-9197-3fe3ccdd5562 (execute_query)",
				FailedJobs:       0,
				Hostname:         "2e21d2a7bae6",
				LastHeartbeat:    "2023-02-14T16:19:51.963",
				Name:             "980ad4d9b3994134abf281d93aaec1b7",
				Pid:              15,
				Queues:           "periodic, emails, default, scheduled_queries, queries, schemas",
				State:            "busy",
				SuccessfulJobs:   44,
				TotalWorkingTime: 62.508043,
			},
		},
	}, res)
}

func Test_Admin_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	outdated, err := client.GetAdminQueriesOutdated(context.Background())
	if err != nil {
		assert.FailNow(err.Error())
	}
	assert.NoError(err)
	assert.NotEmpty(outdated.UpdatedAt)

	rqStatus, err := client.GetAdminQueriesRqStatus(context.Background())
	if err != nil {
		assert.FailNow(err.Error())
	}
	assert.NoError(err)
	assert.Equal("default", rqStatus.Queues.Default.Name)
}
