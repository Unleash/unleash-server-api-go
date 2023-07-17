# VariantSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The variants name. Is unique for this feature toggle | [default to null]
**Weight** | **float64** | The weight is the likelihood of any one user getting this variant. It is a number between 0 and 1000. See the section on [variant weights](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight) for more information | [default to null]
**WeightType** | **string** | Set to fix if this variant must have exactly the weight allocated to it. If the type is variable, the weight will adjust so that the total weight of all variants adds up to 1000 | [optional] [default to null]
**Stickiness** | **string** | [Stickiness](https://docs.getunleash.io/reference/feature-toggle-variants#variant-stickiness) is how Unleash guarantees that the same user gets the same variant every time | [optional] [default to null]
**Payload** | [***VariantSchemaPayload**](variantSchema_payload.md) |  | [optional] [default to null]
**Overrides** | [**[]OverrideSchema**](overrideSchema.md) | Overrides assigning specific variants to specific users. The weighting system automatically assigns users to specific groups for you, but any overrides in this list will take precedence. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

