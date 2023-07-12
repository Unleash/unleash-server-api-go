# BulkRegistrationSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectVia** | [**[]BulkRegistrationSchemaConnectVia**](bulkRegistrationSchema_connectVia.md) | A list of applications this app registration has been registered through. If connected directly to Unleash, this is an empty list.   This can be used in later visualizations to tell how many levels of proxy or Edge instances our SDKs have connected through | [optional] [default to null]
**AppName** | **string** | The name of the application that is evaluating toggles | [default to null]
**Environment** | **string** | Which environment the application is running in | [default to null]
**InstanceId** | **string** | A [(somewhat) unique identifier](https://docs.getunleash.io/reference/sdks/node#advanced-usage) for the application | [default to null]
**Interval** | **float64** | How often (in seconds) the application refreshes its features | [optional] [default to null]
**Started** | [***DateSchema**](dateSchema.md) |  | [optional] [default to null]
**Strategies** | **[]string** | Enabled [strategies](https://docs.getunleash.io/reference/activation-strategies) in the application | [optional] [default to null]
**SdkVersion** | **string** | The version the sdk is running. Typically &lt;client&gt;:&lt;version&gt; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

