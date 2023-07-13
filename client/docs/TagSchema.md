# TagSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The value of the tag | 
**Type** | **string** | The [type](https://docs.getunleash.io/reference/tags#tag-types) of the tag | [default to "simple"]

## Methods

### NewTagSchema

`func NewTagSchema(value string, type_ string, ) *TagSchema`

NewTagSchema instantiates a new TagSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSchemaWithDefaults

`func NewTagSchemaWithDefaults() *TagSchema`

NewTagSchemaWithDefaults instantiates a new TagSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *TagSchema) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagSchema) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagSchema) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *TagSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagSchema) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


