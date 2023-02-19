// Code generated from query.go using internal/gen/withoutctx.go; DO NOT EDIT.

package redash

import (
	"context"
	"io"
)

// Auto-generated
func (client *ClientWithoutContext) ListQueries(input *ListQueriesInput) (*QueryPage, error) {
	return client.withCtx.ListQueries(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) GetQuery(id int) (*Query, error) {
	return client.withCtx.GetQuery(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) CreateFavoriteQuery(id int) error {
	return client.withCtx.CreateFavoriteQuery(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) CreateQuery(input *CreateQueryInput) (*Query, error) {
	return client.withCtx.CreateQuery(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) ForkQuery(id int) (*Query, error) {
	return client.withCtx.ForkQuery(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) UpdateQuery(id int, input *UpdateQueryInput) (*Query, error) {
	return client.withCtx.UpdateQuery(context.Background(), id, input)
}

// Auto-generated
func (client *ClientWithoutContext) ArchiveQuery(id int) error {
	return client.withCtx.ArchiveQuery(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) GetQueryResultsJSON(id int, out io.Writer) error {
	return client.withCtx.GetQueryResultsJSON(context.Background(), id, out)
}

// Auto-generated
func (client *ClientWithoutContext) GetQueryResultsCSV(id int, out io.Writer) error {
	return client.withCtx.GetQueryResultsCSV(context.Background(), id, out)
}

// Auto-generated
func (client *ClientWithoutContext) GetQueryResults(id int, ext string, out io.Writer) error {
	return client.withCtx.GetQueryResults(context.Background(), id, ext, out)
}

// Auto-generated
func (client *ClientWithoutContext) ExecQueryJSON(id int, out io.Writer) (*JobResponse, error) {
	return client.withCtx.ExecQueryJSON(context.Background(), id, out)
}

// Auto-generated
func (client *ClientWithoutContext) GetQueryTags() (*QueryTags, error) {
	return client.withCtx.GetQueryTags(context.Background())
}

// Auto-generated
func (client *ClientWithoutContext) RefreshQuery(id int) (*JobResponse, error) {
	return client.withCtx.RefreshQuery(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) SearchQueries(input *SearchQueriesInput) (*QueryPage, error) {
	return client.withCtx.SearchQueries(context.Background(), input)
}
