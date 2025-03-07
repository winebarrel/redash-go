package redash_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go/v2"
)

func Test_Query_WithParamsNum_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	ds, err := client.CreateDataSource(context.Background(), &redash.CreateDataSourceInput{
		Name: "test-postgres-1",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
	})
	require.NoError(err)

	defer func() {
		client.DeleteDataSource(context.Background(), ds.ID) //nolint:errcheck
	}()

	_, err = client.ListQueries(context.Background(), nil)
	require.NoError(err)

	query, err := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "test-query-1",
		Query:        "select {{ num }}",
		Options: &redash.CreateQueryInputOptions{
			Parameters: []redash.QueryOptionsParameter{
				{
					Global: false,
					Type:   "number",
					Name:   "num",
					Value:  123,
					Title:  "my-number",
				},
			},
		},
		Tags: []string{"my-tag-1"},
	})
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)
	assert.Equal([]string{"my-tag-1"}, query.Tags)

	query, err = client.GetQuery(context.Background(), query.ID)
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)
	assert.Equal([]string{"my-tag-1"}, query.Tags)
	assert.Equal("select {{ num }}", query.Query)
	assert.Equal(redash.QueryOptions{
		Parameters: []redash.QueryOptionsParameter{
			{
				Global: false,
				Type:   "number",
				Name:   "num",
				Value:  float64(123),
				Title:  "my-number",
			},
		},
	}, query.Options)

	var buf bytes.Buffer
	input := &redash.ExecQueryJSONInput{
		Parameters: map[string]any{
			"num": 999,
		},
		MaxAge: 1800,
	}
	job, err := client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)

	if job != nil && job.Job.ID != "" {
		for {
			job, err := client.GetJob(context.Background(), job.Job.ID)
			require.NoError(err)

			if job.Job.Status != redash.JobStatusPending && job.Job.Status != redash.JobStatusStarted {
				assert.Equal(redash.JobStatusSuccess, job.Job.Status)
				_, err := client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
				require.NoError(err)
				break
			}

			time.Sleep(1 * time.Second)
		}
	}

	assert.Contains(buf.String(), `"query": "select 999"`)

	buf = bytes.Buffer{}
	input = &redash.ExecQueryJSONInput{
		Parameters: map[string]any{
			"num": 999,
		},
		ApplyAutoLimit: true,
		MaxAge:         1800,
	}
	job, err = client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)

	if job != nil && job.Job.ID != "" {
		for {
			job, err := client.GetJob(context.Background(), job.Job.ID)
			require.NoError(err)

			if job.Job.Status != redash.JobStatusPending && job.Job.Status != redash.JobStatusStarted {
				assert.Equal(redash.JobStatusSuccess, job.Job.Status)
				_, err := client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
				require.NoError(err)
				break
			}

			time.Sleep(1 * time.Second)
		}
	}

	assert.Contains(buf.String(), `"query": "select 999 LIMIT 1000"`)
}

func Test_Query_WithParamsText_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	ds, err := client.CreateDataSource(context.Background(), &redash.CreateDataSourceInput{
		Name: "test-postgres-1",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
	})
	require.NoError(err)

	defer func() {
		client.DeleteDataSource(context.Background(), ds.ID) //nolint:errcheck
	}()

	_, err = client.ListQueries(context.Background(), nil)
	require.NoError(err)

	query, err := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "test-query-1",
		Query:        "select '{{ msg }}'",
		Options: &redash.CreateQueryInputOptions{
			Parameters: []redash.QueryOptionsParameter{
				{
					Global: false,
					Type:   "text",
					Name:   "msg",
					Value:  "hello",
					Title:  "my-text",
				},
			},
		},
	})
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)

	query, err = client.GetQuery(context.Background(), query.ID)
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)
	assert.Equal("select '{{ msg }}'", query.Query)
	assert.Equal(redash.QueryOptions{
		Parameters: []redash.QueryOptionsParameter{
			{
				Global: false,
				Type:   "text",
				Name:   "msg",
				Value:  "hello",
				Title:  "my-text",
			},
		},
	}, query.Options)

	var buf bytes.Buffer
	input := &redash.ExecQueryJSONInput{
		Parameters: map[string]any{
			"msg": "hellohello",
		},
		MaxAge: 1800,
	}
	job, err := client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	client.WaitQueryJSON(context.Background(), query.ID, job, nil, &buf) //nolint:errcheck
	_, err = client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	assert.Contains(buf.String(), `"query": "select 'hellohello'"`)
}

func Test_Query_WithParamsTextPattern_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	ds, err := client.CreateDataSource(context.Background(), &redash.CreateDataSourceInput{
		Name: "test-postgres-1",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
	})
	require.NoError(err)

	defer func() {
		client.DeleteDataSource(context.Background(), ds.ID) //nolint:errcheck
	}()

	_, err = client.ListQueries(context.Background(), nil)
	require.NoError(err)

	query, err := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "test-query-1",
		Query:        "select '{{ textp }}'",
		Options: &redash.CreateQueryInputOptions{
			Parameters: []redash.QueryOptionsParameter{
				{
					Global: false,
					Type:   "text-pattern",
					Name:   "textp",
					Title:  "my-textp",
					Regex:  "ab+c",
				},
			},
		},
	})
	require.NoError(err)

	query, err = client.GetQuery(context.Background(), query.ID)
	require.NoError(err)
	assert.Equal("select '{{ textp }}'", query.Query)
	assert.Equal(redash.QueryOptions{
		Parameters: []redash.QueryOptionsParameter{
			{
				Global: false,
				Type:   "text-pattern",
				Name:   "textp",
				Title:  "my-textp",
				Regex:  "ab+c",
			},
		},
	}, query.Options)

	var buf bytes.Buffer
	input := &redash.ExecQueryJSONInput{
		Parameters: map[string]any{
			"textp": "abbbc",
		},
		MaxAge: 1800,
	}
	job, err := client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	client.WaitQueryJSON(context.Background(), query.ID, job, nil, &buf) //nolint:errcheck
	_, err = client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	assert.Contains(buf.String(), `"query": "select 'abbbc'"`)
}

func Test_Query_WithParamsDropdownList_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	ds, err := client.CreateDataSource(context.Background(), &redash.CreateDataSourceInput{
		Name: "test-postgres-1",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
	})
	require.NoError(err)

	defer func() {
		client.DeleteDataSource(context.Background(), ds.ID) //nolint:errcheck
	}()

	_, err = client.ListQueries(context.Background(), nil)
	require.NoError(err)

	query, err := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "test-query-1",
		Query:        "select '{{ ddlist }}'",
		Options: &redash.CreateQueryInputOptions{
			Parameters: []redash.QueryOptionsParameter{
				{
					Global:      false,
					Type:        "enum",
					Name:        "ddlist",
					Title:       "my-ddlist",
					EnumOptions: "aaa\nbbb\nccc",
				},
			},
		},
	})
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)

	query, err = client.GetQuery(context.Background(), query.ID)
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)
	assert.Equal("select '{{ ddlist }}'", query.Query)
	assert.Equal(redash.QueryOptions{
		Parameters: []redash.QueryOptionsParameter{
			{
				Global:      false,
				Type:        "enum",
				Name:        "ddlist",
				Title:       "my-ddlist",
				EnumOptions: "aaa\nbbb\nccc",
			},
		},
	}, query.Options)

	var buf bytes.Buffer
	input := &redash.ExecQueryJSONInput{
		Parameters: map[string]any{
			"ddlist": "bbb",
		},
		MaxAge: 1800,
	}
	job, err := client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	client.WaitQueryJSON(context.Background(), query.ID, job, nil, &buf) //nolint:errcheck
	_, err = client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	assert.Contains(buf.String(), `"query": "select 'bbb'"`)
}

func Test_Query_WithParamsDropdownListMultiValues_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	ds, err := client.CreateDataSource(context.Background(), &redash.CreateDataSourceInput{
		Name: "test-postgres-1",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
	})
	require.NoError(err)

	defer func() {
		client.DeleteDataSource(context.Background(), ds.ID) //nolint:errcheck
	}()

	_, err = client.ListQueries(context.Background(), nil)
	require.NoError(err)

	query, err := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "test-query-1",
		Query:        "select '{{ ddlist }}'",
		Options: &redash.CreateQueryInputOptions{
			Parameters: []redash.QueryOptionsParameter{
				{
					Global:      false,
					Type:        "enum",
					Name:        "ddlist",
					Title:       "my-ddlist",
					EnumOptions: "aaa\nbbb\nccc",
					MultiValuesOptions: &redash.QueryOptionsParameterMultiValuesOptions{
						Prefix:    `"`,
						Suffix:    `"`,
						Separator: ",",
					},
				},
			},
		},
	})
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)

	query, err = client.GetQuery(context.Background(), query.ID)
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)
	assert.Equal("select '{{ ddlist }}'", query.Query)
	assert.Equal(redash.QueryOptions{
		Parameters: []redash.QueryOptionsParameter{
			{
				Global:      false,
				Type:        "enum",
				Name:        "ddlist",
				Title:       "my-ddlist",
				EnumOptions: "aaa\nbbb\nccc",
				MultiValuesOptions: &redash.QueryOptionsParameterMultiValuesOptions{
					Prefix:    `"`,
					Suffix:    `"`,
					Separator: ",",
				},
			},
		},
	}, query.Options)

	var buf bytes.Buffer
	input := &redash.ExecQueryJSONInput{
		Parameters: map[string]any{
			"ddlist": []string{"aaa", "bbb"},
		},
		MaxAge: 1800,
	}
	job, err := client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	client.WaitQueryJSON(context.Background(), query.ID, job, nil, &buf) //nolint:errcheck
	_, err = client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	assert.Contains(buf.String(), `"query": "select '\"aaa\",\"bbb\"'"`)
}

func Test_Query_WithParamsQueryBasedDropdownList_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)
	ds, err := client.CreateDataSource(context.Background(), &redash.CreateDataSourceInput{
		Name: "test-postgres-1",
		Type: "pg",
		Options: map[string]any{
			"dbname": "postgres",
			"host":   "postgres",
			"port":   5432,
			"user":   "postgres",
		},
	})
	require.NoError(err)

	defer func() {
		client.DeleteDataSource(context.Background(), ds.ID) //nolint:errcheck
	}()

	_, err = client.ListQueries(context.Background(), nil)
	require.NoError(err)

	dlQuery, err := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "test-dl-query-1",
		Query:        "select unnest(array[1,2,3])",
	})
	require.NoError(err)
	assert.Equal("test-dl-query-1", dlQuery.Name)

	var buf bytes.Buffer
	job, err := client.ExecQueryJSON(context.Background(), dlQuery.ID, &redash.ExecQueryJSONInput{}, &buf)
	require.NoError(err)
	err = client.WaitQueryJSON(context.Background(), dlQuery.ID, job, nil, &buf)
	require.NoError(err)
	require.Contains(buf.String(), `"rows": [{"unnest": 1}, {"unnest": 2}, {"unnest": 3}]}`)

	query, err := client.CreateQuery(context.Background(), &redash.CreateQueryInput{
		DataSourceID: ds.ID,
		Name:         "test-query-1",
		Query:        "select '{{ dbddlist }}'",
		Options: &redash.CreateQueryInputOptions{
			Parameters: []redash.QueryOptionsParameter{
				{
					Global:  false,
					Type:    "query",
					Name:    "dbddlist",
					Title:   "my-dbddlist",
					QueryID: dlQuery.ID,
				},
			},
		},
	})
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)

	query, err = client.GetQuery(context.Background(), query.ID)
	require.NoError(err)
	assert.Equal("test-query-1", query.Name)
	assert.Equal("select '{{ dbddlist }}'", query.Query)
	assert.Equal(redash.QueryOptions{
		Parameters: []redash.QueryOptionsParameter{
			{
				Global:  false,
				Type:    "query",
				Name:    "dbddlist",
				Title:   "my-dbddlist",
				QueryID: dlQuery.ID,
			},
		},
	}, query.Options)

	buf = bytes.Buffer{}
	input := &redash.ExecQueryJSONInput{
		Parameters: map[string]any{
			"dbddlist": "2",
		},
	}
	job, err = client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	client.WaitQueryJSON(context.Background(), query.ID, job, nil, &buf) //nolint:errcheck
	_, err = client.ExecQueryJSON(context.Background(), query.ID, input, &buf)
	require.NoError(err)
	assert.Contains(buf.String(), `"query": "select '2'"`)
}
