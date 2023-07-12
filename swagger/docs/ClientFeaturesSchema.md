# ClientFeaturesSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **float64** | A version number for the format used in the response. Most Unleash instances now return version 2, which includes segments as a separate array | [default to null]
**Features** | [**[]ClientFeatureSchema**](clientFeatureSchema.md) | A list of feature toggles with their configuration | [default to null]
**Segments** | [**[]SegmentSchema**](segmentSchema.md) | A list of [Segments](https://docs.getunleash.io/reference/segments) configured for this Unleash instance | [optional] [default to null]
**Query** | [***ClientFeaturesQuerySchema**](clientFeaturesQuerySchema.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

