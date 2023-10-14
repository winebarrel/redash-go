package redash_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/araddon/dateparse"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go"
)

func Test_ListEvents_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/events", func(req *http.Request) (*http.Response, error) {
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
						"action": "create",
						"browser": "Other / Other / Other",
						"created_at": "2023-02-10T01:23:45.000Z",
						"details": {
							"object_id": "44",
							"object_type": "datasource"
						},
						"location": "Unknown",
						"object_id": "44",
						"object_type": "datasource",
						"org_id": 1,
						"user_id": 1,
						"user_name": "admin"
					}
				]
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListEvents(context.Background(), &redash.ListEventsInput{
		Page:     1,
		PageSize: 25,
	})
	assert.NoError(err)
	assert.Equal(&redash.EventPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.Event{
			{
				Action:    "create",
				Browser:   "Other / Other / Other",
				CreatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				Details: map[string]string{
					"object_id":   "44",
					"object_type": "datasource",
				},
				Location:   "Unknown",
				ObjectID:   "44",
				ObjectType: "datasource",
				OrgID:      1,
				UserID:     1,
				UserName:   "admin",
			},
		},
	}, res)
}

func Test_Event_Acc(t *testing.T) {
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

	page, err := client.ListEvents(context.Background(), &redash.ListEventsInput{
		Page:     1,
		PageSize: 25,
	})

	require.NoError(err)
	assert.Equal(1, page.Page)
	assert.Equal(25, page.PageSize)
	assert.NotEqual(0, page.Count)

	if len(page.Results) == 0 {
		assert.FailNow("len(page.Results) == 0")
	}

	assert.NotEqual(0, len(page.Results))
	assert.Equal("admin", page.Results[0].UserName)
}
