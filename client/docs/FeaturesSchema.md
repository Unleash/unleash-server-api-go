# FeaturesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Features** | [**[]FeatureSchema**](FeatureSchema.md) |  | 

## Methods

### NewFeaturesSchema

`func NewFeaturesSchema(version int32, features []FeatureSchema, ) *FeaturesSchema`

NewFeaturesSchema instantiates a new FeaturesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturesSchemaWithDefaults

`func NewFeaturesSchemaWithDefaults() *FeaturesSchema`

NewFeaturesSchemaWithDefaults instantiates a new FeaturesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FeaturesSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeaturesSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeaturesSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetFeatures

`func (o *FeaturesSchema) GetFeatures() []FeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FeaturesSchema) GetFeaturesOk() (*[]FeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FeaturesSchema) SetFeatures(v []FeatureSchema)`

SetFeatures sets Features field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


