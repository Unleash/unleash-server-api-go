package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func resetDefaultConfig(apiClient *client.APIClient, project, environment string) {
	enableChangeRequest := *client.NewUpdateChangeRequestEnvironmentConfigSchemaWithDefaults()
	enableChangeRequest.ChangeRequestsEnabled = false

	apiClient.ChangeRequestsAPI.UpdateProjectChangeRequestConfig(context.Background(), project, environment).UpdateChangeRequestEnvironmentConfigSchema(enableChangeRequest).Execute()
}

func Test_client_ChangeRequestService(t *testing.T) {
	apiClient := testClient()

	if !enterpriseEnvironmentAvailable() {
		t.Skip("Enterprise only feature")
		return
	}

	t.Run("Test ChangeRequestAPIService EnableChangeRequest", func(t *testing.T) {
		project := "default"
		environment := "production"

		var config, httpRes, err = apiClient.ChangeRequestsAPI.GetProjectChangeRequestConfig(context.Background(), project).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, config)
		require.NotNil(t, config[1])

		require.Equal(t, config[1].RequiredApprovals, *client.NewNullableFloat32(nil))
		require.Equal(t, config[1].Type, "production")
		require.Equal(t, config[1].Environment, "production")
		require.Equal(t, config[1].ChangeRequestEnabled, false)

		enableChangeRequest := *client.NewUpdateChangeRequestEnvironmentConfigSchemaWithDefaults()
		enableChangeRequest.ChangeRequestsEnabled = true
		enableChangeRequest.SetRequiredApprovals(2)

		httpRes, err = apiClient.ChangeRequestsAPI.UpdateProjectChangeRequestConfig(context.Background(), project, environment).UpdateChangeRequestEnvironmentConfigSchema(enableChangeRequest).Execute()
		defer resetDefaultConfig(apiClient, project, environment)

		if err != nil {
			fmt.Println(err)
		}

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)

		config, httpRes, err = apiClient.ChangeRequestsAPI.GetProjectChangeRequestConfig(context.Background(), project).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, config)
		require.NotNil(t, config[1])

		requiredApprovals := float32(2)
		require.Equal(t, config[1].RequiredApprovals, *client.NewNullableFloat32(&requiredApprovals))
		require.Equal(t, config[1].Type, "production")
		require.Equal(t, config[1].Environment, "production")
		require.Equal(t, config[1].ChangeRequestEnabled, true)
	})
}
