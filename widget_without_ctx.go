// Code generated from widget.go using internal/gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) CreateWidget(input *CreateWidgetInput) (*Widget, error) {
	return client.withCtx.CreateWidget(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) DeleteWidget(id int) error {
	return client.withCtx.DeleteWidget(context.Background(), id)
}
