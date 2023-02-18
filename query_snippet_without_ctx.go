// Code generated from query_snippet.go using gen/withoutctx.go; DO NOT EDIT.

package redash

import "context"

// Auto-generated
func (client *ClientWithoutContext) ListQuerySnippets() ([]QuerySnippet, error) {
	return client.withCtx.ListQuerySnippets(context.Background())
}

// Auto-generated
func (client *ClientWithoutContext) GetQuerySnippet(id int) (*QuerySnippet, error) {
	return client.withCtx.GetQuerySnippet(context.Background(), id)
}

// Auto-generated
func (client *ClientWithoutContext) CreateQuerySnippet(input *CreateQuerySnippetInput) (*QuerySnippet, error) {
	return client.withCtx.CreateQuerySnippet(context.Background(), input)
}

// Auto-generated
func (client *ClientWithoutContext) UpdateQuerySnippet(id int, input *UpdateQuerySnippetInput) (*QuerySnippet, error) {
	return client.withCtx.UpdateQuerySnippet(context.Background(), id, input)
}

// Auto-generated
func (client *ClientWithoutContext) DeleteQuerySnippet(id int) error {
	return client.withCtx.DeleteQuerySnippet(context.Background(), id)
}
