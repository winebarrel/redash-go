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

type UpdateVisualizationInput struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Options     any    `json:"options,omitempty"`
	Type        string `json:"type,omitempty"`
}

func (client *Client) UpdateVisualization(ctx context.Context, id int, input *UpdateVisualizationInput) (*Visualization, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/visualizations/%d", id), input)
	defer close()

	if err != nil {
		return nil, err
	}

	viz := &Visualization{}

	if err := util.UnmarshalBody(res, &viz); err != nil {
		return nil, err
	}

	return viz, nil
}
