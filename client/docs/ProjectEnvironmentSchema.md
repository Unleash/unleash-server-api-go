# ProjectEnvironmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** | The environment to add to the project | 
**ChangeRequestsEnabled** | Pointer to **bool** | Whether change requests should be enabled or for this environment on the project or not | [optional] 
**DefaultStrategy** | Pointer to [**CreateFeatureStrategySchema**](CreateFeatureStrategySchema.md) |  | [optional] 

## Methods

### NewProjectEnvironmentSchema

`func NewProjectEnvironmentSchema(environment string, ) *ProjectEnvironmentSchema`

NewProjectEnvironmentSchema instantiates a new ProjectEnvironmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectEnvironmentSchemaWithDefaults

`func NewProjectEnvironmentSchemaWithDefaults() *ProjectEnvironmentSchema`

NewProjectEnvironmentSchemaWithDefaults instantiates a new ProjectEnvironmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ProjectEnvironmentSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProjectEnvironmentSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProjectEnvironmentSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetChangeRequestsEnabled

`func (o *ProjectEnvironmentSchema) GetChangeRequestsEnabled() bool`

GetChangeRequestsEnabled returns the ChangeRequestsEnabled field if non-nil, zero value otherwise.

### GetChangeRequestsEnabledOk

`func (o *ProjectEnvironmentSchema) GetChangeRequestsEnabledOk() (*bool, bool)`

GetChangeRequestsEnabledOk returns a tuple with the ChangeRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestsEnabled

`func (o *ProjectEnvironmentSchema) SetChangeRequestsEnabled(v bool)`

SetChangeRequestsEnabled sets ChangeRequestsEnabled field to given value.

### HasChangeRequestsEnabled

`func (o *ProjectEnvironmentSchema) HasChangeRequestsEnabled() bool`

HasChangeRequestsEnabled returns a boolean if a field has been set.

### GetDefaultStrategy

`func (o *ProjectEnvironmentSchema) GetDefaultStrategy() CreateFeatureStrategySchema`

GetDefaultStrategy returns the DefaultStrategy field if non-nil, zero value otherwise.

### GetDefaultStrategyOk

`func (o *ProjectEnvironmentSchema) GetDefaultStrategyOk() (*CreateFeatureStrategySchema, bool)`

GetDefaultStrategyOk returns a tuple with the DefaultStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStrategy

`func (o *ProjectEnvironmentSchema) SetDefaultStrategy(v CreateFeatureStrategySchema)`

SetDefaultStrategy sets DefaultStrategy field to given value.

### HasDefaultStrategy

`func (o *ProjectEnvironmentSchema) HasDefaultStrategy() bool`

HasDefaultStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


