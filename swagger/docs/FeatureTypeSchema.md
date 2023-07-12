# FeatureTypeSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of this feature toggle type. | [default to null]
**Name** | **string** | The display name of this feature toggle type. | [default to null]
**Description** | **string** | A description of what this feature toggle type is intended to be used for. | [default to null]
**LifetimeDays** | **int32** | How many days it takes before a feature toggle of this typed is flagged as [potentially stale](https://docs.getunleash.io/reference/technical-debt#stale-and-potentially-stale-toggles) by Unleash. If this value is &#x60;null&#x60;, Unleash will never mark it as potentially stale. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

