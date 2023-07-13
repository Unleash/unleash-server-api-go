# UpsertContextFieldSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Stickiness** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **float32** |  | [optional] 
**LegalValues** | Pointer to [**[]LegalValueSchema**](LegalValueSchema.md) |  | [optional] 

## Methods

### NewUpsertContextFieldSchema

`func NewUpsertContextFieldSchema(name string, ) *UpsertContextFieldSchema`

NewUpsertContextFieldSchema instantiates a new UpsertContextFieldSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertContextFieldSchemaWithDefaults

`func NewUpsertContextFieldSchemaWithDefaults() *UpsertContextFieldSchema`

NewUpsertContextFieldSchemaWithDefaults instantiates a new UpsertContextFieldSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpsertContextFieldSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertContextFieldSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertContextFieldSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpsertContextFieldSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpsertContextFieldSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpsertContextFieldSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpsertContextFieldSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStickiness

`func (o *UpsertContextFieldSchema) GetStickiness() bool`

GetStickiness returns the Stickiness field if non-nil, zero value otherwise.

### GetStickinessOk

`func (o *UpsertContextFieldSchema) GetStickinessOk() (*bool, bool)`

GetStickinessOk returns a tuple with the Stickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickiness

`func (o *UpsertContextFieldSchema) SetStickiness(v bool)`

SetStickiness sets Stickiness field to given value.

### HasStickiness

`func (o *UpsertContextFieldSchema) HasStickiness() bool`

HasStickiness returns a boolean if a field has been set.

### GetSortOrder

`func (o *UpsertContextFieldSchema) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpsertContextFieldSchema) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpsertContextFieldSchema) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UpsertContextFieldSchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLegalValues

`func (o *UpsertContextFieldSchema) GetLegalValues() []LegalValueSchema`

GetLegalValues returns the LegalValues field if non-nil, zero value otherwise.

### GetLegalValuesOk

`func (o *UpsertContextFieldSchema) GetLegalValuesOk() (*[]LegalValueSchema, bool)`

GetLegalValuesOk returns a tuple with the LegalValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalValues

`func (o *UpsertContextFieldSchema) SetLegalValues(v []LegalValueSchema)`

SetLegalValues sets LegalValues field to given value.

### HasLegalValues

`func (o *UpsertContextFieldSchema) HasLegalValues() bool`

HasLegalValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


