/*
Unleash API

Testing EventsApiService

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

func Test_client_EventsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EventsApiService GetEvents", func(t *testing.T) {

		resp, httpRes, err := apiClient.EventsApi.GetEvents(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GetEventsForToggle", func(t *testing.T) {

		var featureName string

		resp, httpRes, err := apiClient.EventsApi.GetEventsForToggle(context.Background(), featureName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService SearchEvents", func(t *testing.T) {

		resp, httpRes, err := apiClient.EventsApi.SearchEvents(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
