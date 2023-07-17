# AddonsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | [**[]AddonSchema**](AddonSchema.md) | All the addons that exist on this instance of Unleash. | 
**Providers** | [**[]AddonTypeSchema**](AddonTypeSchema.md) | A list of  all available addon providers, along with their parameters and descriptions. | 

## Methods

### NewAddonsSchema

`func NewAddonsSchema(addons []AddonSchema, providers []AddonTypeSchema, ) *AddonsSchema`

NewAddonsSchema instantiates a new AddonsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsSchemaWithDefaults

`func NewAddonsSchemaWithDefaults() *AddonsSchema`

NewAddonsSchemaWithDefaults instantiates a new AddonsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *AddonsSchema) GetAddons() []AddonSchema`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *AddonsSchema) GetAddonsOk() (*[]AddonSchema, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *AddonsSchema) SetAddons(v []AddonSchema)`

SetAddons sets Addons field to given value.


### GetProviders

`func (o *AddonsSchema) GetProviders() []AddonTypeSchema`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *AddonsSchema) GetProvidersOk() (*[]AddonTypeSchema, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *AddonsSchema) SetProviders(v []AddonTypeSchema)`

SetProviders sets Providers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


