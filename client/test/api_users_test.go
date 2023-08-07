/*
Unleash API

Testing UsersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	openapiclient "github.com/Unleash/unleash-server-api-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_client_UsersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersApiService ChangeMyPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UsersApi.ChangeMyPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService ChangeUserPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.UsersApi.ChangeUserPassword(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService CreateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersApi.CreateUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService DeleteUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.UsersApi.DeleteUser(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetAdminCount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersApi.GetAdminCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetBaseUsersAndGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersApi.GetBaseUsersAndGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetMe", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersApi.GetMe(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersApi.GetProfile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.UsersApi.GetUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersApi.GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService ResetUserPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersApi.ResetUserPassword(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService SearchUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersApi.SearchUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService UpdateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.UsersApi.UpdateUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService ValidateUserPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UsersApi.ValidateUserPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
