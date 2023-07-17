# FeatureStrategySchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A uuid for the feature strategy | [optional] [default to null]
**Name** | **string** | The name or type of strategy | [default to null]
**Title** | **string** | A descriptive title for the strategy | [optional] [default to null]
**Disabled** | **bool** | A toggle to disable the strategy. defaults to false. Disabled strategies are not evaluated or returned to the SDKs | [optional] [default to null]
**FeatureName** | **string** | The name or feature the strategy is attached to | [optional] [default to null]
**SortOrder** | **float64** | The order of the strategy in the list | [optional] [default to null]
**Segments** | **[]float64** | A list of segment ids attached to the strategy | [optional] [default to null]
**Constraints** | [**[]ConstraintSchema**](constraintSchema.md) | A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints | [optional] [default to null]
**Parameters** | [***ModelMap**](map.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

