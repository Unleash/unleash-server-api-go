# ProxyFeatureSchemaVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Enabled** | **bool** |  | 
**Payload** | Pointer to [**ProxyFeatureSchemaVariantPayload**](ProxyFeatureSchemaVariantPayload.md) |  | [optional] 

## Methods

### NewProxyFeatureSchemaVariant

`func NewProxyFeatureSchemaVariant(name string, enabled bool, ) *ProxyFeatureSchemaVariant`

NewProxyFeatureSchemaVariant instantiates a new ProxyFeatureSchemaVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyFeatureSchemaVariantWithDefaults

`func NewProxyFeatureSchemaVariantWithDefaults() *ProxyFeatureSchemaVariant`

NewProxyFeatureSchemaVariantWithDefaults instantiates a new ProxyFeatureSchemaVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProxyFeatureSchemaVariant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProxyFeatureSchemaVariant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProxyFeatureSchemaVariant) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *ProxyFeatureSchemaVariant) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProxyFeatureSchemaVariant) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProxyFeatureSchemaVariant) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPayload

`func (o *ProxyFeatureSchemaVariant) GetPayload() ProxyFeatureSchemaVariantPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ProxyFeatureSchemaVariant) GetPayloadOk() (*ProxyFeatureSchemaVariantPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ProxyFeatureSchemaVariant) SetPayload(v ProxyFeatureSchemaVariantPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ProxyFeatureSchemaVariant) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


