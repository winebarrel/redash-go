// Code generated from status.go using genzfunc.go; DO NOT EDIT.
package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) GetStatus() (*Status, error) {
	return client.withCtx.GetStatus(context.Background())
}
