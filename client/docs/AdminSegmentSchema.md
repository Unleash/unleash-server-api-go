# AdminSegmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of this segment | 
**Name** | **string** | The name of this segment | 
**Description** | Pointer to **NullableString** | The description for this segment | [optional] 
**Constraints** | [**[]ConstraintSchema**](ConstraintSchema.md) | The list of constraints that are used in this segment | 
**UsedInFeatures** | Pointer to **NullableInt32** | The number of projects that use this segment | [optional] 
**UsedInProjects** | Pointer to **NullableInt32** | The number of projects that use this segment | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **string** | The creator&#39;s email or username | [optional] 
**CreatedAt** | **time.Time** | When the segment was created | 

## Methods

### NewAdminSegmentSchema

`func NewAdminSegmentSchema(id int32, name string, constraints []ConstraintSchema, createdAt time.Time, ) *AdminSegmentSchema`

NewAdminSegmentSchema instantiates a new AdminSegmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminSegmentSchemaWithDefaults

`func NewAdminSegmentSchemaWithDefaults() *AdminSegmentSchema`

NewAdminSegmentSchemaWithDefaults instantiates a new AdminSegmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminSegmentSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminSegmentSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminSegmentSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AdminSegmentSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminSegmentSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminSegmentSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AdminSegmentSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminSegmentSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminSegmentSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminSegmentSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AdminSegmentSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AdminSegmentSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConstraints

`func (o *AdminSegmentSchema) GetConstraints() []ConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AdminSegmentSchema) GetConstraintsOk() (*[]ConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AdminSegmentSchema) SetConstraints(v []ConstraintSchema)`

SetConstraints sets Constraints field to given value.


### GetUsedInFeatures

`func (o *AdminSegmentSchema) GetUsedInFeatures() int32`

GetUsedInFeatures returns the UsedInFeatures field if non-nil, zero value otherwise.

### GetUsedInFeaturesOk

`func (o *AdminSegmentSchema) GetUsedInFeaturesOk() (*int32, bool)`

GetUsedInFeaturesOk returns a tuple with the UsedInFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInFeatures

`func (o *AdminSegmentSchema) SetUsedInFeatures(v int32)`

SetUsedInFeatures sets UsedInFeatures field to given value.

### HasUsedInFeatures

`func (o *AdminSegmentSchema) HasUsedInFeatures() bool`

HasUsedInFeatures returns a boolean if a field has been set.

### SetUsedInFeaturesNil

`func (o *AdminSegmentSchema) SetUsedInFeaturesNil(b bool)`

 SetUsedInFeaturesNil sets the value for UsedInFeatures to be an explicit nil

### UnsetUsedInFeatures
`func (o *AdminSegmentSchema) UnsetUsedInFeatures()`

UnsetUsedInFeatures ensures that no value is present for UsedInFeatures, not even an explicit nil
### GetUsedInProjects

`func (o *AdminSegmentSchema) GetUsedInProjects() int32`

GetUsedInProjects returns the UsedInProjects field if non-nil, zero value otherwise.

### GetUsedInProjectsOk

`func (o *AdminSegmentSchema) GetUsedInProjectsOk() (*int32, bool)`

GetUsedInProjectsOk returns a tuple with the UsedInProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInProjects

`func (o *AdminSegmentSchema) SetUsedInProjects(v int32)`

SetUsedInProjects sets UsedInProjects field to given value.

### HasUsedInProjects

`func (o *AdminSegmentSchema) HasUsedInProjects() bool`

HasUsedInProjects returns a boolean if a field has been set.

### SetUsedInProjectsNil

`func (o *AdminSegmentSchema) SetUsedInProjectsNil(b bool)`

 SetUsedInProjectsNil sets the value for UsedInProjects to be an explicit nil

### UnsetUsedInProjects
`func (o *AdminSegmentSchema) UnsetUsedInProjects()`

UnsetUsedInProjects ensures that no value is present for UsedInProjects, not even an explicit nil
### GetProject

`func (o *AdminSegmentSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AdminSegmentSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AdminSegmentSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AdminSegmentSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *AdminSegmentSchema) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *AdminSegmentSchema) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetCreatedBy

`func (o *AdminSegmentSchema) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AdminSegmentSchema) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AdminSegmentSchema) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AdminSegmentSchema) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminSegmentSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminSegmentSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminSegmentSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


