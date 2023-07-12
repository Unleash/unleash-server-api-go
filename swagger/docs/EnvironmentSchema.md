# EnvironmentSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment | [default to null]
**Type_** | **string** | The [type of environment](https://docs.getunleash.io/reference/environments#environment-types). | [default to null]
**Enabled** | **bool** | &#x60;true&#x60; if the environment is enabled for the project, otherwise &#x60;false&#x60;. | [default to null]
**Protected** | **bool** | &#x60;true&#x60; if the environment is protected, otherwise &#x60;false&#x60;. A *protected* environment can not be deleted. | [default to null]
**SortOrder** | **int32** | Priority of the environment in a list of environments, the lower the value, the higher up in the list the environment will appear. Needs to be an integer | [default to null]
**ProjectCount** | **int32** | The number of projects with this environment | [optional] [default to null]
**ApiTokenCount** | **int32** | The number of API tokens for the project environment | [optional] [default to null]
**EnabledToggleCount** | **int32** | The number of enabled toggles for the project environment | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

