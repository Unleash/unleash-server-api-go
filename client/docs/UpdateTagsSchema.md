# UpdateTagsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedTags** | [**[]TagSchema**](TagSchema.md) | Tags to add to the feature. | 
**RemovedTags** | [**[]TagSchema**](TagSchema.md) | Tags to remove from the feature. | 

## Methods

### NewUpdateTagsSchema

`func NewUpdateTagsSchema(addedTags []TagSchema, removedTags []TagSchema, ) *UpdateTagsSchema`

NewUpdateTagsSchema instantiates a new UpdateTagsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTagsSchemaWithDefaults

`func NewUpdateTagsSchemaWithDefaults() *UpdateTagsSchema`

NewUpdateTagsSchemaWithDefaults instantiates a new UpdateTagsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedTags

`func (o *UpdateTagsSchema) GetAddedTags() []TagSchema`

GetAddedTags returns the AddedTags field if non-nil, zero value otherwise.

### GetAddedTagsOk

`func (o *UpdateTagsSchema) GetAddedTagsOk() (*[]TagSchema, bool)`

GetAddedTagsOk returns a tuple with the AddedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedTags

`func (o *UpdateTagsSchema) SetAddedTags(v []TagSchema)`

SetAddedTags sets AddedTags field to given value.


### GetRemovedTags

`func (o *UpdateTagsSchema) GetRemovedTags() []TagSchema`

GetRemovedTags returns the RemovedTags field if non-nil, zero value otherwise.

### GetRemovedTagsOk

`func (o *UpdateTagsSchema) GetRemovedTagsOk() (*[]TagSchema, bool)`

GetRemovedTagsOk returns a tuple with the RemovedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedTags

`func (o *UpdateTagsSchema) SetRemovedTags(v []TagSchema)`

SetRemovedTags sets RemovedTags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


