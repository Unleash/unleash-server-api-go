/*
Unleash API

Testing PlaygroundAPIService

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

func Test_client_PlaygroundAPIService(t *testing.T) {

	configuration := client.NewConfiguration()
	apiClient := client.NewAPIClient(configuration)

	t.Run("Test PlaygroundAPIService GetAdvancedPlayground", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PlaygroundAPI.GetAdvancedPlayground(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaygroundAPIService GetPlayground", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PlaygroundAPI.GetPlayground(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
