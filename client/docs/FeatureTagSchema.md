# FeatureTagSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureName** | **string** | The name of the feature this tag is applied to | 
**TagType** | Pointer to **string** | The [type](https://docs.getunleash.io/reference/tags#tag-types tag types) of the tag | [optional] 
**TagValue** | **string** | The value of the tag | 
**Type** | Pointer to **string** | The [type](https://docs.getunleash.io/reference/tags#tag-types tag types) of the tag. This property is deprecated and will be removed in a future version of Unleash. Superseded by the &#x60;tagType&#x60; property. | [optional] 
**Value** | Pointer to **string** | The value of the tag. This property is deprecated and will be removed in a future version of Unleash. Superseded by the &#x60;tagValue&#x60; property. | [optional] 

## Methods

### NewFeatureTagSchema

`func NewFeatureTagSchema(featureName string, tagValue string, ) *FeatureTagSchema`

NewFeatureTagSchema instantiates a new FeatureTagSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureTagSchemaWithDefaults

`func NewFeatureTagSchemaWithDefaults() *FeatureTagSchema`

NewFeatureTagSchemaWithDefaults instantiates a new FeatureTagSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureName

`func (o *FeatureTagSchema) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *FeatureTagSchema) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *FeatureTagSchema) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.


### GetTagType

`func (o *FeatureTagSchema) GetTagType() string`

GetTagType returns the TagType field if non-nil, zero value otherwise.

### GetTagTypeOk

`func (o *FeatureTagSchema) GetTagTypeOk() (*string, bool)`

GetTagTypeOk returns a tuple with the TagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagType

`func (o *FeatureTagSchema) SetTagType(v string)`

SetTagType sets TagType field to given value.

### HasTagType

`func (o *FeatureTagSchema) HasTagType() bool`

HasTagType returns a boolean if a field has been set.

### GetTagValue

`func (o *FeatureTagSchema) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *FeatureTagSchema) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *FeatureTagSchema) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.


### GetType

`func (o *FeatureTagSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeatureTagSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeatureTagSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeatureTagSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *FeatureTagSchema) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FeatureTagSchema) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FeatureTagSchema) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FeatureTagSchema) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


