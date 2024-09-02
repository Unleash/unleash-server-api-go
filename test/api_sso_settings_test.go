/*
Unleash API

Testing ServiceAccountsAPIService

*/

package client

import (
	"context"
	"testing"

	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/require"
)

func Test_client_SSOSettingsAPIService(t *testing.T) {
	apiClient := testClient()

	t.Run("Test AuthAPIService SetSamlSettings", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {

			innerSettings := client.NewSamlSettingsSchemaOneOfWithDefaults()

			innerSettings.SetEnabled(true)
			innerSettings.SetEntityId("https://example.com")
			innerSettings.SetSignOnUrl("https://example.com")
			innerSettings.SetCertificate("test-cert")
			innerSettings.SetSignOutUrl("https://example.com")
			innerSettings.SetSpCertificate("test-cert")
			innerSettings.SetAutoCreate(true)
			innerSettings.SetDefaultRootRole("Admin")
			innerSettings.SetEmailDomains("example.com")
			innerSettings.SetEnableGroupSyncing(false)
			innerSettings.SetGroupJsonPath("test")

			samlSettings := client.SamlSettingsSchema{
				SamlSettingsSchemaOneOf: innerSettings,
			}
			samlSettingsResponse, httpRes, err := apiClient.AuthAPI.SetSamlSettings(context.Background()).SamlSettingsSchema(samlSettings).Execute()

			require.Nil(t, err)
			require.NotNil(t, samlSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test AuthAPIService SetOidcSettings", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {

			innerSettings := client.NewOidcSettingsSchemaOneOfWithDefaults()

			innerSettings.SetEnabled(true)

			// Be warned, this can and will make Unleash make a slow, blocking http call to this listed URL
			// failure to find the expected well known configuration will cause Unleash to freak out and block
			// Should be caught by the mock Python server in the docker compose, but if you're running this outside of docker
			// know that your Unleash instance will be saying hi to this endpoint
			innerSettings.SetDiscoverUrl("http://mock-openid-server:9000/.well-known/openid-configuration")
			innerSettings.SetClientId("test")
			innerSettings.SetSecret("test")
			innerSettings.SetAutoCreate(true)
			innerSettings.SetEnableSingleSignOut(false)
			innerSettings.SetDefaultRootRole("Admin")
			innerSettings.SetDefaultRootRoleId(1)
			innerSettings.SetAddGroupsScope(false)
			innerSettings.SetEnableGroupSyncing(false)

			oidcSettings := client.OidcSettingsSchema{
				OidcSettingsSchemaOneOf: innerSettings,
			}
			oidcSettingsResponse, httpRes, err := apiClient.AuthAPI.SetOidcSettings(context.Background()).OidcSettingsSchema(oidcSettings).Execute()

			require.Nil(t, err)
			require.NotNil(t, oidcSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)

		} else {
			t.Skip("Enterprise only feature")
		}
	})
}
