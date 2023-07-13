/*
Unleash API

Testing StrategiesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func Test_openapi_StrategiesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StrategiesApiService CreateStrategy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StrategiesApi.CreateStrategy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StrategiesApiService DeprecateStrategy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var strategyName string

		httpRes, err := apiClient.StrategiesApi.DeprecateStrategy(context.Background(), strategyName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StrategiesApiService GetAllStrategies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StrategiesApi.GetAllStrategies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StrategiesApiService GetStrategiesByContextField", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contextField string

		resp, httpRes, err := apiClient.StrategiesApi.GetStrategiesByContextField(context.Background(), contextField).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StrategiesApiService GetStrategy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.StrategiesApi.GetStrategy(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StrategiesApiService ReactivateStrategy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var strategyName string

		httpRes, err := apiClient.StrategiesApi.ReactivateStrategy(context.Background(), strategyName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StrategiesApiService RemoveStrategy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.StrategiesApi.RemoveStrategy(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StrategiesApiService UpdateFeatureStrategySegments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StrategiesApi.UpdateFeatureStrategySegments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StrategiesApiService UpdateStrategy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var strategyName string

		httpRes, err := apiClient.StrategiesApi.UpdateStrategy(context.Background(), strategyName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
