/*
Unleash API

Testing APITokensAPIService

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

func Test_client_APITokensAPIService(t *testing.T) {

	configuration := client.NewConfiguration()
	apiClient := client.NewAPIClient(configuration)

	t.Run("Test APITokensAPIService CreateApiToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.APITokensAPI.CreateApiToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APITokensAPIService DeleteApiToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		httpRes, err := apiClient.APITokensAPI.DeleteApiToken(context.Background(), token).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APITokensAPIService GetAllApiTokens", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.APITokensAPI.GetAllApiTokens(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APITokensAPIService GetApiTokensByName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.APITokensAPI.GetApiTokensByName(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APITokensAPIService UpdateApiToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		httpRes, err := apiClient.APITokensAPI.UpdateApiToken(context.Background(), token).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
