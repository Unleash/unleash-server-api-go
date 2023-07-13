# ClientMetricsEnvSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureName** | **string** | Name of the feature checked by the SDK | 
**AppName** | **string** | The name of the application the SDK is being used in | 
**Environment** | **string** | Which environment the SDK is being used in | 
**Timestamp** | Pointer to [**DateSchema**](DateSchema.md) |  | [optional] 
**Yes** | Pointer to **int32** | How many times the toggle evaluated to true | [optional] 
**No** | Pointer to **int32** | How many times the toggle evaluated to false | [optional] 
**Variants** | Pointer to **map[string]int32** | How many times each variant was returned | [optional] 

## Methods

### NewClientMetricsEnvSchema

`func NewClientMetricsEnvSchema(featureName string, appName string, environment string, ) *ClientMetricsEnvSchema`

NewClientMetricsEnvSchema instantiates a new ClientMetricsEnvSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMetricsEnvSchemaWithDefaults

`func NewClientMetricsEnvSchemaWithDefaults() *ClientMetricsEnvSchema`

NewClientMetricsEnvSchemaWithDefaults instantiates a new ClientMetricsEnvSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureName

`func (o *ClientMetricsEnvSchema) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *ClientMetricsEnvSchema) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *ClientMetricsEnvSchema) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.


### GetAppName

`func (o *ClientMetricsEnvSchema) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ClientMetricsEnvSchema) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ClientMetricsEnvSchema) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetEnvironment

`func (o *ClientMetricsEnvSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ClientMetricsEnvSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ClientMetricsEnvSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetTimestamp

`func (o *ClientMetricsEnvSchema) GetTimestamp() DateSchema`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ClientMetricsEnvSchema) GetTimestampOk() (*DateSchema, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ClientMetricsEnvSchema) SetTimestamp(v DateSchema)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ClientMetricsEnvSchema) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetYes

`func (o *ClientMetricsEnvSchema) GetYes() int32`

GetYes returns the Yes field if non-nil, zero value otherwise.

### GetYesOk

`func (o *ClientMetricsEnvSchema) GetYesOk() (*int32, bool)`

GetYesOk returns a tuple with the Yes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYes

`func (o *ClientMetricsEnvSchema) SetYes(v int32)`

SetYes sets Yes field to given value.

### HasYes

`func (o *ClientMetricsEnvSchema) HasYes() bool`

HasYes returns a boolean if a field has been set.

### GetNo

`func (o *ClientMetricsEnvSchema) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *ClientMetricsEnvSchema) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *ClientMetricsEnvSchema) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *ClientMetricsEnvSchema) HasNo() bool`

HasNo returns a boolean if a field has been set.

### GetVariants

`func (o *ClientMetricsEnvSchema) GetVariants() map[string]int32`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ClientMetricsEnvSchema) GetVariantsOk() (*map[string]int32, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ClientMetricsEnvSchema) SetVariants(v map[string]int32)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ClientMetricsEnvSchema) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


