// Code generated from event.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) ListEvents(input *ListEventsInput) (*EventPage, error) {
	return client.withCtx.ListEvents(context.Background(), input)
}
