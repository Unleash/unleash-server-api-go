# UiConfigSchemaFlagsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the variant. Will always be disabled if &#x60;enabled&#x60; is false. | [optional] 
**Enabled** | Pointer to **bool** | Whether the variant is enabled or not. | [optional] 
**Payload** | Pointer to [**VariantFlagSchemaPayload**](VariantFlagSchemaPayload.md) |  | [optional] 

## Methods

### NewUiConfigSchemaFlagsValue

`func NewUiConfigSchemaFlagsValue() *UiConfigSchemaFlagsValue`

NewUiConfigSchemaFlagsValue instantiates a new UiConfigSchemaFlagsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiConfigSchemaFlagsValueWithDefaults

`func NewUiConfigSchemaFlagsValueWithDefaults() *UiConfigSchemaFlagsValue`

NewUiConfigSchemaFlagsValueWithDefaults instantiates a new UiConfigSchemaFlagsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UiConfigSchemaFlagsValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UiConfigSchemaFlagsValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UiConfigSchemaFlagsValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UiConfigSchemaFlagsValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UiConfigSchemaFlagsValue) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UiConfigSchemaFlagsValue) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UiConfigSchemaFlagsValue) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UiConfigSchemaFlagsValue) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPayload

`func (o *UiConfigSchemaFlagsValue) GetPayload() VariantFlagSchemaPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UiConfigSchemaFlagsValue) GetPayloadOk() (*VariantFlagSchemaPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UiConfigSchemaFlagsValue) SetPayload(v VariantFlagSchemaPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UiConfigSchemaFlagsValue) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


