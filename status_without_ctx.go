// Code generated from status.go using internal/gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) GetStatus() (*Status, error) {
	return client.withCtx.GetStatus(context.Background())
}
