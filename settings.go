//go:generate go run internal/gen/withoutctx.go
package redash

import (
	"context"

	"github.com/winebarrel/redash-go/internal/util"
)

type SettingsOrganization struct {
	SettingsOrganizationSettings `json:"settings"`
}

type SettingsOrganizationSettings struct {
	AuthGoogleAppsDomains             []string `json:"auth_google_apps_domains"`
	AuthJwtAuthAlgorithms             []string `json:"auth_jwt_auth_algorithms"`
	AuthJwtAuthAudience               string   `json:"auth_jwt_auth_audience"`
	AuthJwtAuthCookieName             string   `json:"auth_jwt_auth_cookie_name"`
	AuthJwtAuthHeaderName             string   `json:"auth_jwt_auth_header_name"`
	AuthJwtAuthIssuer                 string   `json:"auth_jwt_auth_issuer"`
	AuthJwtAuthPublicCertsURL         string   `json:"auth_jwt_auth_public_certs_url"`
	AuthJwtLoginEnabled               bool     `json:"auth_jwt_login_enabled"`
	AuthPasswordLoginEnabled          bool     `json:"auth_password_login_enabled"`
	AuthSamlEnabled                   bool     `json:"auth_saml_enabled"`
	AuthSamlEntityID                  string   `json:"auth_saml_entity_id"`
	AuthSamlMetadataURL               string   `json:"auth_saml_metadata_url"`
	AuthSamlNameidFormat              string   `json:"auth_saml_nameid_format"`
	AuthSamlSsoURL                    string   `json:"auth_saml_sso_url"`
	AuthSamlType                      string   `json:"auth_saml_type"`
	AuthSamlX509Cert                  string   `json:"auth_saml_x509_cert"`
	BeaconConsent                     bool     `json:"beacon_consent"`
	DateFormat                        string   `json:"date_format"`
	DisablePublicUrls                 bool     `json:"disable_public_urls"`
	FeatureShowPermissionsControl     bool     `json:"feature_show_permissions_control"`
	FloatFormat                       string   `json:"float_format"`
	HidePlotlyModeBar                 bool     `json:"hide_plotly_mode_bar"`
	IntegerFormat                     string   `json:"integer_format"`
	MultiByteSearchEnabled            bool     `json:"multi_byte_search_enabled"`
	SendEmailOnFailedScheduledQueries bool     `json:"send_email_on_failed_scheduled_queries"`
	TimeFormat                        string   `json:"time_format"`
}

func (client *Client) GetSettingsOrganization(ctx context.Context) (*SettingsOrganization, error) {
	res, close, err := client.Get(ctx, "api/settings/organization", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	org := &SettingsOrganization{}

	if err := util.UnmarshalBody(res, &org); err != nil {
		return nil, err
	}

	return org, nil
}
