/*
Unleash API

Testing ServiceAccountsAPIService

*/

package client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cleanupServiceAccount(apiClient *client.APIClient, serviceAccount *client.ServiceAccountSchema) {
	id := fmt.Sprint(serviceAccount.Id)
	httpRes, err := apiClient.ServiceAccountsAPI.DeleteServiceAccount(context.Background(), id).Execute()

	if err != nil || httpRes.StatusCode != 200 {
		fmt.Println("Failed to clean up a service after a test, this means the test is probably fine but your state is dirty and you should run a docker compose rm --force")
	}
}

func createServiceAccount(t *testing.T, name string, userName string, apiClient *client.APIClient) *client.ServiceAccountSchema {
	createProjectSchema := client.CreateServiceAccountSchema{}
	createProjectSchema.SetUsername(userName)
	createProjectSchema.SetName(name)
	createProjectSchema.SetRootRole(1)
	resp, httpRes, err := apiClient.ServiceAccountsAPI.CreateServiceAccount(context.Background()).CreateServiceAccountSchema(createProjectSchema).Execute()

	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, httpRes.StatusCode)

	return resp
}

func Test_client_ServiceAccountAPIService(t *testing.T) {
	apiClient := testClient()

	t.Run("Test ServiceAccountAPIService CreateServiceAccount", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			serviceAccount := createServiceAccount(t, "test-account", "test-username", apiClient)
			defer cleanupServiceAccount(apiClient, serviceAccount)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ServiceAccountAPIService DeleteServiceAccount", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			serviceAccount := createServiceAccount(t, "test-account-to-be-deleted", "to-be-deleted", apiClient)
			id := fmt.Sprint(serviceAccount.Id)

			httpRes, err := apiClient.ServiceAccountsAPI.DeleteServiceAccount(context.Background(), id).Execute()

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ServiceAccountAPIService GetServiceAccounts", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			serviceAccount := createServiceAccount(t, "first-service-account", "first-service-account", apiClient)
			defer cleanupServiceAccount(apiClient, serviceAccount)

			serviceAccount = createServiceAccount(t, "second-service-account", "second-service-account", apiClient)
			defer cleanupServiceAccount(apiClient, serviceAccount)

			resp, httpRes, err := apiClient.ServiceAccountsAPI.GetServiceAccounts(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, 2, len(resp.ServiceAccounts))
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ServiceAccountAPIService UpdateServiceAccount", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			serviceAccount := createServiceAccount(t, "test-account-to-update", "to-update", apiClient)
			defer cleanupServiceAccount(apiClient, serviceAccount)

			id := fmt.Sprint(serviceAccount.Id)
			updateServiceAccountSchema := client.UpdateServiceAccountSchema{}
			updateServiceAccountSchema.SetName("new-name")
			updateServiceAccountSchema.SetRootRole(2)

			updateResp, updateHttpRes, updateErr := apiClient.ServiceAccountsAPI.UpdateServiceAccount(context.Background(), id).UpdateServiceAccountSchema(updateServiceAccountSchema).Execute()

			require.Nil(t, updateErr)
			require.NotNil(t, updateResp)
			assert.Equal(t, 200, updateHttpRes.StatusCode)

			resp, httpRes, err := apiClient.ServiceAccountsAPI.GetServiceAccounts(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "new-name", *resp.ServiceAccounts[0].Name)

		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ServiceAccountAPIService CreateServiceAccountToken", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			serviceAccount := createServiceAccount(t, "test-account-to-create-token", "to-create-token", apiClient)
			defer cleanupServiceAccount(apiClient, serviceAccount)

			patSchema := client.CreatePatSchema{}
			patSchema.Description = "test token"
			patSchema.ExpiresAt = time.Now().Add(time.Hour * 24 * 365)

			resp, httpRes, err := apiClient.ServiceAccountsAPI.CreateServiceAccountToken(context.Background(), fmt.Sprint(serviceAccount.Id)).CreatePatSchema(patSchema).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 201, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ServiceAccountAPIService GetServiceAccountTokens", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			serviceAccount := createServiceAccount(t, "test-account-to-create-token", "to-create-token", apiClient)
			defer cleanupServiceAccount(apiClient, serviceAccount)

			patSchema := client.CreatePatSchema{}
			patSchema.Description = "test token"
			patSchema.ExpiresAt = time.Now().Add(time.Hour * 24 * 365)

			createResp, createHttpRes, createErr := apiClient.ServiceAccountsAPI.CreateServiceAccountToken(context.Background(), fmt.Sprint(serviceAccount.Id)).CreatePatSchema(patSchema).Execute()

			require.Nil(t, createErr)
			require.NotNil(t, createResp)
			assert.Equal(t, 201, createHttpRes.StatusCode)

			resp, httpRes, err := apiClient.ServiceAccountsAPI.GetServiceAccountTokens(context.Background(), fmt.Sprint(serviceAccount.Id)).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)

			assert.Equal(t, 1, len(resp.Pats))
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test ServiceAccountAPIService DeleteServiceAccountToken", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			serviceAccount := createServiceAccount(t, "test-account-to-create-token", "to-create-token", apiClient)
			defer cleanupServiceAccount(apiClient, serviceAccount)

			patSchema := client.CreatePatSchema{}
			patSchema.Description = "test token"
			patSchema.ExpiresAt = time.Now().Add(time.Hour * 24 * 365)

			createResp, _, createErr := apiClient.ServiceAccountsAPI.CreateServiceAccountToken(context.Background(), fmt.Sprint(serviceAccount.Id)).CreatePatSchema(patSchema).Execute()

			require.Nil(t, createErr)
			require.NotNil(t, createResp)

			resp, _, err := apiClient.ServiceAccountsAPI.GetServiceAccountTokens(context.Background(), fmt.Sprint(serviceAccount.Id)).Execute()
			assert.Equal(t, 1, len(resp.Pats))
			require.Nil(t, err)

			deleteRHttpRes, deleteErr := apiClient.ServiceAccountsAPI.DeleteServiceAccountToken(context.Background(), fmt.Sprint(serviceAccount.Id), fmt.Sprint(resp.Pats[0].Id)).Execute()
			require.Nil(t, deleteErr)
			assert.Equal(t, 200, deleteRHttpRes.StatusCode)

			resp, _, err = apiClient.ServiceAccountsAPI.GetServiceAccountTokens(context.Background(), fmt.Sprint(serviceAccount.Id)).Execute()
			assert.Equal(t, 0, len(resp.Pats))
			require.Nil(t, err)

		} else {
			t.Skip("Enterprise only feature")
		}
	})
}
