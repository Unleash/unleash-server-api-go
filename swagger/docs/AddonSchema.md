# AddonSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The addon&#x27;s unique identifier. | [default to null]
**Provider** | **string** | The addon provider, such as \&quot;webhook\&quot; or \&quot;slack\&quot;. | [default to null]
**Description** | **string** | A description of the addon. &#x60;null&#x60; if no description exists. | [default to null]
**Enabled** | **bool** | Whether the addon is enabled or not. | [default to null]
**Parameters** | [**map[string]Object**](.md) | Parameters for the addon provider. This object has different required and optional properties depending on the provider you choose. | [default to null]
**Events** | **[]string** | The event types that trigger this specific addon. | [default to null]
**Projects** | **[]string** | The projects that this addon listens to events from. An empty list means it listens to events from **all** projects. | [optional] [default to null]
**Environments** | **[]string** | The list of environments that this addon listens to events from. An empty list means it listens to events from **all** environments. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

