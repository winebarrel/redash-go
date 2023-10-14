package redash_test

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/redash-go"
)

func Test_ListUsers_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/users", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal("page=1&page_size=25", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"count": 1,
				"page": 1,
				"page_size": 25,
				"results": [
					{
						"active_at": "2023-02-10T01:23:45.000Z",
						"auth_type": "password",
						"created_at": "2023-02-10T01:23:45.000Z",
						"disabled_at": null,
						"email": "admin@example.com",
						"groups": [
							{
								"id": 1,
								"name": "admin"
							},
							{
								"id": 2,
								"name": "default"
							}
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
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.ListUsers(context.Background(), &redash.ListUsersInput{
		Page:     1,
		PageSize: 25,
	})
	assert.NoError(err)
	assert.Equal(&redash.UserPage{
		Count:    1,
		Page:     1,
		PageSize: 25,
		Results: []redash.User{
			{
				ActiveAt:   dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				APIKey:     "",
				AuthType:   "password",
				CreatedAt:  dateparse.MustParse("2023-02-10T01:23:45.000Z"),
				DisabledAt: time.Time{},
				Email:      "admin@example.com",
				Groups: []any{
					map[string]any{"id": float64(1), "name": "admin"},
					map[string]any{"id": float64(2), "name": "default"},
				},
				ID:                  1,
				IsDisabled:          false,
				IsEmailVerified:     true,
				IsInvitationPending: false,
				Name:                "admin",
				ProfileImageURL:     "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40&d=identicon",
				UpdatedAt:           dateparse.MustParse("2023-02-10T01:23:45.000Z"),
			},
		},
	}, res)
}

func Test_GetUser_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/users/1", func(req *http.Request) (*http.Response, error) {
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
				"active_at": "2023-02-10T01:23:45.000Z",
				"api_key": "api_key",
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
	res, err := client.GetUser(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.User{
		ActiveAt:            dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		APIKey:              "api_key",
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

func Test_CreateUser_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/users", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"auth_type":"password","email":"admin@example.com","name":"admin"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"active_at": "2023-02-10T01:23:45.000Z",
				"api_key": "api_key",
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
	res, err := client.CreateUser(context.Background(), &redash.CreateUsersInput{
		AuthType: "password",
		Email:    "admin@example.com",
		Name:     "admin",
	})
	assert.NoError(err)
	assert.Equal(&redash.User{
		ActiveAt:            dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		APIKey:              "api_key",
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

func Test_UpdateUser_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/users/1", func(req *http.Request) (*http.Response, error) {
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
		assert.Equal(`{"email":"admin2@example.com","name":"admin2"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"active_at": "2023-02-10T01:23:45.000Z",
				"api_key": "api_key",
				"auth_type": "password",
				"created_at": "2023-02-10T01:23:45.000Z",
				"disabled_at": null,
				"email": "admin2@example.com",
				"groups": [
					1,
					2
				],
				"id": 1,
				"is_disabled": false,
				"is_email_verified": true,
				"is_invitation_pending": false,
				"name": "admin2",
				"profile_image_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40&d=identicon",
				"updated_at": "2023-02-10T01:23:45.000Z"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.UpdateUser(context.Background(), 1, &redash.UpdateUserInput{
		Email: "admin2@example.com",
		Name:  "admin2",
	})
	assert.NoError(err)
	assert.Equal(&redash.User{
		ActiveAt:            dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		APIKey:              "api_key",
		AuthType:            "password",
		CreatedAt:           dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DisabledAt:          time.Time{},
		Email:               "admin2@example.com",
		Groups:              []any{float64(1), float64(2)},
		ID:                  1,
		IsDisabled:          false,
		IsEmailVerified:     true,
		IsInvitationPending: false,
		Name:                "admin2",
		ProfileImageURL:     "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40&d=identicon",
		UpdatedAt:           dateparse.MustParse("2023-02-10T01:23:45.000Z"),
	}, res)
}

func Test_DeleteUser_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/users/1", func(req *http.Request) (*http.Response, error) {
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
	err := client.DeleteUser(context.Background(), 1)
	assert.NoError(err)
}

func Test_DisableUser_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/users/1/disable", func(req *http.Request) (*http.Response, error) {
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
				"active_at": "2023-02-10T01:23:45.000Z",
				"api_key": "api_key",
				"auth_type": "password",
				"created_at": "2023-02-10T01:23:45.000Z",
				"disabled_at": "2023-02-10T01:23:45.000Z",
				"email": "admin@example.com",
				"groups": [
					1,
					2
				],
				"id": 1,
				"is_disabled": true,
				"is_email_verified": true,
				"is_invitation_pending": false,
				"name": "admin",
				"profile_image_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40&d=identicon",
				"updated_at": "2023-02-10T01:23:45.000Z"
			}
		`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, err := client.DisableUser(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.User{
		ActiveAt:            dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		APIKey:              "api_key",
		AuthType:            "password",
		CreatedAt:           dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		DisabledAt:          dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		Email:               "admin@example.com",
		Groups:              []any{float64(1), float64(2)},
		ID:                  1,
		IsDisabled:          true,
		IsEmailVerified:     true,
		IsInvitationPending: false,
		Name:                "admin",
		ProfileImageURL:     "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40&d=identicon",
		UpdatedAt:           dateparse.MustParse("2023-02-10T01:23:45.000Z"),
	}, res)
}

func Test_EnableUser_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/users/1/disable", func(req *http.Request) (*http.Response, error) {
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
				"active_at": "2023-02-10T01:23:45.000Z",
				"api_key": "api_key",
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
	res, err := client.EnableUser(context.Background(), 1)
	assert.NoError(err)
	assert.Equal(&redash.User{
		ActiveAt:            dateparse.MustParse("2023-02-10T01:23:45.000Z"),
		APIKey:              "api_key",
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

func Test_User_Acc(t *testing.T) {
	if !testAcc {
		t.Skip()
	}

	assert := assert.New(t)
	require := require.New(t)
	client, _ := redash.NewClient(testRedashEndpoint, testRedashAPIKey)

	uuidObj, _ := uuid.NewUUID()
	name := uuidObj.String()
	email := name + "@example.com"

	_, err := client.ListUsers(context.Background(), nil)
	require.NoError(err)

	user, err := client.CreateUser(context.Background(), &redash.CreateUsersInput{
		AuthType: "password",
		Email:    email,
		Name:     name,
	})
	require.NoError(err)
	assert.Equal(email, user.Email)

	user, err = client.GetUser(context.Background(), user.ID)
	require.NoError(err)
	assert.Equal(name, user.Name)
	assert.Equal(email, user.Email)

	newName := name + "2"
	newEmail := email + ".uk"

	user, err = client.UpdateUser(context.Background(), user.ID, &redash.UpdateUserInput{
		Email: newEmail,
		Name:  newName,
	})
	require.NoError(err)
	assert.Equal(newName, user.Name)
	assert.Equal(newEmail, user.Email)

	user, err = client.DisableUser(context.Background(), user.ID)
	require.NoError(err)
	assert.Equal(newName, user.Name)
	assert.Equal(newEmail, user.Email)
	assert.True(user.IsDisabled)

	user, err = client.EnableUser(context.Background(), user.ID)
	require.NoError(err)
	assert.Equal(newName, user.Name)
	assert.Equal(newEmail, user.Email)
	assert.False(user.IsDisabled)

	err = client.DeleteUser(context.Background(), user.ID)
	require.NoError(err)

	_, err = client.GetUser(context.Background(), user.ID)
	assert.Error(err)
}
