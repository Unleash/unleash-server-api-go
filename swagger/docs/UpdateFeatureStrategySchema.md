# UpdateFeatureStrategySchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the strategy type | [optional] [default to null]
**SortOrder** | **float64** | The order of the strategy in the list in feature environment configuration | [optional] [default to null]
**Constraints** | [**[]ConstraintSchema**](constraintSchema.md) | A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints | [optional] [default to null]
**Title** | **string** | A descriptive title for the strategy | [optional] [default to null]
**Disabled** | **bool** | A toggle to disable the strategy. defaults to true. Disabled strategies are not evaluated or returned to the SDKs | [optional] [default to null]
**Parameters** | [***ModelMap**](map.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

