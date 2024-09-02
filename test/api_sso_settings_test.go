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

			// innerSettings := client.NewOidcSettingsSchema()

			// innerSettings.SetEnabled(true)
			// innerSettings.SetDiscoverUrl("test.com")
			// innerSettings.SetClientId("test")
			// innerSettings.SetSecret("test")
			// innerSettings.SetAutoCreate(true)
			// innerSettings.SetEnableSingleSignOut(false)
			// innerSettings.SetDefaultRootRole("Admin")
			// innerSettings.SetDefaultRootRoleId(1)
			// innerSettings.SetEmailDomains("example.com")
			// innerSettings.SetAcrValues("test")
			// innerSettings.SetIdTokenSigningAlgorithm("RSA256")
			// innerSettings.SetEnableGroupSyncing(false)
			// innerSettings.SetGroupJsonPath(".groups")
			// innerSettings.SetAddGroupsScope(false)

			// // oidcSettings := client.OidcSettingsSchema{
			// // 	OidcSettingsSchemaOneOf: innerSettings,
			// // }
			// oidcSettingsResponse, httpRes, err := apiClient.AuthAPI.SetOidcSettings(context.Background()).OidcSettingsResponseSchema(*innerSettings).Execute()

			// require.Nil(t, err)
			// require.NotNil(t, oidcSettingsResponse)
			// require.Equal(t, 200, httpRes.StatusCode)

		} else {
			t.Skip("Enterprise only feature")
		}
	})
}
