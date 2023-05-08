// Code generated from widget.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) CreateWidget(input *CreateWidgetInput) (*Widget, error) {
	return client.withCtx.CreateWidget(context.Background(), input)
}

func (client *ClientWithoutContext) DeleteWidget(id int) error {
	return client.withCtx.DeleteWidget(context.Background(), id)
}
