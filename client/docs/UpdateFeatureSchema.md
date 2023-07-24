# UpdateFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Stale** | Pointer to **bool** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ImpressionData** | Pointer to **bool** |  | [optional] 
**Constraints** | Pointer to [**[]ConstraintSchema**](ConstraintSchema.md) |  | [optional] 

## Methods

### NewUpdateFeatureSchema

`func NewUpdateFeatureSchema(name string, ) *UpdateFeatureSchema`

NewUpdateFeatureSchema instantiates a new UpdateFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFeatureSchemaWithDefaults

`func NewUpdateFeatureSchemaWithDefaults() *UpdateFeatureSchema`

NewUpdateFeatureSchemaWithDefaults instantiates a new UpdateFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateFeatureSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFeatureSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFeatureSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFeatureSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *UpdateFeatureSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateFeatureSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateFeatureSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateFeatureSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStale

`func (o *UpdateFeatureSchema) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *UpdateFeatureSchema) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *UpdateFeatureSchema) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *UpdateFeatureSchema) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetArchived

`func (o *UpdateFeatureSchema) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateFeatureSchema) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateFeatureSchema) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateFeatureSchema) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UpdateFeatureSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateFeatureSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateFeatureSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdateFeatureSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetImpressionData

`func (o *UpdateFeatureSchema) GetImpressionData() bool`

GetImpressionData returns the ImpressionData field if non-nil, zero value otherwise.

### GetImpressionDataOk

`func (o *UpdateFeatureSchema) GetImpressionDataOk() (*bool, bool)`

GetImpressionDataOk returns a tuple with the ImpressionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressionData

`func (o *UpdateFeatureSchema) SetImpressionData(v bool)`

SetImpressionData sets ImpressionData field to given value.

### HasImpressionData

`func (o *UpdateFeatureSchema) HasImpressionData() bool`

HasImpressionData returns a boolean if a field has been set.

### GetConstraints

`func (o *UpdateFeatureSchema) GetConstraints() []ConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *UpdateFeatureSchema) GetConstraintsOk() (*[]ConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *UpdateFeatureSchema) SetConstraints(v []ConstraintSchema)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *UpdateFeatureSchema) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


