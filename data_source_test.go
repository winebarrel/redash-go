package redash_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/redash-go"
)

func Test_ListDataSources_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/data_sources", func(req *http.Request) (*http.Response, error) {
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
					"id": 1,
					"name": "postgres",
					"pause_reason": null,
					"paused": 0,
					"syntax": "sql",
					"type": "pg",
					"view_only": false
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListDataSources(context.Background())
	assert.NoError(err)
	assert.Equal([]redash.DataSource{
		{
			Groups:             nil,
			ID:                 1,
			Name:               "postgres",
			Options:            nil,
			Paused:             0,
			PauseReason:        "",
			QueueName:          "",
			ScheduledQueueName: "",
			Syntax:             "sql",
			Type:               "pg",
			ViewOnly:           false,
		},
	}, res)
}

func Test_GetDataSource_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/data_sources/1", func(req *http.Request) (*http.Response, error) {
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
				"groups": {
					"2": false
				},
				"id": 1,
				"name": "postgres",
				"options": {
					"dbname": "postgres",
					"host": "postgres",
					"port": 5432,
					"user": "postgres"
				},
				"pause_reason": null,
				"paused": 0,
				"queue_name": "queries",
				"scheduled_queue_name": "scheduled_queries",
				"syntax": "sql",
				"type": "pg"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetDataSource(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.DataSource{
		Groups: map[int]bool{2: false},
		ID:     1,
		Name:   "postgres",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   float64(5432),
			"user":   "postgres",
		},
		Paused:             0,
		PauseReason:        "",
		QueueName:          "queries",
		ScheduledQueueName: "scheduled_queries",
		Syntax:             "sql",
		Type:               "pg",
		ViewOnly:           false,
	}, res)
}

func Test_CreateDataSource_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/data_sources", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"name":"postgres","options":{"dbname":"postgres","host":"postgres","port":5432,"user":"postgres"},"type":"pg"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"groups": {
					"2": false
				},
				"id": 1,
				"name": "postgres",
				"options": {
					"dbname": "postgres",
					"host": "postgres",
					"port": 5432,
					"user": "postgres"
				},
				"pause_reason": null,
				"paused": 0,
				"queue_name": "queries",
				"scheduled_queue_name": "scheduled_queries",
				"syntax": "sql",
				"type": "pg"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.CreateDataSource(context.Background(), &redash.CreateDataSourceInput{
		Name: "postgres",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
		Type: "pg",
	})
	assert.NoError(err)
	assert.Equal(&redash.DataSource{
		Groups: map[int]bool{2: false},
		ID:     1,
		Name:   "postgres",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   float64(5432),
			"user":   "postgres",
		},
		Paused:             0,
		PauseReason:        "",
		QueueName:          "queries",
		ScheduledQueueName: "scheduled_queries",
		Syntax:             "sql",
		Type:               "pg",
		ViewOnly:           false,
	}, res)
}

func Test_UpdateDataSource_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/data_sources/1", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"name":"postgres","options":{"dbname":"postgres","host":"postgres","port":5432,"user":"postgres"},"type":"pg"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"groups": {
					"2": false
				},
				"id": 1,
				"name": "postgres",
				"options": {
					"dbname": "postgres",
					"host": "postgres",
					"port": 5432,
					"user": "postgres"
				},
				"pause_reason": null,
				"paused": 0,
				"queue_name": "queries",
				"scheduled_queue_name": "scheduled_queries",
				"syntax": "sql",
				"type": "pg"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.UpdateDataSource(context.Background(), 1, &redash.UpdateDataSourceInput{
		Name: "postgres",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
		Type: "pg",
	})
	assert.NoError(err)
	assert.Equal(&redash.DataSource{
		Groups: map[int]bool{2: false},
		ID:     1,
		Name:   "postgres",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   float64(5432),
			"user":   "postgres",
		},
		Paused:             0,
		PauseReason:        "",
		QueueName:          "queries",
		ScheduledQueueName: "scheduled_queries",
		Syntax:             "sql",
		Type:               "pg",
		ViewOnly:           false,
	}, res)
}

func Test_PauseDataSource_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/data_sources/1/pause", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"reason":"this is reason"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"id": 1,
				"name": "postgres",
				"pause_reason": "this is reason",
				"paused": 1,
				"supports_auto_limit": true,
				"syntax": "sql",
				"type": "pg"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.PauseDataSource(context.Background(), 1, &redash.PauseDataSourceInput{
		Reason: "this is reason",
	})
	assert.NoError(err)
	assert.Equal(&redash.DataSource{
		ID:          1,
		Name:        "postgres",
		Paused:      1,
		PauseReason: "this is reason",
		Syntax:      "sql",
		Type:        "pg",
	}, res)
}

func Test_ResumeDataSource_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/data_sources/1/pause", func(req *http.Request) (*http.Response, error) {
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
				"id": 1,
				"name": "postgres",
				"pause_reason": null,
				"paused": 0,
				"supports_auto_limit": true,
				"syntax": "sql",
				"type": "pg"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ResumeDataSource(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.DataSource{
		ID:          1,
		Name:        "postgres",
		Paused:      0,
		PauseReason: "",
		Syntax:      "sql",
		Type:        "pg",
	}, res)
}

func Test_DeleteDataSource_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/data_sources/1", func(req *http.Request) (*http.Response, error) {
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
	err := client.DeleteDataSource(context.Background(), 1)
	assert.NoError(err)
}

func Test_TestDataSource_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/data_sources/1/test", func(req *http.Request) (*http.Response, error) {
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
				"message": "success",
				"ok": true
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.TestDataSource(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.TestDataSourceOutput{
		Message: "success",
		Ok:      true,
	}, res)
}

func Test_DataSource_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)

	_, err := client.ListDataSources(context.Background())
	assert.NoError(err)

	ds, err := client.CreateDataSource(context.Background(), &redash.CreateDataSourceInput{
		Name: "test-postgres-1",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   testRedashPgHost,
			"port":   5432,
			"user":   "postgres",
		},
	})
	assert.NoError(err)
	assert.Equal("test-postgres-1", ds.Name)

	output, err := client.TestDataSource(context.Background(), ds.ID)
	assert.NoError(err)
	assert.True(output.Ok)

	ds, err = client.GetDataSource(context.Background(), ds.ID)
	assert.NoError(err)
	assert.Equal("test-postgres-1", ds.Name)

	ds, err = client.UpdateDataSource(context.Background(), ds.ID, &redash.UpdateDataSourceInput{
		Name: "test-postgres-2",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   testRedashPgHost,
			"port":   5432,
			"user":   "postgres",
		},
	})
	assert.NoError(err)
	assert.Equal("test-postgres-2", ds.Name)

	ds, err = client.PauseDataSource(context.Background(), ds.ID, &redash.PauseDataSourceInput{
		Reason: "this is reason",
	})
	assert.NoError(err)
	assert.Equal(1, ds.Paused)
	assert.Equal("this is reason", ds.PauseReason)

	ds, err = client.ResumeDataSource(context.Background(), ds.ID)
	assert.NoError(err)
	assert.Equal(0, ds.Paused)
	assert.Equal("", ds.PauseReason)

	err = client.DeleteDataSource(context.Background(), ds.ID)
	assert.NoError(err)

	_, err = client.GetDataSource(context.Background(), ds.ID)
	assert.Error(err)
}
