# CreateStrategySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the strategy type. Must be unique. | 
**Description** | Pointer to **string** | A description of the strategy type. | [optional] 
**Parameters** | [**[]CreateStrategySchemaParametersInner**](CreateStrategySchemaParametersInner.md) | The parameter list lets you pass arguments to your custom activation strategy. These will be made available to your custom strategy implementation. | 

## Methods

### NewCreateStrategySchema

`func NewCreateStrategySchema(name string, parameters []CreateStrategySchemaParametersInner, ) *CreateStrategySchema`

NewCreateStrategySchema instantiates a new CreateStrategySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStrategySchemaWithDefaults

`func NewCreateStrategySchemaWithDefaults() *CreateStrategySchema`

NewCreateStrategySchemaWithDefaults instantiates a new CreateStrategySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateStrategySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStrategySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStrategySchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateStrategySchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStrategySchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStrategySchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStrategySchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *CreateStrategySchema) GetParameters() []CreateStrategySchemaParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateStrategySchema) GetParametersOk() (*[]CreateStrategySchemaParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateStrategySchema) SetParameters(v []CreateStrategySchemaParametersInner)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


