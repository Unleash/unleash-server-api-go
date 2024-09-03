# PasswordAuthSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Is username/password authentication disabled | [optional] 

## Methods

### NewPasswordAuthSchema

`func NewPasswordAuthSchema() *PasswordAuthSchema`

NewPasswordAuthSchema instantiates a new PasswordAuthSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordAuthSchemaWithDefaults

`func NewPasswordAuthSchemaWithDefaults() *PasswordAuthSchema`

NewPasswordAuthSchemaWithDefaults instantiates a new PasswordAuthSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *PasswordAuthSchema) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PasswordAuthSchema) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PasswordAuthSchema) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PasswordAuthSchema) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


