# ProxyClientSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | Name of the application using Unleash | [default to null]
**InstanceId** | **string** | Instance id for this application (typically hostname, podId or similar) | [optional] [default to null]
**SdkVersion** | **string** | Optional field that describes the sdk version (name:version) | [optional] [default to null]
**Environment** | **string** | deprecated | [optional] [default to null]
**Interval** | **float64** | At which interval, in milliseconds, will this client be expected to send metrics | [default to null]
**Started** | [***OneOfproxyClientSchemaStarted**](OneOfproxyClientSchemaStarted.md) | When this client started. Should be reported as ISO8601 time. | [default to null]
**Strategies** | **[]string** | List of strategies implemented by this application | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

