// Code generated from job.go using genzfunc.go; DO NOT EDIT.
package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) GetJob(id string) (*JobResponse, error) {
	return client.withCtx.GetJob(context.Background(), id)
}
