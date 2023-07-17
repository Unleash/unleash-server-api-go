# AdvancedPlaygroundFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The feature&#39;s name. | 
**ProjectId** | **string** | The ID of the project that contains this feature. | 
**Environments** | [**map[string][]AdvancedPlaygroundEnvironmentFeatureSchema**](array.md) | The lists of features that have been evaluated grouped by environment. | 

## Methods

### NewAdvancedPlaygroundFeatureSchema

`func NewAdvancedPlaygroundFeatureSchema(name string, projectId string, environments map[string][]AdvancedPlaygroundEnvironmentFeatureSchema, ) *AdvancedPlaygroundFeatureSchema`

NewAdvancedPlaygroundFeatureSchema instantiates a new AdvancedPlaygroundFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedPlaygroundFeatureSchemaWithDefaults

`func NewAdvancedPlaygroundFeatureSchemaWithDefaults() *AdvancedPlaygroundFeatureSchema`

NewAdvancedPlaygroundFeatureSchemaWithDefaults instantiates a new AdvancedPlaygroundFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdvancedPlaygroundFeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvancedPlaygroundFeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvancedPlaygroundFeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *AdvancedPlaygroundFeatureSchema) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AdvancedPlaygroundFeatureSchema) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AdvancedPlaygroundFeatureSchema) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetEnvironments

`func (o *AdvancedPlaygroundFeatureSchema) GetEnvironments() map[string][]AdvancedPlaygroundEnvironmentFeatureSchema`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *AdvancedPlaygroundFeatureSchema) GetEnvironmentsOk() (*map[string][]AdvancedPlaygroundEnvironmentFeatureSchema, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *AdvancedPlaygroundFeatureSchema) SetEnvironments(v map[string][]AdvancedPlaygroundEnvironmentFeatureSchema)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


