# UpsertSegmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the segment | 
**Description** | Pointer to **NullableString** | A description of what the segment is for | [optional] 
**Project** | Pointer to **NullableString** | The project the segment belongs to if any. | [optional] 
**Constraints** | [**[]ConstraintSchema**](ConstraintSchema.md) | The list of constraints that make up this segment | 

## Methods

### NewUpsertSegmentSchema

`func NewUpsertSegmentSchema(name string, constraints []ConstraintSchema, ) *UpsertSegmentSchema`

NewUpsertSegmentSchema instantiates a new UpsertSegmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertSegmentSchemaWithDefaults

`func NewUpsertSegmentSchemaWithDefaults() *UpsertSegmentSchema`

NewUpsertSegmentSchemaWithDefaults instantiates a new UpsertSegmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpsertSegmentSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertSegmentSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertSegmentSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpsertSegmentSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpsertSegmentSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpsertSegmentSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpsertSegmentSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpsertSegmentSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpsertSegmentSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProject

`func (o *UpsertSegmentSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *UpsertSegmentSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *UpsertSegmentSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *UpsertSegmentSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *UpsertSegmentSchema) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *UpsertSegmentSchema) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetConstraints

`func (o *UpsertSegmentSchema) GetConstraints() []ConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *UpsertSegmentSchema) GetConstraintsOk() (*[]ConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *UpsertSegmentSchema) SetConstraints(v []ConstraintSchema)`

SetConstraints sets Constraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


