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

func Test_ListDashboards_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards", func(req *http.Request) (*http.Response, error) {
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
						"created_at": "2023-02-10T01:23:45.000Z",
						"dashboard_filters_enabled": false,
						"id": 1,
						"is_archived": false,
						"is_draft": false,
						"is_favorite": false,
						"layout": [],
						"name": "name",
						"slug": "name",
						"tags": [],
						"updated_at": "2023-02-10T01:23:45.000Z",
						"user": {},
						"user_id": 1,
						"version": 2,
						"widgets": null
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListDashboards(context.Background(), &redash.ListDashboardsInput{
		OnlyFavorites: false,
		Page:          1,
		PageSize:      25,
	})
	assert.NoError(err)
	assert.Equal(&redash.DashboardPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Dashboard{
			{
				CanEdit:                 false,
				CreatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DashboardFiltersEnabled: false,
				ID:                      1,
				IsArchived:              false,
				IsDraft:                 false,
				IsFavorite:              false,
				Layout:                  []any{},
				Name:                    "name",
				Slug:                    "name",
				Tags:                    []string{},
				UpdatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				User:                    redash.User{},
				UserID:                  1,
				Version:                 2,
				Widgets:                 nil,
			},
		},
	}, res)
}

func Test_ListDashboards_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListDashboards(context.Background(), &redash.ListDashboardsInput{
		OnlyFavorites: false,
		Page:          1,
		PageSize:      25,
	})
	assert.ErrorContains(err, "GET api/dashboards failed: HTTP status code not OK: 503\nerror")
}

func Test_ListDashboards_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListDashboards(context.Background(), &redash.ListDashboardsInput{
		OnlyFavorites: false,
		Page:          1,
		PageSize:      25,
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_ListDashboards_WithQ(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal("page=1&page_size=25&q=name", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"count": 1,
				"page": 1,
				"page_size": 25,
				"results": [
					{
						"created_at": "2023-02-10T01:23:45.000Z",
						"dashboard_filters_enabled": false,
						"id": 1,
						"is_archived": false,
						"is_draft": false,
						"is_favorite": false,
						"layout": [],
						"name": "name",
						"slug": "name",
						"tags": [],
						"updated_at": "2023-02-10T01:23:45.000Z",
						"user": {},
						"user_id": 1,
						"version": 2,
						"widgets": null
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListDashboards(context.Background(), &redash.ListDashboardsInput{
		OnlyFavorites: false,
		Page:          1,
		PageSize:      25,
		Q:             "name",
	})
	assert.NoError(err)
	assert.Equal(&redash.DashboardPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Dashboard{
			{
				CanEdit:                 false,
				CreatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DashboardFiltersEnabled: false,
				ID:                      1,
				IsArchived:              false,
				IsDraft:                 false,
				IsFavorite:              false,
				Layout:                  []any{},
				Name:                    "name",
				Slug:                    "name",
				Tags:                    []string{},
				UpdatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				User:                    redash.User{},
				UserID:                  1,
				Version:                 2,
				Widgets:                 nil,
			},
		},
	}, res)
}

func Test_GetDashboard_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards/1", func(req *http.Request) (*http.Response, error) {
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
				"can_edit": true,
				"created_at": "2023-02-10T01:23:45.000Z",
				"dashboard_filters_enabled": false,
				"id": 1,
				"is_archived": false,
				"is_draft": false,
				"is_favorite": false,
				"layout": [],
				"name": "name",
				"slug": "name",
				"tags": [],
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {},
				"user_id": 1,
				"version": 2,
				"widgets": []
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetDashboard(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.Dashboard{
		CanEdit:                 true,
		CreatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DashboardFiltersEnabled: false,
		ID:                      1,
		IsArchived:              false,
		IsDraft:                 false,
		IsFavorite:              false,
		Layout:                  []any{},
		Name:                    "name",
		Slug:                    "name",
		Tags:                    []string{},
		UpdatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:                    redash.User{},
		UserID:                  1,
		Version:                 2,
		Widgets:                 []redash.Widget{},
	}, res)
}

func Test_GetDashboard_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetDashboard(context.Background(), 1)
	assert.ErrorContains(err, "GET api/dashboards/1 failed: HTTP status code not OK: 503\nerror")
}

func Test_GetDashboard_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards/1", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetDashboard(context.Background(), 1)
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_CreateFavoriteDashboard_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards/1/favorite", func(req *http.Request) (*http.Response, error) {
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
	err := client.CreateFavoriteDashboard(context.Background(), 1)
	assert.NoError(err)
}

func Test_CreateFavoriteDashboard_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards/1/favorite", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.CreateFavoriteDashboard(context.Background(), 1)
	assert.ErrorContains(err, "POST api/dashboards/1/favorite failed: HTTP status code not OK: 503\nerror")
}

func Test_CreateDashboard_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"name":"name"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"can_edit": true,
				"created_at": "2023-02-10T01:23:45.000Z",
				"dashboard_filters_enabled": false,
				"id": 1,
				"is_archived": false,
				"is_draft": false,
				"is_favorite": false,
				"layout": [],
				"name": "name",
				"slug": "name",
				"tags": [],
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {},
				"user_id": 1,
				"version": 2,
				"widgets": []
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.CreateDashboard(context.Background(), &redash.CreateDashboardInput{
		Name: "name",
	})
	assert.NoError(err)
	assert.Equal(&redash.Dashboard{
		CanEdit:                 true,
		CreatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DashboardFiltersEnabled: false,
		ID:                      1,
		IsArchived:              false,
		IsDraft:                 false,
		IsFavorite:              false,
		Layout:                  []any{},
		Name:                    "name",
		Slug:                    "name",
		Tags:                    []string{},
		UpdatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:                    redash.User{},
		UserID:                  1,
		Version:                 2,
		Widgets:                 []redash.Widget{},
	}, res)
}

func Test_CreateDashboard_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.CreateDashboard(context.Background(), &redash.CreateDashboardInput{
		Name: "name",
	})
	assert.ErrorContains(err, "POST api/dashboards failed: HTTP status code not OK: 503\nerror")
}

func Test_CreateDashboard_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.CreateDashboard(context.Background(), &redash.CreateDashboardInput{
		Name: "name",
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_UpdateDashboard_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards/1", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"dashboard_filters_enabled":true,"is_archived":true,"is_draft":true,"name":"name","tags":["foo","bar"],"version":1}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"can_edit": true,
				"created_at": "2023-02-10T01:23:45.000Z",
				"dashboard_filters_enabled": false,
				"id": 1,
				"is_archived": false,
				"is_draft": false,
				"is_favorite": false,
				"layout": [],
				"name": "name",
				"slug": "name",
				"tags": [],
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {},
				"user_id": 1,
				"version": 2,
				"widgets": []
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.UpdateDashboard(context.Background(), 1, &redash.UpdateDashboardInput{
		DashboardFiltersEnabled: true,
		IsArchived:              true,
		IsDraft:                 true,
		Layout:                  []any{},
		Name:                    "name",
		Options:                 nil,
		Tags:                    &[]string{"foo", "bar"},
		Version:                 1,
	})
	assert.NoError(err)
	assert.Equal(&redash.Dashboard{
		CanEdit:                 true,
		CreatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DashboardFiltersEnabled: false,
		ID:                      1,
		IsArchived:              false,
		IsDraft:                 false,
		IsFavorite:              false,
		Layout:                  []any{},
		Name:                    "name",
		Slug:                    "name",
		Tags:                    []string{},
		UpdatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:                    redash.User{},
		UserID:                  1,
		Version:                 2,
		Widgets:                 []redash.Widget{},
	}, res)
}

func Test_UpdateDashboard_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.UpdateDashboard(context.Background(), 1, &redash.UpdateDashboardInput{
		DashboardFiltersEnabled: true,
		IsArchived:              true,
		IsDraft:                 true,
		Layout:                  []any{},
		Name:                    "name",
		Options:                 nil,
		Tags:                    &[]string{"foo", "bar"},
		Version:                 1,
	})
	assert.ErrorContains(err, "POST api/dashboards/1 failed: HTTP status code not OK: 503")
}

func Test_UpdateDashboard_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards/1", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.UpdateDashboard(context.Background(), 1, &redash.UpdateDashboardInput{
		DashboardFiltersEnabled: true,
		IsArchived:              true,
		IsDraft:                 true,
		Layout:                  []any{},
		Name:                    "name",
		Options:                 nil,
		Tags:                    &[]string{"foo", "bar"},
		Version:                 1,
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_ArchiveDashboard_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/dashboards/1", func(req *http.Request) (*http.Response, error) {
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
	err := client.ArchiveDashboard(context.Background(), 1)
	assert.NoError(err)
}

func Test_ArchiveDashboard_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/dashboards/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.ArchiveDashboard(context.Background(), 1)
	assert.ErrorContains(err, "DELETE api/dashboards/1 failed: HTTP status code not OK: 503\nerror")
}

func Test_ListMyDashboards_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards/my", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal("page=1&page_size=25&q=name", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"count": 1,
				"page": 1,
				"page_size": 25,
				"results": [
					{
						"created_at": "2023-02-10T01:23:45.000Z",
						"dashboard_filters_enabled": false,
						"id": 1,
						"is_archived": false,
						"is_draft": false,
						"is_favorite": false,
						"layout": [],
						"name": "name",
						"slug": "name",
						"tags": [],
						"updated_at": "2023-02-10T01:23:45.000Z",
						"user": {},
						"user_id": 1,
						"version": 2,
						"widgets": null
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListMyDashboards(context.Background(), &redash.ListMyDashboardsInput{
		Page:     1,
		PageSize: 25,
		Q:        "name",
	})
	assert.NoError(err)
	assert.Equal(&redash.DashboardPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Dashboard{
			{
				CanEdit:                 false,
				CreatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DashboardFiltersEnabled: false,
				ID:                      1,
				IsArchived:              false,
				IsDraft:                 false,
				IsFavorite:              false,
				Layout:                  []any{},
				Name:                    "name",
				Slug:                    "name",
				Tags:                    []string{},
				UpdatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				User:                    redash.User{},
				UserID:                  1,
				Version:                 2,
				Widgets:                 nil,
			},
		},
	}, res)
}

func Test_ListMyDashboards_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards/my", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListMyDashboards(context.Background(), &redash.ListMyDashboardsInput{
		Page:     1,
		PageSize: 25,
		Q:        "name",
	})
	assert.ErrorContains(err, "GET api/dashboards/my failed: HTTP status code not OK: 503\nerror")
}

func Test_ListMyDashboards_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards/my", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListMyDashboards(context.Background(), &redash.ListMyDashboardsInput{
		Page:     1,
		PageSize: 25,
		Q:        "name",
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_ListFavoriteDashboards_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards/favorites", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal("page=1&page_size=25&q=name", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"count": 1,
				"page": 1,
				"page_size": 25,
				"results": [
					{
						"created_at": "2023-02-10T01:23:45.000Z",
						"dashboard_filters_enabled": false,
						"id": 1,
						"is_archived": false,
						"is_draft": false,
						"is_favorite": false,
						"layout": [],
						"name": "name",
						"slug": "name",
						"tags": [],
						"updated_at": "2023-02-10T01:23:45.000Z",
						"user": {},
						"user_id": 1,
						"version": 2,
						"widgets": null
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListFavoriteDashboards(context.Background(), &redash.ListFavoriteDashboardsInput{
		Page:     1,
		PageSize: 25,
		Q:        "name",
	})
	assert.NoError(err)
	assert.Equal(&redash.DashboardPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Dashboard{
			{
				CanEdit:                 false,
				CreatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DashboardFiltersEnabled: false,
				ID:                      1,
				IsArchived:              false,
				IsDraft:                 false,
				IsFavorite:              false,
				Layout:                  []any{},
				Name:                    "name",
				Slug:                    "name",
				Tags:                    []string{},
				UpdatedAt:               dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				User:                    redash.User{},
				UserID:                  1,
				Version:                 2,
				Widgets:                 nil,
			},
		},
	}, res)
}

func Test_ListFavoriteDashboards_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards/favorites", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListFavoriteDashboards(context.Background(), &redash.ListFavoriteDashboardsInput{
		Page:     1,
		PageSize: 25,
		Q:        "name",
	})
	assert.ErrorContains(err, "GET api/dashboards/favorites failed: HTTP status code not OK: 503\nerror")
}

func Test_ListFavoriteDashboards_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/dashboards/favorites", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListFavoriteDashboards(context.Background(), &redash.ListFavoriteDashboardsInput{
		Page:     1,
		PageSize: 25,
		Q:        "name",
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_ShareDashboard_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards/1/share", func(req *http.Request) (*http.Response, error) {
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
				"api_key": "a7f60fafc2bea325bbd69ccc1b2846cb",
				"public_url": "https://redash.example.com/public/dashboards/a7f60fafc2bea325bbd69ccc1b2846cb?org_slug=default"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ShareDashboard(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.ShareDashboardOutput{
		APIKey:    "a7f60fafc2bea325bbd69ccc1b2846cb",
		PublicURL: "https://redash.example.com/public/dashboards/a7f60fafc2bea325bbd69ccc1b2846cb?org_slug=default",
	}, res)
}

func Test_ShareDashboard_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards/1/share", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ShareDashboard(context.Background(), 1)
	assert.ErrorContains(err, "POST api/dashboards/1/share failed: HTTP status code not OK: 503\nerror")
}

func Test_ShareDashboard_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/dashboards/1/share", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ShareDashboard(context.Background(), 1)
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_UnshareDashboard_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/dashboards/1/share", func(req *http.Request) (*http.Response, error) {
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
	err := client.UnshareDashboard(context.Background(), 1)
	assert.NoError(err)
}

func Test_UnshareDashboard_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/dashboards/1/share", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.UnshareDashboard(context.Background(), 1)
	assert.ErrorContains(err, "DELETE api/dashboards/1/share failed: HTTP status code not OK: 503\nerror")
}

func Test_Dashboard_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	_, err := client.ListDashboards(context.Background(), nil)
	require.NoError(err)

	dashboard, err := client.CreateDashboard(context.Background(), &redash.CreateDashboardInput{
		Name: "test-dashboard-1",
	})
	require.NoError(err)
	assert.Equal("test-dashboard-1", dashboard.Name)

	// NOTE: for v8
	// dashboard, err = client.GetDashboard(context.Background(), dashboard.Slug)
	dashboard, err = client.GetDashboard(context.Background(), dashboard.ID)
	require.NoError(err)
	assert.Equal("test-dashboard-1", dashboard.Name)

	dashboard, err = client.UpdateDashboard(context.Background(), dashboard.ID, &redash.UpdateDashboardInput{
		Name:    "test-dashboard-2",
		Tags:    &[]string{"foo"},
		Version: 0,
	})
	require.NoError(err)
	assert.Equal("test-dashboard-2", dashboard.Name)
	assert.Equal([]string{"foo"}, dashboard.Tags)

	tags, err := client.GetDashboardTags(context.Background())
	require.NoError(err)
	assert.GreaterOrEqual(len(tags.Tags), 1)

	page, err := client.ListDashboards(context.Background(), &redash.ListDashboardsInput{Q: "test-dashboard-2"})
	require.NoError(err)
	assert.GreaterOrEqual(len(page.Results), 1)

	page, err = client.ListMyDashboards(context.Background(), &redash.ListMyDashboardsInput{Q: "test-dashboard-2"})
	require.NoError(err)
	assert.GreaterOrEqual(len(page.Results), 1)

	page, err = client.ListFavoriteDashboards(context.Background(), &redash.ListFavoriteDashboardsInput{Q: "test-dashboard-2"})
	require.NoError(err)
	assert.Zero(len(page.Results))

	// NOTE: for v8
	// err = client.CreateFavoriteDashboard(context.Background(), dashboard.Slug)
	err = client.CreateFavoriteDashboard(context.Background(), dashboard.ID)
	require.NoError(err)
	page, err = client.ListFavoriteDashboards(context.Background(), &redash.ListFavoriteDashboardsInput{Q: "test-dashboard-2"})
	require.NoError(err)
	assert.GreaterOrEqual(len(page.Results), 1)

	share, err := client.ShareDashboard(context.Background(), dashboard.ID)
	require.NoError(err)
	assert.NotEmpty(share.PublicURL)

	err = client.UnshareDashboard(context.Background(), dashboard.ID)
	require.NoError(err)
	// NOTE: for v8
	// err = client.ArchiveDashboard(context.Background(), dashboard.Slug)
	err = client.ArchiveDashboard(context.Background(), dashboard.ID)
	require.NoError(err)
	// NOTE: for v8
	// dashboard, err = client.GetDashboard(context.Background(), dashboard.Slug)
	dashboard, err = client.GetDashboard(context.Background(), dashboard.ID)
	require.NoError(err)
	assert.Equal("test-dashboard-2", dashboard.Name)
	assert.True(dashboard.IsArchived)
	assert.True(dashboard.IsFavorite)
}
