//go:generate go run internal/gen/withoutctx.go
package redash

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

const (
	pong = "PONG."
)

func (client *Client) Ping(ctx context.Context) error {
	res, close, err := client.Get(ctx, "ping", nil)
	defer close()

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if !bytes.Equal(body, []byte(pong)) {
		return fmt.Errorf("invalid ping response: %s", body)
	}

	return nil
}
