// Code generated from user.go using internal/gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) ListUsers(input *ListUsersInput) (*UserPage, error) {
	return client.withCtx.ListUsers(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) GetUser(id int) (*User, error) {
	return client.withCtx.GetUser(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) CreateUser(input *CreateUsersInput) (*User, error) {
	return client.withCtx.CreateUser(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) DeleteUser(id int) error {
	return client.withCtx.DeleteUser(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) DisableUser(id int) (*User, error) {
	return client.withCtx.DisableUser(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) EnableUser(id int) (*User, error) {
	return client.withCtx.EnableUser(context.Background(), id)
}
