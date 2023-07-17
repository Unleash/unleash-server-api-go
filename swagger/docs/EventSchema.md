# EventSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the event. An increasing natural number. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time the event happened as a RFC 3339-conformant timestamp. | [default to null]
**Type_** | **string** | What [type](https://docs.getunleash.io/reference/api/legacy/unleash/admin/events#event-type-description) of event this is | [default to null]
**CreatedBy** | **string** | Which user created this event | [default to null]
**Environment** | **string** | The feature toggle environment the event relates to, if applicable. | [optional] [default to null]
**Project** | **string** | The project the event relates to, if applicable. | [optional] [default to null]
**FeatureName** | **string** | The name of the feature toggle the event relates to, if applicable. | [optional] [default to null]
**Data** | [***EventSchemaData**](eventSchema_data.md) |  | [optional] [default to null]
**PreData** | [***EventSchemaPreData**](eventSchema_preData.md) |  | [optional] [default to null]
**Tags** | [**[]TagSchema**](tagSchema.md) | Any tags related to the event, if applicable. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

