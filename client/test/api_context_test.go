/*
Unleash API

Testing ContextApiService

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

func Test_client_ContextApiService(t *testing.T) {

	configuration := client.NewConfiguration()
	apiClient := client.NewAPIClient(configuration)

	t.Run("Test ContextApiService CreateContextField", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ContextApi.CreateContextField(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContextApiService DeleteContextField", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contextField string

		httpRes, err := apiClient.ContextApi.DeleteContextField(context.Background(), contextField).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContextApiService GetContextField", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contextField string

		resp, httpRes, err := apiClient.ContextApi.GetContextField(context.Background(), contextField).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContextApiService GetContextFields", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ContextApi.GetContextFields(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContextApiService UpdateContextField", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contextField string

		httpRes, err := apiClient.ContextApi.UpdateContextField(context.Background(), contextField).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContextApiService Validate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ContextApi.Validate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
