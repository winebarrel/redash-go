package redash_test

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go/v2"
)

func Test_ListGroups_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/groups", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, `
			[
				{
					"created_at": "2023-02-10T01:23:45.000Z",
					"id": 1,
					"name": "admin",
					"permissions": [
						"admin",
						"super_admin"
					],
					"type": "builtin"
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListGroups(context.Background())
	assert.NoError(err)
	assert.Equal([]redash.Group{
		{
			CreatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			ID:        1,
			Name:      "admin",
			Permissions: []string{
				"admin",
				"super_admin",
			},
			Type: "builtin",
		},
	}, res)
}

func Test_ListGroups_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/groups", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListGroups(context.Background())
	assert.ErrorContains(err, "GET api/groups failed: HTTP status code not OK: 503\nerror")
}

func Test_ListGroups_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/groups", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.ListGroups(context.Background())
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_GetGroup_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/groups/1", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"created_at": "2023-02-10T01:23:45.000Z",
				"id": 1,
				"name": "admin",
				"permissions": [
					"admin",
					"super_admin"
				],
				"type": "builtin"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.GetGroup(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.Group{
		CreatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		ID:        1,
		Name:      "admin",
		Permissions: []string{
			"admin",
			"super_admin",
		},
		Type: "builtin",
	}, res)
}

func Test_GetGroup_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/groups/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetGroup(context.Background(), 1)
	assert.ErrorContains(err, "GET api/groups/1 failed: HTTP status code not OK: 503\nerror")
}

func Test_GetGroup_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/groups/1", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.GetGroup(context.Background(), 1)
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_CreateGroup_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/groups", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		if req.Body == nil {
			assert.FailNow("req.Body is nil")
		}
		body, _ := io.ReadAll(req.Body)
		assert.Equal(`{"name":"my-group"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"created_at": "2023-02-10T01:23:45.000Z",
				"id": 2,
				"name": "my-group",
				"permissions": [
					"create_dashboard",
					"create_query",
					"edit_dashboard",
					"edit_query",
					"view_query",
					"view_source",
					"execute_query",
					"list_users",
					"schedule_query",
					"list_dashboards",
					"list_alerts",
					"list_data_sources"
				],
				"type": "regular"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.CreateGroup(context.Background(), &redash.CreateGroupInput{
		Name: "my-group",
	})
	assert.NoError(err)
	assert.Equal(&redash.Group{
		CreatedAt: dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		ID:        2,
		Name:      "my-group",
		Permissions: []string{
			"create_dashboard",
			"create_query",
			"edit_dashboard",
			"edit_query",
			"view_query",
			"view_source",
			"execute_query",
			"list_users",
			"schedule_query",
			"list_dashboards",
			"list_alerts",
			"list_data_sources",
		},
		Type: "regular",
	}, res)
}

func Test_CreateGroup_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/groups", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.CreateGroup(context.Background(), &redash.CreateGroupInput{
		Name: "my-group",
	})
	assert.ErrorContains(err, "POST api/groups failed: HTTP status code not OK: 503\nerror")
}

func Test_CreateGroup_IOErr(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/groups", func(req *http.Request) (*http.Response, error) {
		return testIOErrResp, nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, err := client.CreateGroup(context.Background(), &redash.CreateGroupInput{
		Name: "my-group",
	})
	assert.ErrorContains(err, "Read response body failed: IO error")
}

func Test_DeleteGroup_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/groups/2", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.DeleteGroup(context.Background(), 2)
	assert.NoError(err)
}

func Test_DeleteGroup_Err_5xx(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/groups/2", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusServiceUnavailable, "error"), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.DeleteGroup(context.Background(), 2)
	assert.ErrorContains(err, "DELETE api/groups/2 failed: HTTP status code not OK: 503\nerror")
}

func Test_ListGroupMembers_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/groups/1/members", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, `
			[
				{
					"active_at": "2023-02-10T01:23:45.000Z",
					"auth_type": "password",
					"created_at": "2023-02-10T01:23:45.000Z",
					"disabled_at": null,
					"email": "admin@example.com",
					"groups": [
						1,
						2
					],
					"id": 1,
					"is_disabled": false,
					"is_email_verified": true,
					"is_invitation_pending": false,
					"name": "admin",
					"profile_image_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40&d=identicon",
					"updated_at": "2023-02-10T01:23:45.000Z"
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListGroupMembers(context.Background(), 1)
	assert.NoError(err)
	assert.Equal([]redash.User{
		{
			ActiveAt:            dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			AuthType:            "password",
			CreatedAt:           dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			DisabledAt:          time.Time{},
			Email:               "admin@example.com",
			Groups:              []any{float64(1), float64(2)},
			ID:                  1,
			IsDisabled:          false,
			IsEmailVerified:     true,
			IsInvitationPending: false,
			Name:                "admin",
			ProfileImageURL:     "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40&d=identicon",
			UpdatedAt:           dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		},
	}, res)
}

func Test_AddGroupMember_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/groups/1/members", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		if req.Body == nil {
			assert.FailNow("req.Body is nil")
		}
		body, _ := io.ReadAll(req.Body)
		assert.Equal(`{"user_id":1}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"active_at": "2023-02-10T01:23:45.000Z",
				"auth_type": "password",
				"created_at": "2023-02-10T01:23:45.000Z",
				"disabled_at": null,
				"email": "admin@example.com",
				"groups": [
					1,
					2
				],
				"id": 1,
				"is_disabled": false,
				"is_email_verified": true,
				"is_invitation_pending": false,
				"name": "admin",
				"profile_image_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40&d=identicon",
				"updated_at": "2023-02-10T01:23:45.000Z"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.AddGroupMember(context.Background(), 1, 1)
	assert.NoError(err)
	assert.Equal(&redash.User{
		ActiveAt:            dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		AuthType:            "password",
		CreatedAt:           dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DisabledAt:          time.Time{},
		Email:               "admin@example.com",
		Groups:              []any{float64(1), float64(2)},
		ID:                  1,
		IsDisabled:          false,
		IsEmailVerified:     true,
		IsInvitationPending: false,
		Name:                "admin",
		ProfileImageURL:     "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40&d=identicon",
		UpdatedAt:           dateparse.MustParse("2023-02-10T01:23:45.000Z"),
	}, res)
}

func Test_RemoveGroupMember_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/groups/1/members/2", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.RemoveGroupMember(context.Background(), 1, 2)
	assert.NoError(err)
}

func Test_ListGroupDataSources_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/groups/1/data_sources", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, `
			[
				{
					"id": 1,
					"name": "postgres",
					"pause_reason": null,
					"paused": 0,
					"syntax": "sql",
					"type": "pg",
					"view_only": false
				}
			]
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListGroupDataSources(context.Background(), 1)
	assert.NoError(err)
	assert.Equal([]redash.DataSource{
		{
			Groups:             nil,
			ID:                 1,
			Name:               "postgres",
			Options:            nil,
			Paused:             0,
			PauseReason:        "",
			QueueName:          "",
			ScheduledQueueName: "",
			Syntax:             "sql",
			Type:               "pg",
			ViewOnly:           false,
		},
	}, res)
}

func Test_AddGroupDataSource_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/groups/1/data_sources", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		if req.Body == nil {
			assert.FailNow("req.Body is nil")
		}
		body, _ := io.ReadAll(req.Body)
		assert.Equal(`{"data_source_id":1}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"id": 1,
				"name": "postgres",
				"pause_reason": null,
				"paused": 0,
				"syntax": "sql",
				"type": "pg",
				"view_only": false
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.AddGroupDataSource(context.Background(), 1, 1)
	assert.NoError(err)
	assert.Equal(&redash.DataSource{
		Groups:             nil,
		ID:                 1,
		Name:               "postgres",
		Options:            nil,
		Paused:             0,
		PauseReason:        "",
		QueueName:          "",
		ScheduledQueueName: "",
		Syntax:             "sql",
		Type:               "pg",
		ViewOnly:           false,
	}, res)
}

func Test_RemoveGroupDataSource_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/groups/1/data_sources/2", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	err := client.RemoveGroupDataSource(context.Background(), 1, 2)
	assert.NoError(err)
}

func Test_Group_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)

	_, err := client.ListGroups(context.Background())
	require.NoError(err)

	group, err := client.CreateGroup(context.Background(), &redash.CreateGroupInput{
		Name: "test-group-1",
	})
	require.NoError(err)
	assert.Equal("test-group-1", group.Name)

	group, err = client.GetGroup(context.Background(), group.ID)
	require.NoError(err)
	assert.Equal("test-group-1", group.Name)

	_, err = client.ListGroupMembers(context.Background(), group.ID)
	require.NoError(err)

	_, err = client.ListGroupDataSources(context.Background(), group.ID)
	require.NoError(err)

	err = client.DeleteGroup(context.Background(), group.ID)
	require.NoError(err)

	_, err = client.GetGroup(context.Background(), group.ID)
	assert.Error(err)
}
