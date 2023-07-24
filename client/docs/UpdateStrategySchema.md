# UpdateStrategySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the strategy type. | [optional] 
**Parameters** | [**[]CreateStrategySchemaParametersInner**](CreateStrategySchemaParametersInner.md) | The parameter list lets you pass arguments to your custom activation strategy. These will be made available to your custom strategy implementation. | 

## Methods

### NewUpdateStrategySchema

`func NewUpdateStrategySchema(parameters []CreateStrategySchemaParametersInner, ) *UpdateStrategySchema`

NewUpdateStrategySchema instantiates a new UpdateStrategySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStrategySchemaWithDefaults

`func NewUpdateStrategySchemaWithDefaults() *UpdateStrategySchema`

NewUpdateStrategySchemaWithDefaults instantiates a new UpdateStrategySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateStrategySchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStrategySchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStrategySchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStrategySchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateStrategySchema) GetParameters() []CreateStrategySchemaParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateStrategySchema) GetParametersOk() (*[]CreateStrategySchemaParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateStrategySchema) SetParameters(v []CreateStrategySchemaParametersInner)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


