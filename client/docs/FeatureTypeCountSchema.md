# FeatureTypeCountSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the flag e.g. experiment, kill-switch, release, operational, permission | 
**Count** | **float32** | Number of feature flags of this type | 

## Methods

### NewFeatureTypeCountSchema

`func NewFeatureTypeCountSchema(type_ string, count float32, ) *FeatureTypeCountSchema`

NewFeatureTypeCountSchema instantiates a new FeatureTypeCountSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureTypeCountSchemaWithDefaults

`func NewFeatureTypeCountSchemaWithDefaults() *FeatureTypeCountSchema`

NewFeatureTypeCountSchemaWithDefaults instantiates a new FeatureTypeCountSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FeatureTypeCountSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeatureTypeCountSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeatureTypeCountSchema) SetType(v string)`

SetType sets Type field to given value.


### GetCount

`func (o *FeatureTypeCountSchema) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FeatureTypeCountSchema) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FeatureTypeCountSchema) SetCount(v float32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


