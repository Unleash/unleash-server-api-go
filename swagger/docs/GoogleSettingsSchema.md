# GoogleSettingsSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Is google OIDC enabled | [optional] [default to null]
**ClientId** | **string** | The google client id, used to authenticate against google | [default to null]
**ClientSecret** | **string** | The client secret used to authenticate the OAuth session used to log the user in | [default to null]
**UnleashHostname** | **string** | Name of the host allowed to access the Google authentication flow | [default to null]
**AutoCreate** | **bool** | Should Unleash create users based on the emails coming back in the authentication reply from Google | [optional] [default to null]
**EmailDomains** | **string** | A comma separated list of email domains that Unleash will auto create user accounts for. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

