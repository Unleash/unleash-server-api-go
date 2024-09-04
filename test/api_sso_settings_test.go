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

	t.Run("Test AuthAPIService SetSamlSettings Enable path", func(t *testing.T) {
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

	t.Run("Test AuthAPIService SetSamlSettings Disable path", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {

			innerSettings := client.NewSamlSettingsSchemaOneOf1WithDefaults()

			innerSettings.SetEnabled(false)
			// You still need to set these fields even if you're disabling SAML, because reasons
			// they can't be empty strings though, so... this. It's what the UI does. For reasons
			innerSettings.SetEntityId(" ")
			innerSettings.SetSignOnUrl(" ")
			innerSettings.SetCertificate(" ")

			samlSettings := client.SamlSettingsSchema{
				SamlSettingsSchemaOneOf1: innerSettings,
			}
			samlSettingsResponse, httpRes, err := apiClient.AuthAPI.SetSamlSettings(context.Background()).SamlSettingsSchema(samlSettings).Execute()

			require.Nil(t, err)
			require.NotNil(t, samlSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test AuthAPIService GetSamlSettings", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {

			innerSettings := client.NewSamlSettingsSchemaOneOf1WithDefaults()

			innerSettings.SetEnabled(false)
			innerSettings.SetEntityId("this is a thing")
			innerSettings.SetSignOnUrl("this is also a thing, apparently not a url tho")
			innerSettings.SetCertificate("whee I'm a certificate!")

			samlSettings := client.SamlSettingsSchema{
				SamlSettingsSchemaOneOf1: innerSettings,
			}
			setSamlResponse, httpRes, err := apiClient.AuthAPI.SetSamlSettings(context.Background()).SamlSettingsSchema(samlSettings).Execute()

			require.Nil(t, err)
			require.NotNil(t, setSamlResponse)
			require.Equal(t, 200, httpRes.StatusCode)

			samlSettingsResponse, httpRes, err := apiClient.AuthAPI.GetSamlSettings(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, samlSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)

			require.Equal(t, *samlSettingsResponse.Enabled, false)
			require.Equal(t, *samlSettingsResponse.EntityId, "this is a thing")
			require.Equal(t, *samlSettingsResponse.SignOnUrl, "this is also a thing, apparently not a url tho")
			require.Equal(t, *samlSettingsResponse.Certificate, "whee I'm a certificate!")

		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test AuthAPIService SetOidcSettings Enable path", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {

			innerSettings := client.NewOidcSettingsSchemaOneOfWithDefaults()

			innerSettings.SetEnabled(true)
			innerSettings.SetClientId("test")
			innerSettings.SetSecret("test")
			innerSettings.SetAutoCreate(true)
			innerSettings.SetEnableSingleSignOut(false)
			innerSettings.SetDefaultRootRoleId(1)
			innerSettings.SetAddGroupsScope(false)
			innerSettings.SetEnableGroupSyncing(false)

			// Be warned, this can and will make Unleash make a slow, blocking http call to this listed URL
			// failure to find the expected well known configuration will cause Unleash to freak out and block
			// Should be caught by the mock Python server in the docker compose, but if you're running this outside of docker
			// know that your Unleash instance will be saying hi to this endpoint
			innerSettings.SetDiscoverUrl("http://mock-openid-server:9000/.well-known/openid-configuration")

			// somehow this is required but it doesn't appear to do anything useful. Leaving it but commented out
			// in case the next victim finds a usecase for it and wonders why it wasn't there
			// innerSettings.SetDefaultRootRole("Admin")

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

	t.Run("Test AuthAPIService SetOidcSettings Disable path", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {

			innerSettings := client.NewOidcSettingsSchemaOneOf1WithDefaults()

			innerSettings.SetEnabled(false)

			//These properties are required prior to v6.0 of Unleash, but are not required in v6.0 and later
			innerSettings.SetClientId("test")
			innerSettings.SetSecret("test")
			innerSettings.SetAutoCreate(true)
			innerSettings.SetEnableSingleSignOut(false)
			innerSettings.SetDefaultRootRoleId(1)
			innerSettings.SetAddGroupsScope(false)
			innerSettings.SetEnableGroupSyncing(false)

			innerSettings.SetDiscoverUrl("http://mock-openid-server:9000/.well-known/openid-configuration")

			oidcSettings := client.OidcSettingsSchema{
				OidcSettingsSchemaOneOf1: innerSettings,
			}
			oidcSettingsResponse, httpRes, err := apiClient.AuthAPI.SetOidcSettings(context.Background()).OidcSettingsSchema(oidcSettings).Execute()

			require.Nil(t, err)
			require.NotNil(t, oidcSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)

		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test AuthAPIService GetOidcSettings", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {

			innerSettings := client.NewOidcSettingsSchemaOneOf1WithDefaults()

			innerSettings.SetEnabled(true)
			innerSettings.SetClientId("look ma, I'm an id!")
			innerSettings.SetSecret("shhh, it's a secret")
			innerSettings.SetAutoCreate(true)
			innerSettings.SetEnableSingleSignOut(false)
			innerSettings.SetDefaultRootRoleId(1)
			innerSettings.SetAddGroupsScope(false)
			innerSettings.SetEnableGroupSyncing(false)

			// can't avoid this unless we want to get an empty "isEnabled: false" response back and I
			// want a real response back
			innerSettings.SetDiscoverUrl("http://mock-openid-server:9000/.well-known/openid-configuration")

			oidcSettings := client.OidcSettingsSchema{
				OidcSettingsSchemaOneOf1: innerSettings,
			}
			oidcSettingsResponse, httpRes, err := apiClient.AuthAPI.SetOidcSettings(context.Background()).OidcSettingsSchema(oidcSettings).Execute()

			require.Nil(t, err)
			require.NotNil(t, oidcSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)

			oidcSettingsResponse, httpRes, err = apiClient.AuthAPI.GetOidcSettings(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, oidcSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)

			require.Equal(t, *oidcSettingsResponse.Enabled, true)
			require.Equal(t, *oidcSettingsResponse.ClientId, "look ma, I'm an id!")
			require.Equal(t, *oidcSettingsResponse.Secret, "shhh, it's a secret")
			require.Equal(t, *oidcSettingsResponse.AutoCreate, true)
			require.Equal(t, *oidcSettingsResponse.EnableSingleSignOut, false)
			require.Equal(t, *oidcSettingsResponse.DefaultRootRoleId, float32(1.0))
			require.Equal(t, *oidcSettingsResponse.AddGroupsScope, false)
			require.Equal(t, *oidcSettingsResponse.EnableGroupSyncing, false)

		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test AuthAPIService SetSimpleSettings", func(t *testing.T) {

		if enterpriseEnvironmentAvailable() {
			simpleSettingsRequest := client.NewPasswordAuthSchema()
			simpleSettingsRequest.SetDisabled(false)

			simpleSettingsResponse, httpRes, err := apiClient.AuthAPI.SetSimpleSettings(context.Background()).PasswordAuthSchema(*simpleSettingsRequest).Execute()

			require.Nil(t, err)
			require.NotNil(t, simpleSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)
		} else {
			t.Skip("Enterprise only feature")
		}
	})

	t.Run("Test AuthAPIService GetSimpleSettings", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			simpleSettingsRequest := client.NewPasswordAuthSchema()
			simpleSettingsRequest.SetDisabled(true)

			simpleSettingsResponse, httpRes, err := apiClient.AuthAPI.SetSimpleSettings(context.Background()).PasswordAuthSchema(*simpleSettingsRequest).Execute()

			require.Nil(t, err)
			require.NotNil(t, simpleSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)

			simpleSettingsResponse, httpRes, err = apiClient.AuthAPI.GetSimpleSettings(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, simpleSettingsResponse)
			require.Equal(t, 200, httpRes.StatusCode)

			require.Equal(t, *simpleSettingsResponse.Disabled, true)
		} else {
			t.Skip("Enterprise only feature")
		}
	})
}
