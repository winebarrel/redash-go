package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/winebarrel/redash-go/v2"
)

const (
	testRedashEndpoint = "http://localhost:5001"
	testRedashAPIKey   = "6nh64ZsT66WeVJvNZ6WB5D2JKZULeC2VBdSD68wt"
)

func main() {
	client := redash.MustNewClient(testRedashEndpoint, testRedashAPIKey)
	// client.SetDebug(true)
	ctx := context.Background()

	ds, err := client.CreateDataSource(ctx, &redash.CreateDataSourceInput{
		Name: "postgres",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
	})

	if err != nil {
		panic(err)
	}

	query, err := client.CreateQuery(ctx, &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "my-query1",
		Query:        "select 1",
	})

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	job, err := client.ExecQueryJSON(ctx, query.ID, nil, &buf)

	if err != nil {
		panic(err)
	}

	err = client.WaitQueryJSON(ctx, query.ID, job, nil, &buf)

	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
