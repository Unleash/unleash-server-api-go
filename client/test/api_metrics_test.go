/*
Unleash API

Testing MetricsApiService

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

func Test_client_MetricsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MetricsApiService CreateApplication", func(t *testing.T) {

		var appName string

		httpRes, err := apiClient.MetricsApi.CreateApplication(context.Background(), appName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsApiService DeleteApplication", func(t *testing.T) {

		var appName string

		httpRes, err := apiClient.MetricsApi.DeleteApplication(context.Background(), appName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsApiService GetApplication", func(t *testing.T) {

		var appName string

		resp, httpRes, err := apiClient.MetricsApi.GetApplication(context.Background(), appName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsApiService GetApplications", func(t *testing.T) {

		resp, httpRes, err := apiClient.MetricsApi.GetApplications(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsApiService GetFeatureUsageSummary", func(t *testing.T) {

		var name string

		resp, httpRes, err := apiClient.MetricsApi.GetFeatureUsageSummary(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsApiService GetRawFeatureMetrics", func(t *testing.T) {

		var name string

		resp, httpRes, err := apiClient.MetricsApi.GetRawFeatureMetrics(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
