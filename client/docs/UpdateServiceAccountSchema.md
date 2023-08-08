# UpdateServiceAccountSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the service account | [optional] 
**RootRole** | Pointer to **int32** | The id of the root role for the service account | [optional] 

## Methods

### NewUpdateServiceAccountSchema

`func NewUpdateServiceAccountSchema() *UpdateServiceAccountSchema`

NewUpdateServiceAccountSchema instantiates a new UpdateServiceAccountSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceAccountSchemaWithDefaults

`func NewUpdateServiceAccountSchemaWithDefaults() *UpdateServiceAccountSchema`

NewUpdateServiceAccountSchemaWithDefaults instantiates a new UpdateServiceAccountSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateServiceAccountSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServiceAccountSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServiceAccountSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServiceAccountSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootRole

`func (o *UpdateServiceAccountSchema) GetRootRole() int32`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *UpdateServiceAccountSchema) GetRootRoleOk() (*int32, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *UpdateServiceAccountSchema) SetRootRole(v int32)`

SetRootRole sets RootRole field to given value.

### HasRootRole

`func (o *UpdateServiceAccountSchema) HasRootRole() bool`

HasRootRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


