# FeatureEnvironmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment | 
**FeatureName** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of the environment | [optional] 
**Enabled** | **bool** | &#x60;true&#x60; if the feature is enabled for the environment, otherwise &#x60;false&#x60;. | 
**SortOrder** | Pointer to **float32** | The sort order of the feature environment in the feature environments list | [optional] 
**VariantCount** | Pointer to **float32** |  | [optional] 
**Strategies** | Pointer to [**[]FeatureStrategySchema**](FeatureStrategySchema.md) | A list of activation strategies for the feature environment | [optional] 
**Variants** | Pointer to [**[]VariantSchema**](VariantSchema.md) | A list of variants for the feature environment | [optional] 

## Methods

### NewFeatureEnvironmentSchema

`func NewFeatureEnvironmentSchema(name string, enabled bool, ) *FeatureEnvironmentSchema`

NewFeatureEnvironmentSchema instantiates a new FeatureEnvironmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureEnvironmentSchemaWithDefaults

`func NewFeatureEnvironmentSchemaWithDefaults() *FeatureEnvironmentSchema`

NewFeatureEnvironmentSchemaWithDefaults instantiates a new FeatureEnvironmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureEnvironmentSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureEnvironmentSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureEnvironmentSchema) SetName(v string)`

SetName sets Name field to given value.


### GetFeatureName

`func (o *FeatureEnvironmentSchema) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *FeatureEnvironmentSchema) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *FeatureEnvironmentSchema) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *FeatureEnvironmentSchema) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### GetEnvironment

`func (o *FeatureEnvironmentSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FeatureEnvironmentSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FeatureEnvironmentSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FeatureEnvironmentSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetType

`func (o *FeatureEnvironmentSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeatureEnvironmentSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeatureEnvironmentSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeatureEnvironmentSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *FeatureEnvironmentSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeatureEnvironmentSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeatureEnvironmentSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSortOrder

`func (o *FeatureEnvironmentSchema) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FeatureEnvironmentSchema) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FeatureEnvironmentSchema) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *FeatureEnvironmentSchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetVariantCount

`func (o *FeatureEnvironmentSchema) GetVariantCount() float32`

GetVariantCount returns the VariantCount field if non-nil, zero value otherwise.

### GetVariantCountOk

`func (o *FeatureEnvironmentSchema) GetVariantCountOk() (*float32, bool)`

GetVariantCountOk returns a tuple with the VariantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantCount

`func (o *FeatureEnvironmentSchema) SetVariantCount(v float32)`

SetVariantCount sets VariantCount field to given value.

### HasVariantCount

`func (o *FeatureEnvironmentSchema) HasVariantCount() bool`

HasVariantCount returns a boolean if a field has been set.

### GetStrategies

`func (o *FeatureEnvironmentSchema) GetStrategies() []FeatureStrategySchema`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *FeatureEnvironmentSchema) GetStrategiesOk() (*[]FeatureStrategySchema, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *FeatureEnvironmentSchema) SetStrategies(v []FeatureStrategySchema)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *FeatureEnvironmentSchema) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.

### GetVariants

`func (o *FeatureEnvironmentSchema) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *FeatureEnvironmentSchema) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *FeatureEnvironmentSchema) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *FeatureEnvironmentSchema) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


