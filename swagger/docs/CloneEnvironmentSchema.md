# CloneEnvironmentSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new cloned environment, this cannot be changed later | [default to null]
**Type_** | **string** | Updates the type of environment (i.e. development or production). | [default to null]
**Projects** | **[]string** | A list of projects that should be included in the cloned environment. | [optional] [default to null]
**ClonePermissions** | **bool** | Copies the RBAC permissions from the source environment if true. Defaults to true | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

