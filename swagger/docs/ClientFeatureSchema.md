# ClientFeatureSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of a feature toggle. Is validated to be URL safe on creation | [default to null]
**Type_** | **string** | What kind of feature flag is this. Refer to the documentation on [feature toggle types](https://docs.getunleash.io/reference/feature-toggle-types) for more information | [optional] [default to null]
**Description** | **string** | A description of the toggle | [optional] [default to null]
**Enabled** | **bool** | Whether the feature flag is enabled for the current API key or not. This is ANDed with the evaluation results of the strategies list, so if this is false, the evaluation result will always be false | [default to null]
**Stale** | **bool** | If this is true Unleash believes this feature toggle has been active longer than Unleash expects a toggle of this type to be active | [optional] [default to null]
**ImpressionData** | **bool** | Set to true if SDKs should trigger [impression events](https://docs.getunleash.io/reference/impression-data) when this toggle is evaluated | [optional] [default to null]
**Project** | **string** | Which project this feature toggle belongs to | [optional] [default to null]
**Strategies** | [**[]FeatureStrategySchema**](featureStrategySchema.md) | Evaluation strategies for this toggle. Each entry in this list will be evaluated and ORed together | [optional] [default to null]
**Variants** | [**[]VariantSchema**](variantSchema.md) | [Variants](https://docs.getunleash.io/reference/feature-toggle-variants#what-are-variants) configured for this toggle | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

