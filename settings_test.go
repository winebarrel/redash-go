package redash_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go/v2"
)

func Test_GetSettingsOrganization_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/settings/organization", func(req *http.Request) (*http.Response, error) {
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
				"settings": {
					"auth_google_apps_domains": [],
					"auth_jwt_auth_algorithms": [
						"HS256",
						"RS256",
						"ES256"
					],
					"auth_jwt_auth_audience": "",
					"auth_jwt_auth_cookie_name": "",
					"auth_jwt_auth_header_name": "",
					"auth_jwt_auth_issuer": "",
					"auth_jwt_auth_public_certs_url": "",
					"auth_jwt_login_enabled": false,
					"auth_password_login_enabled": true,
					"auth_saml_enabled": false,
					"auth_saml_entity_id": "",
					"auth_saml_metadata_url": "",
					"auth_saml_nameid_format": "",
					"auth_saml_sso_url": "",
					"auth_saml_type": "",
					"auth_saml_x509_cert": "",
					"beacon_consent": false,
					"date_format": "DD/MM/YY",
					"disable_public_urls": false,
					"feature_show_permissions_control": false,
					"float_format": "0,0.00",
					"hide_plotly_mode_bar": false,
					"integer_format": "0,0",
					"multi_byte_search_enabled": false,
					"send_email_on_failed_scheduled_queries": false,
					"time_format": "HH:mm"
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetSettingsOrganization(context.Background())
	assert.NoError(err)
	assert.Equal(&redash.SettingsOrganization{
		SettingsOrganizationSettings: redash.SettingsOrganizationSettings{
			AuthGoogleAppsDomains:             []string{},
			AuthJwtAuthAlgorithms:             []string{"HS256", "RS256", "ES256"},
			AuthJwtAuthAudience:               "",
			AuthJwtAuthCookieName:             "",
			AuthJwtAuthHeaderName:             "",
			AuthJwtAuthIssuer:                 "",
			AuthJwtAuthPublicCertsURL:         "",
			AuthJwtLoginEnabled:               false,
			AuthPasswordLoginEnabled:          true,
			AuthSamlEnabled:                   false,
			AuthSamlEntityID:                  "",
			AuthSamlMetadataURL:               "",
			AuthSamlNameidFormat:              "",
			AuthSamlSsoURL:                    "",
			AuthSamlType:                      "",
			AuthSamlX509Cert:                  "",
			BeaconConsent:                     false,
			DateFormat:                        "DD/MM/YY",
			DisablePublicUrls:                 false,
			FeatureShowPermissionsControl:     false,
			FloatFormat:                       "0,0.00",
			HidePlotlyModeBar:                 false,
			IntegerFormat:                     "0,0",
			MultiByteSearchEnabled:            false,
			SendEmailOnFailedScheduledQueries: false,
			TimeFormat:                        "HH:mm",
		},
	}, res)
}

func Test_GetSettingsOrganization_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/settings/organization", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetSettingsOrganization(context.Background())
	assert.ErrorContains(err, "GET api/settings/organization failed: HTTP status code not OK: 503\nerror")
}

func Test_GetSettingsOrganization_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/settings/organization", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetSettingsOrganization(context.Background())
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_UpdateSettingsOrganization_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/settings/organization", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"date_format":"YYYY/MM/DD"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"settings": {
					"auth_google_apps_domains": [],
					"auth_jwt_auth_algorithms": [
						"HS256",
						"RS256",
						"ES256"
					],
					"auth_jwt_auth_audience": "",
					"auth_jwt_auth_cookie_name": "",
					"auth_jwt_auth_header_name": "",
					"auth_jwt_auth_issuer": "",
					"auth_jwt_auth_public_certs_url": "",
					"auth_jwt_login_enabled": false,
					"auth_password_login_enabled": true,
					"auth_saml_enabled": false,
					"auth_saml_entity_id": "",
					"auth_saml_metadata_url": "",
					"auth_saml_nameid_format": "",
					"auth_saml_sso_url": "",
					"auth_saml_type": "",
					"auth_saml_x509_cert": "",
					"beacon_consent": false,
					"date_format": "YYYY/MM/DD",
					"disable_public_urls": false,
					"feature_show_permissions_control": false,
					"float_format": "0,0.00",
					"hide_plotly_mode_bar": false,
					"integer_format": "0,0",
					"multi_byte_search_enabled": false,
					"send_email_on_failed_scheduled_queries": false,
					"time_format": "HH:mm"
				}
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.UpdateSettingsOrganization(context.Background(), &redash.UpdateSettingsOrganizationInput{
		DateFormat: "YYYY/MM/DD",
	})
	assert.NoError(err)
	assert.Equal(&redash.SettingsOrganization{
		SettingsOrganizationSettings: redash.SettingsOrganizationSettings{
			AuthGoogleAppsDomains:             []string{},
			AuthJwtAuthAlgorithms:             []string{"HS256", "RS256", "ES256"},
			AuthJwtAuthAudience:               "",
			AuthJwtAuthCookieName:             "",
			AuthJwtAuthHeaderName:             "",
			AuthJwtAuthIssuer:                 "",
			AuthJwtAuthPublicCertsURL:         "",
			AuthJwtLoginEnabled:               false,
			AuthPasswordLoginEnabled:          true,
			AuthSamlEnabled:                   false,
			AuthSamlEntityID:                  "",
			AuthSamlMetadataURL:               "",
			AuthSamlNameidFormat:              "",
			AuthSamlSsoURL:                    "",
			AuthSamlType:                      "",
			AuthSamlX509Cert:                  "",
			BeaconConsent:                     false,
			DateFormat:                        "YYYY/MM/DD",
			DisablePublicUrls:                 false,
			FeatureShowPermissionsControl:     false,
			FloatFormat:                       "0,0.00",
			HidePlotlyModeBar:                 false,
			IntegerFormat:                     "0,0",
			MultiByteSearchEnabled:            false,
			SendEmailOnFailedScheduledQueries: false,
			TimeFormat:                        "HH:mm",
		},
	}, res)
}

func Test_UpdateSettingsOrganization_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/settings/organization", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.UpdateSettingsOrganization(context.Background(), &redash.UpdateSettingsOrganizationInput{
		DateFormat: "YYYY/MM/DD",
	})
	assert.ErrorContains(err, "POST api/settings/organization failed: HTTP status code not OK: 503\nerror")
}

func Test_UpdateSettingsOrganization_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/settings/organization", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.UpdateSettingsOrganization(context.Background(), &redash.UpdateSettingsOrganizationInput{
		DateFormat: "YYYY/MM/DD",
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_Settings_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	settings, err := client.GetSettingsOrganization(context.Background())
	require.NoError(err)
	assert.NotEmpty(settings.DateFormat)

	settings, err = client.UpdateSettingsOrganization(context.Background(), &redash.UpdateSettingsOrganizationInput{
		DateFormat: "YYYY/MM/DD",
	})
	require.NoError(err)
	assert.Equal("YYYY/MM/DD", settings.DateFormat)
}
