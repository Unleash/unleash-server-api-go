# TagsBulkAddSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | **[]string** | The list of features that will be affected by the tag changes. | 
**Tags** | [**UpdateTagsSchema**](UpdateTagsSchema.md) |  | 

## Methods

### NewTagsBulkAddSchema

`func NewTagsBulkAddSchema(features []string, tags UpdateTagsSchema, ) *TagsBulkAddSchema`

NewTagsBulkAddSchema instantiates a new TagsBulkAddSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsBulkAddSchemaWithDefaults

`func NewTagsBulkAddSchemaWithDefaults() *TagsBulkAddSchema`

NewTagsBulkAddSchemaWithDefaults instantiates a new TagsBulkAddSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *TagsBulkAddSchema) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TagsBulkAddSchema) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TagsBulkAddSchema) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetTags

`func (o *TagsBulkAddSchema) GetTags() UpdateTagsSchema`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagsBulkAddSchema) GetTagsOk() (*UpdateTagsSchema, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagsBulkAddSchema) SetTags(v UpdateTagsSchema)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


