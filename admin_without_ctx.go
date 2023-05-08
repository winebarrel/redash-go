// Code generated from admin.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) GetAdminQueriesOutdated() (*AdminQueriesOutdated, error) {
	return client.withCtx.GetAdminQueriesOutdated(context.Background())
}

func (client *ClientWithoutContext) GetAdminQueriesRqStatus() (*AdminQuerisRqStatus, error) {
	return client.withCtx.GetAdminQueriesRqStatus(context.Background())
}
