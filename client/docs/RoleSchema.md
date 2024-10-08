# RoleSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The role id | 
**Type** | **string** | A role can either be a global root role (applies to all projects) or a project role | 
**Name** | **string** | The name of the role | 
**Description** | Pointer to **string** | A more detailed description of the role and what use it&#39;s intended for | [optional] 
**Project** | Pointer to **NullableString** | What project the role belongs to | [optional] 

## Methods

### NewRoleSchema

`func NewRoleSchema(id int32, type_ string, name string, ) *RoleSchema`

NewRoleSchema instantiates a new RoleSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleSchemaWithDefaults

`func NewRoleSchemaWithDefaults() *RoleSchema`

NewRoleSchemaWithDefaults instantiates a new RoleSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *RoleSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleSchema) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RoleSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RoleSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *RoleSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RoleSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RoleSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *RoleSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *RoleSchema) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *RoleSchema) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


