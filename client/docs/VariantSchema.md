# VariantSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The variants name. Is unique for this feature toggle | 
**Weight** | **float32** | The weight is the likelihood of any one user getting this variant. It is a number between 0 and 1000. See the section on [variant weights](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight) for more information | 
**WeightType** | Pointer to **string** | Set to fix if this variant must have exactly the weight allocated to it. If the type is variable, the weight will adjust so that the total weight of all variants adds up to 1000 | [optional] 
**Stickiness** | Pointer to **string** | [Stickiness](https://docs.getunleash.io/reference/feature-toggle-variants#variant-stickiness) is how Unleash guarantees that the same user gets the same variant every time | [optional] 
**Payload** | Pointer to [**VariantSchemaPayload**](VariantSchemaPayload.md) |  | [optional] 
**Overrides** | Pointer to [**[]OverrideSchema**](OverrideSchema.md) | Overrides assigning specific variants to specific users. The weighting system automatically assigns users to specific groups for you, but any overrides in this list will take precedence. | [optional] 

## Methods

### NewVariantSchema

`func NewVariantSchema(name string, weight float32, ) *VariantSchema`

NewVariantSchema instantiates a new VariantSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantSchemaWithDefaults

`func NewVariantSchemaWithDefaults() *VariantSchema`

NewVariantSchemaWithDefaults instantiates a new VariantSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VariantSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariantSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariantSchema) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *VariantSchema) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *VariantSchema) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *VariantSchema) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### GetWeightType

`func (o *VariantSchema) GetWeightType() string`

GetWeightType returns the WeightType field if non-nil, zero value otherwise.

### GetWeightTypeOk

`func (o *VariantSchema) GetWeightTypeOk() (*string, bool)`

GetWeightTypeOk returns a tuple with the WeightType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightType

`func (o *VariantSchema) SetWeightType(v string)`

SetWeightType sets WeightType field to given value.

### HasWeightType

`func (o *VariantSchema) HasWeightType() bool`

HasWeightType returns a boolean if a field has been set.

### GetStickiness

`func (o *VariantSchema) GetStickiness() string`

GetStickiness returns the Stickiness field if non-nil, zero value otherwise.

### GetStickinessOk

`func (o *VariantSchema) GetStickinessOk() (*string, bool)`

GetStickinessOk returns a tuple with the Stickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickiness

`func (o *VariantSchema) SetStickiness(v string)`

SetStickiness sets Stickiness field to given value.

### HasStickiness

`func (o *VariantSchema) HasStickiness() bool`

HasStickiness returns a boolean if a field has been set.

### GetPayload

`func (o *VariantSchema) GetPayload() VariantSchemaPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *VariantSchema) GetPayloadOk() (*VariantSchemaPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *VariantSchema) SetPayload(v VariantSchemaPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *VariantSchema) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetOverrides

`func (o *VariantSchema) GetOverrides() []OverrideSchema`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *VariantSchema) GetOverridesOk() (*[]OverrideSchema, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *VariantSchema) SetOverrides(v []OverrideSchema)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *VariantSchema) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


