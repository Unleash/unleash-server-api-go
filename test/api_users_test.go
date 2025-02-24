/*
Unleash API

Testing UsersAPIService

*/

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
	email := prefix + "-username@getunleash.io"
	username := prefix + "-username"
	sendEmail := false
	rootRoleId := int32(1)

	createUserSchema := *client.NewCreateUserSchemaWithDefaults()
	createUserSchema.Name = &name
	createUserSchema.Email = &email
	createUserSchema.Username = &username
	createUserSchema.RootRole = client.Int32AsCreateUserSchemaRootRole(&rootRoleId)
	createUserSchema.SendEmail = &sendEmail

	resp, httpRes, err := apiClient.UsersAPI.CreateUser(context.Background()).CreateUserSchema(createUserSchema).Execute()

	fmt.Println(err)
	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, httpRes.StatusCode)
	return resp
}
func createRole(name string, description string, roleType string) *client.CreateRoleWithPermissionsSchema {
	var newRole *client.CreateRoleWithPermissionsSchema
	if roleType == "root-custom" {
		inner := client.NewCreateRoleWithPermissionsSchemaAnyOf(name)
		inner.Type = &roleType
		inner.Description = &description
		permission1 := client.NewCreateRoleWithPermissionsSchemaAnyOfPermissionsInner("CREATE_PROJECT")
		permission2 := client.NewCreateRoleWithPermissionsSchemaAnyOfPermissionsInner("UPDATE_APPLICATION")
		permissions := []client.CreateRoleWithPermissionsSchemaAnyOfPermissionsInner{}
		permissions = append(permissions, *permission1)
		permissions = append(permissions, *permission2)
		inner.SetPermissions(permissions)
		wrapper := client.CreateRoleWithPermissionsSchema{}
		wrapper.CreateRoleWithPermissionsSchemaAnyOf = inner
		newRole = &wrapper
	}
	if roleType == "custom" {
		env := "development"
		inner := client.NewCreateRoleWithPermissionsSchemaAnyOf1(name)
		inner.Type = &roleType
		inner.Description = &description
		// permission1 is a project permission
		permission1 := client.NewCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner(float32(2))
		// permission2 is an environment permission
		permission2 := client.NewCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner(float32(37))
		permission2.SetEnvironment(env)
		permissions := []client.CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner{}
		permissions = append(permissions, *permission1)
		permissions = append(permissions, *permission2)
		inner.SetPermissions(permissions)
		wrapper := client.CreateRoleWithPermissionsSchema{}
		wrapper.CreateRoleWithPermissionsSchemaAnyOf1 = inner
		newRole = &wrapper
	}
	return newRole
}

func clean_up_user(user *client.CreateUserResponseSchema, apiClient *client.APIClient) {
	httpRes, err := apiClient.UsersAPI.DeleteUser(context.Background(), user.Id).Execute()
	if err != nil || httpRes.StatusCode != 200 {
		fmt.Println("Failed to clean up a user after a test, this means the test is probably fine but your state is dirty and you should run a docker compose rm --force")
	}
}

func clean_up_role(role *client.RoleWithVersionSchema, apiClient *client.APIClient) {
	roleId := fmt.Sprint(role.Roles.Id)
	httpRes, err := apiClient.UsersAPI.DeleteRole(context.Background(), roleId).Execute()
	if err != nil || httpRes.StatusCode != 200 {
		fmt.Println("Failed to clean up a role after a test, this means the test is probably fine but your state is dirty and you should run a docker compose rm --force")
	}
}

func Test_client_UsersAPIService(t *testing.T) {
	apiClient := testClient()

	t.Run("Test UsersAPIService CreateUser", func(t *testing.T) {
		user := createUser(t, apiClient, "test")
		defer clean_up_user(user, apiClient)
	})

	t.Run("Test UsersAPIService DeleteUser", func(t *testing.T) {

		user := createUser(t, apiClient, "to-be-deleted")

		httpRes, err := apiClient.UsersAPI.DeleteUser(context.Background(), user.Id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test UsersAPIService GetUser", func(t *testing.T) {
		user := createUser(t, apiClient, "to-be-retrieved")
		defer clean_up_user(user, apiClient)

		resp, httpRes, err := apiClient.UsersAPI.GetUser(context.Background(), user.Id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUser", func(t *testing.T) {

		user := createUser(t, apiClient, "to-be-updated")
		defer clean_up_user(user, apiClient)

		role := int32(2)
		rootRole := client.Int32AsCreateUserSchemaRootRole(&role)
		newName := "newName"
		updateUser := client.NewUpdateUserSchema()
		updateUser.Name = &newName
		updateUser.RootRole = &rootRole

		resp, httpRes, err := apiClient.UsersAPI.UpdateUser(context.Background(), user.Id).UpdateUserSchema(*updateUser).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService CreateRole (root-custom)", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			newRole := createRole("test-root-role", "test role description", "root-custom")

			resp, httpRes, err := apiClient.UsersAPI.CreateRole(context.Background()).CreateRoleWithPermissionsSchema(*newRole).Execute()
			defer clean_up_role(resp, apiClient)

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test UsersAPIService CreateRole (custom)", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			newRole := createRole("test-custom-role", "test role description", "custom")

			resp, httpRes, err := apiClient.UsersAPI.CreateRole(context.Background()).CreateRoleWithPermissionsSchema(*newRole).Execute()
			defer clean_up_role(resp, apiClient)

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test UsersAPIService UpdateRole", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			newRole := createRole("to-update-role", "test role description", "root-custom")

			resp, httpRes, err := apiClient.UsersAPI.CreateRole(context.Background()).CreateRoleWithPermissionsSchema(*newRole).Execute()
			defer clean_up_role(resp, apiClient)

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)

			roleId := fmt.Sprint(resp.Roles.Id)
			updatedRole := createRole("updated-role-name", "New role description", "root-custom")
			resp, httpRes, err = apiClient.UsersAPI.UpdateRole(context.Background(), roleId).CreateRoleWithPermissionsSchema(*updatedRole).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test UsersAPIService GetRoles", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {

			newRole := createRole("test-root-role", "test role description", "root-custom")

			createResp, _, err := apiClient.UsersAPI.CreateRole(context.Background()).CreateRoleWithPermissionsSchema(*newRole).Execute()

			require.Nil(t, err)
			defer clean_up_role(createResp, apiClient)

			resp, httpRes, err := apiClient.UsersAPI.GetRoles(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotEmpty(t, resp.Roles)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test UsersAPIService GetRoleById", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {

			roleId := "2"

			resp, httpRes, err := apiClient.UsersAPI.GetRoleById(context.Background(), roleId).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, resp.Name, "Editor")
			assert.Equal(t, resp.Type, "root")
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test UsersAPIService DeleteRole", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			newRole := createRole("role-to-be-deleted", "ephimeral role", "root-custom")
			resp, httpRes, err := apiClient.UsersAPI.CreateRole(context.Background()).CreateRoleWithPermissionsSchema(*newRole).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)

			// convert resp.Roles.Id to string and assign to roleId
			roleId := fmt.Sprint(resp.Roles.Id)

			httpRes, err = apiClient.UsersAPI.DeleteRole(context.Background(), roleId).Execute()

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})
}
