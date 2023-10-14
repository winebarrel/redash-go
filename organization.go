//go:generate go run tools/withoutctx.go
package redash

import (
	"context"

	"github.com/winebarrel/redash-go/v2/internal/util"
)

type OrganizationStatus struct {
	ObjectCounters OrganizationStatusObjectCounters `json:"object_counters"`
}

type OrganizationStatusObjectCounters struct {
	Alerts      int `json:"alerts"`
	Dashboards  int `json:"dashboards"`
	DataSources int `json:"data_sources"`
	Queries     int `json:"queries"`
	Users       int `json:"users"`
}

func (client *Client) GetOrganizationStatus(ctx context.Context) (*OrganizationStatus, error) {
	res, close, err := client.Get(ctx, "api/organization/status", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	status := &OrganizationStatus{}

	if err := util.UnmarshalBody(res, &status); err != nil {
		return nil, err
	}

	return status, nil
}
