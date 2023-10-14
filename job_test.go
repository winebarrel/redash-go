package redash_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/redash-go/v2"
)

func Test_GetJob_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/jobs/623b290a-7fd9-4ea6-a2a6-96f9c9101f51", func(req *http.Request) (*http.Response, error) {
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
					"id": "623b290a-7fd9-4ea6-a2a6-96f9c9101f51",
					"query_result_id": 1,
					"status": 3,
					"updated_at": 0
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetJob(context.Background(), "623b290a-7fd9-4ea6-a2a6-96f9c9101f51")
	assert.NoError(err)
	assert.Equal(&redash.JobResponse{
		Job: redash.Job{
			Error:         "",
			ID:            "623b290a-7fd9-4ea6-a2a6-96f9c9101f51",
			QueryResultID: 1,
			Status:        redash.JobStatusSuccess,
			UpdatedAt:     float64(0),
		},
	}, res)
}
