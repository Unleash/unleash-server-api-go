# SegmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The segment&#39;s id. | 
**Name** | Pointer to **string** | The name of the segment. | [optional] 
**Constraints** | [**[]ConstraintSchema**](ConstraintSchema.md) | List of constraints that determine which users are part of the segment | 
**Description** | Pointer to **NullableString** | The description of the segment. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the segment was created as a RFC 3339-conformant timestamp. | [optional] 
**CreatedBy** | Pointer to **string** | Which user created this segment | [optional] 
**Project** | Pointer to **NullableString** | The project the segment relates to, if applicable. | [optional] 

## Methods

### NewSegmentSchema

`func NewSegmentSchema(id float32, constraints []ConstraintSchema, ) *SegmentSchema`

NewSegmentSchema instantiates a new SegmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentSchemaWithDefaults

`func NewSegmentSchemaWithDefaults() *SegmentSchema`

NewSegmentSchemaWithDefaults instantiates a new SegmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SegmentSchema) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentSchema) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentSchema) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *SegmentSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SegmentSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConstraints

`func (o *SegmentSchema) GetConstraints() []ConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *SegmentSchema) GetConstraintsOk() (*[]ConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *SegmentSchema) SetConstraints(v []ConstraintSchema)`

SetConstraints sets Constraints field to given value.


### GetDescription

`func (o *SegmentSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SegmentSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SegmentSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatedAt

`func (o *SegmentSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SegmentSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SegmentSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SegmentSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SegmentSchema) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SegmentSchema) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SegmentSchema) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SegmentSchema) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetProject

`func (o *SegmentSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SegmentSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SegmentSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SegmentSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *SegmentSchema) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *SegmentSchema) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


