/*
Unleash API

Testing EnvironmentsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	openapiclient "github.com/Unleash/unleash-server-api-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_client_EnvironmentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnvironmentsApiService CloneEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		httpRes, err := apiClient.EnvironmentsApi.CloneEnvironment(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService CreateEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EnvironmentsApi.CreateEnvironment(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService GetAllEnvironments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EnvironmentsApi.GetAllEnvironments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService GetEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.EnvironmentsApi.GetEnvironment(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService GetProjectEnvironments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.EnvironmentsApi.GetProjectEnvironments(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService RemoveEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		httpRes, err := apiClient.EnvironmentsApi.RemoveEnvironment(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService ToggleEnvironmentOff", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		httpRes, err := apiClient.EnvironmentsApi.ToggleEnvironmentOff(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService ToggleEnvironmentOn", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		httpRes, err := apiClient.EnvironmentsApi.ToggleEnvironmentOn(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService UpdateEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.EnvironmentsApi.UpdateEnvironment(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService UpdateSortOrder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.EnvironmentsApi.UpdateSortOrder(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsApiService ValidateEnvironmentName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.EnvironmentsApi.ValidateEnvironmentName(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
