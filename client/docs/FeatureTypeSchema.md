# FeatureTypeSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**LifetimeDays** | **NullableFloat32** |  | 

## Methods

### NewFeatureTypeSchema

`func NewFeatureTypeSchema(id string, name string, description string, lifetimeDays NullableFloat32, ) *FeatureTypeSchema`

NewFeatureTypeSchema instantiates a new FeatureTypeSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureTypeSchemaWithDefaults

`func NewFeatureTypeSchemaWithDefaults() *FeatureTypeSchema`

NewFeatureTypeSchemaWithDefaults instantiates a new FeatureTypeSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeatureTypeSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureTypeSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureTypeSchema) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FeatureTypeSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureTypeSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureTypeSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FeatureTypeSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureTypeSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureTypeSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLifetimeDays

`func (o *FeatureTypeSchema) GetLifetimeDays() float32`

GetLifetimeDays returns the LifetimeDays field if non-nil, zero value otherwise.

### GetLifetimeDaysOk

`func (o *FeatureTypeSchema) GetLifetimeDaysOk() (*float32, bool)`

GetLifetimeDaysOk returns a tuple with the LifetimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeDays

`func (o *FeatureTypeSchema) SetLifetimeDays(v float32)`

SetLifetimeDays sets LifetimeDays field to given value.


### SetLifetimeDaysNil

`func (o *FeatureTypeSchema) SetLifetimeDaysNil(b bool)`

 SetLifetimeDaysNil sets the value for LifetimeDays to be an explicit nil

### UnsetLifetimeDays
`func (o *FeatureTypeSchema) UnsetLifetimeDays()`

UnsetLifetimeDays ensures that no value is present for LifetimeDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


