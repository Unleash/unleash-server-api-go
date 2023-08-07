# UpdateFeatureTypeLifetimeSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifetimeDays** | **NullableInt32** | The new lifetime (in days) that you want to assign to the feature toggle type. If the value is &#x60;null&#x60; or &#x60;0&#x60;, then the feature toggles of that type will never be marked as potentially stale. Otherwise, they will be considered potentially stale after the number of days indicated by this property. | 

## Methods

### NewUpdateFeatureTypeLifetimeSchema

`func NewUpdateFeatureTypeLifetimeSchema(lifetimeDays NullableInt32, ) *UpdateFeatureTypeLifetimeSchema`

NewUpdateFeatureTypeLifetimeSchema instantiates a new UpdateFeatureTypeLifetimeSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFeatureTypeLifetimeSchemaWithDefaults

`func NewUpdateFeatureTypeLifetimeSchemaWithDefaults() *UpdateFeatureTypeLifetimeSchema`

NewUpdateFeatureTypeLifetimeSchemaWithDefaults instantiates a new UpdateFeatureTypeLifetimeSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifetimeDays

`func (o *UpdateFeatureTypeLifetimeSchema) GetLifetimeDays() int32`

GetLifetimeDays returns the LifetimeDays field if non-nil, zero value otherwise.

### GetLifetimeDaysOk

`func (o *UpdateFeatureTypeLifetimeSchema) GetLifetimeDaysOk() (*int32, bool)`

GetLifetimeDaysOk returns a tuple with the LifetimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeDays

`func (o *UpdateFeatureTypeLifetimeSchema) SetLifetimeDays(v int32)`

SetLifetimeDays sets LifetimeDays field to given value.


### SetLifetimeDaysNil

`func (o *UpdateFeatureTypeLifetimeSchema) SetLifetimeDaysNil(b bool)`

 SetLifetimeDaysNil sets the value for LifetimeDays to be an explicit nil

### UnsetLifetimeDays
`func (o *UpdateFeatureTypeLifetimeSchema) UnsetLifetimeDays()`

UnsetLifetimeDays ensures that no value is present for LifetimeDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


