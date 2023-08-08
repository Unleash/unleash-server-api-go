# FeatureTypesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The schema version used to describe the feature toggle types listed in the &#x60;types&#x60; property. | 
**Types** | [**[]FeatureTypeSchema**](FeatureTypeSchema.md) | The list of feature toggle types. | 

## Methods

### NewFeatureTypesSchema

`func NewFeatureTypesSchema(version int32, types []FeatureTypeSchema, ) *FeatureTypesSchema`

NewFeatureTypesSchema instantiates a new FeatureTypesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureTypesSchemaWithDefaults

`func NewFeatureTypesSchemaWithDefaults() *FeatureTypesSchema`

NewFeatureTypesSchemaWithDefaults instantiates a new FeatureTypesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FeatureTypesSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeatureTypesSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeatureTypesSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetTypes

`func (o *FeatureTypesSchema) GetTypes() []FeatureTypeSchema`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *FeatureTypesSchema) GetTypesOk() (*[]FeatureTypeSchema, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *FeatureTypesSchema) SetTypes(v []FeatureTypeSchema)`

SetTypes sets Types field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


