// Code generated from organization.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) GetOrganizationStatus() (*OrganizationStatus, error) {
	return client.withCtx.GetOrganizationStatus(context.Background())
}
