# VariantSchemaPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the value. Commonly used types are string, json and csv. | 
**Value** | **string** | The actual value of payload | 

## Methods

### NewVariantSchemaPayload

`func NewVariantSchemaPayload(type_ string, value string, ) *VariantSchemaPayload`

NewVariantSchemaPayload instantiates a new VariantSchemaPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantSchemaPayloadWithDefaults

`func NewVariantSchemaPayloadWithDefaults() *VariantSchemaPayload`

NewVariantSchemaPayloadWithDefaults instantiates a new VariantSchemaPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VariantSchemaPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VariantSchemaPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VariantSchemaPayload) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *VariantSchemaPayload) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariantSchemaPayload) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariantSchemaPayload) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


