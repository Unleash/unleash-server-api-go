/*
Unleash API

Testing ProjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_client_ProjectsAPIService(t *testing.T) {

	configuration := client.NewConfiguration()
	apiClient := client.NewAPIClient(configuration)

	t.Run("Test ProjectsAPIService AddAccessToProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		httpRes, err := apiClient.ProjectsAPI.AddAccessToProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService AddDefaultStrategyToProjectEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var environment string

		resp, httpRes, err := apiClient.ProjectsAPI.AddDefaultStrategyToProjectEnvironment(context.Background(), projectId, environment).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService AddEnvironmentToProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		httpRes, err := apiClient.ProjectsAPI.AddEnvironmentToProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService AddRoleAccessToProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var roleId string

		httpRes, err := apiClient.ProjectsAPI.AddRoleAccessToProject(context.Background(), projectId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService AddRoleToUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var userId string
		var roleId string

		httpRes, err := apiClient.ProjectsAPI.AddRoleToUser(context.Background(), projectId, userId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ChangeRoleForGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var groupId string
		var roleId string

		httpRes, err := apiClient.ProjectsAPI.ChangeRoleForGroup(context.Background(), projectId, groupId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ChangeRoleForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var userId string
		var roleId string

		httpRes, err := apiClient.ProjectsAPI.ChangeRoleForUser(context.Background(), projectId, userId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService CreateProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.CreateProject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService CreateProjectApiToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.CreateProjectApiToken(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService DeleteProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		httpRes, err := apiClient.ProjectsAPI.DeleteProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService DeleteProjectApiToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var token string

		httpRes, err := apiClient.ProjectsAPI.DeleteProjectApiToken(context.Background(), projectId, token).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetProjectAccess", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.GetProjectAccess(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetProjectApiTokens", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.GetProjectApiTokens(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetProjectDora", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.GetProjectDora(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetProjectHealthReport", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.GetProjectHealthReport(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetProjectOverview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.GetProjectOverview(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetProjectUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.GetProjectUsers(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetProjects", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.GetProjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetRoleProjectAccess", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleId string

		resp, httpRes, err := apiClient.ProjectsAPI.GetRoleProjectAccess(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RemoveEnvironmentFromProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var environment string

		httpRes, err := apiClient.ProjectsAPI.RemoveEnvironmentFromProject(context.Background(), projectId, environment).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RemoveGroupAccess", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var groupId string

		httpRes, err := apiClient.ProjectsAPI.RemoveGroupAccess(context.Background(), projectId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RemoveRoleForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var userId string
		var roleId string

		httpRes, err := apiClient.ProjectsAPI.RemoveRoleForUser(context.Background(), projectId, userId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RemoveRoleFromGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var groupId string
		var roleId string

		httpRes, err := apiClient.ProjectsAPI.RemoveRoleFromGroup(context.Background(), projectId, groupId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RemoveUserAccess", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var userId string

		httpRes, err := apiClient.ProjectsAPI.RemoveUserAccess(context.Background(), projectId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService SetRolesForGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var groupId string

		httpRes, err := apiClient.ProjectsAPI.SetRolesForGroup(context.Background(), projectId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService SetRolesForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var userId string

		httpRes, err := apiClient.ProjectsAPI.SetRolesForUser(context.Background(), projectId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService UpdateProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		httpRes, err := apiClient.ProjectsAPI.UpdateProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ValidateProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ValidateProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
