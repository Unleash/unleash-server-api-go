# ExportResultSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | [**[]FeatureSchema**](featureSchema.md) | All the exported features. | [default to null]
**FeatureStrategies** | [**[]FeatureStrategySchema**](featureStrategySchema.md) | All strategy instances that are used by the exported features in the &#x60;features&#x60; list. | [default to null]
**FeatureEnvironments** | [**[]FeatureEnvironmentSchema**](featureEnvironmentSchema.md) | Environment-specific configuration for all the features in the &#x60;features&#x60; list. Includes data such as whether the feature is enabled in the selected export environment, whether there are any variants assigned, etc. | [optional] [default to null]
**ContextFields** | [**[]ContextFieldSchema**](contextFieldSchema.md) | A list of all the context fields that are in use by any of the strategies in the &#x60;featureStrategies&#x60; list. | [optional] [default to null]
**FeatureTags** | [**[]FeatureTagSchema**](featureTagSchema.md) | A list of all the tags that have been applied to any of the features in the &#x60;features&#x60; list. | [optional] [default to null]
**Segments** | [**[]ExportResultSchemaSegments**](exportResultSchema_segments.md) | A list of all the segments that are used by the strategies in the &#x60;featureStrategies&#x60; list. | [optional] [default to null]
**TagTypes** | [**[]TagTypeSchema**](tagTypeSchema.md) | A list of all of the tag types that are used in the &#x60;featureTags&#x60; list. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

