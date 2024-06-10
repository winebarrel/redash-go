# redash-go

[![CI](https://github.com/winebarrel/redash-go/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/redash-go/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/winebarrel/redash-go/v2.svg)](https://pkg.go.dev/github.com/winebarrel/redash-go/v2)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/winebarrel/redash-go)](https://github.com/winebarrel/redash-go/tags)
[![Go Report Card](https://goreportcard.com/badge/github.com/winebarrel/redash-go/v2)](https://goreportcard.com/report/github.com/winebarrel/redash-go/v2)
[![codecov](https://codecov.io/gh/winebarrel/redash-go/graph/badge.svg?token=9E21C7D54I)](https://codecov.io/gh/winebarrel/redash-go)

## Overview

Redash API client for Go that supports almost all APIs.

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

	// The API prefers to return a cached result.
	// If a cached result is not available then a new execution job begins and the job object is returned.
	// see https://redash.io/help/user-guide/integrations-and-api/api#Queries
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
```

### `max_age=0`

```go
input := &redash.ExecQueryJSONInput{
  WithoutOmittingMaxAge: true,
}

job, err := client.ExecQueryJSON(ctx, query.ID, input, nil)

if err != nil {
  panic(err)
}

err = client.WaitQueryJSON(ctx, query.ID, job, nil, &buf)

if err != nil {
  panic(err)
}

fmt.Println(buf.String())
```

### Set debug mode

```go
client := redash.MustNewClient("https://redash.example.com", "<secret>")
client.SetDebug(true)
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
docker compose up -d
make redash-setup
make redash-upgrade-db
make testacc
```

**NOTE:**
* local Redash URL: http://localhost:5001
* email: `admin@example.com`
* password: `password`

## Related Links

* https://redash.io/help/user-guide/integrations-and-api/api
