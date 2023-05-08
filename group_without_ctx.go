// Code generated from group.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) ListGroups() ([]Group, error) {
	return client.withCtx.ListGroups(context.Background())
}

func (client *ClientWithoutContext) GetGroup(id int) (*Group, error) {
	return client.withCtx.GetGroup(context.Background(), id)
}

func (client *ClientWithoutContext) CreateGroup(input *CreateGroupInput) (*Group, error) {
	return client.withCtx.CreateGroup(context.Background(), input)
}

func (client *ClientWithoutContext) DeleteGroup(id int) error {
	return client.withCtx.DeleteGroup(context.Background(), id)
}

func (client *ClientWithoutContext) ListGroupMembers(id int) ([]User, error) {
	return client.withCtx.ListGroupMembers(context.Background(), id)
}

func (client *ClientWithoutContext) AddGroupMember(id int, userId int) (*User, error) {
	return client.withCtx.AddGroupMember(context.Background(), id, userId)
}

func (client *ClientWithoutContext) RemoveGroupMember(id int, userId int) error {
	return client.withCtx.RemoveGroupMember(context.Background(), id, userId)
}

func (client *ClientWithoutContext) ListGroupDataSources(id int) ([]DataSource, error) {
	return client.withCtx.ListGroupDataSources(context.Background(), id)
}

func (client *ClientWithoutContext) AddGroupDataSource(id int, dsId int) (*DataSource, error) {
	return client.withCtx.AddGroupDataSource(context.Background(), id, dsId)
}

func (client *ClientWithoutContext) RemoveGroupDataSource(id int, dsId int) error {
	return client.withCtx.RemoveGroupDataSource(context.Background(), id, dsId)
}

func (client *ClientWithoutContext) UpdateGroupDataSource(id int, dsId int, input *UpdateGroupDataSourceInput) (*DataSource, error) {
	return client.withCtx.UpdateGroupDataSource(context.Background(), id, dsId, input)
}
