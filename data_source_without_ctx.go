// Code generated from data_source.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) ListDataSources() ([]DataSource, error) {
	return client.withCtx.ListDataSources(context.Background())
}

func (client *ClientWithoutContext) GetDataSource(id int) (*DataSource, error) {
	return client.withCtx.GetDataSource(context.Background(), id)
}

func (client *ClientWithoutContext) CreateDataSource(input *CreateDataSourceInput) (*DataSource, error) {
	return client.withCtx.CreateDataSource(context.Background(), input)
}

func (client *ClientWithoutContext) UpdateDataSource(id int, input *UpdateDataSourceInput) (*DataSource, error) {
	return client.withCtx.UpdateDataSource(context.Background(), id, input)
}

func (client *ClientWithoutContext) DeleteDataSource(id int) error {
	return client.withCtx.DeleteDataSource(context.Background(), id)
}

func (client *ClientWithoutContext) PauseDataSource(id int, input *PauseDataSourceInput) (*DataSource, error) {
	return client.withCtx.PauseDataSource(context.Background(), id, input)
}

func (client *ClientWithoutContext) ResumeDataSource(id int) (*DataSource, error) {
	return client.withCtx.ResumeDataSource(context.Background(), id)
}

func (client *ClientWithoutContext) TestDataSource(id int) (*TestDataSourceOutput, error) {
	return client.withCtx.TestDataSource(context.Background(), id)
}

func (client *ClientWithoutContext) GetDataSourceTypes() ([]DataSourceType, error) {
	return client.withCtx.GetDataSourceTypes(context.Background())
}

func (client *ClientWithoutContext) GetDataSourceSchema(id int) (*DataSourceSchemaOutput, error) {
	return client.withCtx.GetDataSourceSchema(context.Background(), id)
}
