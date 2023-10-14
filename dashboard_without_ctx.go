// Code generated from dashboard.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) ListDashboards(input *ListDashboardsInput) (*DashboardPage, error) {
	return client.withCtx.ListDashboards(context.Background(), input)
}

func (client *ClientWithoutContext) GetDashboard(id int) (*Dashboard, error) {
	return client.withCtx.GetDashboard(context.Background(), id)
}

func (client *ClientWithoutContext) CreateFavoriteDashboard(id int) error {
	return client.withCtx.CreateFavoriteDashboard(context.Background(), id)
}

func (client *ClientWithoutContext) CreateDashboard(input *CreateDashboardInput) (*Dashboard, error) {
	return client.withCtx.CreateDashboard(context.Background(), input)
}

func (client *ClientWithoutContext) UpdateDashboard(id int, input *UpdateDashboardInput) (*Dashboard, error) {
	return client.withCtx.UpdateDashboard(context.Background(), id, input)
}

func (client *ClientWithoutContext) ArchiveDashboard(id int) error {
	return client.withCtx.ArchiveDashboard(context.Background(), id)
}

func (client *ClientWithoutContext) GetDashboardTags() (*DashboardTags, error) {
	return client.withCtx.GetDashboardTags(context.Background())
}

func (client *ClientWithoutContext) ListMyDashboards(input *ListMyDashboardsInput) (*DashboardPage, error) {
	return client.withCtx.ListMyDashboards(context.Background(), input)
}

func (client *ClientWithoutContext) ListFavoriteDashboards(input *ListFavoriteDashboardsInput) (*DashboardPage, error) {
	return client.withCtx.ListFavoriteDashboards(context.Background(), input)
}

func (client *ClientWithoutContext) ShareDashboard(id int) (*ShareDashboardOutput, error) {
	return client.withCtx.ShareDashboard(context.Background(), id)
}

func (client *ClientWithoutContext) UnshareDashboard(id int) error {
	return client.withCtx.UnshareDashboard(context.Background(), id)
}
