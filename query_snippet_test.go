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

func Test_ListQuerySnippets_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/query_snippets", func(req *http.Request) (*http.Response, error) {
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
					"created_at": "2023-02-10T01:23:45.000Z",
					"description": "description",
					"id": 1,
					"snippet": "select 1",
					"trigger": "my-snippet",
					"updated_at": "2023-02-10T01:23:45.000Z",
					"user": {}
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListQuerySnippets(context.Background())
	assert.NoError(err)
	assert.Equal([]redash.QuerySnippet{
		{
			CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			Description: "description",
			ID:          1,
			Snippet:     "select 1",
			Trigger:     "my-snippet",
			UpdatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			User:        redash.User{},
		},
	}, res)
}

func Test_GetQuerySnippets_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/query_snippets/1", func(req *http.Request) (*http.Response, error) {
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
				"created_at": "2023-02-10T01:23:45.000Z",
				"description": "description",
				"id": 1,
				"snippet": "select 1",
				"trigger": "my-snippet",
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetQuerySnippet(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.QuerySnippet{
		CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		Description: "description",
		ID:          1,
		Snippet:     "select 1",
		Trigger:     "my-snippet",
		UpdatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:        redash.User{},
	}, res)
}

func Test_CreateQuerySnippets_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/query_snippets", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"description":"description","snippet":"select 1","trigger":"my-snippet"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"created_at": "2023-02-10T01:23:45.000Z",
				"description": "description",
				"id": 1,
				"snippet": "select 1",
				"trigger": "my-snippet",
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.CreateQuerySnippet(context.Background(), &redash.CreateQuerySnippetInput{
		Description: "description",
		Snippet:     "select 1",
		Trigger:     "my-snippet",
	})
	assert.NoError(err)
	assert.Equal(&redash.QuerySnippet{
		CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		Description: "description",
		ID:          1,
		Snippet:     "select 1",
		Trigger:     "my-snippet",
		UpdatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:        redash.User{},
	}, res)
}

func Test_UpdateQuerySnippets_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/query_snippets/1", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"description":"description","snippet":"select 1","trigger":"my-snippet"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"created_at": "2023-02-10T01:23:45.000Z",
				"description": "description",
				"id": 1,
				"snippet": "select 1",
				"trigger": "my-snippet",
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.UpdateQuerySnippet(context.Background(), 1, &redash.UpdateQuerySnippetInput{
		Description: "description",
		Snippet:     "select 1",
		Trigger:     "my-snippet",
	})
	assert.NoError(err)
	assert.Equal(&redash.QuerySnippet{
		CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		Description: "description",
		ID:          1,
		Snippet:     "select 1",
		Trigger:     "my-snippet",
		UpdatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		User:        redash.User{},
	}, res)
}

func Test_DeleteQuerySnippets_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/query_snippets/1", func(req *http.Request) (*http.Response, error) {
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
				"created_at": "2023-02-10T01:23:45.000Z",
				"description": "description",
				"id": 1,
				"snippet": "select 1",
				"trigger": "my-snippet",
				"updated_at": "2023-02-10T01:23:45.000Z",
				"user": {}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.DeleteQuerySnippet(context.Background(), 1)
	assert.NoError(err)
}

func Test_QuerySnippet_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)

	_, err := client.ListQuerySnippets(context.Background())
	assert.NoError(err)

	snippet, err := client.CreateQuerySnippet(context.Background(), &redash.CreateQuerySnippetInput{
		Description: "description",
		Snippet:     "select 1",
		Trigger:     "my-snippet-1",
	})
	assert.NoError(err)
	assert.Equal("my-snippet-1", snippet.Trigger)

	snippet, err = client.GetQuerySnippet(context.Background(), snippet.ID)
	assert.NoError(err)
	assert.Equal("my-snippet-1", snippet.Trigger)

	snippet, err = client.UpdateQuerySnippet(context.Background(), snippet.ID, &redash.UpdateQuerySnippetInput{
		Snippet: "select 2",
	})
	assert.NoError(err)
	assert.Equal("select 2", snippet.Snippet)

	err = client.DeleteQuerySnippet(context.Background(), snippet.ID)
	assert.NoError(err)

	_, err = client.GetAlert(context.Background(), snippet.ID)
	assert.Error(err)
}
