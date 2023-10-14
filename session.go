//go:generate go run tools/withoutctx.go
package redash

import (
	"context"

	"github.com/winebarrel/redash-go/v2/internal/util"
)

type Session struct {
	ClientConfig ClientConfig `json:"client_config"`
	Messages     []string     `json:"messages"`
	OrgSlug      string       `json:"org_slug"`
	User         User         `json:"user"`
}

func (client *Client) TestCredentials(ctx context.Context) error {
	_, close, err := client.Get(ctx, "api/session", nil)
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) GetSession(ctx context.Context) (*Session, error) {
	res, close, err := client.Get(ctx, "api/session", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	session := &Session{}

	if err := util.UnmarshalBody(res, &session); err != nil {
		return nil, err
	}

	return session, nil
}
