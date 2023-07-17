# PatSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identification number for this Personal Access Token. (This property is set by Unleash when the token is created and cannot be set manually: if you provide a value when creating a PAT, Unleash will ignore it.) | [optional] [default to null]
**Secret** | **string** | The token used for authentication. (This property is set by Unleash when the token is created and cannot be set manually: if you provide a value when creating a PAT, Unleash will ignore it.) | [optional] [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | The token&#x27;s expiration date. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the token was created. (This property is set by Unleash when the token is created and cannot be set manually: if you provide a value when creating a PAT, Unleash will ignore it.) | [optional] [default to null]
**SeenAt** | [**time.Time**](time.Time.md) | When the token was last seen/used to authenticate with. &#x60;null&#x60; if it has not been used yet. (This property is set by Unleash when the token is created and cannot be set manually: if you provide a value when creating a PAT, Unleash will ignore it.) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

