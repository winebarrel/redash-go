package redash

import (
	"context"
)

// https://github.com/getredash/redash-toolbelt/blob/f6d2c40881fcacb411665c75f3afbe570533539d/redash_toolbelt/client.py#L10
func (client *Client) TestCredentials(ctx context.Context) error {
	_, err := client.Get(ctx, "api/session", nil)

	if err != nil {
		return err
	}

	return nil
}
