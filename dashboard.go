//go:generate go run tools/withoutctx.go
package redash

import (
	"context"
	"fmt"
	"time"

	"github.com/winebarrel/redash-go/v2/internal/util"
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
	OnlyFavorites bool   `url:"only_favorites,omitempty"`
	Page          int    `url:"page,omitempty"`
	PageSize      int    `url:"page_size,omitempty"`
	Q             string `url:"q,omitempty"`
}

func (client *Client) ListDashboards(ctx context.Context, input *ListDashboardsInput) (*DashboardPage, error) {
	res, close, err := client.Get(ctx, "api/dashboards", input)
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

func (client *Client) GetDashboard(ctx context.Context, id int) (*Dashboard, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/dashboards/%d", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	dashboard := &Dashboard{}

	if err := util.UnmarshalBody(res, &dashboard); err != nil {
		return nil, err
	}

	return dashboard, nil
}

func (client *Client) CreateFavoriteDashboard(ctx context.Context, id int) error {
	_, close, err := client.Post(ctx, fmt.Sprintf("api/dashboards/%d/favorite", id), nil)
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

	dashboard := &Dashboard{}

	if err := util.UnmarshalBody(res, &dashboard); err != nil {
		return nil, err
	}

	return dashboard, nil
}

type UpdateDashboardInput struct {
	DashboardFiltersEnabled bool      `json:"dashboard_filters_enabled,omitempty"`
	IsArchived              bool      `json:"is_archived,omitempty"`
	IsDraft                 bool      `json:"is_draft,omitempty"`
	Layout                  []any     `json:"layout,omitempty"`
	Name                    string    `json:"name,omitempty"`
	Options                 any       `json:"options,omitempty"`
	Tags                    *[]string `json:"tags,omitempty"`
	Version                 int       `json:"version,omitempty"`
}

func (client *Client) UpdateDashboard(ctx context.Context, id int, input *UpdateDashboardInput) (*Dashboard, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/dashboards/%d", id), input)
	defer close()

	if err != nil {
		return nil, err
	}

	dashboard := &Dashboard{}

	if err := util.UnmarshalBody(res, &dashboard); err != nil {
		return nil, err
	}

	return dashboard, nil
}

func (client *Client) ArchiveDashboard(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/dashboards/%d", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}

type DashboardTags struct {
	Tags []DashboardTagsTag `json:"tags"`
}

type DashboardTagsTag struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

func (client *Client) GetDashboardTags(ctx context.Context) (*DashboardTags, error) {
	res, close, err := client.Get(ctx, "api/dashboards/tags", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	tags := &DashboardTags{}

	if err := util.UnmarshalBody(res, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}

type ListMyDashboardsInput struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Q        string `url:"q,omitempty"`
}

func (client *Client) ListMyDashboards(ctx context.Context, input *ListMyDashboardsInput) (*DashboardPage, error) {
	res, close, err := client.Get(ctx, "api/dashboards/my", input)
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

type ListFavoriteDashboardsInput struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Q        string `url:"q,omitempty"`
}

func (client *Client) ListFavoriteDashboards(ctx context.Context, input *ListFavoriteDashboardsInput) (*DashboardPage, error) {
	res, close, err := client.Get(ctx, "api/dashboards/favorites", input)
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

type ShareDashboardOutput struct {
	APIKey    string `json:"api_key"`
	PublicURL string `json:"public_url"`
}

func (client *Client) ShareDashboard(ctx context.Context, id int) (*ShareDashboardOutput, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/dashboards/%d/share", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	output := &ShareDashboardOutput{}

	if err := util.UnmarshalBody(res, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func (client *Client) UnshareDashboard(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/dashboards/%d/share", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}
