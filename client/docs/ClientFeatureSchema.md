# ClientFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of a feature toggle. Is validated to be URL safe on creation | 
**Type** | Pointer to **string** | What kind of feature flag is this. Refer to the documentation on [feature toggle types](https://docs.getunleash.io/reference/feature-toggle-types) for more information | [optional] 
**Description** | Pointer to **NullableString** | A description of the toggle | [optional] 
**Enabled** | **bool** | Whether the feature flag is enabled for the current API key or not. This is ANDed with the evaluation results of the strategies list, so if this is false, the evaluation result will always be false | 
**Stale** | Pointer to **bool** | If this is true Unleash believes this feature toggle has been active longer than Unleash expects a toggle of this type to be active | [optional] 
**ImpressionData** | Pointer to **NullableBool** | Set to true if SDKs should trigger [impression events](https://docs.getunleash.io/reference/impression-data) when this toggle is evaluated | [optional] 
**Project** | Pointer to **string** | Which project this feature toggle belongs to | [optional] 
**Strategies** | Pointer to [**[]FeatureStrategySchema**](FeatureStrategySchema.md) | Evaluation strategies for this toggle. Each entry in this list will be evaluated and ORed together | [optional] 
**Variants** | Pointer to [**[]VariantSchema**](VariantSchema.md) | [Variants](https://docs.getunleash.io/reference/feature-toggle-variants#what-are-variants) configured for this toggle | [optional] 

## Methods

### NewClientFeatureSchema

`func NewClientFeatureSchema(name string, enabled bool, ) *ClientFeatureSchema`

NewClientFeatureSchema instantiates a new ClientFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientFeatureSchemaWithDefaults

`func NewClientFeatureSchemaWithDefaults() *ClientFeatureSchema`

NewClientFeatureSchemaWithDefaults instantiates a new ClientFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientFeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientFeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientFeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ClientFeatureSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientFeatureSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientFeatureSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientFeatureSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ClientFeatureSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientFeatureSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientFeatureSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientFeatureSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClientFeatureSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClientFeatureSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *ClientFeatureSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClientFeatureSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClientFeatureSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetStale

`func (o *ClientFeatureSchema) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ClientFeatureSchema) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ClientFeatureSchema) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ClientFeatureSchema) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetImpressionData

`func (o *ClientFeatureSchema) GetImpressionData() bool`

GetImpressionData returns the ImpressionData field if non-nil, zero value otherwise.

### GetImpressionDataOk

`func (o *ClientFeatureSchema) GetImpressionDataOk() (*bool, bool)`

GetImpressionDataOk returns a tuple with the ImpressionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressionData

`func (o *ClientFeatureSchema) SetImpressionData(v bool)`

SetImpressionData sets ImpressionData field to given value.

### HasImpressionData

`func (o *ClientFeatureSchema) HasImpressionData() bool`

HasImpressionData returns a boolean if a field has been set.

### SetImpressionDataNil

`func (o *ClientFeatureSchema) SetImpressionDataNil(b bool)`

 SetImpressionDataNil sets the value for ImpressionData to be an explicit nil

### UnsetImpressionData
`func (o *ClientFeatureSchema) UnsetImpressionData()`

UnsetImpressionData ensures that no value is present for ImpressionData, not even an explicit nil
### GetProject

`func (o *ClientFeatureSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ClientFeatureSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ClientFeatureSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ClientFeatureSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetStrategies

`func (o *ClientFeatureSchema) GetStrategies() []FeatureStrategySchema`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *ClientFeatureSchema) GetStrategiesOk() (*[]FeatureStrategySchema, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *ClientFeatureSchema) SetStrategies(v []FeatureStrategySchema)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *ClientFeatureSchema) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.

### GetVariants

`func (o *ClientFeatureSchema) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ClientFeatureSchema) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ClientFeatureSchema) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ClientFeatureSchema) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### SetVariantsNil

`func (o *ClientFeatureSchema) SetVariantsNil(b bool)`

 SetVariantsNil sets the value for Variants to be an explicit nil

### UnsetVariants
`func (o *ClientFeatureSchema) UnsetVariants()`

UnsetVariants ensures that no value is present for Variants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


