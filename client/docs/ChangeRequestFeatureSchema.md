# ChangeRequestFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Conflict** | Pointer to **string** |  | [optional] 
**Changes** | [**[]ChangeRequestEventSchema**](ChangeRequestEventSchema.md) | List of changes inside change request. This list may be empty when listing all change requests for a project. | 
**DefaultChange** | Pointer to [**ChangeRequestDefaultEventSchema**](ChangeRequestDefaultEventSchema.md) |  | [optional] 

## Methods

### NewChangeRequestFeatureSchema

`func NewChangeRequestFeatureSchema(name string, changes []ChangeRequestEventSchema, ) *ChangeRequestFeatureSchema`

NewChangeRequestFeatureSchema instantiates a new ChangeRequestFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestFeatureSchemaWithDefaults

`func NewChangeRequestFeatureSchemaWithDefaults() *ChangeRequestFeatureSchema`

NewChangeRequestFeatureSchemaWithDefaults instantiates a new ChangeRequestFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChangeRequestFeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChangeRequestFeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChangeRequestFeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetConflict

`func (o *ChangeRequestFeatureSchema) GetConflict() string`

GetConflict returns the Conflict field if non-nil, zero value otherwise.

### GetConflictOk

`func (o *ChangeRequestFeatureSchema) GetConflictOk() (*string, bool)`

GetConflictOk returns a tuple with the Conflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflict

`func (o *ChangeRequestFeatureSchema) SetConflict(v string)`

SetConflict sets Conflict field to given value.

### HasConflict

`func (o *ChangeRequestFeatureSchema) HasConflict() bool`

HasConflict returns a boolean if a field has been set.

### GetChanges

`func (o *ChangeRequestFeatureSchema) GetChanges() []ChangeRequestEventSchema`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ChangeRequestFeatureSchema) GetChangesOk() (*[]ChangeRequestEventSchema, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ChangeRequestFeatureSchema) SetChanges(v []ChangeRequestEventSchema)`

SetChanges sets Changes field to given value.


### GetDefaultChange

`func (o *ChangeRequestFeatureSchema) GetDefaultChange() ChangeRequestDefaultEventSchema`

GetDefaultChange returns the DefaultChange field if non-nil, zero value otherwise.

### GetDefaultChangeOk

`func (o *ChangeRequestFeatureSchema) GetDefaultChangeOk() (*ChangeRequestDefaultEventSchema, bool)`

GetDefaultChangeOk returns a tuple with the DefaultChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChange

`func (o *ChangeRequestFeatureSchema) SetDefaultChange(v ChangeRequestDefaultEventSchema)`

SetDefaultChange sets DefaultChange field to given value.

### HasDefaultChange

`func (o *ChangeRequestFeatureSchema) HasDefaultChange() bool`

HasDefaultChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


