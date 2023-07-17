# CreateUserSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The user&#x27;s username. Must be provided if email is not provided. | [optional] [default to null]
**Email** | **string** | The user&#x27;s email address. Must be provided if username is not provided. | [optional] [default to null]
**Name** | **string** | The user&#x27;s name (not the user&#x27;s username). | [optional] [default to null]
**Password** | **string** | Password for the user | [optional] [default to null]
**RootRole** | [***OneOfcreateUserSchemaRootRole**](OneOfcreateUserSchemaRootRole.md) | The role to assign to the user. Can be either the role&#x27;s ID or its unique name. | [default to null]
**SendEmail** | **bool** | Whether to send a welcome email with a login link to the user or not. Defaults to &#x60;true&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

