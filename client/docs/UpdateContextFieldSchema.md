# UpdateContextFieldSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the context field | [optional] 
**Stickiness** | Pointer to **bool** | &#x60;true&#x60; if this field should be available for use with [custom stickiness](https://docs.getunleash.io/reference/stickiness#custom-stickiness), otherwise &#x60;false&#x60; | [optional] 
**SortOrder** | Pointer to **int32** | How this context field should be sorted if no other sort order is selected | [optional] 
**LegalValues** | Pointer to [**[]LegalValueSchema**](LegalValueSchema.md) | A list of allowed values for this context field | [optional] 

## Methods

### NewUpdateContextFieldSchema

`func NewUpdateContextFieldSchema() *UpdateContextFieldSchema`

NewUpdateContextFieldSchema instantiates a new UpdateContextFieldSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContextFieldSchemaWithDefaults

`func NewUpdateContextFieldSchemaWithDefaults() *UpdateContextFieldSchema`

NewUpdateContextFieldSchemaWithDefaults instantiates a new UpdateContextFieldSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateContextFieldSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateContextFieldSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateContextFieldSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateContextFieldSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStickiness

`func (o *UpdateContextFieldSchema) GetStickiness() bool`

GetStickiness returns the Stickiness field if non-nil, zero value otherwise.

### GetStickinessOk

`func (o *UpdateContextFieldSchema) GetStickinessOk() (*bool, bool)`

GetStickinessOk returns a tuple with the Stickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickiness

`func (o *UpdateContextFieldSchema) SetStickiness(v bool)`

SetStickiness sets Stickiness field to given value.

### HasStickiness

`func (o *UpdateContextFieldSchema) HasStickiness() bool`

HasStickiness returns a boolean if a field has been set.

### GetSortOrder

`func (o *UpdateContextFieldSchema) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpdateContextFieldSchema) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpdateContextFieldSchema) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UpdateContextFieldSchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLegalValues

`func (o *UpdateContextFieldSchema) GetLegalValues() []LegalValueSchema`

GetLegalValues returns the LegalValues field if non-nil, zero value otherwise.

### GetLegalValuesOk

`func (o *UpdateContextFieldSchema) GetLegalValuesOk() (*[]LegalValueSchema, bool)`

GetLegalValuesOk returns a tuple with the LegalValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalValues

`func (o *UpdateContextFieldSchema) SetLegalValues(v []LegalValueSchema)`

SetLegalValues sets LegalValues field to given value.

### HasLegalValues

`func (o *UpdateContextFieldSchema) HasLegalValues() bool`

HasLegalValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


