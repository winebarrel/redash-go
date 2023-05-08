// Code generated from alert.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) ListAlerts() ([]Alert, error) {
	return client.withCtx.ListAlerts(context.Background())
}

func (client *ClientWithoutContext) GetAlert(id int) (*Alert, error) {
	return client.withCtx.GetAlert(context.Background(), id)
}

func (client *ClientWithoutContext) CreateAlert(input *CreateAlertInput) (*Alert, error) {
	return client.withCtx.CreateAlert(context.Background(), input)
}

func (client *ClientWithoutContext) UpdateAlert(id int, input *UpdateAlertInput) (*Alert, error) {
	return client.withCtx.UpdateAlert(context.Background(), id, input)
}

func (client *ClientWithoutContext) DeleteAlert(id int) error {
	return client.withCtx.DeleteAlert(context.Background(), id)
}

func (client *ClientWithoutContext) ListAlertSubscriptions(id int) ([]AlertSubscription, error) {
	return client.withCtx.ListAlertSubscriptions(context.Background(), id)
}

func (client *ClientWithoutContext) AddAlertSubscription(id int, destinationId int) (*AlertSubscription, error) {
	return client.withCtx.AddAlertSubscription(context.Background(), id, destinationId)
}

func (client *ClientWithoutContext) RemoveAlertSubscription(id int, subscriptionId int) error {
	return client.withCtx.RemoveAlertSubscription(context.Background(), id, subscriptionId)
}

func (client *ClientWithoutContext) MuteAlert(id int) error {
	return client.withCtx.MuteAlert(context.Background(), id)
}

func (client *ClientWithoutContext) UnmuteAlert(id int) error {
	return client.withCtx.UnmuteAlert(context.Background(), id)
}
