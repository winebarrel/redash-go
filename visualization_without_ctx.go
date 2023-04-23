// Code generated from visualization.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) UpdateVisualization(id int, input *UpdateVisualizationInput) (*Visualization, error) {
	return client.withCtx.UpdateVisualization(context.Background(), id, input)
}
