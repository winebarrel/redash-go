package redash_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go/v2"
)

func Test_TestCredentials_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/session", func(req *http.Request) (*http.Response, error) {
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
		return httpmock.NewStringResponse(http.StatusOK, `{}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.TestCredentials(context.Background())
	assert.NoError(err)
}

func Test_TestCredentials_Err(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/session", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusNotFound, `{"message": "Couldn't find resource. Please login and try again."}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.TestCredentials(context.Background())
	assert.ErrorContains(err, "GET api/session failed: HTTP status code not OK: 404")
}

func Test_TestCredentials_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	err := client.TestCredentials(context.Background())
	assert.NoError(err)
}

func Test_GetSession_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/session", func(req *http.Request) (*http.Response, error) {
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
				"client_config": {},
				"messages": [
					"email-not-verified",
					"using-deprecated-embed-feature"
				],
				"org_slug": "default",
				"user": {}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetSession(context.Background())
	assert.NoError(err)
	assert.Equal(&redash.Session{
		ClientConfig: redash.ClientConfig{},
		Messages: []string{
			"email-not-verified",
			"using-deprecated-embed-feature",
		},
		OrgSlug: "default",
		User:    redash.User{},
	}, res)
}

func Test_GetSession_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/session", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetSession(context.Background())
	assert.ErrorContains(err, "GET api/session failed: HTTP status code not OK: 503 Service Unavailable\nerror")
}

func Test_GetSession_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/session", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetSession(context.Background())
	assert.ErrorContains(err, "read response body failed: IO error")
}

func Test_Session_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	session, err := client.GetSession(context.Background())
	require.NoError(err)
	assert.Equal("admin", session.User.Name)
}
