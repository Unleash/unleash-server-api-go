# StrategySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | An optional title for the strategy | [optional] 
**Name** | **string** | The name or type of the strategy | 
**DisplayName** | **NullableString** | A human friendly name for the strategy | 
**Description** | **string** | A short description for the strategy | 
**Editable** | **bool** | Determines whether the strategy allows for editing | 
**Deprecated** | **bool** |  | 
**Parameters** | [**[]StrategySchemaParametersInner**](StrategySchemaParametersInner.md) | A list of relevant parameters for each strategy | 

## Methods

### NewStrategySchema

`func NewStrategySchema(name string, displayName NullableString, description string, editable bool, deprecated bool, parameters []StrategySchemaParametersInner, ) *StrategySchema`

NewStrategySchema instantiates a new StrategySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategySchemaWithDefaults

`func NewStrategySchemaWithDefaults() *StrategySchema`

NewStrategySchemaWithDefaults instantiates a new StrategySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *StrategySchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StrategySchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StrategySchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StrategySchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *StrategySchema) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *StrategySchema) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetName

`func (o *StrategySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StrategySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StrategySchema) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *StrategySchema) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *StrategySchema) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *StrategySchema) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *StrategySchema) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *StrategySchema) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *StrategySchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StrategySchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StrategySchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEditable

`func (o *StrategySchema) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *StrategySchema) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *StrategySchema) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetDeprecated

`func (o *StrategySchema) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *StrategySchema) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *StrategySchema) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetParameters

`func (o *StrategySchema) GetParameters() []StrategySchemaParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *StrategySchema) GetParametersOk() (*[]StrategySchemaParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *StrategySchema) SetParameters(v []StrategySchemaParametersInner)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


