// Code generated from data_source.go using internal/gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) ListDataSources() ([]DataSource, error) {
	return client.withCtx.ListDataSources(context.Background())
}

// Auto-generated
func (client *ClientWithoutContext) GetDataSource(id int) (*DataSource, error) {
	return client.withCtx.GetDataSource(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) CreateDataSource(input *CreateDataSourceInput) (*DataSource, error) {
	return client.withCtx.CreateDataSource(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) UpdateDataSource(id int, input *UpdateDataSourceInput) (*DataSource, error) {
	return client.withCtx.UpdateDataSource(context.Background(), id, input)
}

// Auto-generated
func (client *ClientWithoutContext) DeleteDataSource(id int) error {
	return client.withCtx.DeleteDataSource(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) PauseDataSource(id int, input *PauseDataSourceInput) (*DataSource, error) {
	return client.withCtx.PauseDataSource(context.Background(), id, input)
}

// Auto-generated
func (client *ClientWithoutContext) ResumeDataSource(id int) (*DataSource, error) {
	return client.withCtx.ResumeDataSource(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) TestDataSource(id int) (*TestDataSourceOutput, error) {
	return client.withCtx.TestDataSource(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) GetDataSourceTypes() ([]DataSourceType, error) {
	return client.withCtx.GetDataSourceTypes(context.Background())
}
