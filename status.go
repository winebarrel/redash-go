//go:generate go run gen/withoutctx.go
package redash

import (
	"context"

	"github.com/winebarrel/redash-go/internal/util"
)

type Status struct {
	DashboardsCount         int                   `json:"dashboards_count"`
	DatabaseMetrics         StatusDatabaseMetrics `json:"database_metrics"`
	Manager                 StatusManager         `json:"manager"`
	QueriesCount            int                   `json:"queries_count"`
	QueryResultsCount       int                   `json:"query_results_count"`
	RedisUsedMemory         int                   `json:"redis_used_memory"`
	RedisUsedMemoryHuman    string                `json:"redis_used_memory_human"`
	UnusedQueryResultsCount int                   `json:"unused_query_results_count"`
	Version                 string                `json:"version"`
	WidgetsCount            int                   `json:"widgets_count"`
	Workers                 []any                 `json:"workers"`
}

type StatusDatabaseMetrics struct {
	Metrics [][]any `json:"metrics"`
}

type StatusManager struct {
	LastRefreshAt        string              `json:"last_refresh_at"`
	OutdatedQueriesCount string              `json:"outdated_queries_count"`
	QueryIds             string              `json:"query_ids"`
	Queues               StatusManagerQueues `json:"queues"`
}

type StatusManagerQueues struct {
	Celery           StatusManagerQueuesCelery           `json:"celery"`
	Queries          StatusManagerQueuesQueries          `json:"queries"`
	ScheduledQueries StatusManagerQueuesScheduledQueries `json:"scheduled_queries"`
}

type StatusManagerQueuesCelery struct {
	Size int `json:"size"`
}

type StatusManagerQueuesQueries struct {
	Size int `json:"size"`
}

type StatusManagerQueuesScheduledQueries struct {
	Size int `json:"size"`
}

func (client *Client) GetStatus(ctx context.Context) (*Status, error) {
	res, close, err := client.Get(ctx, "status.json", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	status := &Status{}

	if err := util.UnmarshalBody(res, &status); err != nil {
		return nil, err
	}

	return status, nil
}
