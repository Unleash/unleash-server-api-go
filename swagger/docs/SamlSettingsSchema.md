# SamlSettingsSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Is SAML authentication enabled | [optional] [default to null]
**EntityId** | **string** | The SAML 2.0 entity ID | [default to null]
**SignOnUrl** | **string** | Which URL to use for Single Sign On | [default to null]
**Certificate** | **string** | The X509 certificate used to validate requests | [default to null]
**SignOutUrl** | **string** | Which URL to use for Single Sign Out | [optional] [default to null]
**SpCertificate** | **string** | Signing certificate for sign out requests | [optional] [default to null]
**AutoCreate** | **bool** | Should Unleash create users based on the emails coming back in the authentication reply from the SAML server | [optional] [default to null]
**EmailDomains** | **string** | A comma separated list of email domains that Unleash will auto create user accounts for. | [optional] [default to null]
**DefaultRootRole** | **string** | Assign this root role to auto created users | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

