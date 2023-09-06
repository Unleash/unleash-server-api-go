/*
Unleash API

Testing APITokensAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"testing"
	"time"

	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_client_APITokensAPIService(t *testing.T) {

	apiClient := testClient()
	var createdTokens []string

	t.Run("Test APITokensAPIService CreateApiToken (admin)", func(t *testing.T) {

		expireDate, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00.000Z")
		require.Nil(t, err)
		token := client.CreateApiTokenSchemaOneOf{
			Type:      "admin",
			TokenName: "test admin token",
			ExpiresAt: &expireDate,
		}
		createTokenSchema := client.CreateApiTokenSchemaOneOfAsCreateApiTokenSchema(&token)
		resp, httpRes, err := apiClient.APITokensAPI.CreateApiToken(context.Background()).CreateApiTokenSchema(createTokenSchema).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 201, httpRes.StatusCode)

		createdTokens = append(createdTokens, resp.Secret)
	})

	t.Run("Test APITokensAPIService CreateApiToken (client)", func(t *testing.T) {

		expireDate, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00.000Z")
		require.Nil(t, err)
		environment := "development"
		token := client.CreateApiTokenSchemaOneOf2{
			Type:        "client",
			TokenName:   "test client token",
			ExpiresAt:   &expireDate,
			Environment: &environment,
		}
		createTokenSchema := client.CreateApiTokenSchemaOneOf2AsCreateApiTokenSchema(&token)
		resp, httpRes, err := apiClient.APITokensAPI.CreateApiToken(context.Background()).CreateApiTokenSchema(createTokenSchema).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 201, httpRes.StatusCode)

		createdTokens = append(createdTokens, resp.Secret)
	})

	t.Run("Test APITokensAPIService CreateApiToken (frontend)", func(t *testing.T) {

		expireDate, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00.000Z")
		require.Nil(t, err)
		environment := "development"
		project := "default"
		token := client.CreateApiTokenSchemaOneOf2{
			Type:        "frontend",
			TokenName:   "test frontend token",
			ExpiresAt:   &expireDate,
			Environment: &environment,
			Project:     &project,
		}
		createTokenSchema := client.CreateApiTokenSchemaOneOf2AsCreateApiTokenSchema(&token)
		resp, httpRes, err := apiClient.APITokensAPI.CreateApiToken(context.Background()).CreateApiTokenSchema(createTokenSchema).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 201, httpRes.StatusCode)

		createdTokens = append(createdTokens, resp.Secret)
	})

	t.Run("Test APITokensAPIService DeleteApiToken", func(t *testing.T) {

		for i := 0; i < len(createdTokens); i++ {
			token := createdTokens[i]

			httpRes, err := apiClient.APITokensAPI.DeleteApiToken(context.Background(), token).Execute()

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		}

	})

	t.Run("Test APITokensAPIService GetAllApiTokens", func(t *testing.T) {

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

	t.Run("Test APITokensAPIService GetApiTokensByName and UpdateApiToken", func(t *testing.T) {

		// get the same token we're using for the test and update it to expire a year from now
		token := apiClient.GetConfig().DefaultHeader["Authorization"]
		newExpireDate := time.Now().Add(time.Hour * 24 * 365 * 1)

		updateTokenExpiration := client.NewUpdateApiTokenSchema(newExpireDate)
		httpRes, err := apiClient.APITokensAPI.UpdateApiToken(context.Background(), token).UpdateApiTokenSchema(*updateTokenExpiration).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
