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

func Test_ListAlerts_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/alerts", func(req *http.Request) (*http.Response, error) {
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
					"user": {},
					"state": "unknown",
					"name": "name",
					"rearm": 1,
					"updated_at": "2023-02-10T01:23:45.000Z",
					"query": {},
					"created_at": "2023-02-10T01:23:45.000Z",
					"last_triggered_at": "2023-02-10T01:23:45.000Z",
					"id": 1,
					"options": {
						"column": "col",
						"value": 0,
						"op": "greater than"
					}
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListAlerts(context.Background())
	assert.NoError(err)
	assert.Equal([]redash.Alert{
		{
			CreatedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			ID:              1,
			LastTriggeredAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			Name:            "name",
			Options: redash.AlertOptions{
				Column: "col",
				Value:  0,
				Op:     "greater than",
			},
			Query:     redash.Query{},
			Rearm:     1,
			State:     "unknown",
			UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			User:      redash.User{},
		},
	}, res)
}

func Test_GetAlert_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/alerts/1", func(req *http.Request) (*http.Response, error) {
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
				"user": {},
				"state": "unknown",
				"name": "name",
				"rearm": 1,
				"updated_at": "2023-02-10T01:23:45.000Z",
				"query": {},
				"created_at": "2023-02-10T01:23:45.000Z",
				"last_triggered_at": "2023-02-10T01:23:45.000Z",
				"id": 1,
				"options": {
					"column": "col",
					"value": 0,
					"op": "greater than"
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetAlert(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.Alert{
		CreatedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		ID:              1,
		LastTriggeredAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		Name:            "name",
		Options: redash.AlertOptions{
			Column: "col",
			Value:  0,
			Op:     "greater than",
		},
		Query:     redash.Query{},
		Rearm:     1,
		State:     "unknown",
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
	}, res)
}

func Test_CreateAlert_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"name":"name","options":{"column":"col","op":"greater than","value":0},"query_id":1,"rearm":1}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
  		{
				"user": {},
				"state": "unknown",
				"name": "name",
				"rearm": 1,
				"updated_at": "2023-02-10T01:23:45.000Z",
				"query": {},
				"created_at": "2023-02-10T01:23:45.000Z",
				"last_triggered_at": "2023-02-10T01:23:45.000Z",
				"id": 1,
				"options": {
					"column": "col",
					"value": 0,
					"op": "greater than"
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.CreateAlert(context.Background(), &redash.CreateAlertInput{
		Name: "name",
		Options: redash.CreateAlertOptions{
			Column: "col",
			Value:  0,
			Op:     "greater than",
		},
		QueryId: 1,
		Rearm:   1,
	})
	assert.NoError(err)
	assert.Equal(&redash.Alert{
		CreatedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		ID:              1,
		LastTriggeredAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		Name:            "name",
		Options: redash.AlertOptions{
			Column: "col",
			Value:  0,
			Op:     "greater than",
		},
		Query:     redash.Query{},
		Rearm:     1,
		State:     "unknown",
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
	}, res)
}

func Test_UpdateAlert_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts/1", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"name":"name","options":{"column":"col","value":0,"op":"greater than"},"query_id":1,"rearm":1}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
  		{
				"user": {},
				"state": "unknown",
				"name": "name",
				"rearm": 1,
				"updated_at": "2023-02-10T01:23:45.000Z",
				"query": {},
				"created_at": "2023-02-10T01:23:45.000Z",
				"last_triggered_at": "2023-02-10T01:23:45.000Z",
				"id": 1,
				"options": {
					"column": "col",
					"value": 0,
					"op": "greater than"
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.UpdateAlert(context.Background(), 1, &redash.UpdateAlertInput{
		Name: "name",
		Options: &redash.UpdateAlertOptions{
			Column: "col",
			Value:  0,
			Op:     "greater than",
		},
		QueryId: 1,
		Rearm:   1,
	})
	assert.NoError(err)
	assert.Equal(&redash.Alert{
		CreatedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		ID:              1,
		LastTriggeredAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		Name:            "name",
		Options: redash.AlertOptions{
			Column: "col",
			Value:  0,
			Op:     "greater than",
		},
		Query:     redash.Query{},
		Rearm:     1,
		State:     "unknown",
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
	}, res)
}

func Test_DeleteAlert_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/alerts/1", func(req *http.Request) (*http.Response, error) {
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
	err := client.DeleteAlert(context.Background(), 1)
	assert.NoError(err)
}

func Test_ListAlertSubscriptions_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/alerts/1/subscriptions", func(req *http.Request) (*http.Response, error) {
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
					"alert_id": 1,
					"destination": {
						"icon": "fa-envelope",
						"id": 1,
						"name": "admin@example.com",
						"type": "email"
					},
					"id": 1,
					"user": {}
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListAlertSubscriptions(context.Background(), 1)
	assert.NoError(err)
	assert.Equal([]redash.AlertSubscription{
		{
			AlertID: 1,
			Destination: &redash.Destination{
				Icon:    "fa-envelope",
				ID:      1,
				Name:    "admin@example.com",
				Options: nil,
				Type:    "email",
			},
			ID:   1,
			User: redash.User{},
		},
	}, res)
}

func Test_AddAlertSubscription_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts/1/subscriptions", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"destination_id":1}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"alert_id": 1,
				"destination": {
					"icon": "fa-envelope",
					"id": 1,
					"name": "admin@example.com",
					"type": "email"
				},
				"id": 1,
				"user": {}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.AddAlertSubscription(context.Background(), 1, 1)
	assert.NoError(err)
	assert.Equal(&redash.AlertSubscription{
		AlertID: 1,
		Destination: &redash.Destination{
			Icon:    "fa-envelope",
			ID:      1,
			Name:    "admin@example.com",
			Options: nil,
			Type:    "email",
		},
		ID:   1,
		User: redash.User{},
	}, res)
}

func Test_RemoveAlertSubscription_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/alerts/1/subscriptions/2", func(req *http.Request) (*http.Response, error) {
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
	err := client.RemoveAlertSubscription(context.Background(), 1, 2)
	assert.NoError(err)
}

func Test_MuteAlert_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts/1/mute", func(req *http.Request) (*http.Response, error) {
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
	err := client.MuteAlert(context.Background(), 1)
	assert.NoError(err)
}

func Test_UnmuteAlert_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/alerts/1/mute", func(req *http.Request) (*http.Response, error) {
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
	err := client.UnmuteAlert(context.Background(), 1)
	assert.NoError(err)
}

func Test_Alert_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
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
	if err != nil {
		assert.FailNow(err.Error())
	}

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

	_, err = client.ListAlerts(context.Background())
	assert.NoError(err)

	alert, err := client.CreateAlert(context.Background(), &redash.CreateAlertInput{
		Name: "test-alert-1",
		Options: redash.CreateAlertOptions{
			Column: "col",
			Op:     "greater than",
			Value:  0,
		},
		QueryId: query.ID,
		Rearm:   0,
	})
	assert.NoError(err)
	assert.Equal("test-alert-1", alert.Name)

	alert, err = client.GetAlert(context.Background(), alert.ID)
	assert.NoError(err)
	assert.Equal("test-alert-1", alert.Name)

	alert, err = client.UpdateAlert(context.Background(), alert.ID, &redash.UpdateAlertInput{
		Name: "test-alert-2",
	})
	assert.NoError(err)
	assert.Equal("test-alert-2", alert.Name)

	_, err = client.ListAlertSubscriptions(context.Background(), alert.ID)
	assert.NoError(err)

	err = client.MuteAlert(context.Background(), alert.ID)
	assert.NoError(err)

	err = client.UnmuteAlert(context.Background(), alert.ID)
	assert.NoError(err)

	err = client.DeleteAlert(context.Background(), alert.ID)
	assert.NoError(err)

	_, err = client.GetAlert(context.Background(), alert.ID)
	assert.Error(err)
}
