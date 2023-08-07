# ProxyFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique feature name. | 
**Enabled** | **bool** | Always set to &#x60;true&#x60;. | 
**ImpressionData** | **bool** | &#x60;true&#x60; if the impression data collection is enabled for the feature, otherwise &#x60;false&#x60;. | 
**Variant** | Pointer to [**ProxyFeatureSchemaVariant**](ProxyFeatureSchemaVariant.md) |  | [optional] 

## Methods

### NewProxyFeatureSchema

`func NewProxyFeatureSchema(name string, enabled bool, impressionData bool, ) *ProxyFeatureSchema`

NewProxyFeatureSchema instantiates a new ProxyFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyFeatureSchemaWithDefaults

`func NewProxyFeatureSchemaWithDefaults() *ProxyFeatureSchema`

NewProxyFeatureSchemaWithDefaults instantiates a new ProxyFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProxyFeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProxyFeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProxyFeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *ProxyFeatureSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProxyFeatureSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProxyFeatureSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetImpressionData

`func (o *ProxyFeatureSchema) GetImpressionData() bool`

GetImpressionData returns the ImpressionData field if non-nil, zero value otherwise.

### GetImpressionDataOk

`func (o *ProxyFeatureSchema) GetImpressionDataOk() (*bool, bool)`

GetImpressionDataOk returns a tuple with the ImpressionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressionData

`func (o *ProxyFeatureSchema) SetImpressionData(v bool)`

SetImpressionData sets ImpressionData field to given value.


### GetVariant

`func (o *ProxyFeatureSchema) GetVariant() ProxyFeatureSchemaVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ProxyFeatureSchema) GetVariantOk() (*ProxyFeatureSchemaVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ProxyFeatureSchema) SetVariant(v ProxyFeatureSchemaVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ProxyFeatureSchema) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


