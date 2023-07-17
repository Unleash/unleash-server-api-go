# ChangeRequestSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** |  | [default to null]
**Title** | **string** |  | [optional] [default to null]
**Environment** | **string** |  | [default to null]
**State** | **string** |  | [default to null]
**MinApprovals** | **float64** |  | [default to null]
**Project** | **string** |  | [default to null]
**Features** | [**[]ChangeRequestFeatureSchema**](changeRequestFeatureSchema.md) |  | [default to null]
**Approvals** | [**[]ChangeRequestApprovalSchema**](changeRequestApprovalSchema.md) |  | [optional] [default to null]
**Comments** | [**[]ChangeRequestCommentSchema**](changeRequestCommentSchema.md) |  | [optional] [default to null]
**CreatedBy** | [***ChangeRequestEventSchemaCreatedBy**](changeRequestEventSchema_createdBy.md) |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

