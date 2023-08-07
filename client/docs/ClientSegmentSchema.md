# ClientSegmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The segment&#39;s id. | 
**Name** | Pointer to **string** | The name of the segment. | [optional] 
**Constraints** | [**[]ConstraintSchema**](ConstraintSchema.md) | List of constraints that determine which users are part of the segment | 

## Methods

### NewClientSegmentSchema

`func NewClientSegmentSchema(id float32, constraints []ConstraintSchema, ) *ClientSegmentSchema`

NewClientSegmentSchema instantiates a new ClientSegmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSegmentSchemaWithDefaults

`func NewClientSegmentSchemaWithDefaults() *ClientSegmentSchema`

NewClientSegmentSchemaWithDefaults instantiates a new ClientSegmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientSegmentSchema) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientSegmentSchema) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientSegmentSchema) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *ClientSegmentSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientSegmentSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientSegmentSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientSegmentSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConstraints

`func (o *ClientSegmentSchema) GetConstraints() []ConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *ClientSegmentSchema) GetConstraintsOk() (*[]ConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *ClientSegmentSchema) SetConstraints(v []ConstraintSchema)`

SetConstraints sets Constraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


