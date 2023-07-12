# FeatureSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique feature name | [default to null]
**Type_** | **string** | Type of the toggle e.g. experiment, kill-switch, release, operational, permission | [optional] [default to null]
**Description** | **string** | Detailed description of the feature | [optional] [default to null]
**Archived** | **bool** | &#x60;true&#x60; if the feature is archived | [optional] [default to null]
**Project** | **string** | Name of the project the feature belongs to | [optional] [default to null]
**Enabled** | **bool** | &#x60;true&#x60; if the feature is enabled, otherwise &#x60;false&#x60;. | [optional] [default to null]
**Stale** | **bool** | &#x60;true&#x60; if the feature is stale based on the age and feature type, otherwise &#x60;false&#x60;. | [optional] [default to null]
**Favorite** | **bool** | &#x60;true&#x60; if the feature was favorited, otherwise &#x60;false&#x60;. | [optional] [default to null]
**ImpressionData** | **bool** | &#x60;true&#x60; if the impression data collection is enabled for the feature, otherwise &#x60;false&#x60;. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date the feature was created | [optional] [default to null]
**ArchivedAt** | [**time.Time**](time.Time.md) | The date the feature was archived | [optional] [default to null]
**LastSeenAt** | [**time.Time**](time.Time.md) | The date when metrics where last collected for the feature | [optional] [default to null]
**Environments** | [**[]FeatureEnvironmentSchema**](featureEnvironmentSchema.md) | The list of environments where the feature can be used | [optional] [default to null]
**Variants** | [**[]VariantSchema**](variantSchema.md) | The list of feature variants | [optional] [default to null]
**Strategies** | [**[]interface{}**](interface{}.md) | This is a legacy field that will be deprecated | [optional] [default to null]
**Tags** | [**[]TagSchema**](tagSchema.md) | The list of feature tags | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

