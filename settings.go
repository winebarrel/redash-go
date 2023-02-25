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
	AuthGoogleAppsDomains             []string `json:"auth_google_apps_domains,omitempty"`
	AuthJwtAuthAlgorithms             []string `json:"auth_jwt_auth_algorithms,omitempty"`
	AuthJwtAuthAudience               string   `json:"auth_jwt_auth_audience,omitempty"`
	AuthJwtAuthCookieName             string   `json:"auth_jwt_auth_cookie_name,omitempty"`
	AuthJwtAuthHeaderName             string   `json:"auth_jwt_auth_header_name,omitempty"`
	AuthJwtAuthIssuer                 string   `json:"auth_jwt_auth_issuer,omitempty"`
	AuthJwtAuthPublicCertsURL         string   `json:"auth_jwt_auth_public_certs_url,omitempty"`
	AuthJwtLoginEnabled               bool     `json:"auth_jwt_login_enabled,omitempty"`
	AuthPasswordLoginEnabled          bool     `json:"auth_password_login_enabled,omitempty"`
	AuthSamlEnabled                   bool     `json:"auth_saml_enabled,omitempty"`
	AuthSamlEntityID                  string   `json:"auth_saml_entity_id,omitempty"`
	AuthSamlMetadataURL               string   `json:"auth_saml_metadata_url,omitempty"`
	AuthSamlNameidFormat              string   `json:"auth_saml_nameid_format,omitempty"`
	AuthSamlSsoURL                    string   `json:"auth_saml_sso_url,omitempty"`
	AuthSamlType                      string   `json:"auth_saml_type,omitempty"`
	AuthSamlX509Cert                  string   `json:"auth_saml_x509_cert,omitempty"`
	BeaconConsent                     bool     `json:"beacon_consent,omitempty"`
	DateFormat                        string   `json:"date_format,omitempty"`
	DisablePublicUrls                 bool     `json:"disable_public_urls,omitempty"`
	FeatureShowPermissionsControl     bool     `json:"feature_show_permissions_control,omitempty"`
	FloatFormat                       string   `json:"float_format,omitempty"`
	HidePlotlyModeBar                 bool     `json:"hide_plotly_mode_bar,omitempty"`
	IntegerFormat                     string   `json:"integer_format,omitempty"`
	MultiByteSearchEnabled            bool     `json:"multi_byte_search_enabled,omitempty"`
	SendEmailOnFailedScheduledQueries bool     `json:"send_email_on_failed_scheduled_queries,omitempty"`
	TimeFormat                        string   `json:"time_format,omitempty"`
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

func (client *Client) UpdateSettingsOrganization(ctx context.Context, input *SettingsOrganizationSettings) (*SettingsOrganization, error) {
	res, close, err := client.Post(ctx, "api/settings/organization", input)
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
