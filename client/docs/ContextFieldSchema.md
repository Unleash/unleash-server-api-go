# ContextFieldSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the context field | 
**Description** | Pointer to **NullableString** | The description of the context field. | [optional] 
**Stickiness** | Pointer to **bool** | Does this context field support being used for [stickiness](https://docs.getunleash.io/reference/stickiness) calculations | [optional] 
**SortOrder** | Pointer to **int32** | Used when sorting a list of context fields. Is also used as a tiebreaker if a list of context fields is sorted alphabetically. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | When this context field was created | [optional] 
**UsedInFeatures** | Pointer to **NullableInt32** | Number of projects where this context field is used in | [optional] 
**UsedInProjects** | Pointer to **NullableInt32** | Number of projects where this context field is used in | [optional] 
**LegalValues** | Pointer to [**[]LegalValueSchema**](LegalValueSchema.md) | Allowed values for this context field schema. Can be used to narrow down accepted input | [optional] 

## Methods

### NewContextFieldSchema

`func NewContextFieldSchema(name string, ) *ContextFieldSchema`

NewContextFieldSchema instantiates a new ContextFieldSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextFieldSchemaWithDefaults

`func NewContextFieldSchemaWithDefaults() *ContextFieldSchema`

NewContextFieldSchemaWithDefaults instantiates a new ContextFieldSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContextFieldSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextFieldSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextFieldSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ContextFieldSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContextFieldSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContextFieldSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContextFieldSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContextFieldSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContextFieldSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStickiness

`func (o *ContextFieldSchema) GetStickiness() bool`

GetStickiness returns the Stickiness field if non-nil, zero value otherwise.

### GetStickinessOk

`func (o *ContextFieldSchema) GetStickinessOk() (*bool, bool)`

GetStickinessOk returns a tuple with the Stickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickiness

`func (o *ContextFieldSchema) SetStickiness(v bool)`

SetStickiness sets Stickiness field to given value.

### HasStickiness

`func (o *ContextFieldSchema) HasStickiness() bool`

HasStickiness returns a boolean if a field has been set.

### GetSortOrder

`func (o *ContextFieldSchema) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ContextFieldSchema) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ContextFieldSchema) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ContextFieldSchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContextFieldSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContextFieldSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContextFieldSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContextFieldSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ContextFieldSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ContextFieldSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUsedInFeatures

`func (o *ContextFieldSchema) GetUsedInFeatures() int32`

GetUsedInFeatures returns the UsedInFeatures field if non-nil, zero value otherwise.

### GetUsedInFeaturesOk

`func (o *ContextFieldSchema) GetUsedInFeaturesOk() (*int32, bool)`

GetUsedInFeaturesOk returns a tuple with the UsedInFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInFeatures

`func (o *ContextFieldSchema) SetUsedInFeatures(v int32)`

SetUsedInFeatures sets UsedInFeatures field to given value.

### HasUsedInFeatures

`func (o *ContextFieldSchema) HasUsedInFeatures() bool`

HasUsedInFeatures returns a boolean if a field has been set.

### SetUsedInFeaturesNil

`func (o *ContextFieldSchema) SetUsedInFeaturesNil(b bool)`

 SetUsedInFeaturesNil sets the value for UsedInFeatures to be an explicit nil

### UnsetUsedInFeatures
`func (o *ContextFieldSchema) UnsetUsedInFeatures()`

UnsetUsedInFeatures ensures that no value is present for UsedInFeatures, not even an explicit nil
### GetUsedInProjects

`func (o *ContextFieldSchema) GetUsedInProjects() int32`

GetUsedInProjects returns the UsedInProjects field if non-nil, zero value otherwise.

### GetUsedInProjectsOk

`func (o *ContextFieldSchema) GetUsedInProjectsOk() (*int32, bool)`

GetUsedInProjectsOk returns a tuple with the UsedInProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInProjects

`func (o *ContextFieldSchema) SetUsedInProjects(v int32)`

SetUsedInProjects sets UsedInProjects field to given value.

### HasUsedInProjects

`func (o *ContextFieldSchema) HasUsedInProjects() bool`

HasUsedInProjects returns a boolean if a field has been set.

### SetUsedInProjectsNil

`func (o *ContextFieldSchema) SetUsedInProjectsNil(b bool)`

 SetUsedInProjectsNil sets the value for UsedInProjects to be an explicit nil

### UnsetUsedInProjects
`func (o *ContextFieldSchema) UnsetUsedInProjects()`

UnsetUsedInProjects ensures that no value is present for UsedInProjects, not even an explicit nil
### GetLegalValues

`func (o *ContextFieldSchema) GetLegalValues() []LegalValueSchema`

GetLegalValues returns the LegalValues field if non-nil, zero value otherwise.

### GetLegalValuesOk

`func (o *ContextFieldSchema) GetLegalValuesOk() (*[]LegalValueSchema, bool)`

GetLegalValuesOk returns a tuple with the LegalValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalValues

`func (o *ContextFieldSchema) SetLegalValues(v []LegalValueSchema)`

SetLegalValues sets LegalValues field to given value.

### HasLegalValues

`func (o *ContextFieldSchema) HasLegalValues() bool`

HasLegalValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


