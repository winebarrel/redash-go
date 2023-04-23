// Code generated from dashboard.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) ListDashboards(input *ListDashboardsInput) (*DashboardPage, error) {
	return client.withCtx.ListDashboards(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) GetDashboard(idOrSlug any) (*Dashboard, error) {
	return client.withCtx.GetDashboard(context.Background(), idOrSlug)
}

// Auto-generated
func (client *ClientWithoutContext) CreateFavoriteDashboard(idOrSlug any) error {
	return client.withCtx.CreateFavoriteDashboard(context.Background(), idOrSlug)
}

// Auto-generated
func (client *ClientWithoutContext) CreateDashboard(input *CreateDashboardInput) (*Dashboard, error) {
	return client.withCtx.CreateDashboard(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) UpdateDashboard(id int, input *UpdateDashboardInput) (*Dashboard, error) {
	return client.withCtx.UpdateDashboard(context.Background(), id, input)
}

// Auto-generated
func (client *ClientWithoutContext) ArchiveDashboard(idOrSlug any) error {
	return client.withCtx.ArchiveDashboard(context.Background(), idOrSlug)
}

// Auto-generated
func (client *ClientWithoutContext) GetDashboardTags() (*DashboardTags, error) {
	return client.withCtx.GetDashboardTags(context.Background())
}

// Auto-generated
func (client *ClientWithoutContext) ListMyDashboards(input *ListMyDashboardsInput) (*DashboardPage, error) {
	return client.withCtx.ListMyDashboards(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) ListFavoriteDashboards(input *ListFavoriteDashboardsInput) (*DashboardPage, error) {
	return client.withCtx.ListFavoriteDashboards(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) ShareDashboard(id int) (*ShareDashboardOutput, error) {
	return client.withCtx.ShareDashboard(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) UnshareDashboard(id int) error {
	return client.withCtx.UnshareDashboard(context.Background(), id)
}
