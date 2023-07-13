# FeatureMetricsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The version of this schema | 
**Maturity** | **string** | The maturity level of this API (alpha, beta, stable, deprecated) | 
**Data** | [**[]FeatureEnvironmentMetricsSchema**](FeatureEnvironmentMetricsSchema.md) | Metrics gathered per environment | 

## Methods

### NewFeatureMetricsSchema

`func NewFeatureMetricsSchema(version int32, maturity string, data []FeatureEnvironmentMetricsSchema, ) *FeatureMetricsSchema`

NewFeatureMetricsSchema instantiates a new FeatureMetricsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureMetricsSchemaWithDefaults

`func NewFeatureMetricsSchemaWithDefaults() *FeatureMetricsSchema`

NewFeatureMetricsSchemaWithDefaults instantiates a new FeatureMetricsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FeatureMetricsSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeatureMetricsSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeatureMetricsSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetMaturity

`func (o *FeatureMetricsSchema) GetMaturity() string`

GetMaturity returns the Maturity field if non-nil, zero value otherwise.

### GetMaturityOk

`func (o *FeatureMetricsSchema) GetMaturityOk() (*string, bool)`

GetMaturityOk returns a tuple with the Maturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturity

`func (o *FeatureMetricsSchema) SetMaturity(v string)`

SetMaturity sets Maturity field to given value.


### GetData

`func (o *FeatureMetricsSchema) GetData() []FeatureEnvironmentMetricsSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FeatureMetricsSchema) GetDataOk() (*[]FeatureEnvironmentMetricsSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FeatureMetricsSchema) SetData(v []FeatureEnvironmentMetricsSchema)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


