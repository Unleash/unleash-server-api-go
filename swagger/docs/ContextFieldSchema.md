# ContextFieldSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the context field | [default to null]
**Description** | **string** | The description of the context field. | [optional] [default to null]
**Stickiness** | **bool** | Does this context field support being used for [stickiness](https://docs.getunleash.io/reference/stickiness) calculations | [optional] [default to null]
**SortOrder** | **int32** | Used when sorting a list of context fields. Is also used as a tiebreaker if a list of context fields is sorted alphabetically. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When this context field was created | [optional] [default to null]
**UsedInFeatures** | **int32** | Number of projects where this context field is used in | [optional] [default to null]
**UsedInProjects** | **int32** | Number of projects where this context field is used in | [optional] [default to null]
**LegalValues** | [**[]LegalValueSchema**](legalValueSchema.md) | Allowed values for this context field schema. Can be used to narrow down accepted input | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

