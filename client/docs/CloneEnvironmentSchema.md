# CloneEnvironmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new cloned environment, this cannot be changed later | 
**Type** | **string** | Updates the type of environment (i.e. development or production). | 
**Projects** | Pointer to **[]string** | A list of projects that should be included in the cloned environment. | [optional] 
**ClonePermissions** | Pointer to **bool** | Copies the RBAC permissions from the source environment if true. Defaults to true | [optional] 

## Methods

### NewCloneEnvironmentSchema

`func NewCloneEnvironmentSchema(name string, type_ string, ) *CloneEnvironmentSchema`

NewCloneEnvironmentSchema instantiates a new CloneEnvironmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneEnvironmentSchemaWithDefaults

`func NewCloneEnvironmentSchemaWithDefaults() *CloneEnvironmentSchema`

NewCloneEnvironmentSchemaWithDefaults instantiates a new CloneEnvironmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloneEnvironmentSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneEnvironmentSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneEnvironmentSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CloneEnvironmentSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloneEnvironmentSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloneEnvironmentSchema) SetType(v string)`

SetType sets Type field to given value.


### GetProjects

`func (o *CloneEnvironmentSchema) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CloneEnvironmentSchema) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CloneEnvironmentSchema) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CloneEnvironmentSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetClonePermissions

`func (o *CloneEnvironmentSchema) GetClonePermissions() bool`

GetClonePermissions returns the ClonePermissions field if non-nil, zero value otherwise.

### GetClonePermissionsOk

`func (o *CloneEnvironmentSchema) GetClonePermissionsOk() (*bool, bool)`

GetClonePermissionsOk returns a tuple with the ClonePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonePermissions

`func (o *CloneEnvironmentSchema) SetClonePermissions(v bool)`

SetClonePermissions sets ClonePermissions field to given value.

### HasClonePermissions

`func (o *CloneEnvironmentSchema) HasClonePermissions() bool`

HasClonePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


