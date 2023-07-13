# GroupUserModelSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JoinedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**User** | [**UserSchema**](UserSchema.md) |  | 

## Methods

### NewGroupUserModelSchema

`func NewGroupUserModelSchema(user UserSchema, ) *GroupUserModelSchema`

NewGroupUserModelSchema instantiates a new GroupUserModelSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserModelSchemaWithDefaults

`func NewGroupUserModelSchemaWithDefaults() *GroupUserModelSchema`

NewGroupUserModelSchemaWithDefaults instantiates a new GroupUserModelSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoinedAt

`func (o *GroupUserModelSchema) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *GroupUserModelSchema) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *GroupUserModelSchema) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *GroupUserModelSchema) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GroupUserModelSchema) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GroupUserModelSchema) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GroupUserModelSchema) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GroupUserModelSchema) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GroupUserModelSchema) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GroupUserModelSchema) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUser

`func (o *GroupUserModelSchema) GetUser() UserSchema`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GroupUserModelSchema) GetUserOk() (*UserSchema, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GroupUserModelSchema) SetUser(v UserSchema)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


