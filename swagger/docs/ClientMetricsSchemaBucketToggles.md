# ClientMetricsSchemaBucketToggles

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Yes** | **float64** | How many times the toggle evaluated to true | [optional] [default to null]
**No** | **int32** | How many times the toggle evaluated to false | [optional] [default to null]
**Variants** | **map[string]int32** | An object describing how many times each variant was returned. Variant names are used as properties, and the number of times they were exposed is the corresponding value (i.e. &#x60;{ [variantName]: number }&#x60;). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

