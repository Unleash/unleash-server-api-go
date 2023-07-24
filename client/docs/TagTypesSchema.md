# TagTypesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**TagTypes** | [**[]TagTypeSchema**](TagTypeSchema.md) |  | 

## Methods

### NewTagTypesSchema

`func NewTagTypesSchema(version int32, tagTypes []TagTypeSchema, ) *TagTypesSchema`

NewTagTypesSchema instantiates a new TagTypesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagTypesSchemaWithDefaults

`func NewTagTypesSchemaWithDefaults() *TagTypesSchema`

NewTagTypesSchemaWithDefaults instantiates a new TagTypesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *TagTypesSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TagTypesSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TagTypesSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetTagTypes

`func (o *TagTypesSchema) GetTagTypes() []TagTypeSchema`

GetTagTypes returns the TagTypes field if non-nil, zero value otherwise.

### GetTagTypesOk

`func (o *TagTypesSchema) GetTagTypesOk() (*[]TagTypeSchema, bool)`

GetTagTypesOk returns a tuple with the TagTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypes

`func (o *TagTypesSchema) SetTagTypes(v []TagTypeSchema)`

SetTagTypes sets TagTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


