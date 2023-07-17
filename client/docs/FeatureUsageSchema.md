# FeatureUsageSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The version of this schema | 
**Maturity** | **string** | The maturity level of this API (alpha, beta, stable, deprecated) | 
**FeatureName** | **string** | The name of the feature | 
**LastHourUsage** | [**[]FeatureEnvironmentMetricsSchema**](FeatureEnvironmentMetricsSchema.md) | Last hour statistics. Accumulated per feature per environment. Contains counts for evaluations to true (yes) and to false (no) | 
**SeenApplications** | **[]string** | A list of applications seen using this feature | 

## Methods

### NewFeatureUsageSchema

`func NewFeatureUsageSchema(version int32, maturity string, featureName string, lastHourUsage []FeatureEnvironmentMetricsSchema, seenApplications []string, ) *FeatureUsageSchema`

NewFeatureUsageSchema instantiates a new FeatureUsageSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureUsageSchemaWithDefaults

`func NewFeatureUsageSchemaWithDefaults() *FeatureUsageSchema`

NewFeatureUsageSchemaWithDefaults instantiates a new FeatureUsageSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FeatureUsageSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeatureUsageSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeatureUsageSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetMaturity

`func (o *FeatureUsageSchema) GetMaturity() string`

GetMaturity returns the Maturity field if non-nil, zero value otherwise.

### GetMaturityOk

`func (o *FeatureUsageSchema) GetMaturityOk() (*string, bool)`

GetMaturityOk returns a tuple with the Maturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturity

`func (o *FeatureUsageSchema) SetMaturity(v string)`

SetMaturity sets Maturity field to given value.


### GetFeatureName

`func (o *FeatureUsageSchema) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *FeatureUsageSchema) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *FeatureUsageSchema) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.


### GetLastHourUsage

`func (o *FeatureUsageSchema) GetLastHourUsage() []FeatureEnvironmentMetricsSchema`

GetLastHourUsage returns the LastHourUsage field if non-nil, zero value otherwise.

### GetLastHourUsageOk

`func (o *FeatureUsageSchema) GetLastHourUsageOk() (*[]FeatureEnvironmentMetricsSchema, bool)`

GetLastHourUsageOk returns a tuple with the LastHourUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHourUsage

`func (o *FeatureUsageSchema) SetLastHourUsage(v []FeatureEnvironmentMetricsSchema)`

SetLastHourUsage sets LastHourUsage field to given value.


### GetSeenApplications

`func (o *FeatureUsageSchema) GetSeenApplications() []string`

GetSeenApplications returns the SeenApplications field if non-nil, zero value otherwise.

### GetSeenApplicationsOk

`func (o *FeatureUsageSchema) GetSeenApplicationsOk() (*[]string, bool)`

GetSeenApplicationsOk returns a tuple with the SeenApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenApplications

`func (o *FeatureUsageSchema) SetSeenApplications(v []string)`

SetSeenApplications sets SeenApplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


