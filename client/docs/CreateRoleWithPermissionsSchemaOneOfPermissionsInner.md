# CreateRoleWithPermissionsSchemaOneOfPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the permission | 
**Environment** | Pointer to **string** | The environments of the permission if the permission is environment specific | [optional] 

## Methods

### NewCreateRoleWithPermissionsSchemaOneOfPermissionsInner

`func NewCreateRoleWithPermissionsSchemaOneOfPermissionsInner(name string, ) *CreateRoleWithPermissionsSchemaOneOfPermissionsInner`

NewCreateRoleWithPermissionsSchemaOneOfPermissionsInner instantiates a new CreateRoleWithPermissionsSchemaOneOfPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleWithPermissionsSchemaOneOfPermissionsInnerWithDefaults

`func NewCreateRoleWithPermissionsSchemaOneOfPermissionsInnerWithDefaults() *CreateRoleWithPermissionsSchemaOneOfPermissionsInner`

NewCreateRoleWithPermissionsSchemaOneOfPermissionsInnerWithDefaults instantiates a new CreateRoleWithPermissionsSchemaOneOfPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironment

`func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


