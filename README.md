# redash-go

[![test](https://github.com/winebarrel/redash-go/actions/workflows/test.yml/badge.svg)](https://github.com/winebarrel/redash-go/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/winebarrel/redash-go/v2.svg)](https://pkg.go.dev/github.com/winebarrel/redash-go/v2)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/winebarrel/redash-go)](https://github.com/winebarrel/redash-go/tags)
[![Go Report Card](https://goreportcard.com/badge/github.com/winebarrel/redash-go/v2)](https://goreportcard.com/report/github.com/winebarrel/redash-go/v2)

## Overview

Redash API client for Go.
Supports almost all Redash APIs.

## Usage

```go
package main

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/winebarrel/redash-go/v2"
)

func main() {
	client := redash.MustNewClient("https://redash.example.com", "<secret>")
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

			if job.Job.Status != redash.JobStatusPending && job.Job.Status != redash.JobStatusStarted {
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

### Set debug mode

```go
client := redash.MustNewClient("https://redash.example.com", "<secret>")
client.Debug = true
client.GetStatus(context.Background())
```

```
% go run example.go
---request begin---
GET /status.json HTTP/1.1
Host: redash.example.com
Authorization: Key <secret>
Content-Type: application/json
User-Agent: redash-go


---request end---
---response begin---
HTTP/1.1 200 OK
...

{"dashboards_count": 0, "database_metrics": {"metrics": [ ...
```

### With custom HTTP client

```go
hc := &http.Client{
  Timeout: 3 * time.Second,
}
client := redash.MustNewClientWithHTTPClient("https://redash.example.com", "<secret>", hc)
client.GetStatus(context.Background())
```

### Without context.Context

```go
client0 := redash.MustNewClient("https://redash.example.com", "<secret>")
client := client0.WithoutContext()
client.GetStatus()
```

### NewClient with error

```go
client, err := redash.NewClient("https://redash.example.com", "<secret>")
```

## Tests

```sh
make test
```

### Acceptance Tests

```sh
docker compose up --build -d
make redash-setup
make redash-upgrade-db
make testacc
```

**NOTE:**
* local Redash URL: http://localhost:5001
* email: `admin@example.com`
* password: `password`
