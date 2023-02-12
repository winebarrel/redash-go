# redash-go

[![test](https://github.com/winebarrel/redash-go/actions/workflows/test.yml/badge.svg)](https://github.com/winebarrel/redash-go/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/winebarrel/redash-go.svg)](https://pkg.go.dev/github.com/winebarrel/redash-go)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/winebarrel/redash-go)](https://github.com/winebarrel/redash-go/tags)

Redash API client in Go.

## Usage

```go
package main

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/winebarrel/redash-go"
)

func main() {
	client, err := redash.NewClient("https://redash.example.com", "<secret>")

	if err != nil {
		panic(err)
	}

	//client.Debug = true

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
	job, err := client.ExecQueryJSON(ctx, query.ID, &buf)

	if err != nil {
		panic(err)
	}

	if job != nil {
		for {
			job, err := client.GetJob(ctx, job.Job.ID)

			if err != nil {
				panic(err)
			}

			if job.Job.Status >= 3 {
				buf = bytes.Buffer{}
				err := client.GetQueryResultsJSON(ctx, query.ID, &buf)

				if err != nil {
					panic(err)
				}

				break
			}

			time.Sleep(1 * time.Second)
		}
	}

	fmt.Println(buf.String())
}
```

## Tests

```sh
make test
```

### Acceptance Tests

```sh
docker compose up -d
make redash-setup
make testacc
```

**NOTE:** local Redash URL: http://localhost:5001
