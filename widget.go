//go:generate go run tools/withoutctx.go
package redash

import (
	"context"
	"fmt"
	"time"

	"github.com/winebarrel/redash-go/internal/util"
)

type Widget struct {
	CreatedAt     time.Time      `json:"created_at"`
	DashboardID   int            `json:"dashboard_id"`
	ID            int            `json:"id"`
	Options       map[string]any `json:"options"`
	Text          string         `json:"text"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Visualization *Visualization `json:"visualization"`
	Width         int            `json:"width"`
}

type CreateWidgetInput struct {
	DashboardID     int            `json:"dashboard_id"`
	Options         map[string]any `json:"options"`
	Text            string         `json:"text,omitempty"`
	VisualizationID int            `json:"visualization_id"`
	Width           int            `json:"width"`
}

func (client *Client) CreateWidget(ctx context.Context, input *CreateWidgetInput) (*Widget, error) {
	// workaround
	if input.Width < 1 {
		input.Width = 1
	}

	res, close, err := client.Post(ctx, "api/widgets", input)
	defer close()

	if err != nil {
		return nil, err
	}

	widget := &Widget{}

	if err := util.UnmarshalBody(res, &widget); err != nil {
		return nil, err
	}

	return widget, nil
}

func (client *Client) DeleteWidget(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/widgets/%d", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}
