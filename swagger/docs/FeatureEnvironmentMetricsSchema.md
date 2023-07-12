# FeatureEnvironmentMetricsSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureName** | **string** | The name of the feature | [optional] [default to null]
**AppName** | **string** | The name of the application the SDK is being used in | [optional] [default to null]
**Environment** | **string** | Which environment the SDK is being used in | [default to null]
**Timestamp** | [***DateSchema**](dateSchema.md) |  | [default to null]
**Yes** | **int32** | How many times the toggle evaluated to true | [default to null]
**No** | **int32** | How many times the toggle evaluated to false | [default to null]
**Variants** | **map[string]int32** | How many times each variant was returned | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

