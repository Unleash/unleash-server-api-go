/*
Unleash API

Testing ProjectsAPIService

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

func createTempUser(t *testing.T, apiClient *client.APIClient, prefix string) *client.CreateUserResponseSchema {
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

func cleanupTempUser(t *testing.T, apiClient *client.APIClient, userId int32) {
	httpRes, err := apiClient.UsersAPI.DeleteUser(context.Background(), userId).Execute()

	require.Nil(t, err)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func cleanupProject(t *testing.T, apiClient *client.APIClient, projectId string) {
	httpRes, err := apiClient.ProjectsAPI.DeleteProject(context.Background(), projectId).Execute()

	require.Nil(t, err)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func Test_client_ProjectsAPIService(t *testing.T) {
	apiClient := testClient()

	t.Run("Test ProjectsAPIService CreateProject", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			createProjectSchema := *client.NewCreateProjectSchema("pet-shop-project")
			createProjectSchema.SetId("pet-shop-project")
			resp, httpRes, err := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProjectSchema(createProjectSchema).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 201, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ProjectsAPIService DeleteProject", func(t *testing.T) {

		if enterpriseEnvironmentAvailable() {
			projectId := "pet-shop-project"

			httpRes, err := apiClient.ProjectsAPI.DeleteProject(context.Background(), projectId).Execute()

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
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
			roles[0].Users = []int32{adminUser}

			roles[1] = *client.NewProjectAccessConfigurationSchemaRolesInnerWithDefaults()
			roles[1].Id = &memberRole
			roles[1].Users = []int32{adminUser}

			setAccess := client.NewProjectAccessConfigurationSchema(roles)
			httpRes, err := apiClient.ProjectsAPI.SetProjectAccess(context.Background(), projectId).ProjectAccessConfigurationSchema(*setAccess).Execute()

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ProjectsAPIService GetProjectAccess", func(t *testing.T) {

		if enterpriseEnvironmentAvailable() {

			projectId := "default"

			resp, httpRes, err := apiClient.ProjectsAPI.GetProjectAccess(context.Background(), projectId).Execute()

			fmt.Println(err)

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "Owner", resp.Roles[0].Name)
			assert.Equal(t, "Member", resp.Roles[1].Name)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ProjectsAPIService GetProjects correctly handles ownership responses", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			projectId := "pet-shop-project"
			createProjectSchema := *client.NewCreateProjectSchema("some-random-project")
			createProjectSchema.SetId(projectId)
			resp, httpRes, err := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProjectSchema(createProjectSchema).Execute()

			defer cleanupProject(t, apiClient, projectId)

			user := createTempUser(t, apiClient, "some-random-user")
			defer cleanupTempUser(t, apiClient, user.Id)

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 201, httpRes.StatusCode)

			rolePayload := *client.NewProjectAccessConfigurationSchemaRolesInnerWithDefaults()
			rolePayload.Users = []int32{user.Id}
			id := int32(4)
			rolePayload.Id = &id

			roles := []client.ProjectAccessConfigurationSchemaRolesInner{rolePayload}

			accessConfiguration := *client.NewProjectAccessConfigurationSchema(roles)

			apiResponse, err := apiClient.ProjectsAPI.SetProjectAccess(context.Background(), projectId).ProjectAccessConfigurationSchema(accessConfiguration).Execute()

			require.Nil(t, err)
			require.NotNil(t, apiResponse)

			projects, projectApiResponse, err := apiClient.ProjectsAPI.GetProjects(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, projectApiResponse)
			require.NotNil(t, projects)

		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ProjectsAPIService UpdateSettings correctly updates the project mode", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			projectId := "pet-shop-project"
			createProjectSchema := *client.NewCreateProjectSchema("some-random-project")
			createProjectSchema.SetId(projectId)
			createResp, createResponse, createErr := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProjectSchema(createProjectSchema).Execute()

			require.Nil(t, createErr)
			require.NotNil(t, createResp)
			assert.Equal(t, 201, createResponse.StatusCode)

			updateProjectSettingsRequest := *client.NewUpdateProjectEnterpriseSettingsSchemaWithDefaults()
			updateProjectSettingsRequest.SetMode("private")

			_, err := apiClient.ProjectsAPI.UpdateProjectEnterpriseSettings(context.Background(), projectId).UpdateProjectEnterpriseSettingsSchema(updateProjectSettingsRequest).Execute()
			require.Nil(t, err)

			defer cleanupProject(t, apiClient, projectId)

			projects, projectApiResponse, err := apiClient.ProjectsAPI.GetProjects(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, projectApiResponse)
			require.NotNil(t, projects)

			var project *client.ProjectSchema
			for _, p := range projects.Projects {
				if p.Id == projectId {
					project = &p
					break
				}
			}

			require.NotNil(t, project)
			assert.Equal(t, "private", *project.Mode)

		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ProjectsAPIService addEnvironmentToProject adds an environment to project", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			projectId := "pet-shop-project"
			createProjectSchema := *client.NewCreateProjectSchema("some-random-project")
			createProjectSchema.SetId(projectId)
			project, createResponse, createProjectErr := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProjectSchema(createProjectSchema).Execute()

			if createProjectErr != nil {
				fmt.Println(createProjectErr)
			}

			require.Nil(t, createProjectErr)
			require.NotNil(t, project)
			assert.Equal(t, 201, createResponse.StatusCode)

			defer cleanupProject(t, apiClient, projectId)

			environmentName := "Production"
			environmentType := "production"
			environmentEnabled := true

			createEnvironmentSchema := *client.NewCreateEnvironmentSchemaWithDefaults()
			createEnvironmentSchema.SetName(environmentName)
			createEnvironmentSchema.SetType(environmentType)
			createEnvironmentSchema.Enabled = &environmentEnabled

			environment, httpResp, createEnvironmentErr := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background()).CreateEnvironmentSchema(createEnvironmentSchema).Execute()

			defer cleanupEnvironment(t, apiClient, environmentName)

			if createEnvironmentErr != nil {
				fmt.Println(createEnvironmentErr)
			}

			require.Nil(t, createEnvironmentErr)
			require.NotNil(t, environment)
			assert.Equal(t, 201, httpResp.StatusCode)

			addEnvironmentToProjectSchema := *client.NewProjectEnvironmentSchemaWithDefaults()
			addEnvironmentToProjectSchema.SetEnvironment(environmentName)

			httpResp, err := apiClient.ProjectsAPI.AddEnvironmentToProject(context.Background(), projectId).ProjectEnvironmentSchema(addEnvironmentToProjectSchema).Execute()

			if err != nil {
				fmt.Println(err)
			}

			require.Nil(t, err)
			assert.Equal(t, 200, httpResp.StatusCode)

			environments, httpResp, err := apiClient.EnvironmentsAPI.GetProjectEnvironments(context.Background(), projectId).Execute()

			if err != nil {
				fmt.Println(err)
			}

			require.Nil(t, err)
			require.NotNil(t, environments)
			assert.Equal(t, 200, httpResp.StatusCode)

			var foundEnvironment *client.EnvironmentProjectSchema
			for _, e := range environments.Environments {
				if e.Name == environmentName {
					foundEnvironment = &e
					break
				}
			}

			require.NotNil(t, foundEnvironment)
			assert.Equal(t, environmentName, foundEnvironment.Name)
			assert.Equal(t, environmentType, foundEnvironment.Type)
			assert.Equal(t, true, *foundEnvironment.Visible)

		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ProjectsAPIService removeEnvironmentToProject adds an environment to project", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			projectId := "pet-shop-project"
			createProjectSchema := *client.NewCreateProjectSchema("some-random-project")
			createProjectSchema.SetId(projectId)
			project, createResponse, createProjectErr := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProjectSchema(createProjectSchema).Execute()

			if createProjectErr != nil {
				fmt.Println(createProjectErr)
			}

			require.Nil(t, createProjectErr)
			require.NotNil(t, project)
			assert.Equal(t, 201, createResponse.StatusCode)

			defer cleanupProject(t, apiClient, projectId)

			environmentName := "Production"
			environmentType := "production"
			environmentEnabled := true

			createEnvironmentSchema := *client.NewCreateEnvironmentSchemaWithDefaults()
			createEnvironmentSchema.SetName(environmentName)
			createEnvironmentSchema.SetType(environmentType)
			createEnvironmentSchema.Enabled = &environmentEnabled

			environment, createHttpResp, createEnvironmentErr := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background()).CreateEnvironmentSchema(createEnvironmentSchema).Execute()

			defer cleanupEnvironment(t, apiClient, environmentName)

			if createEnvironmentErr != nil {
				fmt.Println(createEnvironmentErr)
			}

			require.Nil(t, createEnvironmentErr)
			require.NotNil(t, environment)
			assert.Equal(t, 201, createHttpResp.StatusCode)

			addEnvironmentToProjectSchema := *client.NewProjectEnvironmentSchemaWithDefaults()
			addEnvironmentToProjectSchema.SetEnvironment(environmentName)

			var httpResp, err = apiClient.ProjectsAPI.AddEnvironmentToProject(context.Background(), projectId).ProjectEnvironmentSchema(addEnvironmentToProjectSchema).Execute()

			if err != nil {
				fmt.Println(err)
			}

			require.Nil(t, err)
			assert.Equal(t, 200, httpResp.StatusCode)

			httpResp, err = apiClient.ProjectsAPI.RemoveEnvironmentFromProject(context.Background(), projectId, environmentName).Execute()

			if err != nil {
				fmt.Println(err)
			}

			require.Nil(t, err)
			assert.Equal(t, 200, httpResp.StatusCode)

			environments, httpResp, err := apiClient.EnvironmentsAPI.GetProjectEnvironments(context.Background(), projectId).Execute()

			if err != nil {
				fmt.Println(err)
			}

			require.Nil(t, err)
			require.NotNil(t, environments)
			assert.Equal(t, 200, httpResp.StatusCode)

			var foundEnvironment *client.EnvironmentProjectSchema
			for _, e := range environments.Environments {
				if e.Name == environmentName {
					foundEnvironment = &e
					break
				}
			}

			require.NotNil(t, foundEnvironment)
			assert.Equal(t, environmentName, foundEnvironment.Name)
			assert.Equal(t, environmentType, foundEnvironment.Type)
			assert.Equal(t, false, *foundEnvironment.Visible)
		} else {
			t.Skip("Enterprise only feature")
		}
	})
}
