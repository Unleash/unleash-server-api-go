# RoleWithPermissionsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Permissions** | [**[]AdminPermissionSchema**](AdminPermissionSchema.md) |  | 

## Methods

### NewRoleWithPermissionsSchema

`func NewRoleWithPermissionsSchema(id float32, type_ string, name string, permissions []AdminPermissionSchema, ) *RoleWithPermissionsSchema`

NewRoleWithPermissionsSchema instantiates a new RoleWithPermissionsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithPermissionsSchemaWithDefaults

`func NewRoleWithPermissionsSchemaWithDefaults() *RoleWithPermissionsSchema`

NewRoleWithPermissionsSchemaWithDefaults instantiates a new RoleWithPermissionsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleWithPermissionsSchema) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleWithPermissionsSchema) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleWithPermissionsSchema) SetId(v float32)`

SetId sets Id field to given value.


### GetType

`func (o *RoleWithPermissionsSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleWithPermissionsSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleWithPermissionsSchema) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RoleWithPermissionsSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleWithPermissionsSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleWithPermissionsSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RoleWithPermissionsSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleWithPermissionsSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleWithPermissionsSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleWithPermissionsSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *RoleWithPermissionsSchema) GetPermissions() []AdminPermissionSchema`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleWithPermissionsSchema) GetPermissionsOk() (*[]AdminPermissionSchema, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleWithPermissionsSchema) SetPermissions(v []AdminPermissionSchema)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


