// Code generated from session.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) TestCredentials() error {
	return client.withCtx.TestCredentials(context.Background())
}

func (client *ClientWithoutContext) GetSession() (*Session, error) {
	return client.withCtx.GetSession(context.Background())
}
