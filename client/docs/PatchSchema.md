# PatchSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**Op** | **string** |  | 
**From** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchSchema

`func NewPatchSchema(path string, op string, ) *PatchSchema`

NewPatchSchema instantiates a new PatchSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSchemaWithDefaults

`func NewPatchSchemaWithDefaults() *PatchSchema`

NewPatchSchemaWithDefaults instantiates a new PatchSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PatchSchema) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchSchema) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchSchema) SetPath(v string)`

SetPath sets Path field to given value.


### GetOp

`func (o *PatchSchema) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchSchema) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchSchema) SetOp(v string)`

SetOp sets Op field to given value.


### GetFrom

`func (o *PatchSchema) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PatchSchema) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PatchSchema) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PatchSchema) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetValue

`func (o *PatchSchema) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchSchema) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchSchema) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchSchema) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PatchSchema) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PatchSchema) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


