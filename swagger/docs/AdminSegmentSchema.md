# AdminSegmentSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of this segment | [default to null]
**Name** | **string** | The name of this segment | [default to null]
**Description** | **string** | The description for this segment | [optional] [default to null]
**Constraints** | [**[]ConstraintSchema**](constraintSchema.md) | The list of constraints that are used in this segment | [default to null]
**UsedInFeatures** | **int32** | The number of projects that use this segment | [optional] [default to null]
**UsedInProjects** | **int32** | The number of projects that use this segment | [optional] [default to null]
**Project** | **string** |  | [optional] [default to null]
**CreatedBy** | **string** | The creator&#x27;s email or username | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the segment was created | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

