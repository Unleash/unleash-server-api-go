# UsersGroupsBaseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]GroupSchema**](GroupSchema.md) |  | [optional] 
**Users** | Pointer to [**[]UserSchema**](UserSchema.md) |  | [optional] 

## Methods

### NewUsersGroupsBaseSchema

`func NewUsersGroupsBaseSchema() *UsersGroupsBaseSchema`

NewUsersGroupsBaseSchema instantiates a new UsersGroupsBaseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersGroupsBaseSchemaWithDefaults

`func NewUsersGroupsBaseSchemaWithDefaults() *UsersGroupsBaseSchema`

NewUsersGroupsBaseSchemaWithDefaults instantiates a new UsersGroupsBaseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *UsersGroupsBaseSchema) GetGroups() []GroupSchema`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UsersGroupsBaseSchema) GetGroupsOk() (*[]GroupSchema, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UsersGroupsBaseSchema) SetGroups(v []GroupSchema)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UsersGroupsBaseSchema) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *UsersGroupsBaseSchema) GetUsers() []UserSchema`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersGroupsBaseSchema) GetUsersOk() (*[]UserSchema, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersGroupsBaseSchema) SetUsers(v []UserSchema)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UsersGroupsBaseSchema) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


