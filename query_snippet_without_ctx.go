// Code generated from query_snippet.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import "context"

func (client *ClientWithoutContext) ListQuerySnippets() ([]QuerySnippet, error) {
	return client.withCtx.ListQuerySnippets(context.Background())
}

func (client *ClientWithoutContext) GetQuerySnippet(id int) (*QuerySnippet, error) {
	return client.withCtx.GetQuerySnippet(context.Background(), id)
}

func (client *ClientWithoutContext) CreateQuerySnippet(input *CreateQuerySnippetInput) (*QuerySnippet, error) {
	return client.withCtx.CreateQuerySnippet(context.Background(), input)
}

func (client *ClientWithoutContext) UpdateQuerySnippet(id int, input *UpdateQuerySnippetInput) (*QuerySnippet, error) {
	return client.withCtx.UpdateQuerySnippet(context.Background(), id, input)
}

func (client *ClientWithoutContext) DeleteQuerySnippet(id int) error {
	return client.withCtx.DeleteQuerySnippet(context.Background(), id)
}
