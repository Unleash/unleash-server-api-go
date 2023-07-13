# FeatureVariantsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Variants** | [**[]VariantSchema**](VariantSchema.md) |  | 

## Methods

### NewFeatureVariantsSchema

`func NewFeatureVariantsSchema(version int32, variants []VariantSchema, ) *FeatureVariantsSchema`

NewFeatureVariantsSchema instantiates a new FeatureVariantsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureVariantsSchemaWithDefaults

`func NewFeatureVariantsSchemaWithDefaults() *FeatureVariantsSchema`

NewFeatureVariantsSchemaWithDefaults instantiates a new FeatureVariantsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FeatureVariantsSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeatureVariantsSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeatureVariantsSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVariants

`func (o *FeatureVariantsSchema) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *FeatureVariantsSchema) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *FeatureVariantsSchema) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


