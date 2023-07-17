# ServiceAccountSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** | The service account id | [default to null]
**IsAPI** | **bool** | Deprecated: for internal use only, should not be exposed through the API | [optional] [default to null]
**Name** | **string** | The name of the service account | [optional] [default to null]
**Email** | **string** | Deprecated: service accounts don&#x27;t have emails associated with them | [optional] [default to null]
**Username** | **string** | The service account username | [optional] [default to null]
**ImageUrl** | **string** | The service account image url | [optional] [default to null]
**InviteLink** | **string** | Deprecated: service accounts cannot be invited via an invitation link | [optional] [default to null]
**LoginAttempts** | **float64** | Deprecated: service accounts cannot log in to Unleash | [optional] [default to null]
**EmailSent** | **bool** | Deprecated: internal use only | [optional] [default to null]
**RootRole** | **int32** | The root role id associated with the service account | [optional] [default to null]
**SeenAt** | [**time.Time**](time.Time.md) | Deprecated. This property is always &#x60;null&#x60;. To find out when a service account was last seen, check its &#x60;tokens&#x60; list and refer to each token&#x27;s &#x60;lastSeen&#x60; property instead. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The service account creation date | [optional] [default to null]
**Tokens** | [**[]PatSchema**](patSchema.md) | The list of tokens associated with the service account | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

