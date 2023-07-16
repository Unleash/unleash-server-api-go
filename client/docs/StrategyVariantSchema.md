# StrategyVariantSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The variant name. Must be unique for this feature toggle | 
**Weight** | **int32** | The weight is the likelihood of any one user getting this variant. It is an integer between 0 and 1000. See the section on [variant weights](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight) for more information | 
**WeightType** | **string** | Set to &#x60;fix&#x60; if this variant must have exactly the weight allocated to it. If the type is &#x60;variable&#x60;, the weight will adjust so that the total weight of all variants adds up to 1000. Refer to the [variant weight documentation](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight). | 
**Stickiness** | **string** | The [stickiness](https://docs.getunleash.io/reference/feature-toggle-variants#variant-stickiness) to use for distribution of this variant. Stickiness is how Unleash guarantees that the same user gets the same variant every time | 
**Payload** | Pointer to [**VariantSchemaPayload**](VariantSchemaPayload.md) |  | [optional] 

## Methods

### NewStrategyVariantSchema

`func NewStrategyVariantSchema(name string, weight int32, weightType string, stickiness string, ) *StrategyVariantSchema`

NewStrategyVariantSchema instantiates a new StrategyVariantSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyVariantSchemaWithDefaults

`func NewStrategyVariantSchemaWithDefaults() *StrategyVariantSchema`

NewStrategyVariantSchemaWithDefaults instantiates a new StrategyVariantSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StrategyVariantSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StrategyVariantSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StrategyVariantSchema) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *StrategyVariantSchema) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *StrategyVariantSchema) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *StrategyVariantSchema) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetWeightType

`func (o *StrategyVariantSchema) GetWeightType() string`

GetWeightType returns the WeightType field if non-nil, zero value otherwise.

### GetWeightTypeOk

`func (o *StrategyVariantSchema) GetWeightTypeOk() (*string, bool)`

GetWeightTypeOk returns a tuple with the WeightType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightType

`func (o *StrategyVariantSchema) SetWeightType(v string)`

SetWeightType sets WeightType field to given value.


### GetStickiness

`func (o *StrategyVariantSchema) GetStickiness() string`

GetStickiness returns the Stickiness field if non-nil, zero value otherwise.

### GetStickinessOk

`func (o *StrategyVariantSchema) GetStickinessOk() (*string, bool)`

GetStickinessOk returns a tuple with the Stickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickiness

`func (o *StrategyVariantSchema) SetStickiness(v string)`

SetStickiness sets Stickiness field to given value.


### GetPayload

`func (o *StrategyVariantSchema) GetPayload() VariantSchemaPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *StrategyVariantSchema) GetPayloadOk() (*VariantSchemaPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *StrategyVariantSchema) SetPayload(v VariantSchemaPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *StrategyVariantSchema) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


