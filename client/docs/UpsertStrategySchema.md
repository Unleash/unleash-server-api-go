# UpsertStrategySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to [**[]UpsertStrategySchemaParametersInner**](UpsertStrategySchemaParametersInner.md) |  | [optional] 

## Methods

### NewUpsertStrategySchema

`func NewUpsertStrategySchema(name string, ) *UpsertStrategySchema`

NewUpsertStrategySchema instantiates a new UpsertStrategySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertStrategySchemaWithDefaults

`func NewUpsertStrategySchemaWithDefaults() *UpsertStrategySchema`

NewUpsertStrategySchemaWithDefaults instantiates a new UpsertStrategySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpsertStrategySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertStrategySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertStrategySchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpsertStrategySchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpsertStrategySchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpsertStrategySchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpsertStrategySchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditable

`func (o *UpsertStrategySchema) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *UpsertStrategySchema) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *UpsertStrategySchema) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *UpsertStrategySchema) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetParameters

`func (o *UpsertStrategySchema) GetParameters() []UpsertStrategySchemaParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpsertStrategySchema) GetParametersOk() (*[]UpsertStrategySchemaParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpsertStrategySchema) SetParameters(v []UpsertStrategySchemaParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpsertStrategySchema) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


