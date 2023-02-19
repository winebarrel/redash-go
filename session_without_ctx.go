// Code generated from session.go using internal/gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) TestCredentials() error {
	return client.withCtx.TestCredentials(context.Background())
}

// Auto-generated
func (client *ClientWithoutContext) GetSession() (*Session, error) {
	return client.withCtx.GetSession(context.Background())
}
