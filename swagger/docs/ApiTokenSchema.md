# ApiTokenSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | **string** | The token used for authentication. | [default to null]
**Username** | **string** | This property was deprecated in Unleash v5. Prefer the &#x60;tokenName&#x60; property instead. | [optional] [default to null]
**TokenName** | **string** | A unique name for this particular token | [default to null]
**Type_** | **string** | The type of API token | [default to null]
**Environment** | **string** | The environment the token has access to. &#x60;*&#x60; if it has access to all environments. | [optional] [default to null]
**Project** | **string** | The project this token belongs to. | [default to null]
**Projects** | **[]string** | The list of projects this token has access to. If the token has access to specific projects they will be listed here. If the token has access to all projects it will be represented as &#x60;[*]&#x60; | [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | The token&#x27;s expiration date. NULL if the token doesn&#x27;t have an expiration set. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the token was created. | [default to null]
**SeenAt** | [**time.Time**](time.Time.md) | When the token was last seen/used to authenticate with. NULL if the token has not yet been used for authentication. | [optional] [default to null]
**Alias** | **string** | Alias is no longer in active use and will often be NULL. It&#x27;s kept around as a way of allowing old proxy tokens created with the old metadata format to keep working. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

