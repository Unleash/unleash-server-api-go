# PlaygroundFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The feature&#39;s name. | 
**ProjectId** | **string** | The ID of the project that contains this feature. | 
**Strategies** | [**PlaygroundFeatureSchemaStrategies**](PlaygroundFeatureSchemaStrategies.md) |  | 
**IsEnabledInCurrentEnvironment** | **bool** | Whether the feature is active and would be evaluated in the provided environment in a normal SDK context. | 
**IsEnabled** | **bool** | Whether this feature is enabled or not in the current environment.                           If a feature can&#39;t be fully evaluated (that is, &#x60;strategies.result&#x60; is &#x60;unknown&#x60;),                           this will be &#x60;false&#x60; to align with how client SDKs treat unresolved feature states. | 
**Variant** | [**NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant**](AdvancedPlaygroundEnvironmentFeatureSchemaVariant.md) |  | 
**Variants** | [**[]VariantSchema**](VariantSchema.md) | The feature variants. | 

## Methods

### NewPlaygroundFeatureSchema

`func NewPlaygroundFeatureSchema(name string, projectId string, strategies PlaygroundFeatureSchemaStrategies, isEnabledInCurrentEnvironment bool, isEnabled bool, variant NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant, variants []VariantSchema, ) *PlaygroundFeatureSchema`

NewPlaygroundFeatureSchema instantiates a new PlaygroundFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundFeatureSchemaWithDefaults

`func NewPlaygroundFeatureSchemaWithDefaults() *PlaygroundFeatureSchema`

NewPlaygroundFeatureSchemaWithDefaults instantiates a new PlaygroundFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlaygroundFeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlaygroundFeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlaygroundFeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *PlaygroundFeatureSchema) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PlaygroundFeatureSchema) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PlaygroundFeatureSchema) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetStrategies

`func (o *PlaygroundFeatureSchema) GetStrategies() PlaygroundFeatureSchemaStrategies`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *PlaygroundFeatureSchema) GetStrategiesOk() (*PlaygroundFeatureSchemaStrategies, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *PlaygroundFeatureSchema) SetStrategies(v PlaygroundFeatureSchemaStrategies)`

SetStrategies sets Strategies field to given value.


### GetIsEnabledInCurrentEnvironment

`func (o *PlaygroundFeatureSchema) GetIsEnabledInCurrentEnvironment() bool`

GetIsEnabledInCurrentEnvironment returns the IsEnabledInCurrentEnvironment field if non-nil, zero value otherwise.

### GetIsEnabledInCurrentEnvironmentOk

`func (o *PlaygroundFeatureSchema) GetIsEnabledInCurrentEnvironmentOk() (*bool, bool)`

GetIsEnabledInCurrentEnvironmentOk returns a tuple with the IsEnabledInCurrentEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledInCurrentEnvironment

`func (o *PlaygroundFeatureSchema) SetIsEnabledInCurrentEnvironment(v bool)`

SetIsEnabledInCurrentEnvironment sets IsEnabledInCurrentEnvironment field to given value.


### GetIsEnabled

`func (o *PlaygroundFeatureSchema) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PlaygroundFeatureSchema) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PlaygroundFeatureSchema) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetVariant

`func (o *PlaygroundFeatureSchema) GetVariant() AdvancedPlaygroundEnvironmentFeatureSchemaVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *PlaygroundFeatureSchema) GetVariantOk() (*AdvancedPlaygroundEnvironmentFeatureSchemaVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *PlaygroundFeatureSchema) SetVariant(v AdvancedPlaygroundEnvironmentFeatureSchemaVariant)`

SetVariant sets Variant field to given value.


### SetVariantNil

`func (o *PlaygroundFeatureSchema) SetVariantNil(b bool)`

 SetVariantNil sets the value for Variant to be an explicit nil

### UnsetVariant
`func (o *PlaygroundFeatureSchema) UnsetVariant()`

UnsetVariant ensures that no value is present for Variant, not even an explicit nil
### GetVariants

`func (o *PlaygroundFeatureSchema) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *PlaygroundFeatureSchema) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *PlaygroundFeatureSchema) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


