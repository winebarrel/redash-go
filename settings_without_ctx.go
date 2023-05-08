// Code generated from settings.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) GetSettingsOrganization() (*SettingsOrganization, error) {
	return client.withCtx.GetSettingsOrganization(context.Background())
}

func (client *ClientWithoutContext) UpdateSettingsOrganization(input *UpdateSettingsOrganizationInput) (*SettingsOrganization, error) {
	return client.withCtx.UpdateSettingsOrganization(context.Background(), input)
}
