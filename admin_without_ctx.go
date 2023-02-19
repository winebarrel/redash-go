// Code generated from admin.go using internal/gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) GetAdminQueriesOutdated() (*AdminQueriesOutdated, error) {
	return client.withCtx.GetAdminQueriesOutdated(context.Background())
}

// Auto-generated
func (client *ClientWithoutContext) GetAdminQueriesRqStatus() (*AdminQuerisRqStatus, error) {
	return client.withCtx.GetAdminQueriesRqStatus(context.Background())
}
