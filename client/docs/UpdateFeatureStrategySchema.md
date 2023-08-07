# UpdateFeatureStrategySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the strategy type | [optional] 
**SortOrder** | Pointer to **float32** | The order of the strategy in the list in feature environment configuration | [optional] 
**Constraints** | Pointer to [**[]ConstraintSchema**](ConstraintSchema.md) | A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints | [optional] 
**Title** | Pointer to **NullableString** | A descriptive title for the strategy | [optional] 
**Disabled** | Pointer to **NullableBool** | A toggle to disable the strategy. defaults to true. Disabled strategies are not evaluated or returned to the SDKs | [optional] 
**Parameters** | Pointer to **map[string]string** | A list of parameters for a strategy | [optional] 

## Methods

### NewUpdateFeatureStrategySchema

`func NewUpdateFeatureStrategySchema() *UpdateFeatureStrategySchema`

NewUpdateFeatureStrategySchema instantiates a new UpdateFeatureStrategySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFeatureStrategySchemaWithDefaults

`func NewUpdateFeatureStrategySchemaWithDefaults() *UpdateFeatureStrategySchema`

NewUpdateFeatureStrategySchemaWithDefaults instantiates a new UpdateFeatureStrategySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFeatureStrategySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFeatureStrategySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFeatureStrategySchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFeatureStrategySchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSortOrder

`func (o *UpdateFeatureStrategySchema) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpdateFeatureStrategySchema) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpdateFeatureStrategySchema) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UpdateFeatureStrategySchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetConstraints

`func (o *UpdateFeatureStrategySchema) GetConstraints() []ConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *UpdateFeatureStrategySchema) GetConstraintsOk() (*[]ConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *UpdateFeatureStrategySchema) SetConstraints(v []ConstraintSchema)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *UpdateFeatureStrategySchema) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateFeatureStrategySchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateFeatureStrategySchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateFeatureStrategySchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateFeatureStrategySchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpdateFeatureStrategySchema) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateFeatureStrategySchema) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDisabled

`func (o *UpdateFeatureStrategySchema) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UpdateFeatureStrategySchema) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UpdateFeatureStrategySchema) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UpdateFeatureStrategySchema) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *UpdateFeatureStrategySchema) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *UpdateFeatureStrategySchema) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetParameters

`func (o *UpdateFeatureStrategySchema) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateFeatureStrategySchema) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateFeatureStrategySchema) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateFeatureStrategySchema) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


