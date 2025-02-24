# CreateRoleWithPermissionsSchemaAnyOfPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the permission | 
**Environment** | Pointer to **NullableString** | The environments of the permission if the permission is environment specific | [optional] 

## Methods

### NewCreateRoleWithPermissionsSchemaAnyOfPermissionsInner

`func NewCreateRoleWithPermissionsSchemaAnyOfPermissionsInner(name string, ) *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner`

NewCreateRoleWithPermissionsSchemaAnyOfPermissionsInner instantiates a new CreateRoleWithPermissionsSchemaAnyOfPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleWithPermissionsSchemaAnyOfPermissionsInnerWithDefaults

`func NewCreateRoleWithPermissionsSchemaAnyOfPermissionsInnerWithDefaults() *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner`

NewCreateRoleWithPermissionsSchemaAnyOfPermissionsInnerWithDefaults instantiates a new CreateRoleWithPermissionsSchemaAnyOfPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironment

`func (o *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CreateRoleWithPermissionsSchemaAnyOfPermissionsInner) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


