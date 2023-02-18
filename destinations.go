//go:generate go run gen/withoutctx.go
package redash

import (
	"context"
	"fmt"

	"github.com/winebarrel/redash-go/internal/util"
)

type Destination struct {
	Icon    string         `json:"icon"`
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Options map[string]any `json:"options"`
	Type    string         `json:"type"`
}

func (client *Client) ListDestinations(ctx context.Context) ([]Destination, error) {
	res, close, err := client.Get(ctx, "api/destinations", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	destinations := []Destination{}

	if err := util.UnmarshalBody(res, &destinations); err != nil {
		return nil, err
	}

	return destinations, nil
}

func (client *Client) GetDestination(ctx context.Context, id int) (*Destination, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/destinations/%d", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	dest := &Destination{}

	if err := util.UnmarshalBody(res, &dest); err != nil {
		return nil, err
	}

	return dest, nil
}

type CreateDestinationInput struct {
	Name    string         `json:"name"`
	Options map[string]any `json:"options"`
	Type    string         `json:"type"`
}

func (client *Client) CreateDestination(ctx context.Context, input *CreateDestinationInput) (*Destination, error) {
	res, close, err := client.Post(ctx, "api/destinations", input)
	defer close()

	if err != nil {
		return nil, err
	}

	dest := &Destination{}

	if err := util.UnmarshalBody(res, &dest); err != nil {
		return nil, err
	}

	return dest, nil
}

func (client *Client) DeleteDestination(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/destinations/%d", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}
