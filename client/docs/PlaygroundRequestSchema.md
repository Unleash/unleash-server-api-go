# PlaygroundRequestSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** | The environment to evaluate toggles in. | 
**Projects** | Pointer to [**AdvancedPlaygroundRequestSchemaProjects**](AdvancedPlaygroundRequestSchemaProjects.md) |  | [optional] 
**Context** | [**SdkContextSchema**](SdkContextSchema.md) |  | 

## Methods

### NewPlaygroundRequestSchema

`func NewPlaygroundRequestSchema(environment string, context SdkContextSchema, ) *PlaygroundRequestSchema`

NewPlaygroundRequestSchema instantiates a new PlaygroundRequestSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundRequestSchemaWithDefaults

`func NewPlaygroundRequestSchemaWithDefaults() *PlaygroundRequestSchema`

NewPlaygroundRequestSchemaWithDefaults instantiates a new PlaygroundRequestSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *PlaygroundRequestSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PlaygroundRequestSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PlaygroundRequestSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetProjects

`func (o *PlaygroundRequestSchema) GetProjects() AdvancedPlaygroundRequestSchemaProjects`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PlaygroundRequestSchema) GetProjectsOk() (*AdvancedPlaygroundRequestSchemaProjects, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PlaygroundRequestSchema) SetProjects(v AdvancedPlaygroundRequestSchemaProjects)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PlaygroundRequestSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetContext

`func (o *PlaygroundRequestSchema) GetContext() SdkContextSchema`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PlaygroundRequestSchema) GetContextOk() (*SdkContextSchema, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PlaygroundRequestSchema) SetContext(v SdkContextSchema)`

SetContext sets Context field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


