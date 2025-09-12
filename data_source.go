//go:generate go run tools/withoutctx.go
package redash

import (
	"context"
	"fmt"

	"github.com/winebarrel/redash-go/v2/internal/util"
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
		return nil, err
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

type PauseDataSourceInput struct {
	Reason string `json:"reason,omitempty"`
}

func (client *Client) PauseDataSource(ctx context.Context, id int, input *PauseDataSourceInput) (*DataSource, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/data_sources/%d/pause", id), input)
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

func (client *Client) ResumeDataSource(ctx context.Context, id int) (*DataSource, error) {
	res, close, err := client.Delete(ctx, fmt.Sprintf("api/data_sources/%d/pause", id))
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

type TestDataSourceOutput struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func (client *Client) TestDataSource(ctx context.Context, id int) (*TestDataSourceOutput, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/data_sources/%d/test", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	output := &TestDataSourceOutput{}

	if err := util.UnmarshalBody(res, &output); err != nil {
		return nil, err
	}

	return output, nil
}

type DataSourceType struct {
	ConfigurationSchema DataSourceTypeConfigurationSchema `json:"configuration_schema"`
	Name                string                            `json:"name"`
	Type                string                            `json:"type"`
}

type DataSourceTypeConfigurationSchema struct {
	ExtraOptions []string                                             `json:"extra_options"`
	Order        []string                                             `json:"order"`
	Properties   map[string]DataSourceTypeConfigurationSchemaProperty `json:"properties"`
	Required     []string                                             `json:"required"`
	Secret       []string                                             `json:"secret"`
	Type         string                                               `json:"type"`
}

type DataSourceTypeConfigurationSchemaProperty struct {
	Default any    `json:"default"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

func (client *Client) GetDataSourceTypes(ctx context.Context) ([]DataSourceType, error) {
	res, close, err := client.Get(ctx, "api/data_sources/types", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	types := []DataSourceType{}

	if err := util.UnmarshalBody(res, &types); err != nil {
		return nil, err
	}

	return types, nil
}

type DataSourceSchemaOutput struct {
	Schema []DataSourceSchema `json:"schema"`
}

type DataSourceSchema struct {
	Name    string                   `json:"name"`
	Columns []DataSourceSchemaColumn `json:"columns"`
}

type DataSourceSchemaColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (client *Client) GetDataSourceSchema(ctx context.Context, id int) (*DataSourceSchemaOutput, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/data_sources/%d/schema", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	output := &DataSourceSchemaOutput{}

	if err := util.UnmarshalBody(res, &output); err != nil {
		return nil, err
	}

	return output, nil
}
