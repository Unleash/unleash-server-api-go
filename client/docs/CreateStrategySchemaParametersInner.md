# CreateStrategySchemaParametersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the parameter | 
**Type** | **string** | The [type of the parameter](https://docs.getunleash.io/reference/custom-activation-strategies#parameter-types) | 
**Description** | Pointer to **string** | A description of this strategy parameter. Use this to indicate to the users what the parameter does. | [optional] 
**Required** | Pointer to **bool** | Whether this parameter must be configured when using the strategy. Defaults to &#x60;false&#x60; | [optional] 

## Methods

### NewCreateStrategySchemaParametersInner

`func NewCreateStrategySchemaParametersInner(name string, type_ string, ) *CreateStrategySchemaParametersInner`

NewCreateStrategySchemaParametersInner instantiates a new CreateStrategySchemaParametersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStrategySchemaParametersInnerWithDefaults

`func NewCreateStrategySchemaParametersInnerWithDefaults() *CreateStrategySchemaParametersInner`

NewCreateStrategySchemaParametersInnerWithDefaults instantiates a new CreateStrategySchemaParametersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateStrategySchemaParametersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStrategySchemaParametersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStrategySchemaParametersInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateStrategySchemaParametersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStrategySchemaParametersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStrategySchemaParametersInner) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *CreateStrategySchemaParametersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStrategySchemaParametersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStrategySchemaParametersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStrategySchemaParametersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *CreateStrategySchemaParametersInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CreateStrategySchemaParametersInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CreateStrategySchemaParametersInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CreateStrategySchemaParametersInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


