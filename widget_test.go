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

func Test_CreateWidget_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/widgets", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"dashboard_id":1,"options":{"isHidden":false},"text":"text","visualization_id":1,"width":1}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"created_at": "2023-02-10T01:23:45.000Z",
				"dashboard_id": 1,
				"id": 1,
				"options": {
					"isHidden": false
				},
				"text": "text",
				"updated_at": "2023-02-10T01:23:45.000Z",
				"visualization": {},
				"width": 1
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.CreateWidget(context.Background(), &redash.CreateWidgetInput{
		DashboardID: 1,
		Options: map[string]any{
			"isHidden": false,
		},
		Text:            "text",
		VisualizationID: 1,
		Width:           1,
	})
	assert.NoError(err)
	assert.Equal(&redash.Widget{
		CreatedAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DashboardID: 1,
		ID:          1,
		Options: map[string]any{
			"isHidden": false,
		},
		Text:          "text",
		UpdatedAt:     dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		Visualization: redash.Visualization{},
		Width:         1,
	}, res)
}
