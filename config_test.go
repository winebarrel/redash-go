package redash_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go/v2"
)

func Test_GetConfig_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/config", func(req *http.Request) (*http.Response, error) {
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
				"client_config": {
					"allowCustomJSVisualizations": false,
					"allowScriptsInUserInput": false,
					"autoPublishNamedQueries": true,
					"basePath": "http://localhost:5000/",
					"dashboardRefreshIntervals": [
						60,
						300,
						600,
						1800,
						3600,
						43200,
						86400
					],
					"dateFormat": "DD/MM/YY",
					"dateFormatList": [
						"DD/MM/YY",
						"YYYY-MM-DD",
						"MM/DD/YY"
					],
					"dateTimeFormat": "DD/MM/YY HH:mm",
					"extendedAlertOptions": false,
					"floatFormat": "0,0.00",
					"googleLoginEnabled": false,
					"integerFormat": "0,0",
					"mailSettingsMissing": false,
					"newVersionAvailable": false,
					"pageSize": 20,
					"pageSizeOptions": [
						5,
						10,
						20,
						50,
						100
					],
					"queryRefreshIntervals": [
						60,
						300,
						600,
						900,
						1800,
						3600,
						7200,
						10800,
						14400,
						18000,
						21600,
						25200,
						28800,
						32400,
						36000,
						39600,
						43200,
						86400,
						604800,
						1209600,
						2592000
					],
					"showBeaconConsentMessage": true,
					"showPermissionsControl": false,
					"tableCellMaxJSONSize": 50000,
					"timeFormatList": [
						"HH:mm:ss",
						"HH:mm",
						"HH:mm:ss.SSS"
					],
					"version": "8.0.0+b32245"
				},
				"org_slug": "default"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetConfig(context.Background())
	assert.NoError(err)
	assert.Equal(&redash.Config{
		ClientConfig: redash.ClientConfig{
			AllowCustomJSVisualizations: false,
			AllowScriptsInUserInput:     false,
			AutoPublishNamedQueries:     true,
			BasePath:                    "http://localhost:5000/",
			DashboardRefreshIntervals: []int{
				60,
				300,
				600,
				1800,
				3600,
				43200,
				86400,
			},
			DateFormat: "DD/MM/YY",
			DateFormatList: []string{
				"DD/MM/YY",
				"YYYY-MM-DD",
				"MM/DD/YY",
			},
			DateTimeFormat:       "DD/MM/YY HH:mm",
			ExtendedAlertOptions: false,
			FloatFormat:          "0,0.00",
			GoogleLoginEnabled:   false,
			IntegerFormat:        "0,0",
			MailSettingsMissing:  false,
			NewVersionAvailable:  false,
			PageSize:             20,
			PageSizeOptions: []int{
				5,
				10,
				20,
				50,
				100,
			},
			QueryRefreshIntervals: []int{
				60,
				300,
				600,
				900,
				1800,
				3600,
				7200,
				10800,
				14400,
				18000,
				21600,
				25200,
				28800,
				32400,
				36000,
				39600,
				43200,
				86400,
				604800,
				1209600,
				2592000,
			},
			ShowBeaconConsentMessage: true,
			ShowPermissionsControl:   false,
			TableCellMaxJSONSize:     50000,
			TimeFormatList: []string{
				"HH:mm:ss",
				"HH:mm",
				"HH:mm:ss.SSS",
			},
			Version: "8.0.0+b32245",
		},
		OrgSlug: "default",
	}, res)
}

func Test_GetConfig_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/config", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetConfig(context.Background())
	assert.ErrorContains(err, "GET api/config failed: HTTP status code not OK: 503\nerror")
}

func Test_GetConfig_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/config", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetConfig(context.Background())
	assert.ErrorContains(err, "read response body failed: IO error")
}

func Test_Config_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)

	// NOTE: No authentication required
	client, _ := redash.NewClient(testRedashEndpoint, "")
	config, err := client.GetConfig(context.Background())
	require.NoError(err)
	assert.Equal("default", config.OrgSlug)
	assert.Equal(50000, config.ClientConfig.TableCellMaxJSONSize)
	assert.True(strings.HasPrefix(config.ClientConfig.BasePath, testRedashEndpoint))
}
