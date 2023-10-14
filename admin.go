//go:generate go run tools/withoutctx.go
package redash

import (
	"context"

	"github.com/winebarrel/redash-go/v2/internal/util"
)

type AdminQueriesOutdated struct {
	Queries   []Query `json:"queries"`
	UpdatedAt string  `json:"updated_at"`
}

func (client *Client) GetAdminQueriesOutdated(ctx context.Context) (*AdminQueriesOutdated, error) {
	res, close, err := client.Get(ctx, "/api/admin/queries/outdated", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	outdated := &AdminQueriesOutdated{}

	if err := util.UnmarshalBody(res, &outdated); err != nil {
		return nil, err
	}

	return outdated, nil
}

type AdminQuerisRqStatus struct {
	Queues  AdminQuerisRqStatusQueues   `json:"queues"`
	Workers []AdminQuerisRqStatusWorker `json:"workers"`
}

type AdminQuerisRqStatusQueues struct {
	Default  *AdminQuerisRqStatusDefault  `json:"default"`
	Emails   *AdminQuerisRqStatusEmails   `json:"emails"`
	Periodic *AdminQuerisRqStatusPeriodic `json:"periodic"`
	Queries  *AdminQuerisRqStatusQueries  `json:"queries"`
	Schemas  *AdminQuerisRqStatusSchemas  `json:"schemas"`
}

type AdminQuerisRqStatusDefault struct {
	Name    string `json:"name"`
	Queued  int    `json:"queued"`
	Started []any  `json:"started"`
}

type AdminQuerisRqStatusEmails struct {
	Name    string `json:"name"`
	Queued  int    `json:"queued"`
	Started []any  `json:"started"`
}

type AdminQuerisRqStatusPeriodic struct {
	Name    string `json:"name"`
	Queued  int    `json:"queued"`
	Started []any  `json:"started"`
}

type AdminQuerisRqStatusQueries struct {
	Name    string `json:"name"`
	Queued  int    `json:"queued"`
	Started []any  `json:"started"`
}

type AdminQuerisRqStatusSchemas struct {
	Name    string `json:"name"`
	Queued  int    `json:"queued"`
	Started []any  `json:"started"`
}

type AdminQuerisRqStatusWorker struct {
	BirthDate        string  `json:"birth_date"`
	CurrentJob       string  `json:"current_job"`
	FailedJobs       int     `json:"failed_jobs"`
	Hostname         string  `json:"hostname"`
	LastHeartbeat    string  `json:"last_heartbeat"`
	Name             string  `json:"name"`
	Pid              int     `json:"pid"`
	Queues           string  `json:"queues"`
	State            string  `json:"state"`
	SuccessfulJobs   int     `json:"successful_jobs"`
	TotalWorkingTime float64 `json:"total_working_time"`
}

func (client *Client) GetAdminQueriesRqStatus(ctx context.Context) (*AdminQuerisRqStatus, error) {
	res, close, err := client.Get(ctx, "/api/admin/queries/rq_status", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	rqStatus := &AdminQuerisRqStatus{}

	if err := util.UnmarshalBody(res, &rqStatus); err != nil {
		return nil, err
	}

	return rqStatus, nil
}
