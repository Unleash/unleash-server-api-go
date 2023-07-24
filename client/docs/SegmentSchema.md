# SegmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The segment&#39;s id. | 
**Name** | Pointer to **string** | The name of the segment. | [optional] 
**Description** | Pointer to **NullableString** | The description of the segment. | [optional] 
**Constraints** | [**[]ConstraintSchema**](ConstraintSchema.md) | List of constraints that determine which users are part of the segment | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


