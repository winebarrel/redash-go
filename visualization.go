// TODO: This API is not well tested :_(
package redash

import (
	"context"
	"fmt"
	"time"

	"github.com/winebarrel/redash-go/internal/util"
)

type Visualization struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Options     any       `json:"options"`
	Query       Query     `json:"query"`
	Type        string    `json:"type"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// https://github.com/getredash/redash/blob/5cf13afafe4a13c8db9da645d9667bc26fd118c5/redash/handlers/visualizations.py#L29
type UpdateVisualizationInput struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Options     any    `json:"options,omitempty"`
	Type        string `json:"type,omitempty"`
}

// https://github.com/getredash/redash-toolbelt/blob/f6d2c40881fcacb411665c75f3afbe570533539d/redash_toolbelt/client.py#L98
func (client *Client) UpdateVisualization(ctx context.Context, id int, input *UpdateVisualizationInput) (*Visualization, error) {
	res, err := client.Post(ctx, fmt.Sprintf("api/visualizations/%d", id), input)

	if err != nil {
		return nil, err
	}

	viz := &Visualization{}

	if err := util.UnmarshalBody(res, &viz); err != nil {
		return nil, err
	}

	return viz, nil
}
