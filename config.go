//go:generate go run gen/withoutctx.go
package redash

import (
	"context"

	"github.com/winebarrel/redash-go/internal/util"
)

type Config struct {
	ClientConfig ClientConfig `json:"client_config"`
	OrgSlug      string       `json:"org_slug"`
}

type ClientConfig struct {
	AllowCustomJSVisualizations bool     `json:"allowCustomJSVisualizations"`
	AllowScriptsInUserInput     bool     `json:"allowScriptsInUserInput"`
	AutoPublishNamedQueries     bool     `json:"autoPublishNamedQueries"`
	BasePath                    string   `json:"basePath"`
	DashboardRefreshIntervals   []int    `json:"dashboardRefreshIntervals"`
	DateFormat                  string   `json:"dateFormat"`
	DateFormatList              []string `json:"dateFormatList"`
	DateTimeFormat              string   `json:"dateTimeFormat"`
	ExtendedAlertOptions        bool     `json:"extendedAlertOptions"`
	FloatFormat                 string   `json:"floatFormat"`
	GoogleLoginEnabled          bool     `json:"googleLoginEnabled"`
	IntegerFormat               string   `json:"integerFormat"`
	MailSettingsMissing         bool     `json:"mailSettingsMissing"`
	NewVersionAvailable         bool     `json:"newVersionAvailable"`
	PageSize                    int      `json:"pageSize"`
	PageSizeOptions             []int    `json:"pageSizeOptions"`
	QueryRefreshIntervals       []int    `json:"queryRefreshIntervals"`
	ShowBeaconConsentMessage    bool     `json:"showBeaconConsentMessage"`
	ShowPermissionsControl      bool     `json:"showPermissionsControl"`
	TableCellMaxJSONSize        int      `json:"tableCellMaxJSONSize"`
	TimeFormatList              []string `json:"timeFormatList"`
	Version                     string   `json:"version"`
}

func (client *Client) GetConfig(ctx context.Context) (*Config, error) {
	res, close, err := client.Get(ctx, "api/config", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err := util.UnmarshalBody(res, &config); err != nil {
		return nil, err
	}

	return config, nil
}
