/*
Unleash API

Testing UsersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"testing"

	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_client_UsersAPIService(t *testing.T) {
	apiClient := testClient()

	t.Run("Test UsersAPIService ChangeMyPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UsersAPI.ChangeMyPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ChangeUserPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.UsersAPI.ChangeUserPassword(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService CreateUser", func(t *testing.T) {

		name := "A name"
		email := "test-username@getunleash.io"
		username := "test-username"
		sendEmail := false
		rootRoleId := int32(1)

		createUserSchema := *client.NewCreateUserSchemaWithDefaults()
		createUserSchema.Name = &name
		createUserSchema.Email = &email
		createUserSchema.Username = &username
		createUserSchema.RootRole = client.Int32AsCreateUserSchemaRootRole(&rootRoleId)
		createUserSchema.SendEmail = &sendEmail

		resp, httpRes, err := apiClient.UsersAPI.CreateUser(context.Background()).CreateUserSchema(createUserSchema).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 201, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DeleteUser", func(t *testing.T) {

		var id = "2"

		httpRes, err := apiClient.UsersAPI.DeleteUser(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetAdminCount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.GetAdminCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetBaseUsersAndGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.GetBaseUsersAndGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetMe", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersAPI.GetMe(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.GetProfile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUser", func(t *testing.T) {
		var id string

		resp, httpRes, err := apiClient.UsersAPI.GetUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUsers", func(t *testing.T) {

		resp, httpRes, err := apiClient.UsersAPI.GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test UsersAPIService ResetUserPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.ResetUserPassword(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService SearchUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.SearchUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.UsersAPI.UpdateUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ValidateUserPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UsersAPI.ValidateUserPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
