# ProjectAccessConfigurationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | [**[]ProjectAccessConfigurationSchemaRolesInner**](ProjectAccessConfigurationSchemaRolesInner.md) | A list of roles that are available within this project. | 

## Methods

### NewProjectAccessConfigurationSchema

`func NewProjectAccessConfigurationSchema(roles []ProjectAccessConfigurationSchemaRolesInner, ) *ProjectAccessConfigurationSchema`

NewProjectAccessConfigurationSchema instantiates a new ProjectAccessConfigurationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAccessConfigurationSchemaWithDefaults

`func NewProjectAccessConfigurationSchemaWithDefaults() *ProjectAccessConfigurationSchema`

NewProjectAccessConfigurationSchemaWithDefaults instantiates a new ProjectAccessConfigurationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *ProjectAccessConfigurationSchema) GetRoles() []ProjectAccessConfigurationSchemaRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ProjectAccessConfigurationSchema) GetRolesOk() (*[]ProjectAccessConfigurationSchemaRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ProjectAccessConfigurationSchema) SetRoles(v []ProjectAccessConfigurationSchemaRolesInner)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


