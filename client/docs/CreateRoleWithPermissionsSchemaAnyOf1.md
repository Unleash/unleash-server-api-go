# CreateRoleWithPermissionsSchemaAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the custom role | 
**Description** | Pointer to **string** | A more detailed description of the custom role and what use it&#39;s intended for | [optional] 
**Type** | Pointer to **string** | [Custom project roles](https://docs.getunleash.io/reference/rbac#custom-project-roles) contain a specific set of permissions for project resources. | [optional] 
**Permissions** | Pointer to [**[]CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner**](CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner.md) | A list of permissions assigned to this role | [optional] 

## Methods

### NewCreateRoleWithPermissionsSchemaAnyOf1

`func NewCreateRoleWithPermissionsSchemaAnyOf1(name string, ) *CreateRoleWithPermissionsSchemaAnyOf1`

NewCreateRoleWithPermissionsSchemaAnyOf1 instantiates a new CreateRoleWithPermissionsSchemaAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleWithPermissionsSchemaAnyOf1WithDefaults

`func NewCreateRoleWithPermissionsSchemaAnyOf1WithDefaults() *CreateRoleWithPermissionsSchemaAnyOf1`

NewCreateRoleWithPermissionsSchemaAnyOf1WithDefaults instantiates a new CreateRoleWithPermissionsSchemaAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetPermissions() []CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetPermissionsOk() (*[]CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) SetPermissions(v []CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateRoleWithPermissionsSchemaAnyOf1) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


