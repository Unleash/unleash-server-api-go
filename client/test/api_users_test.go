/*
Unleash API

Testing UsersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	openapiclient "github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_client_UsersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersApiService ChangeMyPassword", func(t *testing.T) {

		httpRes, err := apiClient.UsersApi.ChangeMyPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService ChangeUserPassword", func(t *testing.T) {

		var id string

		httpRes, err := apiClient.UsersApi.ChangeUserPassword(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService CreateGroup", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.CreateGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService CreateRole", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.CreateRole(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService CreateUser", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.CreateUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService DeleteGroup", func(t *testing.T) {

		var groupId string

		httpRes, err := apiClient.UsersApi.DeleteGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService DeleteRole", func(t *testing.T) {

		var roleId string

		httpRes, err := apiClient.UsersApi.DeleteRole(context.Background(), roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService DeleteUser", func(t *testing.T) {

		var id string

		httpRes, err := apiClient.UsersApi.DeleteUser(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetAdminCount", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.GetAdminCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetBaseUsersAndGroups", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.GetBaseUsersAndGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetGroup", func(t *testing.T) {

		var groupId string

		resp, httpRes, err := apiClient.UsersApi.GetGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetGroups", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.GetGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetMe", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.GetMe(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetProfile", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.GetProfile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetRoleById", func(t *testing.T) {

		var roleId string

		resp, httpRes, err := apiClient.UsersApi.GetRoleById(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetRoles", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.GetRoles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetUser", func(t *testing.T) {

		var id string

		resp, httpRes, err := apiClient.UsersApi.GetUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService GetUsers", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService ResetUserPassword", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.ResetUserPassword(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService SearchUsers", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersApi.SearchUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService UpdateGroup", func(t *testing.T) {

		var groupId string

		resp, httpRes, err := apiClient.UsersApi.UpdateGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService UpdateRole", func(t *testing.T) {

		var roleId string

		resp, httpRes, err := apiClient.UsersApi.UpdateRole(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService UpdateUser", func(t *testing.T) {

		var id string

		resp, httpRes, err := apiClient.UsersApi.UpdateUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService ValidateRole", func(t *testing.T) {

		httpRes, err := apiClient.UsersApi.ValidateRole(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApiService ValidateUserPassword", func(t *testing.T) {

		httpRes, err := apiClient.UsersApi.ValidateUserPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
