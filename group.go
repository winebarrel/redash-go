package redash

import (
	"context"
	"fmt"
	"time"

	"github.com/winebarrel/redash-go/internal/util"
)

type Group struct {
	CreatedAt   time.Time `json:"created_at"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Permissions []string  `json:"permissions"`
	Type        string    `json:"type"`
}

func (client *Client) ListGroups(ctx context.Context) ([]Group, error) {
	res, close, err := client.Get(ctx, "api/groups", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	groups := []Group{}

	if err := util.UnmarshalBody(res, &groups); err != nil {
		return nil, err
	}

	return groups, nil
}

func (client *Client) GetGroup(ctx context.Context, id int) (*Group, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/groups/%d", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	group := &Group{}

	if err := util.UnmarshalBody(res, &group); err != nil {
		return nil, err
	}

	return group, nil
}

type CreateGroupInput struct {
	Name string `json:"name"`
}

func (client *Client) CreateGroup(ctx context.Context, input *CreateGroupInput) (*Group, error) {
	res, close, err := client.Post(ctx, "api/groups", input)
	defer close()

	if err != nil {
		return nil, err
	}

	group := &Group{}

	if err := util.UnmarshalBody(res, &group); err != nil {
		return nil, err
	}

	return group, nil
}

func (client *Client) DeleteGroup(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/groups/%d", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) ListGroupMembers(ctx context.Context, id int) ([]User, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/groups/%d/members", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	users := []User{}

	if err := util.UnmarshalBody(res, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (client *Client) AddGroupMember(ctx context.Context, id int, userId int) (*User, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/groups/%d/members", id), map[string]int{"user_id": userId})
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

func (client *Client) RemoveGroupMember(ctx context.Context, id int, userId int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/groups/%d/members/%d", id, userId))
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) ListGroupDataSources(ctx context.Context, id int) ([]DataSource, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/groups/%d/data_sources", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	dataSources := []DataSource{}

	if err := util.UnmarshalBody(res, &dataSources); err != nil {
		return nil, err
	}

	return dataSources, nil
}

func (client *Client) AddGroupDataSource(ctx context.Context, id int, dsId int) (*DataSource, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/groups/%d/data_sources", id), map[string]int{"data_source_id": dsId})
	defer close()

	if err != nil {
		return nil, err
	}

	dataSource := &DataSource{}

	if err := util.UnmarshalBody(res, &dataSource); err != nil {
		return nil, err
	}

	return dataSource, nil
}

func (client *Client) RemoveGroupDataSource(ctx context.Context, id int, dsId int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/groups/%d/data_sources/%d", id, dsId))
	defer close()

	if err != nil {
		return err
	}

	return nil
}

type UpdateGroupDataSourceInput struct {
	ViewOnly bool `json:"view_only"`
}

func (client *Client) UpdateGroupDataSource(ctx context.Context, id int, dsId int, input *UpdateGroupDataSourceInput) (*DataSource, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/groups/%d/data_sources/%d", id, dsId), input)
	defer close()

	if err != nil {
		return nil, err
	}

	dataSource := &DataSource{}

	if err := util.UnmarshalBody(res, &dataSource); err != nil {
		return nil, err
	}

	return dataSource, nil
}
