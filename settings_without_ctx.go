// Code generated from settings.go using internal/gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) GetSettingsOrganization() (*SettingsOrganization, error) {
	return client.withCtx.GetSettingsOrganization(context.Background())
}