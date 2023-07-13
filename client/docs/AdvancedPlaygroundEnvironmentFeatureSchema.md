# AdvancedPlaygroundEnvironmentFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The feature&#39;s name. | 
**Environment** | **string** | The feature&#39;s environment. | 
**Context** | [**SdkContextSchema**](SdkContextSchema.md) |  | 
**ProjectId** | **string** | The ID of the project that contains this feature. | 
**Strategies** | [**AdvancedPlaygroundEnvironmentFeatureSchemaStrategies**](AdvancedPlaygroundEnvironmentFeatureSchemaStrategies.md) |  | 
**IsEnabledInCurrentEnvironment** | **bool** | Whether the feature is active and would be evaluated in the provided environment in a normal SDK context. | 
**IsEnabled** | **bool** | Whether this feature is enabled or not in the current environment.                           If a feature can&#39;t be fully evaluated (that is, &#x60;strategies.result&#x60; is &#x60;unknown&#x60;),                           this will be &#x60;false&#x60; to align with how client SDKs treat unresolved feature states. | 
**Variant** | [**NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant**](AdvancedPlaygroundEnvironmentFeatureSchemaVariant.md) |  | 
**Variants** | [**[]VariantSchema**](VariantSchema.md) | The feature variants. | 

## Methods

### NewAdvancedPlaygroundEnvironmentFeatureSchema

`func NewAdvancedPlaygroundEnvironmentFeatureSchema(name string, environment string, context SdkContextSchema, projectId string, strategies AdvancedPlaygroundEnvironmentFeatureSchemaStrategies, isEnabledInCurrentEnvironment bool, isEnabled bool, variant NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant, variants []VariantSchema, ) *AdvancedPlaygroundEnvironmentFeatureSchema`

NewAdvancedPlaygroundEnvironmentFeatureSchema instantiates a new AdvancedPlaygroundEnvironmentFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedPlaygroundEnvironmentFeatureSchemaWithDefaults

`func NewAdvancedPlaygroundEnvironmentFeatureSchemaWithDefaults() *AdvancedPlaygroundEnvironmentFeatureSchema`

NewAdvancedPlaygroundEnvironmentFeatureSchemaWithDefaults instantiates a new AdvancedPlaygroundEnvironmentFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironment

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetContext

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetContext() SdkContextSchema`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetContextOk() (*SdkContextSchema, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetContext(v SdkContextSchema)`

SetContext sets Context field to given value.


### GetProjectId

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetStrategies

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetStrategies() AdvancedPlaygroundEnvironmentFeatureSchemaStrategies`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetStrategiesOk() (*AdvancedPlaygroundEnvironmentFeatureSchemaStrategies, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetStrategies(v AdvancedPlaygroundEnvironmentFeatureSchemaStrategies)`

SetStrategies sets Strategies field to given value.


### GetIsEnabledInCurrentEnvironment

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetIsEnabledInCurrentEnvironment() bool`

GetIsEnabledInCurrentEnvironment returns the IsEnabledInCurrentEnvironment field if non-nil, zero value otherwise.

### GetIsEnabledInCurrentEnvironmentOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetIsEnabledInCurrentEnvironmentOk() (*bool, bool)`

GetIsEnabledInCurrentEnvironmentOk returns a tuple with the IsEnabledInCurrentEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledInCurrentEnvironment

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetIsEnabledInCurrentEnvironment(v bool)`

SetIsEnabledInCurrentEnvironment sets IsEnabledInCurrentEnvironment field to given value.


### GetIsEnabled

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetVariant

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetVariant() AdvancedPlaygroundEnvironmentFeatureSchemaVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetVariantOk() (*AdvancedPlaygroundEnvironmentFeatureSchemaVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetVariant(v AdvancedPlaygroundEnvironmentFeatureSchemaVariant)`

SetVariant sets Variant field to given value.


### SetVariantNil

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetVariantNil(b bool)`

 SetVariantNil sets the value for Variant to be an explicit nil

### UnsetVariant
`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) UnsetVariant()`

UnsetVariant ensures that no value is present for Variant, not even an explicit nil
### GetVariants

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *AdvancedPlaygroundEnvironmentFeatureSchema) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


