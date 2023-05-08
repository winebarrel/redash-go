// Code generated from config.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) GetConfig() (*Config, error) {
	return client.withCtx.GetConfig(context.Background())
}
