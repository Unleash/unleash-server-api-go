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

	t.Run("Test ProjectsAPIService SetAccessToProject", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			projectId := "default"
			adminUser := int32(1)
			ownerRole := int32(4)
			memberRole := int32(5)
			roles := make([]client.ProjectAccessConfigurationSchemaRolesInner, 2)
			roles[0] = *client.NewProjectAccessConfigurationSchemaRolesInnerWithDefaults()
			roles[0].Id = &ownerRole
			ownerUsers := make([]int32, 1)
			ownerUsers = append(ownerUsers, adminUser)
			roles[0].Users = ownerUsers

			roles[1] = *client.NewProjectAccessConfigurationSchemaRolesInnerWithDefaults()
			roles[1].Id = &memberRole
			memberUsers := make([]int32, 1)
			memberUsers = append(memberUsers, adminUser)
			roles[0].Users = memberUsers

			setAccess := client.NewProjectAccessConfigurationSchema(roles)
			httpRes, err := apiClient.ProjectsAPI.SetProjectAccess(context.Background(), projectId).ProjectAccessConfigurationSchema(*setAccess).Execute()

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		}
	})

}
