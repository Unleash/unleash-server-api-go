# EventSchemaData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the feature toggle/strategy/environment that this event relates to | [optional] [default to null]
**Description** | **string** | The description of the object this event relates to | [optional] [default to null]
**Type_** | **string** | If this event relates to a feature toggle, the type of feature toggle. | [optional] [default to null]
**Project** | **string** | The project this event relates to | [optional] [default to null]
**Stale** | **bool** | Is the feature toggle this event relates to stale | [optional] [default to null]
**Variants** | [**[]VariantSchema**](variantSchema.md) | Variants configured for this toggle | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time the event happened as a RFC 3339-conformant timestamp. | [optional] [default to null]
**LastSeenAt** | [**time.Time**](time.Time.md) | The time the feature was last seen | [optional] [default to null]
**ImpressionData** | **bool** | Should [impression events](https://docs.getunleash.io/reference/impression-data) activate for this feature toggle | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

