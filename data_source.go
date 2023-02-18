//go:generate go run gen/withoutctx.go
package redash

import (
	"context"
	"fmt"

	"github.com/winebarrel/redash-go/internal/util"
)

type DataSource struct {
	Groups             map[int]bool   `json:"groups"`
	ID                 int            `json:"id"`
	Name               string         `json:"name"`
	Options            map[string]any `json:"options"`
	Paused             int            `json:"paused"`
	PauseReason        string         `json:"pause_reason"`
	QueueName          string         `json:"queue_name"`
	ScheduledQueueName string         `json:"scheduled_queue_name"`
	Syntax             string         `json:"syntax"`
	Type               string         `json:"type"`
	ViewOnly           bool           `json:"view_only"`
}

func (client *Client) ListDataSources(ctx context.Context) ([]DataSource, error) {
	res, close, err := client.Get(ctx, "api/data_sources", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	dataSources := []DataSource{}

	if err := util.UnmarshalBody(res, &dataSources); err != nil {
		return nil, err
	}

	return dataSources, nil
}

func (client *Client) GetDataSource(ctx context.Context, id int) (*DataSource, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/data_sources/%d", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	dataSource := &DataSource{}

	if err := util.UnmarshalBody(res, &dataSource); err != nil {
		return nil, fmt.Errorf("Unmarshal response body failed: %w", err)
	}

	return dataSource, nil
}

type CreateDataSourceInput struct {
	Name    string         `json:"name"`
	Options map[string]any `json:"options"`
	Type    string         `json:"type"`
}

func (client *Client) CreateDataSource(ctx context.Context, input *CreateDataSourceInput) (*DataSource, error) {
	res, close, err := client.Post(ctx, "api/data_sources", input)
	defer close()

	if err != nil {
		return nil, err
	}

	dataSource := &DataSource{}

	if err := util.UnmarshalBody(res, &dataSource); err != nil {
		return nil, err
	}

	return dataSource, nil
}

type UpdateDataSourceInput struct {
	Name    string         `json:"name"`
	Options map[string]any `json:"options"`
	Type    string         `json:"type"`
}

func (client *Client) UpdateDataSource(ctx context.Context, id int, input *UpdateDataSourceInput) (*DataSource, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/data_sources/%d", id), input)
	defer close()

	if err != nil {
		return nil, err
	}

	dataSource := &DataSource{}

	if err := util.UnmarshalBody(res, &dataSource); err != nil {
		return nil, err
	}

	return dataSource, nil
}

func (client *Client) DeleteDataSource(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/data_sources/%d", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}
