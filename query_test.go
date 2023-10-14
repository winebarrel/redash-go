package redash_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go"
)

func Test_ListQueries_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal("page=1&page_size=25", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"count": 1,
				"page": 1,
				"page_size": 25,
				"results": [
					{
						"api_key": "api_key",
						"created_at": "2023-02-10T01:23:45.000Z",
						"data_source_id": 1,
						"description": "description",
						"id": 1,
						"is_archived": false,
						"is_draft": false,
						"is_favorite": false,
						"is_safe": true,
						"last_modified_by_id": 1,
						"latest_query_data_id": 1,
						"name": "my-query",
						"options": {
							"parameters": []
						},
						"query": "select 1",
						"query_hash": "query_hash",
						"retrieved_at": "2023-02-10T01:23:45.000Z",
						"runtime": 0.1,
						"schedule": {
							"day_of_week": null,
							"interval": 60,
							"time": null,
							"until": "2023-02-11"
						},
						"tags": [],
						"updated_at": "2023-02-10T01:23:45.000Z",
						"user": {},
						"version": 1
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListQueries(context.Background(), &redash.ListQueriesInput{
		OnlyFavorites: false,
		Page:          1,
		PageSize:      25,
	})
	assert.NoError(err)
	assert.Equal(&redash.QueryPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Query{
			{
				APIKey:            "api_key",
				CanEdit:           false,
				CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DataSourceID:      1,
				Description:       "description",
				ID:                1,
				IsArchived:        false,
				IsDraft:           false,
				IsFavorite:        false,
				IsSafe:            true,
				LastModifiedBy:    nil,
				LastModifiedByID:  1,
				LatestQueryDataID: 1,
				Name:              "my-query",
				Options:           redash.QueryOptions{Parameters: []map[string]any{}},
				Query:             "select 1",
				QueryHash:         "query_hash",
				RetrievedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Runtime:           0.1,
				Schedule: &redash.QueueSchedule{
					DayOfWeek: "",
					Interval:  60,
					Time:      "",
					Until:     "2023-02-11",
				},
				Tags:           []string{},
				UpdatedAt:      dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				User:           redash.User{},
				Version:        1,
				Visualizations: nil,
			},
		},
	}, res)
}

func Test_ListQueries_WithQ(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal("page=1&page_size=25&q=my-query", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"count": 1,
				"page": 1,
				"page_size": 25,
				"results": [
					{
						"api_key": "api_key",
						"created_at": "2023-02-10T01:23:45.000Z",
						"data_source_id": 1,
						"description": "description",
						"id": 1,
						"is_archived": false,
						"is_draft": false,
						"is_favorite": false,
						"is_safe": true,
						"last_modified_by_id": 1,
						"latest_query_data_id": 1,
						"name": "my-query",
						"options": {
							"parameters": []
						},
						"query": "select 1",
						"query_hash": "query_hash",
						"retrieved_at": "2023-02-10T01:23:45.000Z",
						"runtime": 0.1,
						"schedule": {
							"day_of_week": null,
							"interval": 60,
							"time": null,
							"until": "2023-02-11"
						},
						"tags": [],
						"updated_at": "2023-02-10T01:23:45.000Z",
						"user": {},
						"version": 1
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListQueries(context.Background(), &redash.ListQueriesInput{
		OnlyFavorites: false,
		Page:          1,
		PageSize:      25,
		Q:             "my-query",
	})
	assert.NoError(err)
	assert.Equal(&redash.QueryPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Query{
			{
				APIKey:            "api_key",
				CanEdit:           false,
				CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DataSourceID:      1,
				Description:       "description",
				ID:                1,
				IsArchived:        false,
				IsDraft:           false,
				IsFavorite:        false,
				IsSafe:            true,
				LastModifiedBy:    nil,
				LastModifiedByID:  1,
				LatestQueryDataID: 1,
				Name:              "my-query",
				Options:           redash.QueryOptions{Parameters: []map[string]any{}},
				Query:             "select 1",
				QueryHash:         "query_hash",
				RetrievedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Runtime:           0.1,
				Schedule: &redash.QueueSchedule{
					DayOfWeek: "",
					Interval:  60,
					Time:      "",
					Until:     "2023-02-11",
				},
				Tags:           []string{},
				UpdatedAt:      dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				User:           redash.User{},
				Version:        1,
				Visualizations: nil,
			},
		},
	}, res)
}

func Test_GetQuery_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
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
				"api_key": "api_key",
				"can_edit": true,
				"created_at": "2023-02-10T01:23:45.000Z",
				"data_source_id": 1,
				"description": "description",
				"id": 1,
				"is_archived": false,
				"is_draft": false,
				"is_favorite": false,
				"is_safe": true,
				"last_modified_by": {},
				"latest_query_data_id": 1,
				"name": "my-query",
				"options": {
					"parameters": []
				},
				"query": "select 1",
				"query_hash": "query_hash",
				"schedule": {
					"day_of_week": null,
					"interval": 60,
					"time": null,
					"until": "2023-02-11"
				},
				"tags": [],
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {},
				"version": 1,
				"visualizations": [
					{
						"created_at": "2023-02-10T01:23:45.000Z",
						"description": "description",
						"id": 1,
						"name": "Table",
						"options": {},
						"type": "TABLE",
						"updated_at": "2023-02-10T01:23:45.000Z"
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetQuery(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.Query{
		APIKey:            "api_key",
		CanEdit:           true,
		CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DataSourceID:      1,
		Description:       "description",
		ID:                1,
		IsArchived:        false,
		IsDraft:           false,
		IsFavorite:        false,
		IsSafe:            true,
		LastModifiedBy:    &redash.User{},
		LastModifiedByID:  0,
		LatestQueryDataID: 1,
		Name:              "my-query",
		Options:           redash.QueryOptions{Parameters: []map[string]any{}},
		Query:             "select 1",
		QueryHash:         "query_hash",
		RetrievedAt:       time.Time{},
		Runtime:           0,
		Schedule: &redash.QueueSchedule{
			DayOfWeek: "",
			Interval:  60,
			Time:      "",
			Until:     "2023-02-11",
		},
		Tags:      []string{},
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
		Version:   1,
		Visualizations: []redash.Visualization{
			{
				CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Description: "description",
				ID:          1,
				Name:        "Table",
				Options:     map[string]any{},
				Query:       redash.Query{},
				Type:        "TABLE",
				UpdatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			},
		},
	}, res)
}

func Test_CreateQuery_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"data_source_id":1,"description":"description","name":"my-query","query":"select 1","schedule":{"interval":60,"time":null,"until":null,"day_of_week":null}}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"api_key": "api_key",
				"can_edit": true,
				"created_at": "2023-02-10T01:23:45.000Z",
				"data_source_id": 1,
				"description": "description",
				"id": 1,
				"is_archived": false,
				"is_draft": false,
				"is_favorite": false,
				"is_safe": true,
				"last_modified_by": {},
				"latest_query_data_id": 1,
				"name": "my-query",
				"options": {
					"parameters": []
				},
				"query": "select 1",
				"query_hash": "query_hash",
				"schedule": {
					"day_of_week": null,
					"interval": 60,
					"time": null,
					"until": "2023-02-11"
				},
				"tags": [],
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {},
				"version": 1,
				"visualizations": [
					{
						"created_at": "2023-02-10T01:23:45.000Z",
						"description": "description",
						"id": 1,
						"name": "Table",
						"options": {},
						"type": "TABLE",
						"updated_at": "2023-02-10T01:23:45.000Z"
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: 1,
		Description:  "description",
		Name:         "my-query",
		Query:        "select 1",
		Schedule: &redash.CreateQueryInputSchedule{
			Interval: 60,
		},
	})
	assert.NoError(err)
	assert.Equal(&redash.Query{
		APIKey:            "api_key",
		CanEdit:           true,
		CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DataSourceID:      1,
		Description:       "description",
		ID:                1,
		IsArchived:        false,
		IsDraft:           false,
		IsFavorite:        false,
		IsSafe:            true,
		LastModifiedBy:    &redash.User{},
		LastModifiedByID:  0,
		LatestQueryDataID: 1,
		Name:              "my-query",
		Options:           redash.QueryOptions{Parameters: []map[string]any{}},
		Query:             "select 1",
		QueryHash:         "query_hash",
		RetrievedAt:       time.Time{},
		Runtime:           0,
		Schedule: &redash.QueueSchedule{
			DayOfWeek: "",
			Interval:  60,
			Time:      "",
			Until:     "2023-02-11",
		},
		Tags:      []string{},
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
		Version:   1,
		Visualizations: []redash.Visualization{
			{
				CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Description: "description",
				ID:          1,
				Name:        "Table",
				Options:     map[string]any{},
				Query:       redash.Query{},
				Type:        "TABLE",
				UpdatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			},
		},
	}, res)
}

func Test_UpdateQuery_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"data_source_id":1,"description":"description","name":"my-query","query":"select 1","schedule":{"interval":60,"time":null,"until":null,"day_of_week":null}}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"api_key": "api_key",
				"can_edit": true,
				"created_at": "2023-02-10T01:23:45.000Z",
				"data_source_id": 1,
				"description": "description",
				"id": 1,
				"is_archived": false,
				"is_draft": false,
				"is_favorite": false,
				"is_safe": true,
				"last_modified_by": {},
				"latest_query_data_id": 1,
				"name": "my-query",
				"options": {
					"parameters": []
				},
				"query": "select 1",
				"query_hash": "query_hash",
				"schedule": {
					"day_of_week": null,
					"interval": 60,
					"time": null,
					"until": "2023-02-11"
				},
				"tags": [],
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {},
				"version": 1,
				"visualizations": [
					{
						"created_at": "2023-02-10T01:23:45.000Z",
						"description": "description",
						"id": 1,
						"name": "Table",
						"options": {},
						"type": "TABLE",
						"updated_at": "2023-02-10T01:23:45.000Z"
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.UpdateQuery(context.Background(), 1, &redash.UpdateQueryInput{
		DataSourceID: 1,
		Description:  "description",
		Name:         "my-query",
		Query:        "select 1",
		Schedule: &redash.UpdateQueryInputSchedule{
			Interval: 60,
		},
	})
	assert.NoError(err)
	assert.Equal(&redash.Query{
		APIKey:            "api_key",
		CanEdit:           true,
		CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DataSourceID:      1,
		Description:       "description",
		ID:                1,
		IsArchived:        false,
		IsDraft:           false,
		IsFavorite:        false,
		IsSafe:            true,
		LastModifiedBy:    &redash.User{},
		LastModifiedByID:  0,
		LatestQueryDataID: 1,
		Name:              "my-query",
		Options:           redash.QueryOptions{Parameters: []map[string]any{}},
		Query:             "select 1",
		QueryHash:         "query_hash",
		RetrievedAt:       time.Time{},
		Runtime:           0,
		Schedule: &redash.QueueSchedule{
			DayOfWeek: "",
			Interval:  60,
			Time:      "",
			Until:     "2023-02-11",
		},
		Tags:      []string{},
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
		Version:   1,
		Visualizations: []redash.Visualization{
			{
				CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Description: "description",
				ID:          1,
				Name:        "Table",
				Options:     map[string]any{},
				Query:       redash.Query{},
				Type:        "TABLE",
				UpdatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			},
		},
	}, res)
}

func Test_ArchiveQuery_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
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
	err := client.ArchiveQuery(context.Background(), 1)
	assert.NoError(err)
}

func Test_CreateFavoriteQuery_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/1/favorite", func(req *http.Request) (*http.Response, error) {
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
	err := client.CreateFavoriteQuery(context.Background(), 1)
	assert.NoError(err)
}

func Test_ForkQuery_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/1/fork", func(req *http.Request) (*http.Response, error) {
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
				"api_key": "api_key",
				"can_edit": true,
				"created_at": "2023-02-10T01:23:45.000Z",
				"data_source_id": 1,
				"description": "description",
				"id": 1,
				"is_archived": false,
				"is_draft": false,
				"is_favorite": false,
				"is_safe": true,
				"last_modified_by": {},
				"latest_query_data_id": 1,
				"name": "my-query",
				"options": {
					"parameters": []
				},
				"query": "select 1",
				"query_hash": "query_hash",
				"schedule": {
					"day_of_week": null,
					"interval": 60,
					"time": null,
					"until": "2023-02-11"
				},
				"tags": [],
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {},
				"version": 1,
				"visualizations": [
					{
						"created_at": "2023-02-10T01:23:45.000Z",
						"description": "description",
						"id": 1,
						"name": "Table",
						"options": {},
						"type": "TABLE",
						"updated_at": "2023-02-10T01:23:45.000Z"
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ForkQuery(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.Query{
		APIKey:            "api_key",
		CanEdit:           true,
		CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DataSourceID:      1,
		Description:       "description",
		ID:                1,
		IsArchived:        false,
		IsDraft:           false,
		IsFavorite:        false,
		IsSafe:            true,
		LastModifiedBy:    &redash.User{},
		LastModifiedByID:  0,
		LatestQueryDataID: 1,
		Name:              "my-query",
		Options:           redash.QueryOptions{Parameters: []map[string]any{}},
		Query:             "select 1",
		QueryHash:         "query_hash",
		RetrievedAt:       time.Time{},
		Runtime:           0,
		Schedule: &redash.QueueSchedule{
			DayOfWeek: "",
			Interval:  60,
			Time:      "",
			Until:     "2023-02-11",
		},
		Tags:      []string{},
		UpdatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:      redash.User{},
		Version:   1,
		Visualizations: []redash.Visualization{
			{
				CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Description: "description",
				ID:          1,
				Name:        "Table",
				Options:     map[string]any{},
				Query:       redash.Query{},
				Type:        "TABLE",
				UpdatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			},
		},
	}, res)
}

func Test_GetQueryResultsJSON_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1/results.json", func(req *http.Request) (*http.Response, error) {
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
		return httpmock.NewStringResponse(http.StatusOK, `{"foo":"bar"}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	var buf bytes.Buffer
	err := client.GetQueryResultsJSON(context.Background(), 1, &buf)
	assert.NoError(err)
	assert.Equal(`{"foo":"bar"}`, buf.String())
}

func Test_GetQueryResultsCSV_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1/results.csv", func(req *http.Request) (*http.Response, error) {
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
		return httpmock.NewStringResponse(http.StatusOK, `foo,bar`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	var buf bytes.Buffer
	err := client.GetQueryResultsCSV(context.Background(), 1, &buf)
	assert.NoError(err)
	assert.Equal(`foo,bar`, buf.String())
}

func Test_GetQueryResults_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1/results.json", func(req *http.Request) (*http.Response, error) {
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
		return httpmock.NewStringResponse(http.StatusOK, `{"foo":"bar"}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	var buf bytes.Buffer
	err := client.GetQueryResults(context.Background(), 1, "json", &buf)
	assert.NoError(err)
	assert.Equal(`{"foo":"bar"}`, buf.String())
}

func Test_GetQueryResults_OK_WithNil(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1/results.json", func(req *http.Request) (*http.Response, error) {
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
		return httpmock.NewStringResponse(http.StatusOK, `{"foo":"bar"}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.GetQueryResults(context.Background(), 1, "json", nil)
	assert.ErrorContains(err, "out(io.Writer) is nil")
}

func Test_ExecQueryJSON_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/1/results", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"filetype":"json"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `{"foo":"bar"}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	var buf bytes.Buffer
	jobId, err := client.ExecQueryJSON(context.Background(), 1, &buf)
	assert.NoError(err)
	assert.Equal(`{"foo":"bar"}`, buf.String())
	assert.Empty(jobId)
}

func Test_ExecQueryJSON_OK_WithNil(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/1/results", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"filetype":"json"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `{"foo":"bar"}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	jobId, err := client.ExecQueryJSON(context.Background(), 1, nil)
	assert.NoError(err)
	assert.Empty(jobId)
}

func Test_ExecQueryJSON_ReturnJob(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/1/results", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"filetype":"json"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `{"job": {"status": 1, "error": "", "id": "623b290a-7fd9-4ea6-a2a6-96f9c9101f51", "query_result_id": null,	"status": 1, "updated_at": 0}}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	job, err := client.ExecQueryJSON(context.Background(), 1, nil)
	assert.NoError(err)
	assert.Equal(&redash.JobResponse{
		Job: redash.Job{
			Error:         "",
			ID:            "623b290a-7fd9-4ea6-a2a6-96f9c9101f51",
			QueryResultID: 0,
			Status:        1,
			UpdatedAt:     float64(0),
		},
	}, job)
}

func Test_GetQueryTags_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/tags", func(req *http.Request) (*http.Response, error) {
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
				"tags": [
					{
						"count": 1,
						"name": "my-tag"
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetQueryTags(context.Background())
	assert.NoError(err)
	assert.Equal(&redash.QueryTags{
		Tags: []redash.QueryTagsTag{
			{
				Name:  "my-tag",
				Count: 1,
			},
		},
	}, res)
}

func Test_RefreshQuery_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/1/refresh", func(req *http.Request) (*http.Response, error) {
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
				"job": {
					"error": "",
					"id": "baaf5b97-6419-4db3-a60c-ef8b4e290376",
					"query_result_id": null,
					"result": null,
					"status": 1,
					"updated_at": 0
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	job, err := client.RefreshQuery(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.JobResponse{
		Job: redash.Job{
			Error:         "",
			ID:            "baaf5b97-6419-4db3-a60c-ef8b4e290376",
			QueryResultID: 0,
			Status:        1,
			UpdatedAt:     float64(0),
		},
	}, job)
}

func Test_SearchQueries_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/search", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal("q=my-query", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"count": 1,
				"page": 1,
				"page_size": 25,
				"results": [
					{
						"api_key": "api_key",
						"created_at": "2023-02-10T01:23:45.000Z",
						"data_source_id": 1,
						"description": "description",
						"id": 1,
						"is_archived": false,
						"is_draft": false,
						"is_favorite": false,
						"is_safe": true,
						"last_modified_by_id": 1,
						"latest_query_data_id": 1,
						"name": "my-query",
						"options": {
							"parameters": []
						},
						"query": "select 1",
						"query_hash": "query_hash",
						"retrieved_at": "2023-02-10T01:23:45.000Z",
						"runtime": 0.1,
						"schedule": {
							"day_of_week": null,
							"interval": 60,
							"time": null,
							"until": "2023-02-11"
						},
						"tags": [],
						"updated_at": "2023-02-10T01:23:45.000Z",
						"user": {},
						"version": 1
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.SearchQueries(context.Background(), &redash.SearchQueriesInput{
		Q: "my-query",
	})
	assert.NoError(err)
	assert.Equal(&redash.QueryPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Query{
			{
				APIKey:            "api_key",
				CanEdit:           false,
				CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DataSourceID:      1,
				Description:       "description",
				ID:                1,
				IsArchived:        false,
				IsDraft:           false,
				IsFavorite:        false,
				IsSafe:            true,
				LastModifiedBy:    nil,
				LastModifiedByID:  1,
				LatestQueryDataID: 1,
				Name:              "my-query",
				Options:           redash.QueryOptions{Parameters: []map[string]any{}},
				Query:             "select 1",
				QueryHash:         "query_hash",
				RetrievedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Runtime:           0.1,
				Schedule: &redash.QueueSchedule{
					DayOfWeek: "",
					Interval:  60,
					Time:      "",
					Until:     "2023-02-11",
				},
				Tags:           []string{},
				UpdatedAt:      dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				User:           redash.User{},
				Version:        1,
				Visualizations: nil,
			},
		},
	}, res)
}

func Test_ListMyQueries_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/my", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal("page=1&page_size=25&q=my-query", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"count": 1,
				"page": 1,
				"page_size": 25,
				"results": [
					{
						"api_key": "api_key",
						"created_at": "2023-02-10T01:23:45.000Z",
						"data_source_id": 1,
						"description": "description",
						"id": 1,
						"is_archived": false,
						"is_draft": false,
						"is_favorite": false,
						"is_safe": true,
						"last_modified_by_id": 1,
						"latest_query_data_id": 1,
						"name": "my-query",
						"options": {
							"parameters": []
						},
						"query": "select 1",
						"query_hash": "query_hash",
						"retrieved_at": "2023-02-10T01:23:45.000Z",
						"runtime": 0.1,
						"schedule": {
							"day_of_week": null,
							"interval": 60,
							"time": null,
							"until": "2023-02-11"
						},
						"tags": [],
						"updated_at": "2023-02-10T01:23:45.000Z",
						"user": {},
						"version": 1
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListMyQueries(context.Background(), &redash.ListMyQueriesInput{
		Page:     1,
		PageSize: 25,
		Q:        "my-query",
	})
	assert.NoError(err)
	assert.Equal(&redash.QueryPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Query{
			{
				APIKey:            "api_key",
				CanEdit:           false,
				CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DataSourceID:      1,
				Description:       "description",
				ID:                1,
				IsArchived:        false,
				IsDraft:           false,
				IsFavorite:        false,
				IsSafe:            true,
				LastModifiedBy:    nil,
				LastModifiedByID:  1,
				LatestQueryDataID: 1,
				Name:              "my-query",
				Options:           redash.QueryOptions{Parameters: []map[string]any{}},
				Query:             "select 1",
				QueryHash:         "query_hash",
				RetrievedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Runtime:           0.1,
				Schedule: &redash.QueueSchedule{
					DayOfWeek: "",
					Interval:  60,
					Time:      "",
					Until:     "2023-02-11",
				},
				Tags:           []string{},
				UpdatedAt:      dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				User:           redash.User{},
				Version:        1,
				Visualizations: nil,
			},
		},
	}, res)
}

func Test_ListFavoriteQueries_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/favorites", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal("page=1&page_size=25&q=my-query", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"count": 1,
				"page": 1,
				"page_size": 25,
				"results": [
					{
						"api_key": "api_key",
						"created_at": "2023-02-10T01:23:45.000Z",
						"data_source_id": 1,
						"description": "description",
						"id": 1,
						"is_archived": false,
						"is_draft": false,
						"is_favorite": false,
						"is_safe": true,
						"last_modified_by_id": 1,
						"latest_query_data_id": 1,
						"name": "my-query",
						"options": {
							"parameters": []
						},
						"query": "select 1",
						"query_hash": "query_hash",
						"retrieved_at": "2023-02-10T01:23:45.000Z",
						"runtime": 0.1,
						"schedule": {
							"day_of_week": null,
							"interval": 60,
							"time": null,
							"until": "2023-02-11"
						},
						"tags": [],
						"updated_at": "2023-02-10T01:23:45.000Z",
						"user": {},
						"version": 1
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListFavoriteQueries(context.Background(), &redash.ListFavoriteQueriesInput{
		Page:     1,
		PageSize: 25,
		Q:        "my-query",
	})
	assert.NoError(err)
	assert.Equal(&redash.QueryPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Query{
			{
				APIKey:            "api_key",
				CanEdit:           false,
				CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DataSourceID:      1,
				Description:       "description",
				ID:                1,
				IsArchived:        false,
				IsDraft:           false,
				IsFavorite:        false,
				IsSafe:            true,
				LastModifiedBy:    nil,
				LastModifiedByID:  1,
				LatestQueryDataID: 1,
				Name:              "my-query",
				Options:           redash.QueryOptions{Parameters: []map[string]any{}},
				Query:             "select 1",
				QueryHash:         "query_hash",
				RetrievedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Runtime:           0.1,
				Schedule: &redash.QueueSchedule{
					DayOfWeek: "",
					Interval:  60,
					Time:      "",
					Until:     "2023-02-11",
				},
				Tags:           []string{},
				UpdatedAt:      dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				User:           redash.User{},
				Version:        1,
				Visualizations: nil,
			},
		},
	}, res)
}

func Test_FormatQuery_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/format", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"query":"select 1 from dual"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"query": "SELECT 1\nFROM dual"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.FormatQuery(context.Background(), "select 1 from dual")
	assert.NoError(err)
	assert.Equal(&redash.FormatQueryOutput{
		Query: "SELECT 1\nFROM dual",
	}, res)
}

func Test_ListRecentQueries_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/recent", func(req *http.Request) (*http.Response, error) {
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
					"api_key": "api_key",
					"created_at": "2023-02-10T01:23:45.000Z",
					"data_source_id": 1,
					"description": "description",
					"id": 1,
					"is_archived": false,
					"is_draft": false,
					"is_favorite": false,
					"is_safe": true,
					"last_modified_by_id": 1,
					"latest_query_data_id": 1,
					"name": "my-query",
					"options": {
						"parameters": []
					},
					"query": "select 1",
					"query_hash": "query_hash",
					"retrieved_at": "2023-02-10T01:23:45.000Z",
					"runtime": 0.1,
					"schedule": {
						"day_of_week": null,
						"interval": 60,
						"time": null,
						"until": "2023-02-11"
					},
					"tags": [],
					"updated_at": "2023-02-10T01:23:45.000Z",
					"user": {},
					"version": 1
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListRecentQueries(context.Background())
	assert.NoError(err)
	assert.Equal([]redash.Query{
		{
			APIKey:            "api_key",
			CanEdit:           false,
			CreatedAt:         dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			DataSourceID:      1,
			Description:       "description",
			ID:                1,
			IsArchived:        false,
			IsDraft:           false,
			IsFavorite:        false,
			IsSafe:            true,
			LastModifiedBy:    nil,
			LastModifiedByID:  1,
			LatestQueryDataID: 1,
			Name:              "my-query",
			Options:           redash.QueryOptions{Parameters: []map[string]any{}},
			Query:             "select 1",
			QueryHash:         "query_hash",
			RetrievedAt:       dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			Runtime:           0.1,
			Schedule: &redash.QueueSchedule{
				DayOfWeek: "",
				Interval:  60,
				Time:      "",
				Until:     "2023-02-11",
			},
			Tags:           []string{},
			UpdatedAt:      dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			User:           redash.User{},
			Version:        1,
			Visualizations: nil,
		},
	}, res)
}
func Test_Query_Acc(t *testing.T) {
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

	_, err = client.ListQueries(context.Background(), nil)
	require.NoError(err)

	query, err := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "test-query-1",
		Query:        "select 1",
		Tags:         []string{"my-tag-1"},
	})
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)
	assert.Equal([]string{"my-tag-1"}, query.Tags)

	query, err = client.UpdateQuery(context.Background(), query.ID, &redash.UpdateQueryInput{
		Schedule: &redash.UpdateQueryInputSchedule{
			Interval: 600,
		},
		Tags: &[]string{"my-tag-2"},
	})
	require.NoError(err)
	assert.Equal(&redash.QueueSchedule{Interval: 600}, query.Schedule)
	assert.Equal([]string{"my-tag-2"}, query.Tags)

	query, err = client.UpdateQuery(context.Background(), query.ID, &redash.UpdateQueryInput{
		Schedule: &redash.UpdateQueryInputSchedule{
			Interval: 600,
		},
	})
	require.NoError(err)
	assert.Equal(&redash.QueueSchedule{Interval: 600}, query.Schedule)
	assert.Equal([]string{"my-tag-2"}, query.Tags)

	tags, err := client.GetQueryTags(context.Background())
	require.NoError(err)
	assert.GreaterOrEqual(len(tags.Tags), 1)

	query, err = client.UpdateQuery(context.Background(), query.ID, &redash.UpdateQueryInput{
		Tags: &[]string{},
	})
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)
	assert.Equal(&redash.QueueSchedule{Interval: 600}, query.Schedule)
	assert.Equal([]string{}, query.Tags)

	tags, err = client.GetQueryTags(context.Background())
	require.NoError(err)
	assert.GreaterOrEqual(len(tags.Tags), 0)

	query, err = client.GetQuery(context.Background(), query.ID)
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)

	page, err := client.SearchQueries(context.Background(), &redash.SearchQueriesInput{
		Q: "test-query-1",
	})
	require.NoError(err)
	assert.GreaterOrEqual(len(page.Results), 1)

	page, err = client.ListQueries(context.Background(), &redash.ListQueriesInput{
		Q: "test-query-1",
	})
	require.NoError(err)
	assert.GreaterOrEqual(len(page.Results), 1)

	page, err = client.ListMyQueries(context.Background(), &redash.ListMyQueriesInput{
		Q: "test-query-1",
	})
	require.NoError(err)
	assert.GreaterOrEqual(len(page.Results), 1)

	page, err = client.ListFavoriteQueries(context.Background(), &redash.ListFavoriteQueriesInput{
		Q: "test-query-1",
	})
	require.NoError(err)
	assert.Zero(len(page.Results))

	err = client.CreateFavoriteQuery(context.Background(), query.ID)
	require.NoError(err)

	page, err = client.ListFavoriteQueries(context.Background(), &redash.ListFavoriteQueriesInput{
		Q: "test-query-1",
	})
	require.NoError(err)
	assert.GreaterOrEqual(len(page.Results), 1)

	var buf bytes.Buffer
	job, err := client.ExecQueryJSON(context.Background(), query.ID, &buf)
	require.NoError(err)

	if job != nil && job.Job.ID != "" {
		for {
			job, err := client.GetJob(context.Background(), job.Job.ID)
			require.NoError(err)

			if job.Job.Status != redash.JobStatusPending && job.Job.Status != redash.JobStatusStarted {
				assert.Equal(redash.JobStatusSuccess, job.Job.Status)
				buf = bytes.Buffer{}
				err = client.GetQueryResultsJSON(context.Background(), query.ID, &buf)
				require.NoError(err)
				break
			}

			time.Sleep(1 * time.Second)
		}
	}

	assert.True(strings.HasPrefix(buf.String(), `{"query_result"`))

	_, err = client.ExecQueryJSON(context.Background(), query.ID, nil)
	require.NoError(err)

	buf = bytes.Buffer{}
	err = client.GetQueryResultsCSV(context.Background(), query.ID, &buf)
	require.NoError(err)
	assert.Equal("?column?\r\n1\r\n", buf.String())

	job, err = client.RefreshQuery(context.Background(), query.ID)
	require.NoError(err)

	if job != nil && job.Job.ID != "" {
		for {
			job, err := client.GetJob(context.Background(), job.Job.ID)
			require.NoError(err)

			if job.Job.Status != redash.JobStatusPending && job.Job.Status != redash.JobStatusStarted {
				assert.Equal(redash.JobStatusSuccess, job.Job.Status)
				buf = bytes.Buffer{}
				err = client.GetQueryResultsJSON(context.Background(), query.ID, &buf)
				require.NoError(err)
				break
			}

			time.Sleep(1 * time.Second)
		}
	}

	assert.True(strings.HasPrefix(buf.String(), `{"query_result"`))

	queries, err := client.ListRecentQueries(context.Background())
	require.NoError(err)
	assert.GreaterOrEqual(len(queries), 1)

	err = client.ArchiveQuery(context.Background(), query.ID)
	require.NoError(err)

	query, err = client.GetQuery(context.Background(), query.ID)
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)
	assert.True(query.IsArchived)
	assert.True(query.IsFavorite)

	formatted, err := client.FormatQuery(context.Background(), "select 1 from dual")
	require.NoError(err)
	assert.Equal("SELECT 1\nFROM dual", formatted.Query)
}
