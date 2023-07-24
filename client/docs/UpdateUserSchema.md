# UpdateUserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RootRole** | Pointer to **float32** |  | [optional] 

## Methods

### NewUpdateUserSchema

`func NewUpdateUserSchema() *UpdateUserSchema`

NewUpdateUserSchema instantiates a new UpdateUserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserSchemaWithDefaults

`func NewUpdateUserSchemaWithDefaults() *UpdateUserSchema`

NewUpdateUserSchemaWithDefaults instantiates a new UpdateUserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateUserSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserSchema) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserSchema) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UpdateUserSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootRole

`func (o *UpdateUserSchema) GetRootRole() float32`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *UpdateUserSchema) GetRootRoleOk() (*float32, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *UpdateUserSchema) SetRootRole(v float32)`

SetRootRole sets RootRole field to given value.

### HasRootRole

`func (o *UpdateUserSchema) HasRootRole() bool`

HasRootRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


