# PlaygroundSegmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The segment&#39;s id. | 
**Name** | **string** | The name of the segment. | 
**Result** | **bool** | Whether this was evaluated as true or false. | 
**Constraints** | [**[]PlaygroundConstraintSchema**](PlaygroundConstraintSchema.md) | The list of constraints in this segment. | 

## Methods

### NewPlaygroundSegmentSchema

`func NewPlaygroundSegmentSchema(id int32, name string, result bool, constraints []PlaygroundConstraintSchema, ) *PlaygroundSegmentSchema`

NewPlaygroundSegmentSchema instantiates a new PlaygroundSegmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundSegmentSchemaWithDefaults

`func NewPlaygroundSegmentSchemaWithDefaults() *PlaygroundSegmentSchema`

NewPlaygroundSegmentSchemaWithDefaults instantiates a new PlaygroundSegmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlaygroundSegmentSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaygroundSegmentSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaygroundSegmentSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PlaygroundSegmentSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlaygroundSegmentSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlaygroundSegmentSchema) SetName(v string)`

SetName sets Name field to given value.


### GetResult

`func (o *PlaygroundSegmentSchema) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PlaygroundSegmentSchema) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PlaygroundSegmentSchema) SetResult(v bool)`

SetResult sets Result field to given value.


### GetConstraints

`func (o *PlaygroundSegmentSchema) GetConstraints() []PlaygroundConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *PlaygroundSegmentSchema) GetConstraintsOk() (*[]PlaygroundConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *PlaygroundSegmentSchema) SetConstraints(v []PlaygroundConstraintSchema)`

SetConstraints sets Constraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


