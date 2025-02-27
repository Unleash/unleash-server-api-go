package client

import (
	"context"
	"testing"

	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/require"
)

func cleanupEnvironment(t *testing.T, apiClient *client.APIClient, name string) {
	_, err := apiClient.EnvironmentsAPI.RemoveEnvironment(context.Background(), name).Execute()

	if err != nil {
		t.Log(err)
	}

	require.Nil(t, err)
}

func Test_client_EnvironmentAPIService(t *testing.T) {
	apiClient := testClient()

	if !enterpriseEnvironmentAvailable() {
		// I think early exiting here is probably okay. While GetEnvironment does exist on OSS
		// it's not useful without the ability to CRUD environments.
		t.Skip("Enterprise only feature")
		return
	}

	t.Run("Test EnvironmentAPIService CreateEnvironment", func(t *testing.T) {
		envName := "RainForest"

		rainForestEnvironment := *client.NewCreateEnvironmentSchemaWithDefaults()
		rainForestEnvironment.Name = envName
		rainForestEnvironment.SetType("production")
		rainForestEnvironment.SetEnabled(true)

		resp, httpRes, err := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background()).CreateEnvironmentSchema(rainForestEnvironment).Execute()

		if err != nil {
			t.Log(err)
		}

		defer cleanupEnvironment(t, apiClient, envName)

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 201, httpRes.StatusCode)
	})

	t.Run("Test EnvironmentAPIService GetEnvironment", func(t *testing.T) {
		envName := "RainForest"

		rainForestEnvironment := *client.NewCreateEnvironmentSchemaWithDefaults()
		rainForestEnvironment.Name = envName
		rainForestEnvironment.SetType("production")
		rainForestEnvironment.SetEnabled(true)

		_, _, err := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background()).CreateEnvironmentSchema(rainForestEnvironment).Execute()

		if err != nil {
			t.Log(err)
		}

		defer cleanupEnvironment(t, apiClient, envName)

		resp, httpRes, err := apiClient.EnvironmentsAPI.GetEnvironment(context.Background(), envName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test EnvironmentAPIService ListEnvironments", func(t *testing.T) {
		rainForestEnvironment := *client.NewCreateEnvironmentSchemaWithDefaults()
		rainForestEnvironment.Name = "RainForest"
		rainForestEnvironment.SetType("production")
		rainForestEnvironment.SetEnabled(true)

		_, _, err := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background()).CreateEnvironmentSchema(rainForestEnvironment).Execute()

		if err != nil {
			t.Log(err)
		}

		defer cleanupEnvironment(t, apiClient, "RainForest")

		desertEnvironment := *client.NewCreateEnvironmentSchemaWithDefaults()
		desertEnvironment.Name = "Desert"
		desertEnvironment.SetType("production")
		desertEnvironment.SetEnabled(true)

		_, _, err = apiClient.EnvironmentsAPI.CreateEnvironment(context.Background()).CreateEnvironmentSchema(desertEnvironment).Execute()

		if err != nil {
			t.Log(err)
		}

		defer cleanupEnvironment(t, apiClient, "Desert")

		resp, httpRes, err := apiClient.EnvironmentsAPI.GetAllEnvironments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)
		require.Len(t, resp.Environments, 5) // we expect the two we created + the 3 built in ones
	})

	t.Run("Test EnvironmentAPIService UpdateEnvironment", func(t *testing.T) {
		envName := "RainForest"

		rainForestEnvironment := *client.NewCreateEnvironmentSchemaWithDefaults()
		rainForestEnvironment.Name = envName
		rainForestEnvironment.SetType("production")
		rainForestEnvironment.SetEnabled(true)

		_, _, err := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background()).CreateEnvironmentSchema(rainForestEnvironment).Execute()

		if err != nil {
			t.Log(err)
		}

		defer cleanupEnvironment(t, apiClient, envName)

		updateEnvironmentSchema := client.UpdateEnvironmentSchema{}
		updateEnvironmentSchema.SetType("development")

		resp, httpRes, err := apiClient.EnvironmentsAPI.UpdateEnvironment(context.Background(), envName).UpdateEnvironmentSchema(updateEnvironmentSchema).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)

		getResp, _, err := apiClient.EnvironmentsAPI.GetEnvironment(context.Background(), envName).Execute()

		require.Nil(t, err)
		require.NotNil(t, getResp)
		require.Equal(t, "development", getResp.Type)
	})

	t.Run("Test EnvironmentAPIService ToggleOnOff", func(t *testing.T) {
		envName := "RainForest"

		rainForestEnvironment := *client.NewCreateEnvironmentSchemaWithDefaults()
		rainForestEnvironment.Name = envName
		rainForestEnvironment.SetType("production")
		rainForestEnvironment.SetEnabled(true)

		_, _, err := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background()).CreateEnvironmentSchema(rainForestEnvironment).Execute()

		if err != nil {
			t.Log(err)
		}

		defer cleanupEnvironment(t, apiClient, envName)

		httpRes, err := apiClient.EnvironmentsAPI.ToggleEnvironmentOff(context.Background(), envName).Execute()

		require.Nil(t, err)
		require.Equal(t, 204, httpRes.StatusCode)

		//validate the env is off

		getResp, _, err := apiClient.EnvironmentsAPI.GetEnvironment(context.Background(), envName).Execute()

		require.Nil(t, err)
		require.NotNil(t, getResp)
		require.False(t, getResp.Enabled)

		//toggle it back on

		httpRes, err = apiClient.EnvironmentsAPI.ToggleEnvironmentOn(context.Background(), envName).Execute()

		require.Nil(t, err)
		require.Equal(t, 204, httpRes.StatusCode)

		//validate the env is on

		getResp, _, err = apiClient.EnvironmentsAPI.GetEnvironment(context.Background(), envName).Execute()

		require.Nil(t, err)
		require.NotNil(t, getResp)
		require.True(t, getResp.Enabled)
	})
}
