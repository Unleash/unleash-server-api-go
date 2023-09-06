# RoleWithVersionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The version of this schema | 
**Roles** | [**RoleSchema**](RoleSchema.md) |  | 

## Methods

### NewRoleWithVersionSchema

`func NewRoleWithVersionSchema(version int32, roles RoleSchema, ) *RoleWithVersionSchema`

NewRoleWithVersionSchema instantiates a new RoleWithVersionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithVersionSchemaWithDefaults

`func NewRoleWithVersionSchemaWithDefaults() *RoleWithVersionSchema`

NewRoleWithVersionSchemaWithDefaults instantiates a new RoleWithVersionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *RoleWithVersionSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RoleWithVersionSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RoleWithVersionSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetRoles

`func (o *RoleWithVersionSchema) GetRoles() RoleSchema`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleWithVersionSchema) GetRolesOk() (*RoleSchema, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleWithVersionSchema) SetRoles(v RoleSchema)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


