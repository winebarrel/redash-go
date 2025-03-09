//go:generate go run tools/withoutctx.go
package redash

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"time"

	"github.com/winebarrel/redash-go/v2/internal/util"
)

type QueryPage struct {
	Count    int     `json:"count"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
	Results  []Query `json:"results"`
}

type Query struct {
	APIKey            string          `json:"api_key"`
	CanEdit           bool            `json:"can_edit"`
	CreatedAt         time.Time       `json:"created_at"`
	DataSourceID      int             `json:"data_source_id"`
	Description       string          `json:"description"`
	ID                int             `json:"id"`
	IsArchived        bool            `json:"is_archived"`
	IsDraft           bool            `json:"is_draft"`
	IsFavorite        bool            `json:"is_favorite"`
	IsSafe            bool            `json:"is_safe"`
	LastModifiedBy    *User           `json:"last_modified_by"`
	LastModifiedByID  int             `json:"last_modified_by_id"`
	LatestQueryDataID int             `json:"latest_query_data_id"`
	Name              string          `json:"name"`
	Options           QueryOptions    `json:"options"`
	Query             string          `json:"query"`
	QueryHash         string          `json:"query_hash"`
	RetrievedAt       time.Time       `json:"retrieved_at"`
	Runtime           float64         `json:"runtime"`
	Schedule          *QueueSchedule  `json:"schedule"`
	Tags              []string        `json:"tags"`
	UpdatedAt         time.Time       `json:"updated_at"`
	User              User            `json:"user"`
	Version           int             `json:"version"`
	Visualizations    []Visualization `json:"visualizations"`
}

type QueryOptions struct {
	Parameters []QueryOptionsParameter `json:"parameters"`
}

type QueryOptionsParameter struct {
	Global             bool                                     `json:"global"`
	Type               string                                   `json:"type"`
	Name               string                                   `json:"name"`
	Value              any                                      `json:"value,omitempty"`
	Title              string                                   `json:"title"`
	Regex              string                                   `json:"regex,omitempty"`
	EnumOptions        string                                   `json:"enumOptions,omitempty"`
	MultiValuesOptions *QueryOptionsParameterMultiValuesOptions `json:"multiValuesOptions,omitempty"`
	QueryID            int                                      `json:"queryId,omitempty"`
}

type QueryOptionsParameterMultiValuesOptions struct {
	Prefix    string `json:"prefix"`
	Suffix    string `json:"suffix"`
	Separator string `json:"separator"`
}

type QueueSchedule struct {
	DayOfWeek string `json:"day_of_week"`
	Interval  int    `json:"interval"`
	Time      string `json:"time"`
	Until     string `json:"until"`
}

type ListQueriesInput struct {
	OnlyFavorites bool   `url:"only_favorites,omitempty"`
	Page          int    `url:"page,omitempty"`
	PageSize      int    `url:"page_size,omitempty"`
	Q             string `url:"q,omitempty"`
}

type GetQueryResultsOutput struct {
	QueryResult GetQueryResultsOutputQueryResult `json:"query_result"`
}

type GetQueryResultsOutputQueryResult struct {
	ID           int                                  `json:"id"`
	QueryHash    string                               `json:"query_hash"`
	Query        string                               `json:"query"`
	Data         GetQueryResultsOutputQueryResultData `json:"data"`
	DataSourceID int                                  `json:"data_source_id"`
	Runtime      float64                              `json:"runtime"`
	RetrievedAt  time.Time                            `json:"retrieved_at"`
}

type GetQueryResultsOutputQueryResultData struct {
	Columns []GetQueryResultsOutputQueryResultDataColumn `json:"columns"`
	Rows    []map[string]any                             `json:"rows"`
}

type GetQueryResultsOutputQueryResultDataColumn struct {
	Name         string `json:"name"`
	FriendlyName string `json:"friendly_name"`
	Type         string `json:"type"`
}

func (client *Client) ListQueries(ctx context.Context, input *ListQueriesInput) (*QueryPage, error) {
	res, close, err := client.Get(ctx, "api/queries", input)
	defer close()

	if err != nil {
		return nil, err
	}

	page := &QueryPage{}

	if err := util.UnmarshalBody(res, &page); err != nil {
		return nil, err
	}

	return page, nil
}

func (client *Client) GetQuery(ctx context.Context, id int) (*Query, error) {
	res, close, err := client.Get(ctx, fmt.Sprintf("api/queries/%d", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	query := &Query{}

	if err := util.UnmarshalBody(res, &query); err != nil {
		return nil, err
	}

	return query, nil
}

func (client *Client) CreateFavoriteQuery(ctx context.Context, id int) error {
	_, close, err := client.Post(ctx, fmt.Sprintf("api/queries/%d/favorite", id), nil)
	defer close()

	if err != nil {
		return err
	}

	return nil
}

type CreateQueryInput struct {
	DataSourceID int                       `json:"data_source_id"`
	Description  string                    `json:"description,omitempty"`
	Name         string                    `json:"name"`
	Options      *CreateQueryInputOptions  `json:"options,omitempty"`
	Query        string                    `json:"query"`
	Schedule     *CreateQueryInputSchedule `json:"schedule,omitempty"`
	Tags         []string                  `json:"tags,omitempty"`
}

type CreateQueryInputOptions struct {
	Parameters []QueryOptionsParameter `json:"parameters"`
}

type CreateQueryInputSchedule struct {
	Interval  int     `json:"interval"`
	Time      *string `json:"time"`
	Until     *string `json:"until"`
	DayOfWeek *string `json:"day_of_week"`
}

func (client *Client) CreateQuery(ctx context.Context, input *CreateQueryInput) (*Query, error) {
	res, close, err := client.Post(ctx, "api/queries", input)
	defer close()

	if err != nil {
		return nil, err
	}

	query := &Query{}

	if err := util.UnmarshalBody(res, &query); err != nil {
		return nil, err
	}

	return query, nil
}

func (client *Client) ForkQuery(ctx context.Context, id int) (*Query, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/queries/%d/fork", id), nil)
	defer close()

	if err != nil {
		return nil, err
	}

	query := &Query{}

	if err := util.UnmarshalBody(res, &query); err != nil {
		return nil, err
	}

	return query, nil
}

type UpdateQueryInput struct {
	DataSourceID int                       `json:"data_source_id,omitempty"`
	Description  string                    `json:"description,omitempty"`
	Name         string                    `json:"name,omitempty"`
	Options      *UpdateQueryInputOptions  `json:"options,omitempty"`
	Query        string                    `json:"query,omitempty"`
	Schedule     *UpdateQueryInputSchedule `json:"schedule,omitempty"`
	Tags         *[]string                 `json:"tags,omitempty"`
}

type UpdateQueryInputOptions struct {
	Parameters []QueryOptionsParameter `json:"parameters"`
}

type UpdateQueryInputSchedule struct {
	Interval  int     `json:"interval"`
	Time      *string `json:"time"`
	Until     *string `json:"until"`
	DayOfWeek *string `json:"day_of_week"`
}

func (client *Client) UpdateQuery(ctx context.Context, id int, input *UpdateQueryInput) (*Query, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/queries/%d", id), input)
	defer close()

	if err != nil {
		return nil, err
	}

	query := &Query{}

	if err := util.UnmarshalBody(res, &query); err != nil {
		return nil, err
	}

	return query, nil
}

type PublishQueryInput struct {
	IsDraft bool `json:"is_draft"`
}

func (client *Client) PublishQuery(ctx context.Context, id int) error {
	input := &PublishQueryInput{
		IsDraft: false,
	}

	_, close, err := client.Post(ctx, fmt.Sprintf("api/queries/%d", id), input)
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) UnpublishQuery(ctx context.Context, id int) error {
	input := &PublishQueryInput{
		IsDraft: true,
	}

	_, close, err := client.Post(ctx, fmt.Sprintf("api/queries/%d", id), input)
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) ArchiveQuery(ctx context.Context, id int) error {
	_, close, err := client.Delete(ctx, fmt.Sprintf("api/queries/%d", id))
	defer close()

	if err != nil {
		return err
	}

	return nil
}

func (client *Client) GetQueryResultsJSON(ctx context.Context, id int, out io.Writer) error {
	return client.GetQueryResults(ctx, id, "json", out)
}

func JsonToGetQueryResultsOutput(bs []byte) (*GetQueryResultsOutput, error) {
	var out *GetQueryResultsOutput
	err := json.Unmarshal(bs, &out)

	if err != nil {
		return nil, err
	}

	return out, err
}

func (client *Client) GetQueryResultsStruct(ctx context.Context, id int) (*GetQueryResultsOutput, error) {
	var buf bytes.Buffer
	err := client.GetQueryResultsJSON(ctx, id, &buf)

	if err != nil {
		return nil, err
	}

	out, err := JsonToGetQueryResultsOutput(buf.Bytes())

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (client *Client) GetQueryResultsCSV(ctx context.Context, id int, out io.Writer) error {
	return client.GetQueryResults(ctx, id, "csv", out)
}

func (client *Client) GetQueryResults(ctx context.Context, id int, ext string, out io.Writer) error {
	if out == nil {
		return errors.New("out(io.Writer) is nil")
	}

	res, close, err := client.Get(ctx, fmt.Sprintf("api/queries/%d/results.%s", id, ext), nil)
	defer close()

	if err != nil {
		return err
	}

	_, err = io.Copy(out, res.Body)

	return err
}

func (client *Client) GetQueryResultByID(ctx context.Context, queryResultId int, ext string, out *bytes.Buffer) error {
	if out == nil {
		return errors.New("out(io.Writer) is nil")
	}

	path := fmt.Sprintf("api/query_results/%d", queryResultId)

	if ext != "" {
		path += "." + ext
	}

	res, close, err := client.Get(ctx, path, nil)
	defer close()

	if err != nil {
		return err
	}

	_, err = io.Copy(out, res.Body)

	return err
}

type ExecQueryJSONInput struct {
	Parameters            map[string]any `json:"parameters,omitempty"`
	ApplyAutoLimit        bool           `json:"apply_auto_limit,omitempty"`
	MaxAge                int            `json:"max_age,omitempty"`
	WithoutOmittingMaxAge bool           `json:"-"`
}

type execQueryJSONInputWithMaxAge struct {
	Parameters     map[string]any `json:"parameters,omitempty"`
	ApplyAutoLimit bool           `json:"apply_auto_limit,omitempty"`
	MaxAge         int            `json:"max_age"`
}

func (client *Client) ExecQueryJSON(ctx context.Context, id int, input *ExecQueryJSONInput, out io.Writer) (*JobResponse, error) {
	if out == nil {
		out = io.Discard
	}

	var body any = input

	if input != nil && input.WithoutOmittingMaxAge {
		body = &execQueryJSONInputWithMaxAge{
			Parameters:     input.Parameters,
			ApplyAutoLimit: input.ApplyAutoLimit,
			MaxAge:         input.MaxAge,
		}
	}

	res, close, err := client.Post(ctx, fmt.Sprintf("api/queries/%d/results", id), body)
	defer close()

	if err != nil {
		return nil, err
	}

	magic := []byte(`{"job":`)
	head := make([]byte, len(magic))
	_, err = io.ReadFull(res.Body, head)

	if err != nil {
		return nil, err
	}

	buf := io.MultiReader(bytes.NewReader(head), res.Body)

	if bytes.Equal(head, magic) {
		job := &JobResponse{}
		err := json.NewDecoder(buf).Decode(&job)
		return job, err
	}

	_, err = io.Copy(out, buf)

	return nil, err
}

var (
	defaultWaitQueryJSONOptionWaitStatuses = []int{
		JobStatusPending,
		JobStatusStarted,
	}
)

const (
	defaultWaitQueryJSONOptionInterval = 1 * time.Second
)

type WaitQueryJSONOption struct {
	WaitStatuses []int
	Interval     time.Duration
}

func (client *Client) WaitQueryJSON(ctx context.Context, queryId int, job *JobResponse, option *WaitQueryJSONOption, out io.Writer) error {
	if job == nil || job.Job.ID == "" {
		return nil
	}

	waitStatus := defaultWaitQueryJSONOptionWaitStatuses
	interval := defaultWaitQueryJSONOptionInterval

	if option != nil {
		if len(option.WaitStatuses) > 0 {
			waitStatus = option.WaitStatuses
		}

		if option.Interval > 0 {
			interval = option.Interval
		}
	}

	for {
		job, err := client.GetJob(ctx, job.Job.ID)

		if err != nil {
			return err
		}

		if !slices.Contains(waitStatus, job.Job.Status) {
			err := client.GetQueryResultsJSON(ctx, queryId, out)

			if err != nil {
				return err
			}

			break
		}

		time.Sleep(interval)
	}

	return nil
}

func (client *Client) WaitQueryStruct(ctx context.Context, queryId int, job *JobResponse, option *WaitQueryJSONOption, buf *bytes.Buffer) (*GetQueryResultsOutput, error) {
	err := client.WaitQueryJSON(ctx, queryId, job, option, buf)

	if err != nil {
		return nil, err
	}

	out, err := JsonToGetQueryResultsOutput(buf.Bytes())

	if err != nil {
		return nil, err
	}

	return out, nil
}

type QueryTags struct {
	Tags []QueryTagsTag `json:"tags"`
}

type QueryTagsTag struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

func (client *Client) GetQueryTags(ctx context.Context) (*QueryTags, error) {
	res, close, err := client.Get(ctx, "api/queries/tags", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	tags := &QueryTags{}

	if err := util.UnmarshalBody(res, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}

type RefreshQueryInput struct {
	ApplyAutoLimit bool `json:"apply_auto_limit"`
}

func (client *Client) RefreshQuery(ctx context.Context, id int, input *RefreshQueryInput) (*JobResponse, error) {
	res, close, err := client.Post(ctx, fmt.Sprintf("api/queries/%d/refresh", id), input)
	defer close()

	if err != nil {
		return nil, err
	}

	job := &JobResponse{}

	if err := util.UnmarshalBody(res, &job); err != nil {
		return nil, err
	}

	return job, nil
}

type SearchQueriesInput struct {
	Q string `url:"q"`
}

func (client *Client) SearchQueries(ctx context.Context, input *SearchQueriesInput) (*QueryPage, error) {
	res, close, err := client.Get(ctx, "api/queries/search", input)
	defer close()

	if err != nil {
		return nil, err
	}

	page := &QueryPage{}

	if err := util.UnmarshalBody(res, &page); err != nil {
		return nil, err
	}

	return page, nil
}

type ListMyQueriesInput struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Q        string `url:"q,omitempty"`
}

func (client *Client) ListMyQueries(ctx context.Context, input *ListMyQueriesInput) (*QueryPage, error) {
	res, close, err := client.Get(ctx, "api/queries/my", input)
	defer close()

	if err != nil {
		return nil, err
	}

	page := &QueryPage{}

	if err := util.UnmarshalBody(res, &page); err != nil {
		return nil, err
	}

	return page, nil
}

type ListFavoriteQueriesInput struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Q        string `url:"q,omitempty"`
}

func (client *Client) ListFavoriteQueries(ctx context.Context, input *ListFavoriteQueriesInput) (*QueryPage, error) {
	res, close, err := client.Get(ctx, "api/queries/favorites", input)
	defer close()

	if err != nil {
		return nil, err
	}

	page := &QueryPage{}

	if err := util.UnmarshalBody(res, &page); err != nil {
		return nil, err
	}

	return page, nil
}

type FormatQueryOutput struct {
	Query string `json:"query"`
}

func (client *Client) FormatQuery(ctx context.Context, query string) (*FormatQueryOutput, error) {
	res, close, err := client.Post(ctx, "api/queries/format", map[string]string{"query": query})
	defer close()

	if err != nil {
		return nil, err
	}

	output := &FormatQueryOutput{}

	if err := util.UnmarshalBody(res, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func (client *Client) ListRecentQueries(ctx context.Context) ([]Query, error) {
	res, close, err := client.Get(ctx, "api/queries/recent", nil)
	defer close()

	if err != nil {
		return nil, err
	}

	queries := []Query{}

	if err := util.UnmarshalBody(res, &queries); err != nil {
		return nil, err
	}

	return queries, nil
}
