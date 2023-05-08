// Code generated from ping.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) Ping() error { return client.withCtx.Ping(context.Background()) }
