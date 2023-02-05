package redash

import (
	"context"
	"fmt"

	"github.com/winebarrel/redash-go/internal/util"
)

type JobResponse struct {
	Job Job `json:"job"`
}

type Job struct {
	Error         string `json:"error"`
	ID            string `json:"id"`
	QueryResultID int    `json:"query_result_id"`
	Status        int    `json:"status"`
	UpdatedAt     int    `json:"updated_at"`
}

func (client *Client) GetJob(ctx context.Context, id string) (*JobResponse, error) {
	res, err := client.Get(ctx, fmt.Sprintf("api/jobs/%s", id), nil)

	if err != nil {
		return nil, err
	}

	jobRes := &JobResponse{}

	if err := util.UnmarshalBody(res, &jobRes); err != nil {
		return nil, err
	}

	return jobRes, nil
}
