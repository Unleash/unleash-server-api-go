/*
Unleash API

Testing UsersAPIService

*/

package client

import (
	"context"
	"fmt"
	"strconv"
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

func createGroup(t *testing.T, apiCient *client.APIClient, prefix string) *client.GroupSchema {
	createGroupSchema := client.CreateGroupSchema{}
	createGroupSchema.SetName(prefix + "-group")
	createGroupSchema.SetDescription("Test group description")

	resp, httpResp, err := apiCient.UsersAPI.CreateGroup(context.Background()).CreateGroupSchema(createGroupSchema).Execute()

	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, httpResp.StatusCode)
	return resp
}

func createGroupWithUsers(t *testing.T, apiCient *client.APIClient, prefix string, uids []int32) *client.GroupSchema {
	// Populate Users array
	var users []client.CreateGroupSchemaUsersInner
	for _, uid := range uids {
		users = append(users, client.CreateGroupSchemaUsersInner{
			User: client.CreateGroupSchemaUsersInnerUser{
				Id: uid,
			},
		})
	}
	createGroupSchema := client.CreateGroupSchema{}
	createGroupSchema.SetName(prefix + "-group")
	createGroupSchema.SetDescription("Test group with users description")
	createGroupSchema.SetUsers(users)

	resp, httpResp, err := apiCient.UsersAPI.CreateGroup(context.Background()).CreateGroupSchema(createGroupSchema).Execute()

	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, httpResp.StatusCode)
	return resp
}

func createGroupWithSSO(t *testing.T, apiCient *client.APIClient, prefix string, mappingSSO []string) *client.GroupSchema {
	// Set up group with SSO
	createGroupSchema := client.CreateGroupSchema{}
	createGroupSchema.SetName(prefix + "-group")
	createGroupSchema.SetDescription("Test group with SSO mappings description")
	createGroupSchema.SetMappingsSSO(mappingSSO)

	resp, httpResp, err := apiCient.UsersAPI.CreateGroup(context.Background()).CreateGroupSchema(createGroupSchema).Execute()

	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, httpResp.StatusCode)
	return resp
}

func cleanUpGroup(apiClient *client.APIClient, group *client.GroupSchema) {
	groupID := fmt.Sprint(group.Id)
	httpResp, err := apiClient.UsersAPI.DeleteGroup(context.Background(), groupID).Execute()
	if err != nil || httpResp.StatusCode != 200 {
		fmt.Println("Failed to clean up a service after a test, this means the test is probably fine but your state is dirty and you should run a docker compose rm --force")
	}
}

func Test_client_UsersAPIService(t *testing.T) {
	apiClient := testClient()
	t.Run("Test UsersAPIService Minimal CreateGroup", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			group := createGroup(t, apiClient, "test")
			defer cleanUpGroup(apiClient, group)
		} else {
			t.Skip("Enterprise only feature")
		}
	})
	t.Run("Test UsersAPIService Users CreateGroup", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			user1 := createUser(t, apiClient, "test-user-1")
			defer clean_up_user(user1, apiClient)
			user2 := createUser(t, apiClient, "test-user-2")
			defer clean_up_user(user2, apiClient)

			group := createGroupWithUsers(t, apiClient, "test", []int32{user1.Id, user2.Id})
			defer cleanUpGroup(apiClient, group)
			assert.NotNil(t, group.Users)
			assert.Equal(t, 2, len(group.Users))
		} else {
			t.Skip("Enterprise only feature")
		}
	})
	t.Run("Test UsersAPIService SSO CreateGroup", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			SSOs := [...]string{"SSOGroup1", "SSOGroup2"}
			group := createGroupWithSSO(t, apiClient, "test", SSOs[:])
			defer cleanUpGroup(apiClient, group)
			assert.NotNil(t, group.MappingsSSO)
			assert.Equal(t, 2, len(group.MappingsSSO))
			assert.Contains(t, group.MappingsSSO, "SSOGroup1")
			assert.Contains(t, group.MappingsSSO, "SSOGroup2")
		} else {
			t.Skip("Enterprise only feature")
		}
	})
	t.Run("Test UsersAPIService DeleteGroup", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			group := createGroup(t, apiClient, "to-be-deleted")
			gID := string(strconv.Itoa(int(*group.Id)))
			httpRes, err := apiClient.UsersAPI.DeleteGroup(context.Background(), gID).Execute()

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})
	t.Run("Test UsersAPIService GetGroup", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			group := createGroup(t, apiClient, "to-be-retrieved")
			gID := fmt.Sprint(group.Id)
			defer cleanUpGroup(apiClient, group)

			resp, httpRes, err := apiClient.UsersAPI.GetGroup(context.Background(), gID).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, group.Id, resp.Id)
			assert.Equal(t, group.Name, resp.Name)
		} else {
			t.Skip("Enterprise only feature")
		}
	})
	t.Run("Test UsersAPIService GetGroups", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			group1 := createGroup(t, apiClient, "to-be-retrieved-1")
			group2 := createGroup(t, apiClient, "to-be-retrieved-2")

			defer cleanUpGroup(apiClient, group1)
			defer cleanUpGroup(apiClient, group2)

			resp, httpRes, err := apiClient.UsersAPI.GetGroups(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, resp.Groups)
			assert.NotEmpty(t, resp.Groups)
		} else {
			t.Skip("Enterprise only feature")
		}
	})
	t.Run("Test UsersAPIService UpdateGroup", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			group := createGroup(t, apiClient, "to-be-updated")
			gID := fmt.Sprint(group.Id)
			defer cleanUpGroup(apiClient, group)

			newName := "newTestName"
			newDescription := "newTestDescription"
			updateGroup := client.CreateGroupSchema{}

			updateGroup.SetName(newName)
			updateGroup.SetDescription(newDescription)

			resp, httpRes, err := apiClient.UsersAPI.UpdateGroup(context.Background(), gID).CreateGroupSchema(updateGroup).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, newName, resp.Name)
			assert.Equal(t, newDescription, *resp.Description.Get())
		} else {
			t.Skip("Enterprise only feature")
		}
	})

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

	t.Run("Test UsersAPIService SearchUsers", func(t *testing.T) {
		user := createUser(t, apiClient, "searchable")
		defer clean_up_user(user, apiClient)

		testCases := []struct {
			name        string
			query       string
			description string
			expected    bool
		}{
			{
				name:        "Search by email",
				query:       "searchable-username@getunleash.io",
				description: "Should find the user in search results by email",
				expected:    true,
			},
			{
				name:        "Search by username",
				query:       "searchable-username",
				description: "Should find the user in search results by username",
				expected:    true,
			},
			{
				name:        "Search by username",
				query:       "sear",
				description: "Should find the user in search results by part of username",
				expected:    true,
			},
			{
				name:        "Search by part of username in the middle",
				query:       "able-user",
				description: "Should not find the user as search is from the beginning",
				expected:    false,
			},
			{
				name:        "Search by less than 2 characters",
				query:       "s",
				description: "Should not find the user as search requires at least 2 characters",
				expected:    false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				resp, httpRes, err := apiClient.UsersAPI.SearchUsers(context.Background()).Q(tc.query).Execute()

				require.Nil(t, err)
				require.NotNil(t, resp)
				assert.Equal(t, 200, httpRes.StatusCode)
				if tc.expected {
					assert.NotEmpty(t, resp.Users)
				} else {
					assert.Empty(t, resp.Users)
				}

				found := false
				for _, searchedUser := range resp.Users {
					if searchedUser.Id == user.Id {
						found = true
						break
					}
				}
				if tc.expected {
					assert.True(t, found, tc.description)
				} else {
					assert.False(t, found, tc.description)
				}
			})
		}
	})
}
