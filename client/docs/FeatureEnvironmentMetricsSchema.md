# FeatureEnvironmentMetricsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureName** | Pointer to **string** | The name of the feature | [optional] 
**AppName** | Pointer to **string** | The name of the application the SDK is being used in | [optional] 
**Environment** | **string** | Which environment the SDK is being used in | 
**Timestamp** | [**DateSchema**](DateSchema.md) |  | 
**Yes** | **int32** | How many times the toggle evaluated to true | 
**No** | **int32** | How many times the toggle evaluated to false | 
**Variants** | Pointer to **map[string]int32** | How many times each variant was returned | [optional] 

## Methods

### NewFeatureEnvironmentMetricsSchema

`func NewFeatureEnvironmentMetricsSchema(environment string, timestamp DateSchema, yes int32, no int32, ) *FeatureEnvironmentMetricsSchema`

NewFeatureEnvironmentMetricsSchema instantiates a new FeatureEnvironmentMetricsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureEnvironmentMetricsSchemaWithDefaults

`func NewFeatureEnvironmentMetricsSchemaWithDefaults() *FeatureEnvironmentMetricsSchema`

NewFeatureEnvironmentMetricsSchemaWithDefaults instantiates a new FeatureEnvironmentMetricsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureName

`func (o *FeatureEnvironmentMetricsSchema) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *FeatureEnvironmentMetricsSchema) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *FeatureEnvironmentMetricsSchema) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *FeatureEnvironmentMetricsSchema) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### GetAppName

`func (o *FeatureEnvironmentMetricsSchema) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *FeatureEnvironmentMetricsSchema) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *FeatureEnvironmentMetricsSchema) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *FeatureEnvironmentMetricsSchema) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetEnvironment

`func (o *FeatureEnvironmentMetricsSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FeatureEnvironmentMetricsSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FeatureEnvironmentMetricsSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetTimestamp

`func (o *FeatureEnvironmentMetricsSchema) GetTimestamp() DateSchema`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FeatureEnvironmentMetricsSchema) GetTimestampOk() (*DateSchema, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FeatureEnvironmentMetricsSchema) SetTimestamp(v DateSchema)`

SetTimestamp sets Timestamp field to given value.


### GetYes

`func (o *FeatureEnvironmentMetricsSchema) GetYes() int32`

GetYes returns the Yes field if non-nil, zero value otherwise.

### GetYesOk

`func (o *FeatureEnvironmentMetricsSchema) GetYesOk() (*int32, bool)`

GetYesOk returns a tuple with the Yes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYes

`func (o *FeatureEnvironmentMetricsSchema) SetYes(v int32)`

SetYes sets Yes field to given value.


### GetNo

`func (o *FeatureEnvironmentMetricsSchema) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *FeatureEnvironmentMetricsSchema) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *FeatureEnvironmentMetricsSchema) SetNo(v int32)`

SetNo sets No field to given value.


### GetVariants

`func (o *FeatureEnvironmentMetricsSchema) GetVariants() map[string]int32`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *FeatureEnvironmentMetricsSchema) GetVariantsOk() (*map[string]int32, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *FeatureEnvironmentMetricsSchema) SetVariants(v map[string]int32)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *FeatureEnvironmentMetricsSchema) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


