# redash-go

[![test](https://github.com/winebarrel/redash-go/actions/workflows/test.yml/badge.svg)](https://github.com/winebarrel/redash-go/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/winebarrel/redash-go.svg)](https://pkg.go.dev/github.com/winebarrel/redash-go)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/winebarrel/redash-go)

Redash API client in Go.

## Usage

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/winebarrel/redash-go"
)

func main() {
	client, _ := redash.NewClient("https://redash.example.com", "<secret>")
	//client.Debug = true

	ctx := context.Background()

	query, _ := client.CreateQuery(ctx, &redash.CreateQueryInput{
		DataSourceID: 1,
		Name:         "my-query1",
		Query:        "select 1",
	})

	rs, jobId, _ := client.ExecQueryJSON(ctx, query.ID)

	if jobId != "" {
		for {
			job, _ := client.GetJob(ctx, jobId)

			if job.Job.Status >= 3 {
				rs, _ = client.GetQueryResultsJSON(ctx, query.ID)
				break
			}

			time.Sleep(1 * time.Second)
		}
	}

	fmt.Println(string(rs))
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
