package redash

import (
	"context"
	"fmt"
	"time"

	"github.com/winebarrel/redash-go/internal/util"
)

type QuerySnippet struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
	Snippet     string    `json:"snippet"`
	Trigger     string    `json:"trigger"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `json:"user"`
}

func (client *Client) ListQuerySnippets(ctx context.Context) ([]QuerySnippet, error) {
	res, close, err := client.Get(ctx, "api/query_snippets", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	querySnippets := []QuerySnippet{}

	if err := util.UnmarshalBody(res, &querySnippets); err != nil {
		return nil, err
	}

	return querySnippets, nil
}

func (client *Client) GetQuerySnippet(ctx context.Context, id int) (*QuerySnippet, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/query_snippets/%d", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	querySnippet := &QuerySnippet{}

	if err := util.UnmarshalBody(res, &querySnippet); err != nil {
		return nil, err
	}

	return querySnippet, nil
}

type CreateQuerySnippetInput struct {
	Description string `json:"description"`
	Snippet     string `json:"snippet"`
	Trigger     string `json:"trigger"`
}

func (client *Client) CreateQuerySnippet(ctx context.Context, input *CreateQuerySnippetInput) (*QuerySnippet, error) {
	res, close, err := client.Post(ctx, "api/query_snippets", input)
	defer close()

	if err != nil {
		return nil, err
	}

	querySnippet := &QuerySnippet{}

	if err := util.UnmarshalBody(res, &querySnippet); err != nil {
		return nil, err
	}

	return querySnippet, nil
}

type UpdateQuerySnippetInput struct {
	Description string `json:"description,omitempty"`
	Snippet     string `json:"snippet,omitempty"`
	Trigger     string `json:"trigger,omitempty"`
}

func (client *Client) UpdateQuerySnippet(ctx context.Context, id int, input *UpdateQuerySnippetInput) (*QuerySnippet, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/query_snippets/%d", id), input)
	defer close()

	if err != nil {
		return nil, err
	}

	querySnippet := &QuerySnippet{}

	if err := util.UnmarshalBody(res, &querySnippet); err != nil {
		return nil, err
	}

	return querySnippet, nil
}

func (client *Client) DeleteQuerySnippet(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/query_snippets/%d", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}
