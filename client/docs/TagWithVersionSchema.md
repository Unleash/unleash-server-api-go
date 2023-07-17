# TagWithVersionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The version of the schema used to model the tag. | 
**Tag** | [**TagSchema**](TagSchema.md) |  | 

## Methods

### NewTagWithVersionSchema

`func NewTagWithVersionSchema(version int32, tag TagSchema, ) *TagWithVersionSchema`

NewTagWithVersionSchema instantiates a new TagWithVersionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithVersionSchemaWithDefaults

`func NewTagWithVersionSchemaWithDefaults() *TagWithVersionSchema`

NewTagWithVersionSchemaWithDefaults instantiates a new TagWithVersionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *TagWithVersionSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TagWithVersionSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TagWithVersionSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetTag

`func (o *TagWithVersionSchema) GetTag() TagSchema`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TagWithVersionSchema) GetTagOk() (*TagSchema, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TagWithVersionSchema) SetTag(v TagSchema)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


