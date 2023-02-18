//go:generate go run gen/withoutctx.go
package redash

import (
	"context"
	"fmt"
	"time"

	"github.com/winebarrel/redash-go/internal/util"
)

type Alert struct {
	CreatedAt       time.Time    `json:"created_at"`
	ID              int          `json:"id"`
	LastTriggeredAt time.Time    `json:"last_triggered_at"`
	Name            string       `json:"name"`
	Options         AlertOptions `json:"options"`
	Query           Query        `json:"query"`
	Rearm           int          `json:"rearm"`
	State           string       `json:"state"`
	UpdatedAt       time.Time    `json:"updated_at"`
	User            User         `json:"user"`
}

type AlertOptions struct {
	Column string `json:"column"`
	Op     string `json:"op"`
	Value  int    `json:"value"`
}

type AlertSubscription struct {
	AlertID     int          `json:"alert_id"`
	Destination *Destination `json:"destination"`
	ID          int          `json:"id"`
	User        User         `json:"user"`
}

func (client *Client) ListAlerts(ctx context.Context) ([]Alert, error) {
	res, close, err := client.Get(ctx, "api/alerts", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	alerts := []Alert{}

	if err := util.UnmarshalBody(res, &alerts); err != nil {
		return nil, err
	}

	return alerts, nil
}

func (client *Client) GetAlert(ctx context.Context, id int) (*Alert, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/alerts/%d", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	alert := &Alert{}

	if err := util.UnmarshalBody(res, &alert); err != nil {
		return nil, err
	}

	return alert, nil
}

type CreateAlertInput struct {
	Name    string             `json:"name"`
	Options CreateAlertOptions `json:"options"`
	QueryId int                `json:"query_id"`
	Rearm   int                `json:"rearm,omitempty"`
}

type CreateAlertOptions struct {
	Column string `json:"column"`
	Op     string `json:"op"`
	Value  int    `json:"value"`
}

func (client *Client) CreateAlert(ctx context.Context, input *CreateAlertInput) (*Alert, error) {
	res, close, err := client.Post(ctx, "api/alerts", input)
	defer close()

	if err != nil {
		return nil, err
	}

	alert := &Alert{}

	if err := util.UnmarshalBody(res, &alert); err != nil {
		return nil, err
	}

	return alert, nil
}

type UpdateAlertInput struct {
	Name    string              `json:"name,omitempty"`
	Options *UpdateAlertOptions `json:"options,omitempty"`
	QueryId int                 `json:"query_id,omitempty"`
	Rearm   int                 `json:"rearm,omitempty"`
}

type UpdateAlertOptions struct {
	Column string `json:"column"`
	Value  int    `json:"value"`
	Op     string `json:"op"`
}

func (client *Client) UpdateAlert(ctx context.Context, id int, input *UpdateAlertInput) (*Alert, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/alerts/%d", id), input)
	defer close()

	if err != nil {
		return nil, err
	}

	alert := &Alert{}

	if err := util.UnmarshalBody(res, &alert); err != nil {
		return nil, err
	}

	return alert, nil
}

func (client *Client) DeleteAlert(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/alerts/%d", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) ListAlertSubscriptions(ctx context.Context, id int) ([]AlertSubscription, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/alerts/%d/subscriptions", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	subs := []AlertSubscription{}

	if err := util.UnmarshalBody(res, &subs); err != nil {
		return nil, err
	}

	return subs, nil
}

func (client *Client) AddAlertSubscription(ctx context.Context, id int, destinationId int) (*AlertSubscription, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/alerts/%d/subscriptions", id), map[string]int{"destination_id": destinationId})
	defer close()

	if err != nil {
		return nil, err
	}

	sub := &AlertSubscription{}

	if err := util.UnmarshalBody(res, &sub); err != nil {
		return nil, err
	}

	return sub, nil
}

func (client *Client) RemoveAlertSubscription(ctx context.Context, id int, subscriptionId int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/alerts/%d/subscriptions/%d", id, subscriptionId))
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) MuteAlert(ctx context.Context, id int) error {
	_, close, err := client.Post(ctx, fmt.Sprintf("api/alerts/%d/mute", id), nil)
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) UnmuteAlert(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/alerts/%d/mute", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}
