# AdvancedPlaygroundRequestSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | **[]string** | The environments to evaluate toggles in. | 
**Projects** | Pointer to [**AdvancedPlaygroundRequestSchemaProjects**](AdvancedPlaygroundRequestSchemaProjects.md) |  | [optional] 
**Context** | [**SdkContextSchema**](SdkContextSchema.md) |  | 

## Methods

### NewAdvancedPlaygroundRequestSchema

`func NewAdvancedPlaygroundRequestSchema(environments []string, context SdkContextSchema, ) *AdvancedPlaygroundRequestSchema`

NewAdvancedPlaygroundRequestSchema instantiates a new AdvancedPlaygroundRequestSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedPlaygroundRequestSchemaWithDefaults

`func NewAdvancedPlaygroundRequestSchemaWithDefaults() *AdvancedPlaygroundRequestSchema`

NewAdvancedPlaygroundRequestSchemaWithDefaults instantiates a new AdvancedPlaygroundRequestSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *AdvancedPlaygroundRequestSchema) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *AdvancedPlaygroundRequestSchema) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *AdvancedPlaygroundRequestSchema) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.


### GetProjects

`func (o *AdvancedPlaygroundRequestSchema) GetProjects() AdvancedPlaygroundRequestSchemaProjects`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AdvancedPlaygroundRequestSchema) GetProjectsOk() (*AdvancedPlaygroundRequestSchemaProjects, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AdvancedPlaygroundRequestSchema) SetProjects(v AdvancedPlaygroundRequestSchemaProjects)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *AdvancedPlaygroundRequestSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetContext

`func (o *AdvancedPlaygroundRequestSchema) GetContext() SdkContextSchema`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AdvancedPlaygroundRequestSchema) GetContextOk() (*SdkContextSchema, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AdvancedPlaygroundRequestSchema) SetContext(v SdkContextSchema)`

SetContext sets Context field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


