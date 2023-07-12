# PlaygroundFeatureSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The feature&#x27;s name. | [default to null]
**ProjectId** | **string** | The ID of the project that contains this feature. | [default to null]
**Strategies** | [***PlaygroundFeatureSchemaStrategies**](playgroundFeatureSchema_strategies.md) |  | [default to null]
**IsEnabledInCurrentEnvironment** | **bool** | Whether the feature is active and would be evaluated in the provided environment in a normal SDK context. | [default to null]
**IsEnabled** | **bool** | Whether this feature is enabled or not in the current environment.                           If a feature can&#x27;t be fully evaluated (that is, &#x60;strategies.result&#x60; is &#x60;unknown&#x60;),                           this will be &#x60;false&#x60; to align with how client SDKs treat unresolved feature states. | [default to null]
**Variant** | [***AdvancedPlaygroundEnvironmentFeatureSchemaVariant**](advancedPlaygroundEnvironmentFeatureSchema_variant.md) |  | [default to null]
**Variants** | [**[]VariantSchema**](variantSchema.md) | The feature variants. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

