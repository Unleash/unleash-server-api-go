/*
Unleash API

Testing ProjectsApiService

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

func Test_client_ProjectsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsApiService AddAccessToProject", func(t *testing.T) {

		var projectId string
		var roleId string

		httpRes, err := apiClient.ProjectsApi.AddAccessToProject(context.Background(), projectId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService AddDefaultStrategyToProjectEnvironment", func(t *testing.T) {

		var projectId string
		var environment string

		resp, httpRes, err := apiClient.ProjectsApi.AddDefaultStrategyToProjectEnvironment(context.Background(), projectId, environment).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService AddEnvironmentToProject", func(t *testing.T) {

		var projectId string

		httpRes, err := apiClient.ProjectsApi.AddEnvironmentToProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService AddRoleToUser", func(t *testing.T) {

		var projectId string
		var userId string
		var roleId string

		httpRes, err := apiClient.ProjectsApi.AddRoleToUser(context.Background(), projectId, userId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ChangeRoleForGroup", func(t *testing.T) {

		var projectId string
		var groupId string
		var roleId string

		httpRes, err := apiClient.ProjectsApi.ChangeRoleForGroup(context.Background(), projectId, groupId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ChangeRoleForUser", func(t *testing.T) {

		var projectId string
		var userId string
		var roleId string

		httpRes, err := apiClient.ProjectsApi.ChangeRoleForUser(context.Background(), projectId, userId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService CreateProject", func(t *testing.T) {

		resp, httpRes, err := apiClient.ProjectsApi.CreateProject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService CreateProjectApiToken", func(t *testing.T) {

		var projectId string

		resp, httpRes, err := apiClient.ProjectsApi.CreateProjectApiToken(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService DeleteProject", func(t *testing.T) {

		var projectId string

		httpRes, err := apiClient.ProjectsApi.DeleteProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService DeleteProjectApiToken", func(t *testing.T) {

		var projectId string
		var token string

		httpRes, err := apiClient.ProjectsApi.DeleteProjectApiToken(context.Background(), projectId, token).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService GetProjectAccess", func(t *testing.T) {

		var projectId string

		resp, httpRes, err := apiClient.ProjectsApi.GetProjectAccess(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService GetProjectApiTokens", func(t *testing.T) {

		var projectId string

		resp, httpRes, err := apiClient.ProjectsApi.GetProjectApiTokens(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService GetProjectHealthReport", func(t *testing.T) {

		var projectId string

		resp, httpRes, err := apiClient.ProjectsApi.GetProjectHealthReport(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService GetProjectOverview", func(t *testing.T) {

		var projectId string

		resp, httpRes, err := apiClient.ProjectsApi.GetProjectOverview(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService GetProjectUsers", func(t *testing.T) {

		var projectId string

		resp, httpRes, err := apiClient.ProjectsApi.GetProjectUsers(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService GetProjects", func(t *testing.T) {

		resp, httpRes, err := apiClient.ProjectsApi.GetProjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService RemoveEnvironmentFromProject", func(t *testing.T) {

		var projectId string
		var environment string

		httpRes, err := apiClient.ProjectsApi.RemoveEnvironmentFromProject(context.Background(), projectId, environment).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService RemoveRoleForUser", func(t *testing.T) {

		var projectId string
		var userId string
		var roleId string

		httpRes, err := apiClient.ProjectsApi.RemoveRoleForUser(context.Background(), projectId, userId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService RemoveRoleFromGroup", func(t *testing.T) {

		var projectId string
		var groupId string
		var roleId string

		httpRes, err := apiClient.ProjectsApi.RemoveRoleFromGroup(context.Background(), projectId, groupId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService UpdateProject", func(t *testing.T) {

		var projectId string

		httpRes, err := apiClient.ProjectsApi.UpdateProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ValidateProject", func(t *testing.T) {

		httpRes, err := apiClient.ProjectsApi.ValidateProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
