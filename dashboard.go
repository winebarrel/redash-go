package redash

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/winebarrel/redash-go/internal/util"
)

type DashboardPage struct {
	Count    int         `json:"count"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Results  []Dashboard `json:"results"`
}

type Dashboard struct {
	CanEdit                 bool      `json:"can_edit"`
	CreatedAt               time.Time `json:"created_at"`
	DashboardFiltersEnabled bool      `json:"dashboard_filters_enabled"`
	ID                      int       `json:"id"`
	IsArchived              bool      `json:"is_archived"`
	IsDraft                 bool      `json:"is_draft"`
	IsFavorite              bool      `json:"is_favorite"`
	Layout                  any       `json:"layout"`
	Name                    string    `json:"name"`
	Slug                    string    `json:"slug"`
	Tags                    []string  `json:"tags"`
	UpdatedAt               time.Time `json:"updated_at"`
	User                    User      `json:"user"`
	UserID                  int       `json:"user_id"`
	Version                 int       `json:"version"`
	Widgets                 []Widget  `json:"widgets"`
}

type ListDashboardsInput struct {
	OnlyFavorites bool
	Page          int
	PageSize      int
}

// https://github.com/getredash/redash-toolbelt/blob/f6d2c40881fcacb411665c75f3afbe570533539d/redash_toolbelt/client.py#L51
func (client *Client) ListDashboards(ctx context.Context, input *ListDashboardsInput) (*DashboardPage, error) {
	params := map[string]string{}

	if input != nil {
		params["page"] = strconv.Itoa(input.Page)
		params["page_size"] = strconv.Itoa(input.PageSize)
		params["only_favorites"] = strconv.FormatBool(input.OnlyFavorites)
	}

	res, close, err := client.Get(ctx, "api/dashboards", params)
	defer close()

	if err != nil {
		return nil, err
	}

	page := &DashboardPage{}

	if err := util.UnmarshalBody(res, &page); err != nil {
		return nil, err
	}

	return page, nil
}

// https://github.com/getredash/redash-toolbelt/blob/f6d2c40881fcacb411665c75f3afbe570533539d/redash_toolbelt/client.py#L58
func (client *Client) GetDashboard(ctx context.Context, slug string) (*Dashboard, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/dashboards/%s", slug), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	dashbaord := &Dashboard{}

	if err := util.UnmarshalBody(res, &dashbaord); err != nil {
		return nil, err
	}

	return dashbaord, nil
}

// https://github.com/getredash/redash-toolbelt/blob/f6d2c40881fcacb411665c75f3afbe570533539d/redash_toolbelt/client.py#L24
func (client *Client) CreateFavoriteDashboard(ctx context.Context, slug string) error {
	_, close, err := client.Post(ctx, fmt.Sprintf("api/dashboards/%s/favorite", slug), nil)
	defer close()

	if err != nil {
		return err
	}

	return nil
}

type CreateDashboardInput struct {
	Name string `json:"name"`
}

// https://github.com/getredash/redash-toolbelt/blob/f6d2c40881fcacb411665c75f3afbe570533539d/redash_toolbelt/client.py#L90
func (client *Client) CreateDashboard(ctx context.Context, input *CreateDashboardInput) (*Dashboard, error) {
	res, close, err := client.Post(ctx, "api/dashboards", input)
	defer close()

	if err != nil {
		return nil, err
	}

	dashbaord := &Dashboard{}

	if err := util.UnmarshalBody(res, &dashbaord); err != nil {
		return nil, err
	}

	return dashbaord, nil
}

// https://github.com/getredash/redash/blob/5cf13afafe4a13c8db9da645d9667bc26fd118c5/redash/handlers/dashboards.py#L239-L247
type UpdateDashboardInput struct {
	DashboardFiltersEnabled bool     `json:"dashboard_filters_enabled,omitempty"`
	IsArchived              bool     `json:"is_archived,omitempty"`
	IsDraft                 bool     `json:"is_draft,omitempty"`
	Layout                  []any    `json:"layout,omitempty"`
	Name                    string   `json:"name,omitempty"`
	Options                 any      `json:"options,omitempty"`
	Tags                    []string `json:"tags,omitempty"`
	Version                 int      `json:"version,omitempty"`
}

// https://github.com/getredash/redash-toolbelt/blob/f6d2c40881fcacb411665c75f3afbe570533539d/redash_toolbelt/client.py#L93
func (client *Client) UpdateDashboard(ctx context.Context, id int, input *UpdateDashboardInput) (*Dashboard, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/dashboards/%d", id), input)
	defer close()

	if err != nil {
		return nil, err
	}

	dashbaord := &Dashboard{}

	if err := util.UnmarshalBody(res, &dashbaord); err != nil {
		return nil, err
	}

	return dashbaord, nil
}

func (client *Client) ArchiveDashboard(ctx context.Context, slug string) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/dashboards/%s", slug))
	defer close()

	if err != nil {
		return err
	}

	return nil
}
