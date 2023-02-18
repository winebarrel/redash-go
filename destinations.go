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

type DestinationType struct {
	ConfigurationSchema DestinationTypeConfigurationSchema `json:"configuration_schema"`
	Icon                string                             `json:"icon"`
	Name                string                             `json:"name"`
	Type                string                             `json:"type"`
}

type DestinationTypeConfigurationSchema struct {
	ExtraOptions []string                                              `json:"extra_options"`
	Properties   map[string]DestinationTypeConfigurationSchemaProperty `json:"properties"`
	Required     []string                                              `json:"required"`
	Secret       any                                                   `json:"secret"`
	Type         string                                                `json:"type"`
}

type DestinationTypeConfigurationSchemaProperty struct {
	Default any    `json:"default"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

func (client *Client) GetDestinationTypes(ctx context.Context) ([]DestinationType, error) {
	res, close, err := client.Get(ctx, "api/destinations/types", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	types := []DestinationType{}

	if err := util.UnmarshalBody(res, &types); err != nil {
		return nil, err
	}

	return types, nil
}
