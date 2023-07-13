# RolesWithVersionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **float32** |  | 
**Roles** | [**[]RoleSchema**](RoleSchema.md) |  | 

## Methods

### NewRolesWithVersionSchema

`func NewRolesWithVersionSchema(version float32, roles []RoleSchema, ) *RolesWithVersionSchema`

NewRolesWithVersionSchema instantiates a new RolesWithVersionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesWithVersionSchemaWithDefaults

`func NewRolesWithVersionSchemaWithDefaults() *RolesWithVersionSchema`

NewRolesWithVersionSchemaWithDefaults instantiates a new RolesWithVersionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *RolesWithVersionSchema) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RolesWithVersionSchema) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RolesWithVersionSchema) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetRoles

`func (o *RolesWithVersionSchema) GetRoles() []RoleSchema`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RolesWithVersionSchema) GetRolesOk() (*[]RoleSchema, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RolesWithVersionSchema) SetRoles(v []RoleSchema)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


