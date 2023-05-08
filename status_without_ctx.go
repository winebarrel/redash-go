// Code generated from status.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) GetStatus() (*Status, error) {
	return client.withCtx.GetStatus(context.Background())
}
