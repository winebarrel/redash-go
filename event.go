package redash

import (
	"context"
	"strconv"
	"time"

	"github.com/winebarrel/redash-go/internal/util"
)

type EventPage struct {
	Count    int     `json:"count"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
	Results  []Event `json:"results"`
}

type Event struct {
	Action     string            `json:"action"`
	Browser    string            `json:"browser"`
	CreatedAt  time.Time         `json:"created_at"`
	Details    map[string]string `json:"details"`
	Location   string            `json:"location"`
	ObjectID   string            `json:"object_id"`
	ObjectType string            `json:"object_type"`
	OrgID      int               `json:"org_id"`
	UserID     int               `json:"user_id"`
	UserName   string            `json:"user_name"`
}

type ListEventsInput struct {
	Page     int
	PageSize int
}

func (client *Client) ListEvents(ctx context.Context, input *ListEventsInput) (*EventPage, error) {
	params := map[string]string{}

	if input != nil {
		params["page"] = strconv.Itoa(input.Page)
		params["page_size"] = strconv.Itoa(input.PageSize)
	}

	res, close, err := client.Get(ctx, "api/events", params)
	defer close()

	if err != nil {
		return nil, err
	}

	page := &EventPage{}

	if err := util.UnmarshalBody(res, &page); err != nil {
		return nil, err
	}

	return page, nil
}
