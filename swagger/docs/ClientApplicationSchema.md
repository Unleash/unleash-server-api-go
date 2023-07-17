# ClientApplicationSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | An identifier for the app that uses the sdk, should be static across SDK restarts | [default to null]
**InstanceId** | **string** | A unique identifier identifying the instance of the application running the SDK. Often changes based on execution environment. For instance: two pods in Kubernetes will have two different instanceIds | [optional] [default to null]
**SdkVersion** | **string** | An SDK version identifier. Usually formatted as \&quot;unleash-client-&lt;language&gt;:&lt;version&gt;\&quot; | [optional] [default to null]
**Environment** | **string** | The SDK&#x27;s configured &#x27;environment&#x27; property. Deprecated. This property  does **not** control which Unleash environment the SDK gets toggles for. To control Unleash environments, use the SDKs API key. | [optional] [default to null]
**Interval** | **float64** | How often (in seconds) does the client refresh its toggles | [default to null]
**Started** | [***OneOfclientApplicationSchemaStarted**](OneOfclientApplicationSchemaStarted.md) | Either an RFC-3339 timestamp or a unix timestamp in seconds | [default to null]
**Strategies** | **[]string** | Which strategies the SDKs runtime knows about | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

