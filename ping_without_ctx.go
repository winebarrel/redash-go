// Code generated from ping.go using gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) Ping() error { return client.withCtx.Ping(context.Background()) }
