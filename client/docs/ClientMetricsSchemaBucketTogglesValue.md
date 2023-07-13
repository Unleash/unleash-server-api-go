# ClientMetricsSchemaBucketTogglesValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Yes** | Pointer to **float32** | How many times the toggle evaluated to true | [optional] 
**No** | Pointer to **int32** | How many times the toggle evaluated to false | [optional] 
**Variants** | Pointer to **map[string]int32** | An object describing how many times each variant was returned. Variant names are used as properties, and the number of times they were exposed is the corresponding value (i.e. &#x60;{ [variantName]: number }&#x60;). | [optional] 

## Methods

### NewClientMetricsSchemaBucketTogglesValue

`func NewClientMetricsSchemaBucketTogglesValue() *ClientMetricsSchemaBucketTogglesValue`

NewClientMetricsSchemaBucketTogglesValue instantiates a new ClientMetricsSchemaBucketTogglesValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMetricsSchemaBucketTogglesValueWithDefaults

`func NewClientMetricsSchemaBucketTogglesValueWithDefaults() *ClientMetricsSchemaBucketTogglesValue`

NewClientMetricsSchemaBucketTogglesValueWithDefaults instantiates a new ClientMetricsSchemaBucketTogglesValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYes

`func (o *ClientMetricsSchemaBucketTogglesValue) GetYes() float32`

GetYes returns the Yes field if non-nil, zero value otherwise.

### GetYesOk

`func (o *ClientMetricsSchemaBucketTogglesValue) GetYesOk() (*float32, bool)`

GetYesOk returns a tuple with the Yes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYes

`func (o *ClientMetricsSchemaBucketTogglesValue) SetYes(v float32)`

SetYes sets Yes field to given value.

### HasYes

`func (o *ClientMetricsSchemaBucketTogglesValue) HasYes() bool`

HasYes returns a boolean if a field has been set.

### GetNo

`func (o *ClientMetricsSchemaBucketTogglesValue) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *ClientMetricsSchemaBucketTogglesValue) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *ClientMetricsSchemaBucketTogglesValue) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *ClientMetricsSchemaBucketTogglesValue) HasNo() bool`

HasNo returns a boolean if a field has been set.

### GetVariants

`func (o *ClientMetricsSchemaBucketTogglesValue) GetVariants() map[string]int32`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ClientMetricsSchemaBucketTogglesValue) GetVariantsOk() (*map[string]int32, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ClientMetricsSchemaBucketTogglesValue) SetVariants(v map[string]int32)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ClientMetricsSchemaBucketTogglesValue) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


