# CreateRoleWithPermissionsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the custom role | 
**Description** | Pointer to **string** | A more detailed description of the custom role and what use it&#39;s intended for | [optional] 
**Type** | Pointer to **string** | [Custom root roles](https://docs.getunleash.io/concepts/rbac#custom-root-roles) (type&#x3D;root-custom) are root roles with a custom set of permissions. [Custom project roles](https://docs.getunleash.io/concepts/rbac#custom-project-roles) (type&#x3D;custom) contain a specific set of permissions for project resources. | [optional] 
**Permissions** | Pointer to [**[]CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner**](CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner.md) | A list of permissions assigned to this role | [optional] 

## Methods

### NewCreateRoleWithPermissionsSchema

`func NewCreateRoleWithPermissionsSchema(name string, ) *CreateRoleWithPermissionsSchema`

NewCreateRoleWithPermissionsSchema instantiates a new CreateRoleWithPermissionsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleWithPermissionsSchemaWithDefaults

`func NewCreateRoleWithPermissionsSchemaWithDefaults() *CreateRoleWithPermissionsSchema`

NewCreateRoleWithPermissionsSchemaWithDefaults instantiates a new CreateRoleWithPermissionsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleWithPermissionsSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleWithPermissionsSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleWithPermissionsSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRoleWithPermissionsSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRoleWithPermissionsSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRoleWithPermissionsSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRoleWithPermissionsSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateRoleWithPermissionsSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRoleWithPermissionsSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRoleWithPermissionsSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateRoleWithPermissionsSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateRoleWithPermissionsSchema) GetPermissions() []CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateRoleWithPermissionsSchema) GetPermissionsOk() (*[]CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateRoleWithPermissionsSchema) SetPermissions(v []CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateRoleWithPermissionsSchema) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


