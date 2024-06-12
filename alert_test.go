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
						"op": "greater than",
						"custom_subject": "custom_subject",
						"custom_body": "custom_body"
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
				Column:        "col",
				Value:         0,
				Op:            "greater than",
				CustomSubject: "custom_subject",
				CustomBody:    "custom_body",
			},
			Query:     redash.Query{},
			Rearm:     1,
			State:     "unknown",
			UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			User:      redash.User{},
		},
	}, res)
}

func Test_ListAlerts_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/alerts", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListAlerts(context.Background())
	assert.ErrorContains(err, "GET api/alerts failed: HTTP status code not OK: 503\nerror")
}

func Test_ListAlerts_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/alerts", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListAlerts(context.Background())
	assert.ErrorContains(err, "Read response body failed: IO error")
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
					"op": "greater than",
					"custom_subject": "custom_subject",
					"custom_body": "custom_body",
					"muted": true
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
			Column:        "col",
			Value:         0,
			Op:            "greater than",
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
			Muted:         true,
		},
		Query:     redash.Query{},
		Rearm:     1,
		State:     "unknown",
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
	}, res)
}

func Test_GetAlert_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/alerts/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetAlert(context.Background(), 1)
	assert.ErrorContains(err, "GET api/alerts/1 failed: HTTP status code not OK: 503\nerror")
}

func Test_GetAlert_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/alerts/1", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetAlert(context.Background(), 1)
	assert.ErrorContains(err, "Read response body failed: IO error")
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
		assert.Equal(`{"name":"name","options":{"column":"col","op":"greater than","value":0,"custom_subject":"custom_subject","custom_body":"custom_body"},"query_id":1,"rearm":1}`, string(body))
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
					"op": "greater than",
					"custom_subject": "custom_subject",
					"custom_body": "custom_body"
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.CreateAlert(context.Background(), &redash.CreateAlertInput{
		Name: "name",
		Options: redash.CreateAlertOptions{
			Column:        "col",
			Value:         0,
			Op:            "greater than",
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
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
			Column:        "col",
			Value:         0,
			Op:            "greater than",
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
		},
		Query:     redash.Query{},
		Rearm:     1,
		State:     "unknown",
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
	}, res)
}

func Test_CreateAlert_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.CreateAlert(context.Background(), &redash.CreateAlertInput{
		Name: "name",
		Options: redash.CreateAlertOptions{
			Column:        "col",
			Value:         0,
			Op:            "greater than",
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
		},
		QueryId: 1,
		Rearm:   1,
	})
	assert.ErrorContains(err, "POST api/alerts failed: HTTP status code not OK: 503\nerror")
}

func Test_CreateAlert_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.CreateAlert(context.Background(), &redash.CreateAlertInput{
		Name: "name",
		Options: redash.CreateAlertOptions{
			Column:        "col",
			Value:         0,
			Op:            "greater than",
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
		},
		QueryId: 1,
		Rearm:   1,
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
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
		assert.Equal(`{"name":"name","options":{"column":"col","value":0,"op":"greater than","custom_subject":"custom_subject","custom_body":"custom_body"},"query_id":1,"rearm":1}`, string(body))
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
					"op": "greater than",
					"custom_subject": "custom_subject",
					"custom_body": "custom_body"
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.UpdateAlert(context.Background(), 1, &redash.UpdateAlertInput{
		Name: "name",
		Options: &redash.UpdateAlertOptions{
			Column:        "col",
			Value:         0,
			Op:            "greater than",
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
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
			Column:        "col",
			Value:         0,
			Op:            "greater than",
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
		},
		Query:     redash.Query{},
		Rearm:     1,
		State:     "unknown",
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
	}, res)
}

func Test_UpdateAlert_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.UpdateAlert(context.Background(), 1, &redash.UpdateAlertInput{
		Name: "name",
		Options: &redash.UpdateAlertOptions{
			Column:        "col",
			Value:         0,
			Op:            "greater than",
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
		},
		QueryId: 1,
		Rearm:   1,
	})
	assert.ErrorContains(err, "POST api/alerts/1 failed: HTTP status code not OK: 503\nerror")
}

func Test_UpdateAlert_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts/1", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.UpdateAlert(context.Background(), 1, &redash.UpdateAlertInput{
		Name: "name",
		Options: &redash.UpdateAlertOptions{
			Column:        "col",
			Value:         0,
			Op:            "greater than",
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
		},
		QueryId: 1,
		Rearm:   1,
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
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

func Test_DeleteAlert_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/alerts/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.DeleteAlert(context.Background(), 1)
	assert.ErrorContains(err, "DELETE api/alerts/1 failed: HTTP status code not OK: 503\nerror")
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

func Test_ListAlertSubscriptions_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/alerts/1/subscriptions", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListAlertSubscriptions(context.Background(), 1)
	assert.ErrorContains(err, "GET api/alerts/1/subscriptions failed: HTTP status code not OK: 503\nerror")
}

func Test_ListAlertSubscriptions_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/alerts/1/subscriptions", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListAlertSubscriptions(context.Background(), 1)
	assert.ErrorContains(err, "Read response body failed: IO error")
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

func Test_AddAlertSubscription_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts/1/subscriptions", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.AddAlertSubscription(context.Background(), 1, 1)
	assert.ErrorContains(err, "POST api/alerts/1/subscriptions failed: HTTP status code not OK: 503\nerror")
}

func Test_AddAlertSubscription_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts/1/subscriptions", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.AddAlertSubscription(context.Background(), 1, 1)
	assert.ErrorContains(err, "Read response body failed: IO error")
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

func Test_RemoveAlertSubscription_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/alerts/1/subscriptions/2", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.RemoveAlertSubscription(context.Background(), 1, 2)
	assert.ErrorContains(err, "DELETE api/alerts/1/subscriptions/2 failed: HTTP status code not OK: 503\nerror")
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

func Test_MuteAlert_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/alerts/1/mute", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.MuteAlert(context.Background(), 1)
	assert.ErrorContains(err, "POST api/alerts/1/mute failed: HTTP status code not OK: 503\nerror")
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

func Test_UnmuteAlert_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/alerts/1/mute", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.UnmuteAlert(context.Background(), 1)
	assert.ErrorContains(err, "DELETE api/alerts/1/mute failed: HTTP status code not OK: 503\nerror")
}

func Test_Alert_Acc(t *testing.T) {
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

	_, err = client.ListAlerts(context.Background())
	require.NoError(err)

	alert, err := client.CreateAlert(context.Background(), &redash.CreateAlertInput{
		Name: "test-alert-1",
		Options: redash.CreateAlertOptions{
			Column:        "col",
			Op:            "greater than",
			Value:         1,
			CustomSubject: "custom_subject",
			CustomBody:    "custom_body",
		},
		QueryId: query.ID,
		Rearm:   0,
	})
	require.NoError(err)
	assert.Equal("test-alert-1", alert.Name)

	alert, err = client.GetAlert(context.Background(), alert.ID)
	require.NoError(err)
	assert.Equal("test-alert-1", alert.Name)
	assert.Equal(0, alert.Rearm)
	assert.Equal("col", alert.Options.Column)
	assert.Equal("greater than", alert.Options.Op)
	assert.Equal(1, alert.Options.Value)
	assert.Equal("custom_subject", alert.Options.CustomSubject)
	assert.Equal("custom_body", alert.Options.CustomBody)
	assert.False(alert.Options.Muted)

	alert, err = client.UpdateAlert(context.Background(), alert.ID, &redash.UpdateAlertInput{
		Name: "test-alert-2",
	})
	require.NoError(err)
	assert.Equal("test-alert-2", alert.Name)
	assert.Equal(0, alert.Rearm)
	assert.Equal("col", alert.Options.Column)
	assert.Equal("greater than", alert.Options.Op)
	assert.Equal(1, alert.Options.Value)
	assert.Equal("custom_subject", alert.Options.CustomSubject)
	assert.Equal("custom_body", alert.Options.CustomBody)

	alert, err = client.UpdateAlert(context.Background(), alert.ID, &redash.UpdateAlertInput{
		Name:  "test-alert-3",
		Rearm: 300,
		Options: &redash.UpdateAlertOptions{
			Column: "col",
			Op:     "greater than",
			Value:  2,
		},
	})
	require.NoError(err)
	assert.Equal("test-alert-3", alert.Name)
	assert.Equal(300, alert.Rearm)
	assert.Equal("col", alert.Options.Column)
	assert.Equal("greater than", alert.Options.Op)
	assert.Equal(2, alert.Options.Value)
	assert.Empty(alert.Options.CustomSubject)
	assert.Empty(alert.Options.CustomBody)

	alert, err = client.UpdateAlert(context.Background(), alert.ID, &redash.UpdateAlertInput{
		Name:  "test-alert-3",
		Rearm: 0,
	})

	require.NoError(err)
	assert.Equal("test-alert-3", alert.Name)
	assert.Equal(0, alert.Rearm)

	_, err = client.ListAlertSubscriptions(context.Background(), alert.ID)
	require.NoError(err)

	err = client.MuteAlert(context.Background(), alert.ID)
	require.NoError(err)
	alert, err = client.GetAlert(context.Background(), alert.ID)
	require.NoError(err)
	assert.True(alert.Options.Muted)

	err = client.UnmuteAlert(context.Background(), alert.ID)
	require.NoError(err)
	alert, err = client.GetAlert(context.Background(), alert.ID)
	require.NoError(err)
	assert.False(alert.Options.Muted)

	err = client.DeleteAlert(context.Background(), alert.ID)
	require.NoError(err)

	_, err = client.GetAlert(context.Background(), alert.ID)
	assert.Error(err)
}
