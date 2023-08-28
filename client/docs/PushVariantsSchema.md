# PushVariantsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variants** | Pointer to [**[]VariantSchema**](VariantSchema.md) | The variants to write to the provided environments | [optional] 
**Environments** | Pointer to **[]string** | The enviromnents to write the provided variants to | [optional] 

## Methods

### NewPushVariantsSchema

`func NewPushVariantsSchema() *PushVariantsSchema`

NewPushVariantsSchema instantiates a new PushVariantsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushVariantsSchemaWithDefaults

`func NewPushVariantsSchemaWithDefaults() *PushVariantsSchema`

NewPushVariantsSchemaWithDefaults instantiates a new PushVariantsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariants

`func (o *PushVariantsSchema) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *PushVariantsSchema) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *PushVariantsSchema) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *PushVariantsSchema) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetEnvironments

`func (o *PushVariantsSchema) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PushVariantsSchema) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PushVariantsSchema) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PushVariantsSchema) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


