# ClientFeaturesQuerySchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | [**[][]string**](array.md) | Features tagged with one of these tags are included | [optional] [default to null]
**Project** | **[]string** | Features that are part of these projects are included in this response. (DEPRECATED) - Handled by API tokens | [optional] [default to null]
**NamePrefix** | **string** | Features are filtered to only include features whose name starts with this prefix | [optional] [default to null]
**Environment** | **string** | Strategies for the feature toggle configured for this environment are included. (DEPRECATED) - Handled by API tokens | [optional] [default to null]
**InlineSegmentConstraints** | **bool** | Set to true if requesting client does not support Unleash-Client-Specification 4.2.2 or newer. Modern SDKs will have this set to false, since they will be able to merge constraints and segments themselves | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

