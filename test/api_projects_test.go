/*
Unleash API

Testing ProjectsAPIService

*/

package client

import (
	"context"
	"testing"

	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_client_ProjectsAPIService(t *testing.T) {
	apiClient := testClient()

	t.Run("Test ProjectsAPIService CreateProject", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			createProjectSchema := *client.NewCreateProjectSchema("pet-shop-project", "Pet shop project")
			resp, httpRes, err := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProjectSchema(createProjectSchema).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 201, httpRes.StatusCode)
		}
	})

	t.Run("Test ProjectsAPIService DeleteProject", func(t *testing.T) {

		if enterpriseEnvironmentAvailable() {
			projectId := "pet-shop-project"

			httpRes, err := apiClient.ProjectsAPI.DeleteProject(context.Background(), projectId).Execute()

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		}
	})

	t.Run("Test ProjectsAPIService AddAccessToProject", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			projectId := "default"
			adminUser := 1
			memberRole := 5
			roles := make([]int32, 1)
			roles[0] = int32(memberRole)
			users := make([]int32, 1)
			users[0] = int32(adminUser)
			groups := make([]int32, 0)
			addAccess := client.NewProjectAddAccessSchema(roles, groups, users)
			httpRes, err := apiClient.ProjectsAPI.AddAccessToProject(context.Background(), projectId).ProjectAddAccessSchema(*addAccess).Execute()

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		}
	})

}
