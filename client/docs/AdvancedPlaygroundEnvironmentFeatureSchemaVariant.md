# AdvancedPlaygroundEnvironmentFeatureSchemaVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The variant&#39;s name. If there is no variant or if the toggle is disabled, this will be &#x60;disabled&#x60; | 
**Enabled** | **bool** | Whether the variant is enabled or not. If the feature is disabled or if it doesn&#39;t have variants, this property will be &#x60;false&#x60; | 
**Payload** | Pointer to [**AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload**](AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload.md) |  | [optional] 

## Methods

### NewAdvancedPlaygroundEnvironmentFeatureSchemaVariant

`func NewAdvancedPlaygroundEnvironmentFeatureSchemaVariant(name string, enabled bool, ) *AdvancedPlaygroundEnvironmentFeatureSchemaVariant`

NewAdvancedPlaygroundEnvironmentFeatureSchemaVariant instantiates a new AdvancedPlaygroundEnvironmentFeatureSchemaVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedPlaygroundEnvironmentFeatureSchemaVariantWithDefaults

`func NewAdvancedPlaygroundEnvironmentFeatureSchemaVariantWithDefaults() *AdvancedPlaygroundEnvironmentFeatureSchemaVariant`

NewAdvancedPlaygroundEnvironmentFeatureSchemaVariantWithDefaults instantiates a new AdvancedPlaygroundEnvironmentFeatureSchemaVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPayload

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetPayload() AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetPayloadOk() (*AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) SetPayload(v AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

