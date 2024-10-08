# ProjectCreatedSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The project&#39;s identifier. | 
**Name** | **string** | The project&#39;s name. | 
**Description** | Pointer to **NullableString** | The project&#39;s description. | [optional] 
**FeatureLimit** | Pointer to **NullableInt32** | A limit on the number of features allowed in the project. &#x60;null&#x60; if no limit. | [optional] 
**Mode** | Pointer to **string** | A mode of the project affecting what actions are possible in this project | [optional] 
**DefaultStickiness** | Pointer to **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | [optional] 
**Environments** | Pointer to **[]string** | The environments enabled for the project. | [optional] 
**ChangeRequestEnvironments** | Pointer to [**[]ProjectCreatedSchemaChangeRequestEnvironmentsInner**](ProjectCreatedSchemaChangeRequestEnvironmentsInner.md) | The list of environments that have change requests enabled. | [optional] 

## Methods

### NewProjectCreatedSchema

`func NewProjectCreatedSchema(id string, name string, ) *ProjectCreatedSchema`

NewProjectCreatedSchema instantiates a new ProjectCreatedSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreatedSchemaWithDefaults

`func NewProjectCreatedSchemaWithDefaults() *ProjectCreatedSchema`

NewProjectCreatedSchemaWithDefaults instantiates a new ProjectCreatedSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectCreatedSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCreatedSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCreatedSchema) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectCreatedSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCreatedSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCreatedSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectCreatedSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectCreatedSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectCreatedSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectCreatedSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectCreatedSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectCreatedSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFeatureLimit

`func (o *ProjectCreatedSchema) GetFeatureLimit() int32`

GetFeatureLimit returns the FeatureLimit field if non-nil, zero value otherwise.

### GetFeatureLimitOk

`func (o *ProjectCreatedSchema) GetFeatureLimitOk() (*int32, bool)`

GetFeatureLimitOk returns a tuple with the FeatureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimit

`func (o *ProjectCreatedSchema) SetFeatureLimit(v int32)`

SetFeatureLimit sets FeatureLimit field to given value.

### HasFeatureLimit

`func (o *ProjectCreatedSchema) HasFeatureLimit() bool`

HasFeatureLimit returns a boolean if a field has been set.

### SetFeatureLimitNil

`func (o *ProjectCreatedSchema) SetFeatureLimitNil(b bool)`

 SetFeatureLimitNil sets the value for FeatureLimit to be an explicit nil

### UnsetFeatureLimit
`func (o *ProjectCreatedSchema) UnsetFeatureLimit()`

UnsetFeatureLimit ensures that no value is present for FeatureLimit, not even an explicit nil
### GetMode

`func (o *ProjectCreatedSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProjectCreatedSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProjectCreatedSchema) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ProjectCreatedSchema) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDefaultStickiness

`func (o *ProjectCreatedSchema) GetDefaultStickiness() string`

GetDefaultStickiness returns the DefaultStickiness field if non-nil, zero value otherwise.

### GetDefaultStickinessOk

`func (o *ProjectCreatedSchema) GetDefaultStickinessOk() (*string, bool)`

GetDefaultStickinessOk returns a tuple with the DefaultStickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickiness

`func (o *ProjectCreatedSchema) SetDefaultStickiness(v string)`

SetDefaultStickiness sets DefaultStickiness field to given value.

### HasDefaultStickiness

`func (o *ProjectCreatedSchema) HasDefaultStickiness() bool`

HasDefaultStickiness returns a boolean if a field has been set.

### GetEnvironments

`func (o *ProjectCreatedSchema) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ProjectCreatedSchema) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ProjectCreatedSchema) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ProjectCreatedSchema) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetChangeRequestEnvironments

`func (o *ProjectCreatedSchema) GetChangeRequestEnvironments() []ProjectCreatedSchemaChangeRequestEnvironmentsInner`

GetChangeRequestEnvironments returns the ChangeRequestEnvironments field if non-nil, zero value otherwise.

### GetChangeRequestEnvironmentsOk

`func (o *ProjectCreatedSchema) GetChangeRequestEnvironmentsOk() (*[]ProjectCreatedSchemaChangeRequestEnvironmentsInner, bool)`

GetChangeRequestEnvironmentsOk returns a tuple with the ChangeRequestEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestEnvironments

`func (o *ProjectCreatedSchema) SetChangeRequestEnvironments(v []ProjectCreatedSchemaChangeRequestEnvironmentsInner)`

SetChangeRequestEnvironments sets ChangeRequestEnvironments field to given value.

### HasChangeRequestEnvironments

`func (o *ProjectCreatedSchema) HasChangeRequestEnvironments() bool`

HasChangeRequestEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


