//go:generate go run tools/withoutctx.go
package redash

import (
	"context"
	"fmt"

	"github.com/winebarrel/redash-go/v2/internal/util"
)

const (
	// see https://redash.io/help/user-guide/integrations-and-api/api#Jobs
	JobStatusPending   = 1
	JobStatusStarted   = 2
	JobStatusSuccess   = 3
	JobStatusFailure   = 4
	JobStatusCancelled = 5
)

type JobResponse struct {
	Job Job `json:"job"`
}

type Job struct {
	Error         string `json:"error"`
	ID            string `json:"id"`
	QueryResultID int    `json:"query_result_id"`
	Status        int    `json:"status"`
	UpdatedAt     any    `json:"updated_at"`
}

func (client *Client) GetJob(ctx context.Context, id string) (*JobResponse, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/jobs/%s", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	jobRes := &JobResponse{}

	if err := util.UnmarshalBody(res, &jobRes); err != nil {
		return nil, err
	}

	return jobRes, nil
}
