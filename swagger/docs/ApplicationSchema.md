# ApplicationSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | Name of the application | [default to null]
**SdkVersion** | **string** | Which SDK and version the application reporting uses. Typically represented as &#x60;&lt;identifier&gt;:&lt;version&gt;&#x60; | [optional] [default to null]
**Strategies** | **[]string** | Which [strategies](https://docs.getunleash.io/topics/the-anatomy-of-unleash#activation-strategies) the application has loaded. Useful when trying to figure out if your [custom strategy](https://docs.getunleash.io/reference/custom-activation-strategies) has been loaded in the SDK | [optional] [default to null]
**Description** | **string** | Extra information added about the application reporting the metrics. Only present if added via the Unleash Admin interface | [optional] [default to null]
**Url** | **string** | A link to reference the application reporting the metrics. Could for instance be a GitHub link to the repository of the application | [optional] [default to null]
**Color** | **string** | The CSS color that is used to color the application&#x27;s entry in the application list | [optional] [default to null]
**Icon** | **string** | An URL to an icon file to be used for the applications&#x27;s entry in the application list | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

