# FeatureUsageSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The version of this schema | [default to null]
**Maturity** | **string** | The maturity level of this API (alpha, beta, stable, deprecated) | [default to null]
**FeatureName** | **string** | The name of the feature | [default to null]
**LastHourUsage** | [**[]FeatureEnvironmentMetricsSchema**](featureEnvironmentMetricsSchema.md) | Last hour statistics. Accumulated per feature per environment. Contains counts for evaluations to true (yes) and to false (no) | [default to null]
**SeenApplications** | **[]string** | A list of applications seen using this feature | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

