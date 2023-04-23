// Code generated from destinations.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) ListDestinations() ([]Destination, error) {
	return client.withCtx.ListDestinations(context.Background())
}

// Auto-generated
func (client *ClientWithoutContext) GetDestination(id int) (*Destination, error) {
	return client.withCtx.GetDestination(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) CreateDestination(input *CreateDestinationInput) (*Destination, error) {
	return client.withCtx.CreateDestination(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) DeleteDestination(id int) error {
	return client.withCtx.DeleteDestination(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) GetDestinationTypes() ([]DestinationType, error) {
	return client.withCtx.GetDestinationTypes(context.Background())
}
