// Code generated from config.go using gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) GetConfig() (*Config, error) {
	return client.withCtx.GetConfig(context.Background())
}