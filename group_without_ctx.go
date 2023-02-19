// Code generated from group.go using internal/gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) ListGroups() ([]Group, error) {
	return client.withCtx.ListGroups(context.Background())
}

// Auto-generated
func (client *ClientWithoutContext) GetGroup(id int) (*Group, error) {
	return client.withCtx.GetGroup(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) CreateGroup(input *CreateGroupInput) (*Group, error) {
	return client.withCtx.CreateGroup(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) DeleteGroup(id int) error {
	return client.withCtx.DeleteGroup(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) ListGroupMembers(id int) ([]User, error) {
	return client.withCtx.ListGroupMembers(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) AddGroupMember(id int, userId int) (*User, error) {
	return client.withCtx.AddGroupMember(context.Background(), id, userId)
}

// Auto-generated
func (client *ClientWithoutContext) RemoveGroupMember(id int, userId int) error {
	return client.withCtx.RemoveGroupMember(context.Background(), id, userId)
}

// Auto-generated
func (client *ClientWithoutContext) ListGroupDataSources(id int) ([]DataSource, error) {
	return client.withCtx.ListGroupDataSources(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) AddGroupDataSource(id int, dsId int) (*DataSource, error) {
	return client.withCtx.AddGroupDataSource(context.Background(), id, dsId)
}

// Auto-generated
func (client *ClientWithoutContext) RemoveGroupDataSource(id int, dsId int) error {
	return client.withCtx.RemoveGroupDataSource(context.Background(), id, dsId)
}

// Auto-generated
func (client *ClientWithoutContext) UpdateGroupDataSource(id int, dsId int, input *UpdateGroupDataSourceInput) (*DataSource, error) {
	return client.withCtx.UpdateGroupDataSource(context.Background(), id, dsId, input)
}
