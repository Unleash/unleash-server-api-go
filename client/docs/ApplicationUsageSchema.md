# ApplicationUsageSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | **string** | Name of the project | 
**Environments** | **[]string** | Which environments have been accessed in this project. | 

## Methods

### NewApplicationUsageSchema

`func NewApplicationUsageSchema(project string, environments []string, ) *ApplicationUsageSchema`

NewApplicationUsageSchema instantiates a new ApplicationUsageSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUsageSchemaWithDefaults

`func NewApplicationUsageSchemaWithDefaults() *ApplicationUsageSchema`

NewApplicationUsageSchemaWithDefaults instantiates a new ApplicationUsageSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ApplicationUsageSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ApplicationUsageSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ApplicationUsageSchema) SetProject(v string)`

SetProject sets Project field to given value.


### GetEnvironments

`func (o *ApplicationUsageSchema) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ApplicationUsageSchema) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ApplicationUsageSchema) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


