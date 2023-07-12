# OidcSettingsSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | &#x60;true&#x60; if OpenID connect is turned on for this instance, otherwise &#x60;false&#x60; | [optional] [default to null]
**DiscoverUrl** | **string** | The [.well-known OpenID discover URL](https://swagger.io/docs/specification/authentication/openid-connect-discovery/) | [optional] [default to null]
**ClientId** | **string** | The OIDC client ID of this application. | [default to null]
**Secret** | **string** | Shared secret from OpenID server. Used to authenticate login requests | [default to null]
**AutoCreate** | **bool** | Auto create users based on email addresses from login tokens | [optional] [default to null]
**EnableSingleSignOut** | **bool** | Support Single sign out when user clicks logout in Unleash. If &#x60;true&#x60; user is signed out of all OpenID Connect sessions against the clientId they may have active | [optional] [default to null]
**DefaultRootRole** | **string** | [Default role](https://docs.getunleash.io/reference/rbac#standard-roles) granted to users auto-created from email. Only relevant if autoCreate is &#x60;true&#x60; | [optional] [default to null]
**EmailDomains** | **string** | Comma separated list of email domains that are automatically approved for an account in the server. Only relevant if autoCreate is &#x60;true&#x60; | [optional] [default to null]
**AcrValues** | **string** | Authentication Context Class Reference, used to request extra values in the acr claim returned from the server. If multiple values are required, they should be space separated.   Consult [the OIDC reference](https://openid.net/specs/openid-connect-core-1_0.html#AuthorizationEndpoint) for more information   | [optional] [default to null]
**IdTokenSigningAlgorithm** | **string** | The signing algorithm used to sign our token. Refer to the [JWT signatures](https://jwt.io/introduction) documentation for more information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

