# CreateContextFieldSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the context field | [optional] 
**Stickiness** | Pointer to **bool** | &#x60;true&#x60; if this field should be available for use with [custom stickiness](https://docs.getunleash.io/reference/stickiness#custom-stickiness), otherwise &#x60;false&#x60; | [optional] 
**SortOrder** | Pointer to **int32** | How this context field should be sorted if no other sort order is selected | [optional] 
**LegalValues** | Pointer to [**[]LegalValueSchema**](LegalValueSchema.md) | A list of allowed values for this context field | [optional] 
**Name** | **string** | The name of the context field. | 

## Methods

### NewCreateContextFieldSchema

`func NewCreateContextFieldSchema(name string, ) *CreateContextFieldSchema`

NewCreateContextFieldSchema instantiates a new CreateContextFieldSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContextFieldSchemaWithDefaults

`func NewCreateContextFieldSchemaWithDefaults() *CreateContextFieldSchema`

NewCreateContextFieldSchemaWithDefaults instantiates a new CreateContextFieldSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateContextFieldSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateContextFieldSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateContextFieldSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateContextFieldSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStickiness

`func (o *CreateContextFieldSchema) GetStickiness() bool`

GetStickiness returns the Stickiness field if non-nil, zero value otherwise.

### GetStickinessOk

`func (o *CreateContextFieldSchema) GetStickinessOk() (*bool, bool)`

GetStickinessOk returns a tuple with the Stickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickiness

`func (o *CreateContextFieldSchema) SetStickiness(v bool)`

SetStickiness sets Stickiness field to given value.

### HasStickiness

`func (o *CreateContextFieldSchema) HasStickiness() bool`

HasStickiness returns a boolean if a field has been set.

### GetSortOrder

`func (o *CreateContextFieldSchema) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *CreateContextFieldSchema) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *CreateContextFieldSchema) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *CreateContextFieldSchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLegalValues

`func (o *CreateContextFieldSchema) GetLegalValues() []LegalValueSchema`

GetLegalValues returns the LegalValues field if non-nil, zero value otherwise.

### GetLegalValuesOk

`func (o *CreateContextFieldSchema) GetLegalValuesOk() (*[]LegalValueSchema, bool)`

GetLegalValuesOk returns a tuple with the LegalValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalValues

`func (o *CreateContextFieldSchema) SetLegalValues(v []LegalValueSchema)`

SetLegalValues sets LegalValues field to given value.

### HasLegalValues

`func (o *CreateContextFieldSchema) HasLegalValues() bool`

HasLegalValues returns a boolean if a field has been set.

### GetName

`func (o *CreateContextFieldSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContextFieldSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContextFieldSchema) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


