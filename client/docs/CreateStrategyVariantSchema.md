# CreateStrategyVariantSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The variant name. Must be unique for this feature toggle | 
**Weight** | **int32** | The weight is the likelihood of any one user getting this variant. It is an integer between 0 and 1000. See the section on [variant weights](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight) for more information | 
**WeightType** | **string** | Set to &#x60;fix&#x60; if this variant must have exactly the weight allocated to it. If the type is &#x60;variable&#x60;, the weight will adjust so that the total weight of all variants adds up to 1000. Refer to the [variant weight documentation](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight). | 
**Stickiness** | **string** | The [stickiness](https://docs.getunleash.io/reference/feature-toggle-variants#variant-stickiness) to use for distribution of this variant. Stickiness is how Unleash guarantees that the same user gets the same variant every time | 
**Payload** | Pointer to [**StrategyVariantSchemaPayload**](StrategyVariantSchemaPayload.md) |  | [optional] 

## Methods

### NewCreateStrategyVariantSchema

`func NewCreateStrategyVariantSchema(name string, weight int32, weightType string, stickiness string, ) *CreateStrategyVariantSchema`

NewCreateStrategyVariantSchema instantiates a new CreateStrategyVariantSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStrategyVariantSchemaWithDefaults

`func NewCreateStrategyVariantSchemaWithDefaults() *CreateStrategyVariantSchema`

NewCreateStrategyVariantSchemaWithDefaults instantiates a new CreateStrategyVariantSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateStrategyVariantSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStrategyVariantSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStrategyVariantSchema) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *CreateStrategyVariantSchema) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateStrategyVariantSchema) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateStrategyVariantSchema) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetWeightType

`func (o *CreateStrategyVariantSchema) GetWeightType() string`

GetWeightType returns the WeightType field if non-nil, zero value otherwise.

### GetWeightTypeOk

`func (o *CreateStrategyVariantSchema) GetWeightTypeOk() (*string, bool)`

GetWeightTypeOk returns a tuple with the WeightType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightType

`func (o *CreateStrategyVariantSchema) SetWeightType(v string)`

SetWeightType sets WeightType field to given value.


### GetStickiness

`func (o *CreateStrategyVariantSchema) GetStickiness() string`

GetStickiness returns the Stickiness field if non-nil, zero value otherwise.

### GetStickinessOk

`func (o *CreateStrategyVariantSchema) GetStickinessOk() (*string, bool)`

GetStickinessOk returns a tuple with the Stickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickiness

`func (o *CreateStrategyVariantSchema) SetStickiness(v string)`

SetStickiness sets Stickiness field to given value.


### GetPayload

`func (o *CreateStrategyVariantSchema) GetPayload() StrategyVariantSchemaPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CreateStrategyVariantSchema) GetPayloadOk() (*StrategyVariantSchemaPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CreateStrategyVariantSchema) SetPayload(v StrategyVariantSchemaPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CreateStrategyVariantSchema) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


