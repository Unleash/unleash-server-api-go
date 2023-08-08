# AdminPermissionsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**AdminPermissionsSchemaPermissions**](AdminPermissionsSchemaPermissions.md) |  | 
**Version** | **int32** | The api version of this response. A natural increasing number. Only increases if format changes | 

## Methods

### NewAdminPermissionsSchema

`func NewAdminPermissionsSchema(permissions AdminPermissionsSchemaPermissions, version int32, ) *AdminPermissionsSchema`

NewAdminPermissionsSchema instantiates a new AdminPermissionsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminPermissionsSchemaWithDefaults

`func NewAdminPermissionsSchemaWithDefaults() *AdminPermissionsSchema`

NewAdminPermissionsSchemaWithDefaults instantiates a new AdminPermissionsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *AdminPermissionsSchema) GetPermissions() AdminPermissionsSchemaPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AdminPermissionsSchema) GetPermissionsOk() (*AdminPermissionsSchemaPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AdminPermissionsSchema) SetPermissions(v AdminPermissionsSchemaPermissions)`

SetPermissions sets Permissions field to given value.


### GetVersion

`func (o *AdminPermissionsSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdminPermissionsSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdminPermissionsSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


