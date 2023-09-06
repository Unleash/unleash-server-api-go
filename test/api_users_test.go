/*
Unleash API

Testing UsersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T, apiClient *client.APIClient, prefix string) *client.CreateUserResponseSchema {
	name := "A name"
	email := prefix + "username@getunleash.io"
	username := prefix + "username"
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
	return resp
}

func Test_client_UsersAPIService(t *testing.T) {
	apiClient := testClient()

	t.Run("Test UsersAPIService CreateUser", func(t *testing.T) {

		createUser(t, apiClient, "test-")

	})

	t.Run("Test UsersAPIService DeleteUser", func(t *testing.T) {

		user := createUser(t, apiClient, "to-be-deleted-")
		id := fmt.Sprint(user.Id)

		httpRes, err := apiClient.UsersAPI.DeleteUser(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUser", func(t *testing.T) {
		var id string

		resp, httpRes, err := apiClient.UsersAPI.GetUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUser", func(t *testing.T) {

		user := createUser(t, apiClient, "to-be-updated-")
		id := fmt.Sprint(user.Id)

		role := int32(2)
		rootRole := client.Int32AsCreateUserSchemaRootRole(&role)
		newName := "newName"
		updateUser := client.NewUpdateUserSchema()
		updateUser.Name = &newName
		updateUser.RootRole = &rootRole

		resp, httpRes, err := apiClient.UsersAPI.UpdateUser(context.Background(), id).UpdateUserSchema(*updateUser).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
