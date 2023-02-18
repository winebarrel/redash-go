//go:generate go run gen/withoutctx.go
package redash

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/winebarrel/redash-go/internal/util"
)

type UserPage struct {
	Count    int    `json:"count"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Results  []User `json:"results"`
}

type User struct {
	ActiveAt            time.Time `json:"active_at"`
	APIKey              string    `json:"api_key"`
	AuthType            string    `json:"auth_type"`
	CreatedAt           time.Time `json:"created_at"`
	DisabledAt          time.Time `json:"disabled_at"`
	Email               string    `json:"email"`
	Groups              []any     `json:"groups"`
	ID                  int       `json:"id"`
	IsDisabled          bool      `json:"is_disabled"`
	IsEmailVerified     bool      `json:"is_email_verified"`
	IsInvitationPending bool      `json:"is_invitation_pending"`
	Name                string    `json:"name"`
	ProfileImageURL     string    `json:"profile_image_url"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type ListUsersInput struct {
	Page     int
	PageSize int
}

func (client *Client) ListUsers(ctx context.Context, input *ListUsersInput) (*UserPage, error) {
	params := map[string]string{}

	if input != nil {
		params["page"] = strconv.Itoa(input.Page)
		params["page_size"] = strconv.Itoa(input.PageSize)
	}

	res, close, err := client.Get(ctx, "api/users", params)
	defer close()

	if err != nil {
		return nil, err
	}

	page := &UserPage{}

	if err := util.UnmarshalBody(res, &page); err != nil {
		return nil, err
	}

	return page, nil
}

func (client *Client) GetUser(ctx context.Context, id int) (*User, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/users/%d", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	user := &User{}

	if err := util.UnmarshalBody(res, &user); err != nil {
		return nil, err
	}

	return user, nil
}

type CreateUsersInput struct {
	AuthType string `json:"auth_type"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func (client *Client) CreateUser(ctx context.Context, input *CreateUsersInput) (*User, error) {
	res, close, err := client.Post(ctx, "api/users", input)
	defer close()

	if err != nil {
		return nil, err
	}

	user := &User{}

	if err := util.UnmarshalBody(res, &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (client *Client) DeleteUser(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/users/%d", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) DisableUser(ctx context.Context, id int) (*User, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/users/%d/disable", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	user := &User{}

	if err := util.UnmarshalBody(res, &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (client *Client) EnableUser(ctx context.Context, id int) (*User, error) {
	res, close, err := client.Delete(ctx, fmt.Sprintf("api/users/%d/disable", id))
	defer close()

	if err != nil {
		return nil, err
	}

	user := &User{}

	if err := util.UnmarshalBody(res, &user); err != nil {
		return nil, err
	}

	return user, nil
}
