# UsersSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserSchema**](UserSchema.md) | A list of users in the Unleash instance. | 
**RootRoles** | Pointer to [**[]RoleSchema**](RoleSchema.md) | A list of [root roles](https://docs.getunleash.io/reference/rbac#predefined-roles) in the Unleash instance. | [optional] 

## Methods

### NewUsersSchema

`func NewUsersSchema(users []UserSchema, ) *UsersSchema`

NewUsersSchema instantiates a new UsersSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersSchemaWithDefaults

`func NewUsersSchemaWithDefaults() *UsersSchema`

NewUsersSchemaWithDefaults instantiates a new UsersSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UsersSchema) GetUsers() []UserSchema`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersSchema) GetUsersOk() (*[]UserSchema, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersSchema) SetUsers(v []UserSchema)`

SetUsers sets Users field to given value.


### GetRootRoles

`func (o *UsersSchema) GetRootRoles() []RoleSchema`

GetRootRoles returns the RootRoles field if non-nil, zero value otherwise.

### GetRootRolesOk

`func (o *UsersSchema) GetRootRolesOk() (*[]RoleSchema, bool)`

GetRootRolesOk returns a tuple with the RootRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRoles

`func (o *UsersSchema) SetRootRoles(v []RoleSchema)`

SetRootRoles sets RootRoles field to given value.

### HasRootRoles

`func (o *UsersSchema) HasRootRoles() bool`

HasRootRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


