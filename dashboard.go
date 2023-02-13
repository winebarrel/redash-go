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

// idOrSlug:
//
//	v8: int
//	v10: string
func (client *Client) GetDashboard(ctx context.Context, idOrSlug any) (*Dashboard, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/dashboards/%v", idOrSlug), nil)
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

// idOrSlug:
//
//	v8: int
//	v10: string
func (client *Client) CreateFavoriteDashboard(ctx context.Context, idOrSlug any) error {
	_, close, err := client.Post(ctx, fmt.Sprintf("api/dashboards/%v/favorite", idOrSlug), nil)
	defer close()

	if err != nil {
		return err
	}

	return nil
}

type CreateDashboardInput struct {
	Name string `json:"name"`
}

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

// idOrSlug:
//
//	v8: int
//	v10: string
func (client *Client) ArchiveDashboard(ctx context.Context, idOrSlug any) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/dashboards/%v", idOrSlug))
	defer close()

	if err != nil {
		return err
	}

	return nil
}
