# TagsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The version of the schema used to model the tags. | 
**Tags** | [**[]TagSchema**](TagSchema.md) | A list of tags. | 

## Methods

### NewTagsSchema

`func NewTagsSchema(version int32, tags []TagSchema, ) *TagsSchema`

NewTagsSchema instantiates a new TagsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsSchemaWithDefaults

`func NewTagsSchemaWithDefaults() *TagsSchema`

NewTagsSchemaWithDefaults instantiates a new TagsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *TagsSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TagsSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TagsSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetTags

`func (o *TagsSchema) GetTags() []TagSchema`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagsSchema) GetTagsOk() (*[]TagSchema, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagsSchema) SetTags(v []TagSchema)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


