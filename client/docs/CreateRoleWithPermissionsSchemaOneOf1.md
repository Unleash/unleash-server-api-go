# CreateRoleWithPermissionsSchemaOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the custom role | 
**Description** | Pointer to **string** | A more detailed description of the custom role and what use it&#39;s intended for | [optional] 
**Type** | Pointer to **string** | [Custom project roles](https://docs.getunleash.io/reference/rbac#custom-project-roles) contain a specific set of permissions for project resources. | [optional] 
**Permissions** | Pointer to [**[]CreateRoleWithPermissionsSchemaOneOf1PermissionsInner**](CreateRoleWithPermissionsSchemaOneOf1PermissionsInner.md) | A list of permissions assigned to this role | [optional] 

## Methods

### NewCreateRoleWithPermissionsSchemaOneOf1

`func NewCreateRoleWithPermissionsSchemaOneOf1(name string, ) *CreateRoleWithPermissionsSchemaOneOf1`

NewCreateRoleWithPermissionsSchemaOneOf1 instantiates a new CreateRoleWithPermissionsSchemaOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleWithPermissionsSchemaOneOf1WithDefaults

`func NewCreateRoleWithPermissionsSchemaOneOf1WithDefaults() *CreateRoleWithPermissionsSchemaOneOf1`

NewCreateRoleWithPermissionsSchemaOneOf1WithDefaults instantiates a new CreateRoleWithPermissionsSchemaOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleWithPermissionsSchemaOneOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleWithPermissionsSchemaOneOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleWithPermissionsSchemaOneOf1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRoleWithPermissionsSchemaOneOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRoleWithPermissionsSchemaOneOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRoleWithPermissionsSchemaOneOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRoleWithPermissionsSchemaOneOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateRoleWithPermissionsSchemaOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRoleWithPermissionsSchemaOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRoleWithPermissionsSchemaOneOf1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateRoleWithPermissionsSchemaOneOf1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateRoleWithPermissionsSchemaOneOf1) GetPermissions() []CreateRoleWithPermissionsSchemaOneOf1PermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateRoleWithPermissionsSchemaOneOf1) GetPermissionsOk() (*[]CreateRoleWithPermissionsSchemaOneOf1PermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateRoleWithPermissionsSchemaOneOf1) SetPermissions(v []CreateRoleWithPermissionsSchemaOneOf1PermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateRoleWithPermissionsSchemaOneOf1) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


