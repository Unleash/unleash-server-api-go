# FeatureEnvironmentSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment | [default to null]
**FeatureName** | **string** |  | [optional] [default to null]
**Environment** | **string** |  | [optional] [default to null]
**Type_** | **string** | The type of the environment | [optional] [default to null]
**Enabled** | **bool** | &#x60;true&#x60; if the feature is enabled for the environment, otherwise &#x60;false&#x60;. | [default to null]
**SortOrder** | **float64** | The sort order of the feature environment in the feature environments list | [optional] [default to null]
**VariantCount** | **float64** | The number of defined variants | [optional] [default to null]
**Strategies** | [**[]FeatureStrategySchema**](featureStrategySchema.md) | A list of activation strategies for the feature environment | [optional] [default to null]
**Variants** | [**[]VariantSchema**](variantSchema.md) | A list of variants for the feature environment | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

