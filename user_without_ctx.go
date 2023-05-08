// Code generated from user.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) ListUsers(input *ListUsersInput) (*UserPage, error) {
	return client.withCtx.ListUsers(context.Background(), input)
}

func (client *ClientWithoutContext) GetUser(id int) (*User, error) {
	return client.withCtx.GetUser(context.Background(), id)
}

func (client *ClientWithoutContext) CreateUser(input *CreateUsersInput) (*User, error) {
	return client.withCtx.CreateUser(context.Background(), input)
}

func (client *ClientWithoutContext) UpdateUser(id int, input *UpdateUserInput) (*User, error) {
	return client.withCtx.UpdateUser(context.Background(), id, input)
}

func (client *ClientWithoutContext) DeleteUser(id int) error {
	return client.withCtx.DeleteUser(context.Background(), id)
}

func (client *ClientWithoutContext) DisableUser(id int) (*User, error) {
	return client.withCtx.DisableUser(context.Background(), id)
}

func (client *ClientWithoutContext) EnableUser(id int) (*User, error) {
	return client.withCtx.EnableUser(context.Background(), id)
}
