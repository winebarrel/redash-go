package redash_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go"
)

func Test_GetOrganizationStatus_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/organization/status", func(req *http.Request) (*http.Response, error) {
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
				"object_counters": {
					"alerts": 1,
					"dashboards": 2,
					"data_sources": 3,
					"queries": 4,
					"users": 5
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetOrganizationStatus(context.Background())
	assert.NoError(err)
	assert.Equal(&redash.OrganizationStatus{
		ObjectCounters: redash.OrganizationStatusObjectCounters{
			Alerts:      1,
			Dashboards:  2,
			DataSources: 3,
			Queries:     4,
			Users:       5,
		},
	}, res)
}

func Test_GetOrganizationStatus_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)

	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	status, err := client.GetOrganizationStatus(context.Background())
	require.NoError(err)
	assert.NotNil(status)
}
