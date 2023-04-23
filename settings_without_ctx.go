// Code generated from settings.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) GetSettingsOrganization() (*SettingsOrganization, error) {
	return client.withCtx.GetSettingsOrganization(context.Background())
}

// Auto-generated
func (client *ClientWithoutContext) UpdateSettingsOrganization(input *UpdateSettingsOrganizationInput) (*SettingsOrganization, error) {
	return client.withCtx.UpdateSettingsOrganization(context.Background(), input)
}
